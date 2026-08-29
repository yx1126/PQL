import { isStr } from "@/utils/validata";
import type { Component } from "vue";
import { renderIcon as BaseRenderIcon } from "@/utils/renderIcon";

export function renderIcon(icon?: string | Component) {
    if(!icon) return;
    if(isStr(icon)) return BaseRenderIcon(icon);
    return icon;
}