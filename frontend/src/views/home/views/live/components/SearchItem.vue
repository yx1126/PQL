<script setup lang="ts">
import { windowOpen, openURL } from "@/utils/window";
import { toString } from "@/utils/validata";
import { LiveService } from "@bind/service";
import { Clipboard } from "@wailsio/runtime";
import { renderIcon } from "@/utils/renderIcon";
import LiveCard from "./LiveCard.vue";
import type { RoomInfo } from "@bind/parse/live";

defineOptions({
    name: "SearchItem",
});

const {
    type,
    keyword = "",
} = defineProps<{
    type: string;
    keyword?: string;
}>();

const TYPE_MAP: Record<string, string> = {
    1: "斗鱼",
    2: "虎牙",
};

const message = useMessage();
const msgBox = useMessageBox();

const form = ref({
    page: 1,
});

const { data, query } = useService({
    request: () => LiveService.Search({
        keyword: keyword,
        page: toString(form.value.page),
        type,
    }),
    default: {
        data: [],
        total: 0,
    },
    isLayoutLoad: true,
    error(error) {
        const prefix = TYPE_MAP[type] || "";
        return (prefix ? `${prefix}：` : "") + (error as any)?.message || "未知错误！";
    },
});

function onCommand(type: string, data: RoomInfo) {
    if(type === "care") {
        msgBox.confirm("确认要关注吗？").then(async () => {
            try {
                await LiveService.CreateLive({
                    type: "1",
                    roomId: toString(data.roomId) as `${number}`,
                });
            } catch (error) {
                message.error((error as any)?.message || "未知错误！");
            }
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

function onLiveClick(live: RoomInfo) {
    if(!live.isLive) return;
    windowOpen({
        title: live.roomName || "PQL",
        name: "Live",
        path: "/live-play",
        query: {
            roomId: live.roomId,
            type: live.type,
        },
    });
}

function search() {
    form.value.page = 1;
    query();
}

defineExpose({
    search,
    refresh: query,
});
</script>

<template>
    <div class="size-full">
        <div v-if="data.data && data.data.length > 0" class="live-list">
            <template v-for="live in data.data" :key="live.roomId">
                <live-card :data="live" @click="onLiveClick(live)" @command="onCommand">
                    <template #dropdown>
                        <el-dropdown-item command="care" :icon="renderIcon('heart')">关注</el-dropdown-item>
                        <el-dropdown-item command="copy" icon="ele-CopyDocument">复制房号</el-dropdown-item>
                        <el-dropdown-item command="web" :icon="renderIcon('pc')">打开网页</el-dropdown-item>
                    </template>
                </live-card>
            </template>
        </div>
        <el-empty v-else />
        <div class="mt-[15px]">
            <w-paging
                v-model:page="form.page"
                :size="20"
                :total="data.total"
                :type="[''].includes(type) ? 'default' : 'paging'"
                :is-last="data.data ? data.data.length > 0 && data.data.length < 20 : false"
                @current-change="query"
                @size-change="query"
            />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.live-list {
    display: grid;
    gap: 20px;
    grid-template-columns: repeat(4, minmax(150px, 1fr));
    @media screen and (width > 1600px) {
        grid-template-columns: repeat(5, minmax(160px, 1fr));
    }
}
</style>