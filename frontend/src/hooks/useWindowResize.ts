import { tryOnScopeDispose } from "@vueuse/core";
import { debounce } from "lodash-es";

interface ResizeFn {
    (ev: UIEvent): void;
    cancel?: () => void;
};

const fnList: ResizeFn[] = [];

let windowFn: Nullable<ResizeFn> = null;

interface Options {
    options?: boolean | AddEventListenerOptions;
    lazy?: number;
}

export function useWindowResize(fn: ResizeFn, options?: Options) {
    const { options: _options, lazy } = options || {};

    const width = ref(window.innerWidth);
    const height = ref(window.innerHeight);

    if(lazy) fn = debounce(fn, lazy);

    const index = fnList.findIndex(v => v === fn);
    if(index === -1) fnList.push(fn);

    if(!windowFn) {
        windowFn = (e: UIEvent) => {
            width.value = window.innerWidth;
            height.value = window.innerHeight;
            fnList.forEach(f => f.call(window, e));
        };
        window.addEventListener("resize", windowFn, _options);
    }

    function stop() {
        const index = fnList.findIndex(v => v === fn);
        if(index !== -1) {
            if(fnList[index]?.cancel) {
                fnList[index]?.cancel();
            }
            fnList.splice(index, 1);
        }
        if(fnList.length <= 0 && windowFn) {
            window.removeEventListener("resize", windowFn, _options);
            windowFn = null;
        }
    }

    tryOnScopeDispose(stop);

    return {
        width,
        height,
        stop,
    };
}