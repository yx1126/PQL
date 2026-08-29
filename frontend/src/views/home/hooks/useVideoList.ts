import jp, { createParse } from "@/utils/parse";
import { throttle } from "lodash-es";
import { useState, type VideoType } from "./useState";
import type { PagingType } from "@/components/WUI";

export function useVideoList(type: VideoType) {
    const state = useLoading();

    const {
        typeList,
        dataList,
        form,
        total,
        store,
        maxSize,
        open,
    } = useState(type);

    const types = computed(() => {
        const { config, createRequest } = createParse(store.value, "typeApi");
        return {
            ...config,
            isShow: !!config.url,
            request: createRequest(),
        };
    });

    const datas = computed(() => {
        const { config, getImgUri, createPageRequest } = createParse(store.value, "dataApi");
        return {
            ...config,
            getImgUri,
            pagingType: (config.totalPath ? "paging" : "default") as PagingType,
            request: createPageRequest(),
        };
    });

    onBeforeMount(async () => {
        await getTypeList();
        typeList.value.forEach(item => {
            const v = jp.value(item, types.value.childrenPath);
            const primary = jp.value(v.at(0) || {}, types.value.childPrimaryPath);
            form.value[jp.value(item, types.value.primaryPath)] = primary ?? "";
        });
        getDataList();
    });

    const onRefresh = throttle(() => {
        getTypeList();
        getDataList();
    }, 500);

    async function getTypeList() {
        try {
            state.setLoad(true);
            const res = await types.value.request();
            typeList.value = res.data;
        } finally {
            state.setLoad(false);
        }
    }

    async function getDataList() {
        try {
            state.setLoad(true);
            const res = await datas.value.request(form.value);
            dataList.value = res.data;
            total.value = res.total || 0;
        } finally {
            state.setLoad(false);
        }
    }

    function onTypeClick(sub: any, key?: string) {
        if(!key) return;
        form.value.page = 1;
        form.value[key] = jp.value(sub, types.value.childPrimaryPath);
        getDataList();
    }

    function onVideoClick(item: any) {
        const { namePath, next } = datas.value;
        open(item, namePath, next);
    }

    return {
        typeList,
        dataList,
        form,
        total,
        types,
        datas,
        maxSize,
        onRefresh,
        getDataList,
        onTypeClick,
        onVideoClick,
    };
}