import { acceptHMRUpdate, defineStore } from "pinia";
import { createTheme } from "@/utils/color";
import { SetService } from "@bind/service";
import { getIsDark, setTheme } from "@/utils/theme";
import { Window } from "@wailsio/runtime";
import type { SetState } from "#/stores";
import type { UpdateSettingVo } from "@bind/vo";

const defaultSetting: SetState = {
    windownName: "",
    themeColor: "#ff9f43",
    lang: "zh-cn", // 语言
    closeBehavior: 0,
    theme: 0,
    videoDetailTabActive: "info",
    videoDetailGrid: "default",
    videoDetailSort: "asc",
    videoSourceType: "",
    animeSourceType: "",
    liveShowType: "default",
    liveSpecialShowType: "all",
    animeWeeklyType: "cn",
};

export const useSetStore = defineStore("setting", () => {
    const state: SetState = reactive(Object.assign({}, defaultSetting));

    const primaryColor = computed({
        get: () => state.themeColor,
        set: value => {
            state.themeColor = value || "#ff9f43";
        },
    });

    watch(() => [state.themeColor, state.theme], async ([color, theme]) => {
        const root = document.documentElement;
        const colorMap = createTheme(color as string, await getIsDark(theme as number));
        for(const key in colorMap) {
            root.style.setProperty(key, colorMap[key]);
        }
    }, {
        immediate: true,
    });

    watch(() => [
        state.closeBehavior,
        state.lang,
        state.theme,
        state.videoSourceType,
        state.liveShowType,
        state.liveSpecialShowType,
        state.animeWeeklyType,
    ], async () => {
        if(state.windownName !== import.meta.env.VITE_APP_MAIN_NAME) return;
        await updateConfig();
    });

    watch(() => state.theme, value => {
        setTheme(value);
    }, {
        immediate: true,
    });

    async function updateConfig() {
        await SetService.UpdateConfig(state as UpdateSettingVo);
    }

    async function load() {
        const name = await Window.Name();
        state.windownName = name;
        const res = await SetService.GetConfig();
        state.id = res.id;
        state.closeBehavior = res.closeBehavior;
        state.theme = res.theme;
        state.lang = res.lang;
        state.videoDetailTabActive = res.videoDetailTabActive;
        state.videoDetailGrid = res.videoDetailGrid;
        state.videoDetailSort = res.videoDetailSort;
        state.liveShowType = res.liveShowType;
        state.liveSpecialShowType = res.liveSpecialShowType;
        state.animeWeeklyType = res.animeWeeklyType;
        const store = useParserStore();
        if(!state.videoSourceType && store.videoList.length > 0) {
            state.videoSourceType = store.videoList.at(0)?.subType || "";
        }
        if(!state.animeSourceType && store.animeList.length > 0) {
            state.animeSourceType = store.animeList.at(0)?.subType || "";
        }
    }

    return {
        ...toRefs(state),
        primaryColor,
        load,
        updateConfig,
    };
});

if(import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useSetStore, import.meta.hot));
}