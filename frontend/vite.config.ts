import { fileURLToPath, URL } from "node:url";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";
import wails from "@wailsio/runtime/plugins/vite";
import VueI18nPlugin from "@intlify/unplugin-vue-i18n/vite";
import Uoncss from "unocss/vite";
import AutoImport from "unplugin-auto-import/vite";
import VitePluginSvgIcon from "./scripts/build-icons/index.ts";
import VueDevTools from "vite-plugin-vue-devtools";

function resolve(path: string) {
    return fileURLToPath(new URL(path, import.meta.url));
}

// https://vite.dev/config/
export default defineConfig(() => {
    return {
        assetsInclude: ["/files/**"],
        resolve: {
            alias: {
                "@": resolve("./src"),
                "@bind": resolve("./bindings/pql/pkg"),
                "#": resolve("./types"),
            },
        },
        plugins: [
            vue(),
            vueJsx(),
            AutoImport({
                dts: "./types/auto-imports.d.ts",
                dirs: ["./src/hooks", "./src/enums", "./src/stores/modules"],
                dirsScanOptions: {
                    types: false,
                },
                imports: [
                    "vue",
                    "vue-router",
                    {
                        vue: ["renderSlot", "renderList", "mergeProps", "createVNode", "render"],
                    },
                ],
            }),
            VueI18nPlugin({}),
            Uoncss(),
            VitePluginSvgIcon({
                paths: [resolve("./src/assets/icon/")],
                symbolId: "local-[name]",
                type: "script",
            }),
            wails("./bindings"),
            VueDevTools(),
        ],
        server: {
            host: "127.0.0.1",
            hmr: true,
            port: Number(process.env.WAILS_VITE_PORT) || 9245,
            strictPort: true,
        },
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: `@use "@/styles/global/global" as *; @use "@/styles/global/mixins" as *;`,
                },
            },
        },
        build: {
            chunkSizeWarningLimit: 1500,
            rolldownOptions: {
                output: {
                    entryFileNames: `js/[name].[hash].js`,
                    chunkFileNames: `js/[name].[hash].js`,
                    assetFileNames: chunk => {
                        const isCss = chunk.names.find(v => v.endsWith(".css"));
                        return `${isCss ? "[ext]" : "assets"}/[name].[hash].[ext]`;
                    },
                    codeSplitting: {
                        groups: [
                            {
                                name: "element-plus-vendor",
                                test: /node_modules[\\/]element-plus/,
                                priority: 1,
                            },
                            {
                                name: "flvjs-vendor",
                                test: /node_modules[\\/]flv\.js/,
                                priority: 1,
                            },
                            {
                                name: "hlsjs-vendor",
                                test: /node_modules[\\/]hls\.js/,
                                priority: 1,
                            },
                        ],
                    },
                },
            },
        },
    };
});
