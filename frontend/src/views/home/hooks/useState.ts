import { windowOpen } from "@/utils/window";
import jp from "@/utils/parse";
import type { Parse, RuleItem } from "@/utils/parse";

export interface VideoState {
    typeList: Parse.Data[];
    dataList: Parse.Data[];
    form: Record<string, any>;
    total: number;
}

export type VideoType = "video" | "anime";

export function useState(type: VideoType) {
    const parse = useParserStore();

    const state: VideoState = reactive({
        typeList: [],
        dataList: [],
        form: {
            page: 1,
            size: getStore()?.defaultSize || 25,
        },
        total: 0,
    });

    const store = computed(() => getStore());

    const maxSize = computed<Empty<number>>(() => store.value?.defaultMaxSize);

    function getStore() {
        switch(type) {
        case "video":
            return parse.videoSource?.data || {};
        case "anime":
            return parse.animeSource?.data || {};
        default:
            return {};
        }
    }

    function open(item: any, namePath?: RuleItem["namePath"], next?: RuleItem["next"]) {
        windowOpen({
            title: jp.value(item, namePath) || "PQL",
            name: "Video",
            path: "/video-play",
            query: {
                ...jp.params(item, next),
                type,
            },
        });
    }
    return {
        ...toRefs(state),
        store,
        maxSize,
        getStore,
        open,
    };
}