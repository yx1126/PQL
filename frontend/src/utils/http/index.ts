import Http from "./http";

Http.setConfig({
    timeout: 60 * 1000, // request timeout
});

Http.request.use(config => {
    return config;
});

const message = useMessage();

const jphttp = new Http({
    baseURL: "https://japi.zxfmj.com/api",
});

jphttp.interceptors.response.use(response => {
    const { data } = response;
    if(data.code !== 1) {
        message.error(data.msg || "Error");
        return Promise.reject(data);
    }
    return data;
}, error => {
    let { message } = error;
    if(message === "Network Error") {
        message = "后端接口连接异常";
    } else if(message.includes("timeout")) {
        message = "系统接口请求超时";
    } else if(message.includes("Request failed with status code")) {
        message = "系统接口" + message.substr(message.length - 3) + "异常";
    }
    message.error(message || "Error");
    return Promise.reject(error);
});

export {
    jphttp,
};

export default Http;