<script setup lang="ts">
import Player from "@/components/Player";
import LinkList from "@/components/LinkList";
import jp, { createParse, type Parse } from "@/utils/parse";
import { isArray } from "@/utils/validata";
import { Clipboard } from "@wailsio/runtime";
import { createEmun } from "@/utils/emunSwitch";
import type { PagingType } from "@/components/WUI";

defineOptions({
    name: "VideoPlay",
});

const playItemRefs = useTemplateRef("playItemRefs");

const route = useRoute();
const set = useSetStore();
const state = useLoading();
const store = useParserStore();
const message = useMessage();

const current = ref<Nullable<string | number>>(null);
const sourceId = ref<Nullable<string | number>>(null);
const data = ref<Record<string, any>>({});
const sourceReqList = ref<any[]>([]);
const episodeReqList = ref<any[]>([]);
const playReqUrl = ref("");

const form = ref({
    page: 1,
    size: getStore()?.defaultSize || 100,
});

const total = ref(0);

const storeSource = computed(() => getStore());

const maxSize = computed<Empty<number>>(() => storeSource.value?.defaultMaxSize);

// 详情
const details = computed(() => {
    const { config, createRequest } = createParse(storeSource.value, "detailApi");
    return {
        ...config,
        request: createRequest(),
    };
});

// 源
const source = computed(() => {
    const { config, createRequest } = createParse(storeSource.value, "sourceApi");
    return {
        ...config,
        isHasSource: !!storeSource.value?.sourceApi,
        isRequest: !!config.url,
        request: createRequest(),
    };
});

// 集
const episode = computed(() => {
    const { config, createRequest, createPageRequest } = createParse(storeSource.value, "episodeApi");
    return {
        ...config,
        pagingType: (config.totalPath ? "paging" : "default") as PagingType,
        isRequest: !!config.url,
        request: createRequest(),
        requestPage: createPageRequest(),
    };
});

// 视频链接
const playUri = computed(() => {
    const { config, createRequest } = createParse(storeSource.value, "playUrlApi");
    return {
        ...config,
        isRequest: !!config.url,
        request: createRequest(),
    };
});

// 源数据
const sourceList = computed<any[]>(() => {
    const { isRequest, dataPath, parseType } = source.value;
    if(isRequest) return sourceReqList.value;
    const v = jp.parse(data.value, dataPath, parseType);
    if(!isArray(v)) return [];
    return v;
});

// 当前源
const sourceItem = computed(() => {
    const { primaryPath } = source.value;
    return sourceList.value?.find(v => jp.value(v, primaryPath) === sourceId.value);
});

// 集数据
const playList = computed<any[]>(() => {
    const { isRequest, dataPath, parseType } = episode.value;
    const data: any[] = [];
    if(isRequest) {
        data.push(...episodeReqList.value);
    } else {
        const { isHasSource } = source.value;
        const v = jp.parse(isHasSource ? sourceItem.value : data, dataPath, parseType) || [];
        if(isArray(v)) {
            data.push(...v);
        }
    }
    return set.videoDetailSort === "asc" ? data : data.reverse();
});

// 当前集
const playItem = computed(() => {
    const { primaryPath } = episode.value;
    return playList.value?.find(v => jp.value(v, primaryPath) === current.value);
});

// 小标题
const subTitle = computed(() => jp.value(playItem.value, episode.value.namePath));
// 播放路径
const playUrl = computed(() => {
    if(playUri.value.isRequest) {
        return playReqUrl.value;
    }
    return jp.value(playItem.value, episode.value.srcPath);
});

watch(() => [
    set.videoDetailTabActive,
    set.videoDetailGrid,
    set.videoDetailSort,
], async () => {
    await set.updateConfig();
});

watch(() => route.query, query => {
    if(Object.keys(query).length <= 0) return;
    getDetailData();
}, {
    immediate: true,
    deep: true,
});

function getStore() {
    switch(route.query.type) {
    case "video":
        return store.videoSource?.data || {};
    case "anime":
        return store.animeSource?.data || {};
    default:
        return {};
    }
}

function onSourceClick(item: any) {
    const v = jp.value(item, source.value.primaryPath);
    if(sourceId.value === v) return;
    sourceId.value = v;
    setCurrent(jp.value(item, episode.value.dataPath) || []);
}

function onPlayClick(item: any) {
    const v = jp.value(item, episode.value.primaryPath);
    if(current.value === v) return;
    current.value = v;
    getPlayUrl(item);
}

async function getDetailData() {
    try {
        state.setLoad(true);
        const res = await details.value.request(route.query);
        data.value = res.data;
        if(source.value.isRequest) {
            await getSourceData();
        }
        if(episode.value.isRequest) {
            await getEpisodeData();
        }
        if(playUri.value.isRequest) {
            await getPlayUrl(playList.value.at(0));
        }
        const sItem = sourceList.value?.at(0);
        sourceId.value = jp.value(sItem, source.value.primaryPath) || null;
        setCurrent(playList.value);
    } catch (error) {
        console.error(error);
    } finally {
        state.setLoad(false);
    }
}

async function getSourceData() {
    try {
        const res = await source.value.request(route.query);
        sourceReqList.value = res.data;
    } catch (error) {
        console.error(error);
    }
}

async function getEpisodeData() {
    try {
        const { paging, request, requestPage } = episode.value;
        if(paging) {
            const res = await requestPage({
                ...route.query,
                ...form.value,
            });
            episodeReqList.value = res.data;
            total.value = res.total || 0;
        } else {
            const res = await request(route.query);
            episodeReqList.value = res.data;
        }
    } catch (error) {
        console.error(error);
    }
}

async function getPlayUrl<T extends object>(params: T) {
    if(!params) return;
    try {
        const res = await playUri.value.request(jp.params(params, episode.value.next));
        playReqUrl.value = res.data;
    } catch (error) {
        console.error(error);
    }
}

function setCurrent(data: any[]) {
    const index = set.videoDetailSort === "asc" ? 0 : -1;
    current.value = jp.value(data.at(index), episode.value.primaryPath) || null;
}

function onFocus() {
    const dom = playItemRefs.value?.find(v => v?.$el.dataset.id == current.value);
    dom?.$el?.scrollIntoView({
        behavior: "smooth",
        block: "center",
        inline: "center",
    });
}

const gridList = createEmun<string>(["default", "full", "half", "four"]);
function onChangeGrid() {
    set.videoDetailGrid = gridList.next(set.videoDetailGrid);
}

const sortList = createEmun<string>(["asc", "desc"]);
function onSort() {
    set.videoDetailSort = sortList.next(set.videoDetailSort);
}

async function onCopy() {
    if(playUrl.value) {
        await Clipboard.SetText(playUrl.value);
        message.success({
            message: "复制成功！",
            duration: 1000,
        });
    }
}
</script>

<template>
    <w-player-layout :title="jp.value(data, details.titlePath)" :sub-title="subTitle">
        <div class="w-box">
            <div class="videoplay-box">
                <Player v-if="playUrl" :src="playUrl" />
            </div>
        </div>
        <w-card v-if="source.isHasSource" title="播放源：" is-list>
            <template
                v-for="item, i in sourceList"
                :key="jp.value(item, source.primaryPath) || i"
            >
                <w-button
                    size="medium"
                    border
                    :type="sourceId === jp.value(item, source.primaryPath) ? 'primary' : ''"
                    :title="jp.value(item, source.namePath)"
                    @click="onSourceClick(item)"
                >
                    {{ jp.value(item, source.namePath) }}
                </w-button>
            </template>
        </w-card>
        <template v-for="item, i in details.descriptionList" :key="i">
            <w-card :title="item.label">
                <link-list
                    :options="jp.parse(data, item.path, item as Parse.JsonParse)"
                    :template="item.template"
                    :disabled="!item.copy"
                    :separator="item.separator"
                />
            </w-card>
        </template>
        <template #sider>
            <div class="videoplay-sider">
                <el-tabs v-model="set.videoDetailTabActive" stretch>
                    <el-tab-pane label="详情" name="info">
                        <el-descriptions :column="1" border label-width="80px">
                            <template v-for="item, i in details.detailList" :key="i">
                                <el-descriptions-item :label="item.label">
                                    <link-list
                                        :options="jp.parse(data, item.path, item as Parse.JsonParse)"
                                        :template="item.template"
                                        :disabled="!item.copy"
                                        :separator="item.separator"
                                    />
                                </el-descriptions-item>
                            </template>
                        </el-descriptions>
                    </el-tab-pane>
                    <el-tab-pane name="list">
                        <template #label>
                            <span>列表</span>
                            <span v-if="episode.paging">（{{ total }}）</span>
                            <span v-else>（{{ playList?.length || 0 }}）</span>
                        </template>
                        <div class="flex flex-col gap-[10px]">
                            <div class="w-box flex items-center justify-between p-[15px] pos-sticky top-0">
                                <div class="flex items-center gap-3" :style="`${!episode.paging ? 'flex: 1;justify-content: flex-end;' : ''}`">
                                    <el-link icon="focus" style="--el-link-font-size: 18px;" title="定位" @click="onFocus" />
                                    <el-link icon="ele-Sort" title="排序" style="--el-link-font-size: 18px;" @click="onSort" />
                                    <el-link icon="grid" title="列布局" @click="onChangeGrid" />
                                    <el-link icon="ele-CopyDocument" title="复制视频地址" @click="onCopy" />
                                </div>
                                <div v-if="episode.paging">
                                    <w-paging
                                        v-model:page="form.page"
                                        v-model:size="form.size"
                                        class="!w-auto"
                                        button-size="small"
                                        :total="total"
                                        :type="episode.pagingType"
                                        layout="prev, next"
                                        :is-last="maxSize ? data.length < maxSize : false"
                                        @current-change="getEpisodeData"
                                        @size-change="getEpisodeData"
                                    />
                                </div>
                            </div>
                            <div class="grid grid-cols-3 gap-[10px]" :class="`is-${set.videoDetailGrid}`">
                                <template v-for="item, i in playList" :key="jp.value(item, episode.primaryPath) || i">
                                    <w-button
                                        ref="playItemRefs"
                                        border
                                        size="large"
                                        :type="current === jp.value(item, episode.primaryPath) ? 'primary' : ''"
                                        :title="jp.value(item, episode.namePath)"
                                        :data-id="jp.value(item, episode.primaryPath)"
                                        :data-sort="item.sort"
                                        @click="onPlayClick(item)"
                                    >
                                        {{ jp.value(item, episode.namePath) }}
                                    </w-button>
                                </template>
                            </div>
                        </div>
                    </el-tab-pane>
                </el-tabs>
            </div>
        </template>
    </w-player-layout>
</template>

<style lang="scss" scoped>
.videoplay {
    &-box {
        width: 100%;
        aspect-ratio: 16 / 9;
        border-radius: var(--w-border-radius);
        overflow: hidden;
        border: 1px solid var(--el-color-primary-light-9);
        box-sizing: var(--el-box-shadow-dark-white);
        background-image: url("@/assets/video/poster.png");
        background-size: 100% 100%;
        background-repeat: no-repeat;
    }
    &-sider {
        width: 100%;
        height: 100%;
        overflow-x: hidden;
        overflow-y: auto;
        position: relative;
        @include hidden-scroll;
        :deep(.el-descriptions) {
            --el-fill-color-blank: var(--w-box-bg);
            --el-descriptions-item-bordered-label-background: var(--w-box-bg);
            --el-descriptions-table-border: 1px solid var(--w-border-color);
            .el-descriptions__label {
                text-align: center;
            }
        }
        :deep(.el-tabs) {
            height: 100%;
            .el-tab-pane {
                height: 100%;
                overflow-x: hidden;
                overflow-y: auto;
                position: relative;
                @include hidden-scroll;
            }
        }
        .is-full {
            grid-template-columns: repeat(1, 1fr) !important;
        }
        .is-half {
            grid-template-columns: repeat(2, 1fr) !important;
        }
        .is-four {
            grid-template-columns: repeat(4, 1fr) !important;
        }
    }
}
</style>