import { AppService } from "@bind/service";

const state: Record<string, string> = reactive({});

// 适用于单个窗口的状态 多窗口数据不会同步
export function useState(key: string) {
    onBeforeMount(async () => {
        if(!Object.hasOwn(state, key)) {
            const v = await AppService.GetStore(key);
            state[key] = v;
        }
    });

    function set(v: string) {
        state[key] = v;
        AppService.SetStore(key, v);
    }

    const value = computed(() => state[key] || "");

    return [value, set];
}