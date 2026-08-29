import mitt, { type Emitter } from "mitt";
import { tryOnScopeDispose } from "@vueuse/core";

type EventType = {
    "*": any;
    "video-type": [];
};

const emitter = mitt<EventType>();

type EventFunc<T extends keyof EventType> = (...args: EventType[T]) => void;

interface MittBack<T extends keyof EventType> {
    on(fn: EventFunc<T>): void;
    emit(...args: EventType[T]): void;
    off(fn: EventFunc<T>): void;
    once(fn: EventFunc<T>): void;
}

type MittBacks<K extends (keyof EventType)[]> = K extends [infer Key, ...infer Other]
    ? Key extends keyof EventType
        ? Other extends (keyof EventType)[]
            ? MittBacks<Other> extends never
                ? [MittBack<Key>]
                : [MittBack<Key>, ...MittBacks<Other>]
            : never
        : never
    : never;

function create<T extends keyof EventType>(key: T): MittBack<T> {
    function on(fn: EventFunc<T>) {
        emitter.on(key, fn);
        // 自动销毁
        tryOnScopeDispose(() => {
            emitter.off(key, fn);
        });
    }
    function emit(...args: EventType[T]) {
        emitter.emit.apply(null, [key].concat(args) as any);
    }
    function off(fn: EventFunc<T>) {
        emitter.off(key, fn);
    }
    function once(fn: EventFunc<T>) {
        function oncefn(...args: any) {
            fn(...args);
            emitter.off(key, oncefn);
        }
        emitter.on(key, oncefn);
    }
    return {
        on,
        off,
        emit,
        once,
    };
}

export function useMitt(): Emitter<EventType>;
export function useMitt<T extends keyof EventType>(key: T): MittBack<T>;
export function useMitt<T extends (keyof EventType)[]>(...keys: T): MittBacks<T>;
export function useMitt<T extends keyof EventType>(...keys: T[]) {
    if(keys.length === 0) {
        return emitter;
    }
    if(keys.length === 1) {
        return create(keys[0]);
    }
    return keys.map(key => create(key));
}

export default useMitt;