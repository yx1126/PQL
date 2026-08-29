import jp from "jsonpath";
import { isArray, isObject, isStr } from "./validata";
import { AppService } from "@bind/service";
import { Res } from "./response";
import type { Result } from "#/axios";
import type { ParseOption, RuleItem } from "@bind/parse/video";

export {
    type ParseOption,
    type RuleItem,
};

// eslint-disable-next-line @typescript-eslint/no-namespace
export namespace Parse {
    export type Data<T = any> = Record<string, T>;

    export interface ResRule {
        codeok?: string | number;
        codePath?: string;
        dataPath?: string;
        messagePath?: string;
    }

    /**
     * @see https://github.com/dchester/jsonpath#methods
     * type=auto 数组[query] 其他[value]
     */
    export interface JsonParse {
        type?: string;
        extra?: any[];
    }
}

// jsonpath 数组解析
export function parseArray<T = any>(data: unknown, parse?: string, ...args: any[]): T[] {
    if(!parse) return [];
    try {
        return jp.query(data, parse, ...args);
    } catch (error) {
        console.error(error);
        return [];
    }
}

// jsonpath 值解析
export function parseValue<T = any>(data: unknown, parse?: string): Noable<T> {
    if(!parse) return;
    try {
        return jp.value(data, parse);
    } catch (error) {
        console.error(error);
    }
}

// jsonpath parent解析
export function parseParent<T = any>(data: unknown, parse?: string): Noable<T> {
    if(!parse) return;
    try {
        return jp.parent(data, parse);
    } catch (error) {
        console.error(error);
    }
}

// jsonpath 值解析
export function parseAuto<T = any>(data: unknown, parse?: string, ...args: any[]): Noable<T> {
    if(isArray(data)) return parseArray(data, parse, ...args) as T;
    return parseValue(data, parse);
}

export function parse<T = any>(data?: unknown, path?: string, config?: string | Parse.JsonParse): Noable<T> {
    if(!path) return data as T;
    let type = "";
    let extra: Parse.JsonParse["extra"];
    if(isObject<Parse.JsonParse>(config)) {
        type = config["type"] || "";
        extra = config["extra"];
    } else if(isStr(config)) {
        type = config;
    }
    try {
        switch(type) {
        case "auto":
            return parseAuto(data, path, ...extra || []);
        case "query":
            return parseArray(data, path, ...extra || []) as T;
        case "value":
            return parseValue(data, path);
        case "parent":
            return parseParent(data, path);
        default:
            return parseAuto(data, path, ...extra || []);
        }
    } catch (error) {
        console.error(error);
        return data as T;
    }
}

// 模板解析
export function parseTemp(data: Parse.Data, template: string) {
    return template.replace(/\{(.*?)(\|(.*))?\}/g, (_, key, _2, defValue) => {
        return data[key] || defValue || "";
    });
}

/**
 * 路径参数解析
 *
 * @example
 *
 * const data = { page: 1, size: 10 };
 *
 * parseUri("https://baidu.com?page=@page&size=@size&a=c&b=@@page", data)
 * // https://baidu.com?page=1&size=10&a=c&b=@page
 *
 * parseUri("https://baidu.com/@page/@size", data)
 * // https://baidu.com/1/10
 */
export function parseUri(url: string, params?: Parse.Data, omitempty?: boolean) {
    try {
        if(!params || !Object.keys(params).length) return url;
        const [prefix, str] = url.split("?");

        // 解析路径参数
        const newUrl = !prefix?.includes("@")
            ? prefix
            : prefix.split("/").map(v => {
                if(!v.startsWith("@")) return v;
                return params[v.substring(1)] ?? v;
            }).join("/");

        if(!str || !str?.includes("@")) return newUrl + "?" + (str || "");

        // 解析?后参数
        let suffix = "?";
        str.split("&").forEach((v, i) => {
            if(i > 0) suffix += "&";
            const [k, value] = v.split("=");
            if(!value.startsWith("@")) {
                suffix += v;
                return;
            }
            if(value.startsWith("@@")) {
                suffix += `${k}=${value.substring(1)}`;
                return;
            }
            const kv = params[value.substring(1)] || "";
            if(omitempty && !kv) {
                return;
            }
            suffix += `${k}=${kv}`;
        });
        return newUrl + suffix;
    } catch (error) {
        console.error(error);
        return url;
    }
}

/**
 * 请求体参数解析
 *
 * @example
 * const params = { p: 1, s: 10 }
 * const rules = { page: "@p", size: "@s", a: 1, b: "@b" }
 * parseBodyData(params, data)
 * // {page: 1, size: 10, a: 1, b: "@b"}
 */
export function parseBodyData(params?: Parse.Data, rules?: RuleItem["data"], omitempty?: boolean) {
    if(!params || !Object.keys(params).length) return rules;
    return Object.keys(rules || {}).reduce((pre, key) => {
        const path = rules![key];
        if(!isStr(path) || !path?.startsWith("@")) {
            pre[key] = path;
        } else {
            const pk = path.substring(1);
            if(path.startsWith("@@")) {
                pre[key] = pk;
            } else {
                const v = params[pk];
                if(omitempty) {
                    // eslint-disable-next-line no-extra-boolean-cast
                    if(!!v) pre[key] = v;
                } else {
                    pre[key] = v;
                }
            }
        }
        return pre;
    }, {} as Parse.Data);
}

export function parseParams(data?: Parse.Data, params?: Nullable<Parse.Data<string>>) {
    if(!data || !params) return {};
    return Object.keys(params).reduce((pre, k) => {
        pre[k] = parseValue(data, params[k]);
        return pre;
    }, {} as Parse.Data);
}

export function jsonParse(value?: string) {
    if(!value) return null;
    try {
        const result = JSON.parse(value);
        return result;
    } catch (error) {
        console.error(error);
        return null;
    }
}

function getObjValue(data: Parse.Data, key: string) {
    return Object.hasOwn(data, key) ? data[key] : undefined;
}

type ApiValue = "loginApi" | "typeApi" | "dataApi" | "searchTypeApi" | "searchApi" | "detailApi" | "sourceApi" | "episodeApi" | "playUrlApi" | "scheduleApi";

interface RequestOptions {
    bodyRule?: RuleItem["data"];
    omitempty?: boolean;
}

const tokenMap = new Map<string, Array<(value: string | PromiseLike<string>) => void>>();

export function createParse<T extends ApiValue>(data: ParseOption, key: T) {
    const loginApi = data?.loginApi || {};
    const isLogin = !!loginApi.url;
    const parser = useParserStore();

    const baseConfig = (data[key] || {}) as RuleItem;

    function login() {
        return new Promise<string>(async (resolve, reject) => {
            const { type, subType } = data;
            const key = `${type}-${subType}`;
            if(tokenMap.has(key)) {
                tokenMap.set(key, [...tokenMap.get(key) || [], resolve]);
            } else {
                tokenMap.set(key, [resolve]);
                const loginRequest = createRequest(loginApi);
                const res = await loginRequest(loginApi.data!);
                if(res.isOk) {
                    tokenMap.get(key)?.forEach(r => r(res.data));
                    tokenMap.delete(key);
                } else {
                    reject(res);
                }
            }
        });
    }

    function getImgUri(path?: string) {
        if(!path) return "";
        const prefix = data?.imgdomain || "";
        return path?.startsWith(prefix) ? path : (prefix + path);
    }

    function getResponse<T>(res: Record<string, any>): Result<T> {
        const { codeok, codePath, dataPath, messagePath } = (data?.response || {}) as Parse.ResRule;
        const code = parseValue(res, codePath) ?? getObjValue(res, "code");
        return {
            ...res,
            isOk: code == (codeok ?? 1),
            code,
            data: parseValue(res, dataPath) ?? getObjValue(res, "data"),
            // total:
            msg: (parseValue(res, messagePath) ?? getObjValue(res, "message")) || "",
            source: res,
        };
    }

    function getUri(url: string) {
        if(url.startsWith("http")) return url;
        const domain = data?.domain || "";
        return domain + url;
    }

    function request<T extends object, D = any>(
        url: string,
        method = "get",
        params?: T,
        options?: RequestOptions,
    ) {
        return new Promise<Result<D>>(async (resolve, reject) => {
            const { bodyRule, omitempty } = options || {};
            const token = parser.getToken(data);
            const headers: Record<string, any> = {
                ...data?.headers,
            };
            if(
                isLogin
                && data.authorization
                && baseConfig.isAuth
                && token
                && data.response?.codeExpired
            ) {
                headers[data.authorization] = token;
            }
            const isGet = ["", "GET", "HEAD", "OPTIONS"].some(v => method?.toUpperCase() === v);
            const response = await AppService.Request({
                url: parseUri(getUri(url), params, omitempty),
                method,
                params: null,
                data: isGet ? {} : { ...parseBodyData(params, bodyRule, omitempty) },
                headers,
            });
            const res = getResponse<D>(response);
            if(!res.isOk) {
                if(
                    isLogin
                    && data.authorization
                    && baseConfig.isAuth
                    && data.response?.codeExpired
                    && res.code === data.response?.codeExpired
                ) {
                    try {
                        const t = await login();
                        await parser.updateToken(data, t);
                        const res2 = request(url, method, params, options);
                        resolve(res2);
                    } catch (error) {
                        reject(error);
                    }
                } else {
                    useMessage().error(res.msg || "未知错误！");
                    reject(res);
                }
                return;
            }
            resolve(res);
        });
    }

    function createRequest(config = baseConfig) {
        const { url, method, options, data: body, dataPath, parseType, omitempty } = config;
        return async <T extends object>(params?: T): Promise<Result<any>> => {
            if(isArray(options)) return Res.setData(options);
            if(!url) return Res.setData([]);
            const res = await request(url, method, params, {
                bodyRule: body,
                omitempty,
            });
            return {
                ...res,
                data: dataPath ? parse(res.data, dataPath, parseType) : res.data,
            };
        };
    }

    function createPageRequest(config = baseConfig) {
        const { totalPath, requestCount } = config;
        const get = createRequest();
        return async <T extends object>(params?: T): Promise<Result<Parse.Data[]> & { total: number }> => {
            if(!requestCount || requestCount <= 1) {
                const res = await get(params);
                return {
                    ...res,
                    total: totalPath ? parseValue(res.source, totalPath) || 0 : 0,
                };
            }
            const oldpage = (params as Parse.Data)?.page as number;
            const current = (oldpage - 1) * requestCount + 1;
            const requestList = Array.from({ length: requestCount }, (_, i) => {
                const page = i + current;
                return get({ ...params, page });
            });
            const result: Parse.Data[] = [];
            const resList = await Promise.allSettled(requestList);
            let total = 0;
            resList.forEach(item => {
                if(item.status === "fulfilled") {
                    result.push(...item.value.data);
                    total = totalPath ? parseValue(item.value, totalPath) || 0 : 0;
                }
            });
            return {
                ...Res.setData(result),
                total,
            };
        };
    }

    return {
        data,
        getImgUri,
        getUri,
        request,
        createRequest,
        createPageRequest,
        config: baseConfig,
    };
}

export default {
    // jsonpath
    query: parseArray,
    value: parseValue,
    parent: parseParent,
    auto: parseAuto,
    parse,

    // custom parse
    uri: parseUri,
    params: parseParams,
    body: parseBodyData,
};