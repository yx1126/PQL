<script setup lang="ts">
import KeepRouteView from "@/components/KeepRouteView";
import VideoBtns from "../../components/VideoBtns.vue";
import type { ParseOption } from "@/utils/parse";
import type { MenuSiderItem } from "@/components/WUI";

defineOptions({
    name: "Anime",
});

const route = useRoute();
const router = useRouter();
const store = useParserStore();
const set = useSetStore();

const menuList: MenuSiderItem[] = [
    { id: 0, label: "全部", icon: "video", path: "/home/anime" },
    { id: 1, label: "搜索", icon: "ele-Search", path: "/home/anime/search" },
    { id: 2, label: "追番周表", icon: "ele-Calendar", path: "/home/anime/schedule" },
];

// 解决源跟新 数据刷新问题 移除缓存
const keepKey = ref(0);

onBeforeMount(() => {
    onReset();
});

onActivated(() => {
    onReset();
});

async function onReset() {
    if(set.animeSourceType != "" && store.videoList.length <= 0) {
        router.replace("/home/anime");
        await nextTick();
        set.animeSourceType = "";
        onChange();
    }
}

async function onSuccess(data: ParseOption[]) {
    await store.load();
    if(!data.length) return;
    await nextTick();
    // 导入后 没有选中源/导入的源有当前源版本跟新  刷新数据
    if(!set.animeSourceType || data.find(v => v.type === "anime" && v.subType === set.animeSourceType)) {
        onChange();
        set.animeSourceType = store.animeList.at(0)?.subType || "";
    }
}

function onChange() {
    keepKey.value++;
}
</script>

<template>
    <w-sider-layout is-wbox>
        <template #sider>
            <div class="video-source">
                <el-select v-model="set.animeSourceType" placeholder="请选择源" @change="onChange">
                    <template v-for="source in store.animeList" :key="source.id">
                        <el-option :label="source.subType" :value="source.subType" />
                    </template>
                </el-select>
            </div>
            <w-sider-menu class="flex-1" :default-active="route.path" :data="menuList" router />
            <video-btns type="anime" @import="onSuccess" />
        </template>
        <keep-route-view :key="keepKey" />
    </w-sider-layout>
</template>

<style lang="scss" scoped>
.video {
    width: 100%;
    height: 100%;
    display: flex;
    gap: 15px;
    padding: 10px;
    &-sider {
        width: var(--w-sider-min-width);
        height: 100%;
        display: flex;
        flex-direction: column;
        gap: 15px;
        & > div {
            background-color: var(--w-box-bg);
            border-radius: var(--w-border-radius);
            border: 1px solid var(--w-border-color);
            padding: 10px;
        }
    }
    &-main {
        width: calc(100% - 15px - var(--w-sider-min-width));
        height: 100%;
    }
}
</style>