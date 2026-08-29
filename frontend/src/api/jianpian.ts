import { jphttp } from "@/utils/http";
import type { ResultPaging } from "#/axios";

export function getUri(path?: string) {
    if(!path) return "";
    return `https://img.shanhuixiu.com${path.startsWith("/") ? path : ("/" + path)}`;
}

export function getTypeList<T extends object>(params?: T) {
    return jphttp.get("/crumb/filterOptions", params);
}

export function getDataList<T extends object>(params?: T) {
    return jphttp.get("/crumb/list", params);
}

export function getSearchTypeList<T extends object>(params?: T) {
    return jphttp.get("/v2/settings/topCategory", params);
}

export function getSearchDataList<T extends object>(params?: T) {
    return jphttp.get<any[], ResultPaging<any[]>>("/v2/search/videoV2", params);
}

export function getVideoDetail(id: string) {
    return jphttp.get("/video/detailv2", { id });
}