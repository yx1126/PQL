import { optimize, type Config, type Output } from "svgo";
import path from "node:path";
import fg from "fast-glob";
import fs from "fs-extra";
import type { SVGLoadOptions, SVGItem, SymbolIdOptions } from "./types";

const assign = Object.assign;

export const SVG_ATTRS: Record<string, string> = {
    xmlns: "http://www.w3.org/2000/svg",
    "xmlns:xlink": "http://www.w3.org/1999/xlink",
    style: "position: absolute; width: 0; height: 0",
};

function createSymbolId({ name, fileId, symbolId }: SymbolIdOptions): string {
    [
        { rege: /\[id\]/g, value: fileId || "" },
        { rege: /\[name\]/g, value: name },
        { rege: /-+/g, value: "-" },
    ].forEach(({ rege, value }) => {
        symbolId = symbolId.replace(rege, value);
    });
    if(symbolId.startsWith("-")) {
        symbolId = symbolId.slice(1);
    }
    if(symbolId.endsWith("-")) {
        symbolId = symbolId.slice(0, -1);
    }
    return symbolId;
}

function parseSVG(file: string, id: string, config?: Config): Output {
    try {
        return optimize(file, assign<Config, Config>({
            plugins: [
                {
                    name: "preset-default",
                    params: {
                        overrides: {
                            cleanupIds: {
                                minify: false,
                            },
                        },
                    },
                },
                "removeDimensions",
                "removeXMLNS",
                "removeStyleElement",
                "removeScripts",
                "removeTitle",
                {
                    name: "removeAttrs",
                    params: {
                        attrs: "class",
                    },
                },
                {
                    name: "addAttributesToSVGElement",
                    params: {
                        attribute: {
                            id,
                        },
                    },
                },
            ],
        }, config!));
    } catch (error) {
        console.error(error);
        return { data: file };
    }
}

export function getSvgFiles(options: SVGLoadOptions) {
    return options.paths.reduce<SVGItem[]>((pre, _path) => {
        const cwdPath = typeof _path === "string" ? _path : _path.path;
        const result = fg.sync("**/*.svg", {
            cwd: cwdPath,
            stats: true,
            absolute: true,
            dot: true,
        }).map<SVGItem>(v => {
            const { name } = path.parse(v.name);
            const id = createSymbolId({
                name,
                fileId: typeof _path === "string" ? "" : _path.id!,
                symbolId: options.symbolId!,
            });
            return {
                name,
                path: v.path,
                svg: "",
                symbolId: id,
                symbol: "",
            };
        });
        pre.push(...result);
        return pre;
    }, []);
}

export function readSvgFiles(options: SVGLoadOptions) {
    const svgList = getSvgFiles(options);
    return svgList.map(v => {
        const svgStr = parseSVG(fs.readFileSync(v.path).toString(), v.symbolId, options.svgoOptions).data;
        return {
            ...v,
            svg: svgStr,
            symbol: svgStr.replace(/svg/g, "symbol"),
        };
    });
}

export function parseHtml(value: string, html: string) {
    const BODY = "<body>";
    const index = html.indexOf(BODY) + BODY.length;
    return html.substring(0, index) + value + html.substring(index);
}

export function parseScript(value: string, options: SVGLoadOptions) {
    return `
    if (typeof window !== 'undefined') {
        function loadSvg() {
            var svg = document.getElementById("${options.SVG_DOM_ID}");
            if (!svg) {
                svg = document.createElementNS("${SVG_ATTRS["xmlns"]}", "svg");
                svg.style.position = "absolute";
                svg.style.width = "0";
                svg.style.height = "0";
                svg.setAttribute("id", "${options.SVG_DOM_ID}");
                svg.setAttribute("xmlns", "${SVG_ATTRS["xmlns"]}");
                svg.setAttribute("xmlns:xlink", "${SVG_ATTRS["xmlns:xlink"]}");
                svg.innerHTML = ${JSON.stringify(value)};
                document.body.insertBefore(svg, document.body.firstChild);
            } else {
                svg.innerHTML = ${JSON.stringify(value)};
            }
        }
        if (document.readyState === "loading") {
            document.addEventListener("DOMContentLoaded", loadSvg);
        } else {
            loadSvg();
        }
    }
    `;
}

export function getOptions(options?: SVGLoadOptions): SVGLoadOptions {
    const opt: any = assign({}, options);
    const def: any = {
        symbolId: "icon-[id]-[name]",
        SVG_DOM_ID: "__svg__icons__dom__",
        type: "html",
    };
    for(const key in def) {
        if(!opt[key]) {
            opt[key] = def[key];
        }
    }
    return opt;
}