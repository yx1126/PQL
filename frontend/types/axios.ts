import "axios";

export interface PageInfo<T = any> {
    readonly page: number;
    readonly size: number;
    readonly list: T;
    readonly total: number;
}

export interface Result<T = any> {
    readonly code: number;
    readonly data: T;
    readonly msg: string;
    readonly isOk: boolean;
    readonly source: unknown;
}

export interface ResultPaging<T = any> extends Result<T> {
    readonly total: number;
}

export interface TableResult<T = any> {
    readonly code: number;
    readonly data: PageInfo<T>;
    readonly msg: string;
}

declare module "axios" {
    export interface AxiosInstance {
        request<T = any, R = Result<T>>(config: AxiosRequestConfig): Promise<R>;
        get<T = any, R = Result<T>>(url: string, config?: AxiosRequestConfig): Promise<R>;
        delete<T = any, R = Result<T>>(url: string, config?: AxiosRequestConfig): Promise<R>;
        head<T = any, R = Result<T>>(url: string, config?: AxiosRequestConfig): Promise<R>;
        options<T = any, R = Result<T>>(url: string, config?: AxiosRequestConfig): Promise<R>;
        post<T = any, R = Result<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>;
        put<T = any, R = Result<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>;
        patch<T = any, R = Result<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>;
    }
}