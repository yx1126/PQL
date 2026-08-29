import { defineConfig, extractorSplit } from "unocss";
import presetUno from "@unocss/preset-wind3";

export default defineConfig({
    presets: [
        presetUno(),
        // presetAttributify(),
        extractorSplit,
    ],
    transformers: [],
    rules: [
        [
            /^(c|text)-(primary|success|info|warning|danger)$/,
            ([,,type]) => ({ color: `var(--el-color-${type})` }),
        ],
        [
            /^rounded-(none|small|base|round)$/,
            ([,type]) => ({ "border-radius": `var(--el-border-radius-${type})` }),
        ],
        [
            /^shadow-(base|light|lighter|dark)$/,
            ([_, type]) => ({ "box-shadow": `var(--el-box-shadow${type != "base" ? "-" + type : ""})` }),
        ],
    ],
});