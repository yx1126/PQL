import { LiveService, WindowService } from "@bind/service";
import qs from "qs";
import type { WindowOptions as BaseWindowOptions } from "@bind/utils/types";

interface WindowOptions extends BaseWindowOptions {
    query?: Record<string, any>;
}

export async function windowOpen(options: WindowOptions) {
    const { name, title, path, query } = options;
    const prefix = window.location.href.split("#").at(0) || "";
    let url = prefix + "#" + path;
    if(query) {
        url += ("?" + qs.stringify(query));
    }
    // await OpenNewWindow(name, prefix + "#" + path);
    await WindowService.OpenNewWindow({
        name,
        title,
        path: url,
    });
    return () => windowClose(name);
}

export function windowClose(name: string) {
    return WindowService.Close(name);
}

export function createWindow(name: string) {
    function open(title: string, path: string) {
        return windowOpen({ name, title, path });
    }
    return {
        open,
        close: () => windowClose(name),
    };
}

export function openURL(roomId: string, type: string) {
    LiveService.OpenWeb(roomId, type);
}

export default createWindow;