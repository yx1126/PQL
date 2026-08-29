import { defineStore } from "pinia";
import type { StatesState } from "#/stores";

export const useStatesStore = defineStore("states", () => {
    const loadCount = ref(0);

    const state: StatesState = reactive({});

    const loading = computed(() => loadCount.value > 0);

    function setLoad(value = true) {
        if(value) {
            loadCount.value++;
        } else {
            loadCount.value = Math.max(loadCount.value - 1, 0);
        }
    }

    return {
        ...toRefs(state),
        loading,
        setLoad,
    };
});