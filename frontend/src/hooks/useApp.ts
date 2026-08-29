import type { ComponentInternalInstance } from "vue";

export function useApp() {
    return getCurrentInstance() as ComponentInternalInstance;
}