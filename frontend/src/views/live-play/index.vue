<script setup lang="ts">
import { LiveService } from "@bind/service";
import Player, { type Selector } from "@/components/Player";
import { Clipboard } from "@wailsio/runtime";
import { toString } from "@/utils/validata";
import { openURL } from "@/utils/window";
import type { LiveVo } from "@bind/vo";

defineOptions({
    name: "LivePlay",
});

const message = useMessage();
const msgbox = useMessageBox();
const route = useRoute();

const data = ref<Nullable<LiveVo>>(null);

const { data: remote, query: onGetLiveInfo } = useService({
    request: () => {
        const { roomId, type } = route.query as Record<string, string>;
        return LiveService.GetLiveRemoteInfo(roomId, type);
    },
    default: null,
    isLayoutLoad: true,
});

const { data: player, query: onGetPlayUrl } = useService({
    request: (data: Record<string, string>) => {
        const { roomId, rid, type } = route.query as Record<string, string>;
        return LiveService.GetPlayInfo(rid || roomId, type, data);
    },
    default: null,
    immediate: false,
    isLayoutLoad: true,
});

const cdns = computed<Selector[]>(() => {
    const { cdnsWithName, rtmpCdn } = player.value || {};
    if(!cdnsWithName) return [];
    return cdnsWithName.map(v => {
        return {
            html: v.name,
            value: v.cdn,
            default: v.cdn === rtmpCdn,
        };
    });
});

const quality = computed<Selector[]>(() => {
    const { multirates, rate } = player.value || {};
    if(!multirates) return [];
    return multirates.map(v => {
        return {
            html: v.name,
            value: v.rate,
            default: v.rate === rate,
        };
    });
});

watch(() => route.query, query => {
    if(!query.roomId || !query.type) return;
    onGetPlayUrl({ cdn: "", rate: "-1" });
    onGetLiveInfo();
    onGetInfo();
}, {
    immediate: true,
    deep: true,
});

async function onCopy(value?: unknown) {
    if(value) {
        await Clipboard.SetText(toString(value));
        message.success({
            message: "复制成功！",
            duration: 1000,
        });
    }
}

function onOpenPc() {
    msgbox.confirm("确认要打开网页端直播间吗？").then(() => {
        const { roomId, type } = route.query as Record<string, string>;
        openURL(roomId, type);
    });
}

function onQualityChange(rate: Selector["value"]) {
    if(rate === player.value?.rate) return;
    onGetPlayUrl({ cdn: player.value?.rtmpCdn || "", rate: toString(rate) || "-1" });
}

function onCdnChange(cdn: Selector["value"]) {
    if(cdn === player.value?.rtmpCdn) return;
    onGetPlayUrl({ cdn: toString(cdn) || "", rate: toString(player.value?.rate) || "-1" });
}

async function onGetInfo() {
    const { roomId, type } = route.query;
    data.value = await LiveService.GetLiveInfo(roomId as string, type as string);
}

function onCare(message: string, value = true) {
    msgbox.confirm(`确认要${message}吗？`).then(async () => {
        const { roomId, type } = route.query as Record<string, string>;
        if(value) {
            await LiveService.CreateLive({
                roomId: roomId,
                type: type,
                isSpecial: 0,
            });
        } else {
            const id = data.value?.id;
            if(id) {
                await LiveService.DeleteLive([id]);
            }
        }
        onGetInfo();
    });
}
</script>

<template>
    <w-player-layout :title="remote?.roomName" sider-width="200">
        <div class="w-box">
            <div class="liveplay-box">
                <Player
                    v-if="player?.url"
                    :src="player?.url"
                    :type="player.videoType || 'flv'"
                    is-live
                    :cdns
                    :quality
                    @quality-change="onQualityChange"
                    @cdn-change="onCdnChange"
                />
            </div>
        </div>
        <w-card v-if="remote?.description" title="直播公告：">{{ remote?.description }}</w-card>
        <w-card title="直播流：" hover>
            <span>{{ player?.url }}</span>
            <template #extra>
                <el-link icon="ele-CopyDocument" title="复制直播地址" @click="onCopy(player?.url)" />
            </template>
        </w-card>
        <template #sider>
            <div class="liveplay-info">
                <w-image class="size-[100px] rounded-[50%]" :src="remote?.ownerAvatar" />
                <div class="text-[18px] text-center">{{ remote?.ownerName }}</div>
                <div class="flex items-center leading-[1] flex-wrap">
                    <span class="whitespace-nowrap">房间号：</span>
                </div>
                <div class="liveplay-roomId">
                    <el-link @click="onCopy(remote?.roomId)">{{ remote?.roomId }}</el-link>
                </div>
                <div class="flex gap-3 mt-2">
                    <el-link v-if="!!data" type="primary" title="取消关注" @click="onCare('取消关注', false)">
                        <Icon icon="heart-fill" size="20" />
                    </el-link>
                    <el-link v-else title="关注" @click="onCare('关注')">
                        <Icon icon="heart" size="20" />
                    </el-link>
                    <el-link title="打开网页直播" @click="onOpenPc">
                        <Icon icon="pc" size="18" />
                    </el-link>
                </div>
            </div>
        </template>
    </w-player-layout>
</template>

<style lang="scss" scoped>
.liveplay {
    &-box {
        width: 100%;
        aspect-ratio: 16 / 9;
        border-radius: var(--w-border-radius);
        border: 1px solid var(--el-color-primary-light-9);
        box-sizing: var(--el-box-shadow-dark-white);
        background-image: url("@/assets/video/poster.png");
        background-size: 100% 100%;
        background-repeat: no-repeat;
        overflow: hidden;
    }
    &-info {
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: var(--w-layout-space-large);
    }
    &-roomId {
        /* 关键换行 */
        word-break: break-all;
        overflow-wrap: break-word;
        white-space: normal;
        text-align: center;
    }
}
</style>