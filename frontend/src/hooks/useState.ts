import { AppService } from "@bind/service";

const state: Record<string, string> = reactive({});

// 适用于单个窗口的状态 多窗口数据不会同步
export function useState(key: string, defaultValue?: string) {
    const loading = ref(false);

    onBeforeMount(async () => {
        if(!Object.hasOwn(state, key)) {
            const v = await AppService.GetStore(key);
            state[key] = v;
        }
    });

    async function set(v: string) {
        try {
            loading.value = true;
            await AppService.SetStore(key, v);
        } finally {
            loading.value = false;
        }
        state[key] = v;
    }

    const data = computed({
        get: () => state[key] || defaultValue,
        set,
    });

    return {
        data,
        loading,
    };
}