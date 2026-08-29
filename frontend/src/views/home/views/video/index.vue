<script setup lang="ts">
import KeepRouteView from "@/components/KeepRouteView";
import VideoBtns from "../../components/VideoBtns.vue";
import type { ParseOption } from "@/utils/parse";
import type { MenuSiderItem } from "@/components/WUI";

defineOptions({
    name: "Video",
});

const route = useRoute();
const router = useRouter();
const store = useParserStore();
const set = useSetStore();

const menuList: MenuSiderItem[] = [
    { id: 0, label: "全部", icon: "video", path: "/home/video" },
    { id: 1, label: "搜索", icon: "ele-Search", path: "/home/video/search" },
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
    if(set.videoSourceType != "" && store.videoList.length <= 0) {
        router.replace("/home/video");
        await nextTick();
        set.videoSourceType = "";
        onChange();
    }
}

async function onSuccess(data: ParseOption[]) {
    await store.load();
    if(!data.length) return;
    await nextTick();
    // 导入后 没有选中源/导入的源有当前源版本跟新  刷新数据
    if(!set.videoSourceType || data.find(v => v.type === "video" && v.subType === set.videoSourceType)) {
        onChange();
        set.videoSourceType = store.videoList.at(0)?.subType || "";
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
                <el-select v-model="set.videoSourceType" placeholder="请选择源" @change="onChange">
                    <template v-for="source in store.videoList" :key="source.id">
                        <el-option :label="source.subType" :value="source.subType" />
                    </template>
                </el-select>
            </div>
            <w-sider-menu class="flex-1" :default-active="route.path" :data="menuList" router />
            <video-btns type="video" @import="onSuccess" />
        </template>
        <keep-route-view :key="keepKey" />
    </w-sider-layout>
</template>