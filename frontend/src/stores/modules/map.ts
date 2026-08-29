import { defineStore } from "pinia";
import type { MapState } from "#/stores";

export const useMapStore = defineStore("map", () => {
    const state: MapState = reactive({
        pointTypeList: [],
    });
    return {
        ...toRefs(state),
    };
});