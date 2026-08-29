import { AppService } from "@bind/service";
import { isBoolean } from "./validata";

export async function getIsDark(theme: number) {
    switch(theme) {
    case 0:
        return true;
    case 1:
        return false;
    case 2:
        return await AppService.GetDarkMode();
    default:
        return false;
    }
}

export function setTheme(theme: number): void;
export function setTheme(dark: boolean): void;
export function setTheme(value: unknown) {
    if(isBoolean(value)) {
        value = value ? 0 : 1;
    }
    const html = document.documentElement;
    if(value === 0) {
        html.classList.add("dark");
    } else {
        html.classList.remove("dark");
    }
}