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

const LiveDialog = defineAsyncComponent(() => import("./components/LiveDialog.vue"));

const liveDialogRef = useTemplateRef("liveDialogRef");
const msgBox = useMessageBox();
const set = useSetStore();

const showTypeList = [
    { label: "全部", value: "all" },
    { label: "已开播", value: "live" },
    { label: "未开播", value: "nolive" },
];

const { data, query } = useService({
    request: LiveService.GetLiveList,
    default: [],
    immediate: true,
    isLayoutLoad: true,
});

const liveData = computed(() => {
    const { liveShowType: lst, liveSpecialShowType: lsst } = set;
    return data.value.reduce((pre, item) => {
        if(item.isSpecial == 1) {
            pre.specialSource.push(item);
            switch(lsst) {
            case "live":
                if(item.isLive) pre.special.push(item);
                break;
            case "nolive":
                if(!item.isLive) pre.special.push(item);
                break;
            default:
                pre.special.push(item);
                break;
            }
        } else {
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
        }
        return pre;
    }, {
        special: [],
        specialSource: [],
        data: [],
        dataSource: [],
    } as Record<"special" | "data" | `${"special" | "data"}Source`, LiveVo[]>);
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
        await LiveService.UpdateLive({
            id: data.id,
            sort: null,
            isSpecial: type === "special" ? 1 : 0,
        });
        query();
        return;
    }
    if(type === "care") {
        msgBox.confirm("确认要取消关注吗？").then(async () => {
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
        <div class="mine-header w-box">
            <el-button-group>
                <el-button type="primary" icon="ele-Plus" @click="liveDialogRef?.open()">添加</el-button>
                <!-- <el-button type="primary" icon="grid" /> -->
            </el-button-group>
        </div>
        <w-card v-if="liveData.specialSource.length > 0" title="特别关注">
            <template #extra>
                <el-segmented v-model="set.liveSpecialShowType" class="is-rounded" :options="showTypeList" size="small" />
            </template>
            <div v-if="liveData.special.length > 0" class="live-list">
                <template v-for="live in liveData.special" :key="live.roomId">
                    <live-card :data="live" @click="onLiveClick(live)" @command="onCommand">
                        <template #dropdown>
                            <el-dropdown-item command="no-special" :icon="renderIcon('heart')">取消特别</el-dropdown-item>
                            <el-dropdown-item command="care" :icon="renderIcon('heart')">取消关注</el-dropdown-item>
                            <el-dropdown-item command="copy" icon="ele-CopyDocument">复制房号</el-dropdown-item>
                            <el-dropdown-item command="web" :icon="renderIcon('pc')">打开网页</el-dropdown-item>
                        </template>
                    </live-card>
                </template>
            </div>
            <el-empty v-else class="h-[160px]" :image-size="80" />
        </w-card>
        <w-card title="关注">
            <template #extra>
                <el-segmented v-model="set.liveShowType" class="is-rounded" :options="showTypeList" size="small" />
            </template>
            <div v-if="liveData.dataSource.length > 0" class="live-list">
                <template v-for="live in liveData.data" :key="live.roomId">
                    <live-card :data="live" @click="onLiveClick(live)" @command="onCommand">
                        <template #dropdown>
                            <el-dropdown-item command="special" :icon="renderIcon('heart-fill')">特别关注</el-dropdown-item>
                            <el-dropdown-item command="care" :icon="renderIcon('heart')">取消关注</el-dropdown-item>
                            <el-dropdown-item command="copy" icon="ele-CopyDocument">复制房号</el-dropdown-item>
                            <el-dropdown-item command="web" :icon="renderIcon('pc')">打开网页</el-dropdown-item>
                        </template>
                    </live-card>
                </template>
            </div>
            <el-empty v-else class="h-[160px]" :image-size="80" />
        </w-card>
        <w-backtop
            target="#liveListTarget"
            :right="20"
            :bottom="20"
            @refresh="query"
        />
        <live-dialog ref="liveDialogRef" @success="query" />
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