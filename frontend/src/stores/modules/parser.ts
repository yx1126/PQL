import { acceptHMRUpdate, defineStore } from "pinia";
import { ParserService } from "@bind/service";
import { jsonParse, type ParseOption } from "@/utils/parse";
import type { ParserState } from "#/stores";

const defaultParser: ParserState = {
    parserList: [],
};

export const useParserStore = defineStore("parser", () => {
    const state: ParserState = reactive(Object.assign({}, defaultParser));

    const videoList = computed(() => state.parserList.filter(v => v.type === "video"));
    const animeList = computed(() => state.parserList.filter(v => v.type === "anime"));

    const set = useSetStore();

    const videoSource = computed(() => {
        const item = videoList.value.find(v => v.subType === set.videoSourceType);
        if(!item) return null;
        return {
            ...item,
            data: jsonParse(item.source) as ParseOption,
        };
    });

    const animeSource = computed(() => {
        const item = animeList.value.find(v => v.subType === set.animeSourceType);
        if(!item) return null;
        return {
            ...item,
            data: jsonParse(item.source) as ParseOption,
        };
    });

    async function load() {
        const res = await ParserService.GetParserList();
        state.parserList = res || [];
    }

    async function updateToken(parser: ParseOption, token: string) {
        const { type, subType } = parser;
        const index = state.parserList.findIndex(v => v.type === type && v.subType === subType);
        if(index === -1) return;
        const v = state.parserList[index];
        await ParserService.UpdateToken(v.id, token);
        state.parserList[index].token = token;
    }

    function getToken(parser: ParseOption) {
        const { type, subType } = parser;
        return state.parserList.find(v => v.type === type && v.subType === subType)?.token;
    }

    return {
        ...toRefs(state),
        videoList,
        videoSource,
        animeList,
        animeSource,
        load,
        updateToken,
        getToken,
    };
});

if(import.meta.hot) {
    import.meta.hot.accept(acceptHMRUpdate(useParserStore, import.meta.hot));
}