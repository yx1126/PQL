import { EffectScope } from "vue";

export function createScope<T extends (...args: any[]) => any>(composable: T) {
    let subscribers = 0;
    let state: Nullable<ReturnType<T>>,
        scope: Nullable<EffectScope>;

    const dispose = () => {
        if(scope && --subscribers <= 0) {
            scope.stop();
            state = scope = null;
        }
    };

    return (...args: Parameters<T>) => {
        subscribers++;
        if(!state) {
            scope = effectScope(true);
            state = scope.run(() => composable(...args));
        }
        onScopeDispose(dispose);
        return state as ReturnType<T>;
    };
}