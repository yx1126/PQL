import { isStr } from "./validata";

export function set(obj: Record<string, any>, path: string | string[], value: any): void {
    const paths = isStr(path) ? path.split(".") : path;
    let current = obj;
    for(let i = 0; i < paths.length; i++) {
        const key = paths[i];
        if(i === paths.length - 1) {
            current[key] = value;
        } else {
            if(!(key in current)) {
                current[key] = {};
            }
            current = current[key];
        }
    }
}