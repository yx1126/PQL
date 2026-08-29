import { TinyColor } from "@ctrl/tinycolor";

function darken(color: TinyColor, amount = 20) {
    return color.mix("#141414", amount).toString();
}

function lighten(color: TinyColor, amount = 20) {
    return color.mix("#fff", amount).toString();
}

export function createTheme(value: string, dark = false) {
    const color = new TinyColor(value);
    const result: Record<string, string> = {};
    result["--el-color-primary"] = value;
    for(let i = 0; i < 10; i++) {
        result[`--el-color-primary-light-${i}`] = dark ? darken(color, i * 10) : lighten(color, i * 10);
    }
    result["--el-color-primary-dark-2"] = color.mix("#000", 20).toString();
    return result;
}