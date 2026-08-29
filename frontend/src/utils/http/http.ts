import axios, {
    mergeConfig,
    type AxiosInstance,
    type AxiosRequestConfig,
    type AxiosResponse,
    type InternalAxiosRequestConfig,
    type AxiosInterceptorManager,
} from "axios";
import { isFn } from "../validata";
import type { Result } from "#/axios";

const configList: AxiosRequestConfig[] = [];

type Request = Pick<AxiosInterceptorManager<InternalAxiosRequestConfig>, "use">;
type Response = Pick<AxiosInterceptorManager<AxiosResponse>, "use">;

export class Http {
    private readonly service: AxiosInstance;

    static setConfig(config: AxiosRequestConfig) {
        configList.push(config);
    }

    private static requestParams: Parameters<Request["use"]> = [];
    private static responseParams: Parameters<Response["use"]> = [];

    static get request(): Request {
        return {
            use(onFulfilled, onRejected, options) {
                Http.requestParams = [onFulfilled, onRejected, options];
                return -1;
            },
        };
    }

    static get response(): Response {
        return {
            use(onFulfilled, onRejected) {
                Http.responseParams = [onFulfilled, onRejected];
                return -1;
            },
        };
    }

    constructor(config?: AxiosRequestConfig | ((config?: AxiosRequestConfig) => Noable<AxiosRequestConfig>)) {
        const _config = configList.reduce((pre, item) => {
            return mergeConfig(pre, item);
        }, {});
        const conf = isFn(config) ? config(_config) : config;
        const service = axios.create(mergeConfig(_config, conf || {}));
        const { requestParams, responseParams } = Http;
        if(requestParams.length) {
            service.interceptors.request.use(...requestParams);
        }
        if(responseParams.length) {
            service.interceptors.response.use(...responseParams);
        }
        this.service = service;
    }

    get interceptors() {
        return this.service.interceptors;
    }

    request<T = any>(params: AxiosRequestConfig) {
        return this.service.request<T>(params);
    }

    abortRequest<T = any>(params: AxiosRequestConfig) {
        const controller = new AbortController();
        const request = this.service.request<T>({
            ...params,
            signal: controller.signal,
        });
        return {
            abort: controller.abort,
            request,
        };
    }

    getUri(config?: AxiosRequestConfig) {
        return this.service.getUri(config);
    }

    head<T = any>(url: string, params?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.head<T>(url, { params, headers });
    }

    options<T = any>(url: string, params?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.options<T>(url, { params, headers });
    }

    patch<T = any>(url: string, data?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.patch<T>(url, data, { headers });
    }

    get<T = any, R = Result<T>>(url: string, params?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.get<T, R>(url, { params, headers });
    }

    post<T = any>(url: string, data?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.post<T>(url, data, { headers });
    }

    postForm<T = any>(url: string, data?: any, headers?: AxiosRequestConfig["headers"]) {
        const form = new FormData();
        if(data) {
            Object.keys(data).forEach(key => {
                form.append(key, data[key]);
            });
        }
        return this.service.post<T>(url, form, {
            headers: {
                "content-type": ContentType.FORM_DATA,
                ...headers,
            },
        });
    }

    down<T = Blob>(url: string, params?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.get<T, T>(url, {
            params,
            headers: {
                ...headers,
                isFile: true,
            },
        });
    }

    download<T = Blob, R = AxiosResponse<T>>(url: string, data?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.post<T, R>(url, data, {
            responseType: "blob",
            ...headers,
        });
    }

    put<T = any>(url: string, data?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.put<T>(url, data, { headers });
    }

    delete<T = any>(url: string, data?: any, headers?: AxiosRequestConfig["headers"]) {
        return this.service.delete<T>(url, { data, headers });
    }
}

export default Http;