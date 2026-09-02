import { onScopeDispose } from "vue";

/**
 *
 * @param {Function} fn
 * @param {Number} delay
 * @returns
 */
export function useInterval<T extends Function>(fn: T) {
    let timer: ReturnType<typeof setInterval> | null = null;

    onScopeDispose(() => {
        stop();
    });

    function stop() {
        if(timer) {
            clearInterval(timer);
            timer = null;
        }
    }

    function start(delay: number) {
        stop();
        timer = setInterval(fn, delay);
    }

    return {
        stop,
        start,
    };
}

/**
 *
 * @param {Function} fn
 * @param {Number} delay
 * @returns
 */
export function useTimeout<T extends Function>(fn: T, delay: number) {
    let timer: ReturnType<typeof setTimeout> | null = null;

    timer = setTimeout(fn, delay);

    onScopeDispose(() => {
        stop();
    });

    function reset(_delay: number = delay) {
        stop();
        timer = setTimeout(fn, _delay);
    }

    function stop() {
        if(timer) {
            clearTimeout(timer);
            timer = null;
        }
    }

    return {
        stop,
        reset,
    };
}