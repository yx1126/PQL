import { isNum, isStr } from "./validata";

export function parseUnit(val?: Empty<string | number>, defaultValue?: string | number) {
    if(!val && val !== 0) val = defaultValue;
    if(isNum(val) || (isStr(val) && /^\d+$/.test(val))) {
        return val + "px";
    }
    return val;
}

export function parseNum(val: string | number, defaultValue = 0) {
    val = parseInt(val as string);
    return isNaN(val) ? defaultValue : val;
}