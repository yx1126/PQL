// 手机号正则验证
const MOBILE_REG = /^1(3[0-9]|4[5,7]|5[0,1,2,3,5,6,7,8,9]|6[2,5,6,7]|7[0,1,7,8]|8[0-9]|9[1,8,9])\d{8}$/;
// 身份证正则验证
const IDCARD_REG
    = /^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$|^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}([0-9]|X)$/;
// 超链接正则验证
const LINK_REG = /^https?:\/\/(([a-zA-Z0-9_-])+(\.)?)*(:\d+)?(\/((\.)?(\?)?=?&?[a-zA-Z0-9_-](\?)?)*)*$/;
// email正则验证
const EMAIL_REG = /\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/;

const PHONE_LIST: readonly string[] = ["Android", "iPhone", "Windows Phone", "SymbianOS", "iPad", "iPod", "BlackBerry", "MeeGo", "HarmonyOS"];

function callToString(val: unknown) {
    return Object.prototype.toString.call(val);
}

export function isNum(val: unknown): val is number {
    return typeof val === "number";
}

export function isObject<T extends object>(val: unknown): val is T {
    return callToString(val) === "[object Object]";
}

export function isStr(val: unknown): val is string {
    return typeof val === "string";
}

export function isArray<T = any>(val: unknown): val is T {
    return Array.isArray(val);
}

export function isFn<T extends Function>(val: unknown): val is T {
    return typeof val === "function";
}

export function isDate(val: unknown): val is Date {
    return callToString(val) === "[object Date]";
}

export function isBoolean(val: unknown): val is boolean {
    return typeof val === "boolean";
}

export function isSymbol(val: unknown): val is symbol {
    return typeof val === "symbol";
}

export function isPromise<T = any>(val: unknown): val is Promise<T> {
    return (isObject(val) || isFn(val)) && isFn((val as any).then) && isFn((val as any).catch);
}

export function isRegExp(val: unknown) {
    return callToString(val) === "[object RegExp]";
}

export function isNull(val: unknown) {
    return val === null;
}

export function isUndefined(val: unknown) {
    return val === void 0;
}

export function isEmpty(val: unknown) {
    return isNull(val) || isUndefined(val);
}

export function isZeroEmpty(val: unknown) {
    if(isNum(val)) {
        return val === 0;
    } if(isArray<any[]>(val) || isStr(val)) {
        return val.length === 0;
    } if(isObject(val)) {
        return Object.keys(val).length === 0;
    } if(isBoolean(val)) {
        return !val;
    }
    return isEmpty(val);
}

export function isPhoneNum(val: string) {
    return MOBILE_REG.test(val);
}

export function isIdCard(val: string) {
    return IDCARD_REG.test(val);
}

export function isLink(val: unknown) {
    if(!isStr(val)) return false;
    return LINK_REG.test(val);
}

export function isEmail(val: string) {
    return EMAIL_REG.test(val);
}

export function isMobile() {
    return PHONE_LIST.some(item => navigator.userAgent.includes(item));
}

export function toString(val: unknown) {
    return val?.toString ? val.toString() : String(val);
}