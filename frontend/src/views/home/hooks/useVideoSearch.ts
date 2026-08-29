import { throttle } from "lodash-es";
import jp, { createParse, type Parse } from "@/utils/parse";
import { useState } from "./useState";
import type { PagingType } from "@/components/WUI";

export function useVideoSearch(type: "video" | "anime") {
    const state = useStatesStore();

    const inputValue = ref("");

    const {
        typeList,
        dataList,
        form,
        total,
        store,
        maxSize,
        open,
    } = useState(type);

    const sTypes = computed(() => {
        const { config, createRequest } = createParse(store.value, "searchTypeApi");
        return {
            ...config,
            isShow: !!config.url,
            request: createRequest(),
        };
    });

    const sDatas = computed(() => {
        const { config, getImgUri, createPageRequest } = createParse(store.value, "searchApi");
        return {
            ...config,
            getImgUri,
            pagingType: (config.totalPath ? "paging" : "default") as PagingType,
            request: createPageRequest(),
        };
    });

    onBeforeMount(async () => {
        await onGetTypeList();
        form.value = {
            ...form.value,
            id: jp.value(typeList.value.at(0), sTypes.value.primaryPath),
            ...jp.params(typeList.value.at(0), sTypes.value.next),
        };
    });

    const onRefresh = throttle(() => {
        onGetTypeList();
        onGetData();
    }, 500);

    async function onGetTypeList() {
        try {
            state.setLoad(true);
            const res = await sTypes.value.request();
            typeList.value = res.data || [];
        } finally {
            state.setLoad(false);
        }
    }

    async function onGetData() {
        try {
            state.setLoad(true);
            const res = await sDatas.value.request({
                ...form.value,
                keyword: inputValue.value,
            });
            dataList.value = res.data || [];
            total.value = res.total || 0;
        } finally {
            state.setLoad(false);
        }
    }

    function onTypeClick(item: Parse.Data) {
        form.value.page = 1;
        form.value = {
            ...form.value,
            id: jp.value(item, sTypes.value.primaryPath),
            ...jp.params(item, sTypes.value.next),
        };
        onGetData();
    }

    function onSearch() {
        form.value.page = 1;
        onGetData();
    }

    function onVideoClick(item: any) {
        const { namePath, next } = sDatas.value;
        open(item, namePath, next);
    }

    return {
        inputValue,
        typeList,
        dataList,
        form,
        total,
        sTypes,
        sDatas,
        maxSize,
        onRefresh,
        onSearch,
        onGetData,
        onTypeClick,
        onVideoClick,
    };
}