import { isArray, isBoolean, isFn, isObject, isStr } from "@/utils/validata";
import type { CancellablePromise } from "@wailsio/runtime";

export type ServiceRequest<P extends any[] = any[], R = any> = (...args: P) => CancellablePromise<R>;

export interface ServiceOptions<
    T = any,
    P extends any[] = any[],
> {
    request: ServiceRequest<P, T>;
    default: T;
    isLayoutLoad?: boolean;
    immediate?: boolean;
    concat?: boolean;
    error?: (error: unknown) => string | boolean;
}

export interface ServiceBack<T, P extends any[] = any[]> {
    loading: Ref<boolean>;
    data: Ref<T & {}>;
    query: (...args: P) => void;
}

export function useService<T = any, P extends any[] = any[]>(config: ServiceOptions<T, P>): ServiceBack<T, P>;
export function useService<T = any, P extends any[] = any[]>(request: ServiceRequest<P, T>, defaultValue: ServiceOptions<T, P>["default"]): ServiceBack<T, P>;
export function useService<
    T = any,
    P extends any[] = any[],
>(request: ServiceOptions<T, P> | ServiceRequest<P, T>, defaultValue?: ServiceOptions<T, P>["default"]): ServiceBack<T, P> {
    const {
        request: req,
        default: dv,
        immediate,
        isLayoutLoad,
        concat = true,
        error: errorFn,
    } = (isObject(request) ? request : { request, default: defaultValue, immediate: true }) as ServiceOptions<T, P>;

    const state = useStatesStore();
    const message = useMessage();

    const initialValue = isFn<() => T>(dv) ? dv() : dv;

    const data = ref(initialValue) as Ref<T & {}>;

    const load = ref(false);

    const loading = computed(() => isLayoutLoad ? state.loading : load.value);

    // 立即执行
    onBeforeMount(() => {
        if(immediate) {
            query(...[] as unknown as P);
        }
    });

    async function query(...args: P) {
        try {
            setLoad();
            const res = await req(...args);
            if(concat && isArray(initialValue)) {
                data.value = (res || []) as (T & {});
            } else if(concat && isObject(initialValue)) {
                data.value = (res || {}) as (T & {});
            } else {
                data.value = res as (T & {});
            }
        } catch (error) {
            console.error(error);
            const run = () => message.error((error as any)?.message || "未知错误！");
            if(errorFn) {
                const v = await errorFn(error);
                if(isBoolean(v)) {
                    if(v) run();
                } else if(isStr(v) && v) {
                    message.error(v);
                } else {
                    run();
                }
            } else {
                run();
            }
        } finally {
            setLoad(false);
        }
    }

    function setLoad(value = true) {
        if(isLayoutLoad) {
            state.setLoad(value);
        } else {
            load.value = value;
        }
    }

    return {
        data,
        loading,
        query,
    };
}