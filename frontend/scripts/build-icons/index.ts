import { normalizePath, type PluginOption, type ResolvedConfig } from "vite";
import { SVG_ATTRS, getOptions, getSvgFiles, readSvgFiles, parseScript, parseHtml } from "./utils.ts";
import type { SVGLoadOptions, SVGItem } from "./types";

function resolveSvgList({ name, symbolId }: SVGItem): Pick<SVGItem, "name" | "symbolId"> {
    return { name, symbolId };
}

export default function vitePluginSvgIcon(options?: SVGLoadOptions): PluginOption {
    const virtualModuleId = "virtual:svg-icons-list";
    const resolvedVirtualModuleId = "\0" + virtualModuleId;

    const registerVirtualModuleId = "virtual:svg-icons-load";
    const resolvedRegisterVirtualModuleId = "\0" + registerVirtualModuleId;

    let svgNames: Array<Pick<SVGItem, "name" | "symbolId">> = [];

    const opt = getOptions(options);

    let viteConfig: ResolvedConfig;

    if(!Array.isArray(opt.paths)) {
        opt.paths = [];
    }

    opt.paths = opt.paths.map(v => {
        if(typeof v === "string") {
            return normalizePath(v);
        }
        return { ...v, path: normalizePath(v.path) };
    });

    return {
        name: "vite:svg-icons-load",
        configResolved(resolvedConfig) {
            viteConfig = resolvedConfig;
        },
        resolveId(id) {
            if(id === virtualModuleId) {
                return resolvedVirtualModuleId;
            }
            if(id === registerVirtualModuleId) {
                return resolvedRegisterVirtualModuleId;
            }
        },
        load(id) {
            if(id === resolvedVirtualModuleId) {
                if(viteConfig.command === "build" && opt.type === "html") {
                    const svgList = getSvgFiles(opt);
                    return `export default ${JSON.stringify(svgList.map(resolveSvgList))}`;
                }
                return `export default ${JSON.stringify(svgNames)}`;
            }
            if(id === resolvedRegisterVirtualModuleId) {
                if(opt.type === "script") {
                    const svgList = readSvgFiles(opt);
                    svgNames = svgList.map(resolveSvgList);
                    return parseScript(svgList.map(v => v.symbol).join(""), opt);
                }
                return "";
            }
        },
        transformIndexHtml(html) {
            if(opt.type === "html") {
                const svgList = readSvgFiles(opt);
                svgNames = svgList.map(resolveSvgList);
                const svgAttrs = Object.keys(SVG_ATTRS).map(key => `${key}="${SVG_ATTRS[key]}"`).join(" ");
                const str = `\n\t<svg ${svgAttrs} id="${opt.SVG_DOM_ID}">\t${svgList.map(v => v.symbol).join("\n\t\t")}\n\t</svg>`;
                return parseHtml(str, html);
            }
        },
    };
}