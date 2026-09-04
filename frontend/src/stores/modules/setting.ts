import { acceptHMRUpdate, defineStore } from "pinia";
import { createTheme } from "@/utils/color";
import { SetService } from "@bind/service";
import { getIsDark, setTheme } from "@/utils/theme";
import { Window, Events } from "@wailsio/runtime";
import type { SetState } from "#/stores";
import type { UpdateSettingVo } from "@bind/vo";

const defaultSetting: SetState = {
    windownName: "",
    colorTheme: "#ff9f43",
    lang: "zh-cn", // 语言
    closeBehavior: 0,
    theme: 0,
    videoDetailTabActive: "info",
    videoDetailGrid: "default",
    videoDetailSort: "asc",
    videoSourceType: "",
    animeSourceType: "",
    liveShowType: "default",
    headerShowTheme: "1",
    animeWeeklyType: "cn",
};

export const useSetStore = defineStore("setting", () => {
    const state: SetState = reactive(Object.assign({}, defaultSetting));

    const primaryColor = computed({
        get: () => state.colorTheme,
        set: value => {
            state.colorTheme = value || "#ff9f43";
        },
    });

    watch(() => [state.colorTheme, state.theme], async () => {
        updateTheme(await getIsDark(state.theme));
    }, {
        immediate: true,
    });

    watch(() => [
        state.closeBehavior,
        state.lang,
        state.theme,
        state.colorTheme,
        state.videoSourceType,
        state.liveShowType,
        state.headerShowTheme,
        state.animeWeeklyType,
    ], async () => {
        // 主窗口更改
        if(state.windownName !== import.meta.env.VITE_APP_MAIN_NAME) return;
        await updateConfig();
        Events.Emit(WailsEvents.AppSetChange, "");
    });

    watch(() => state.theme, value => {
        setTheme(value);
    }, {
        immediate: true,
    });

    async function updateTheme(dark?: boolean) {
        const root = document.documentElement;
        const colorMap = createTheme(state.colorTheme, dark);
        for(const key in colorMap) {
            root.style.setProperty(key, colorMap[key]);
        }
    }

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
        state.headerShowTheme = res.headerShowTheme;
        state.animeWeeklyType = res.animeWeeklyType;
        state.colorTheme = res.colorTheme;
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
        updateTheme,
    };
});

if(import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useSetStore, import.meta.hot));
}