<script setup lang="ts">
import { LiveService } from "@bind/service";
import LiveCard from "./components/LiveCard.vue";
import { openURL, windowOpen } from "@/utils/window";
import { renderIcon } from "@/utils/renderIcon";
import { Clipboard } from "@wailsio/runtime";
import { toString } from "@/utils/validata";
import type { LiveVo } from "@bind/vo";

defineOptions({
    name: "LiveMine",
});

const {
    type = "",
    isSpecial = "",
} = defineProps<{
    type?: string;
    isSpecial?: string;
    title?: string;
}>();

const msgBox = useMessageBox();
const route = useRoute();
const set = useSetStore();
const keepalive = useMitt("keepalive");
const livelistener = useMitt("live:data:refresh");

const showTypeList = [
    { label: "全部", value: "all" },
    { label: "已开播", value: "live" },
    { label: "未开播", value: "nolive" },
];

const { data, query } = useService({
    request: () => LiveService.GetLiveList({ isSpecial: isSpecial, type: type }),
    default: [],
    immediate: true,
    isLayoutLoad: true,
});

const liveData = computed(() => {
    const { liveShowType: lst } = set;
    return data.value.reduce((pre, item) => {
        pre.dataSource.push(item);
        switch(lst) {
        case "live":
            if(item.isLive) pre.data.push(item);
            break;
        case "nolive":
            if(!item.isLive) pre.data.push(item);
            break;
        default:
            pre.data.push(item);
            break;
        }
        return pre;
    }, {
        data: [],
        dataSource: [],
    } as Record<"data" | "dataSource", LiveVo[]>);
});

const routeName = route.name;
onBeforeMount(() => {
    livelistener.on(name => {
        if(routeName === name) {
            query();
        }
    });
});

function onLiveClick(live: LiveVo) {
    if(!live.isLive) return;
    windowOpen({
        title: live.roomName || "PQL",
        name: "Live",
        path: "/live-play",
        query: {
            roomId: live.roomId,
            rid: live.rid,
            type: live.type,
        },
    });
}

async function onCommand(type: string, data: LiveVo) {
    if(type === "special" || type === "no-special") {
        const isSpecial = type === "special" ? 1 : 0;
        await LiveService.UpdateLive({
            id: data.id,
            sort: null,
            isSpecial,
        });
        keepalive.emit("LiveMine");
        data.isSpecial = isSpecial;
        return;
    }
    if(type === "care") {
        msgBox.confirm("确认要删除吗？").then(async () => {
            if(data?.id) await LiveService.DeleteLive([data.id]);
            query();
        });
        return;
    }
    if(type === "copy") {
        Clipboard.SetText(toString(data.roomId));
        return;
    }
    if(type === "web") {
        openURL(data.roomId, data.type);
        return;
    }
}
</script>

<template>
    <div id="liveListTarget" class="mine">
        <w-card :title="title">
            <template #extra>
                <el-segmented v-model="set.liveShowType" class="is-rounded" :options="showTypeList" size="small" />
            </template>
            <div v-if="liveData.data.length > 0" class="live-list">
                <template v-for="live in liveData.data" :key="live.roomId">
                    <live-card :data="live" @click="onLiveClick(live)" @command="onCommand">
                        <template #dropdown>
                            <el-dropdown-item v-if="live.isSpecial == 0" command="special" :icon="renderIcon('heart')">关注</el-dropdown-item>
                            <el-dropdown-item v-else command="no-special" :icon="renderIcon('heart-fill')">取消关注</el-dropdown-item>
                            <el-dropdown-item command="copy" icon="ele-CopyDocument">复制房号</el-dropdown-item>
                            <el-dropdown-item command="web" :icon="renderIcon('pc')">打开网页</el-dropdown-item>
                            <el-dropdown-item command="care" icon="ele-Delete">删除</el-dropdown-item>
                        </template>
                    </live-card>
                </template>
            </div>
            <el-empty v-else class="h-full" :image-size="80" />
        </w-card>
        <w-backtop
            target="#liveListTarget"
            :right="20"
            :bottom="20"
            @refresh="query"
        />
    </div>
</template>

<style lang="scss" scoped>
.mine {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: var(--w-layout-space);
    overflow-x: hidden;
    overflow-y: auto;
    @include hidden-scroll;
    &-header {
        display: flex;
        justify-content: flex-end;
        padding: var(--w-layout-space);
    }
    .live-list {
        display: grid;
        gap: 20px;
        grid-template-columns: repeat(4, minmax(150px, 1fr));
        @media screen and (width > 1600px) {
            grid-template-columns: repeat(5, minmax(160px, 1fr));
        }
    }
}
</style>