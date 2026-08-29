import type { Result } from "#/axios";

function Res<T = any>(option?: Partial<Result<T>>) {
    return Object.assign({
        code: 1,
        data: null as T,
        msg: "",
        isOk: true,
        source: {},
    }, option);
}

Res.setCode = function<T>(value: Result<T>["code"]) {
    return Res({ code: value });
};
Res.setData = function<T>(value: Result<T>["data"]) {
    return Res({ data: value });
};
Res.setMsg = function<T>(value: Result<T>["msg"]) {
    return Res({ msg: value });
};

export {
    Res,
};