import { acceptHMRUpdate, defineStore } from "pinia";
import type { MenuState } from "#/stores";

const defaultMenu: MenuState = {
    keepMap: {},
    menuList: [],
};

export const useMenuStore = defineStore("menu", () => {
    const state: MenuState = reactive(Object.assign({}, defaultMenu));

    return {
        ...toRefs(state),
    };
});

if(import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useSetStore, import.meta.hot));
}