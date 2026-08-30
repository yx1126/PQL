<script setup lang="ts" generic="T extends LiveVo | RoomInfo">
import poster from "@/assets/video/poster.png";
import type { RoomInfo } from "@bind/parse/live";
import type { LiveVo } from "@bind/vo";
import type { VNode } from "vue";

defineOptions({
    name: "LiveCard",
});

const {
    data,
    type = "default",
} = defineProps<{
    data: T;
    type?: string;
}>();

const emit = defineEmits<{
    command: [command: string, data: T];
}>();

const slots = defineSlots<{
    dropdown?: () => VNode[];
}>();

function onCommand(command: string) {
    emit("command", command, data);
}
</script>

<template>
    <div
        class="live-card"
        :class="{
            'is-live': data?.isLive,
            [`is-${type}`]: !!type
        }"
        :title="data?.roomName"
    >
        <div class="live-card__img">
            <w-image class="live-img" :src="data?.roomPic || poster" />
            <div class="live-owner" :title="data?.ownerName">
                <w-image class="live-avatar" :src="data?.ownerAvatar" />
                <div class="live-name">{{ data?.ownerName }}</div>
            </div>
            <div class="live-type" :class="`is-type${data?.type}`"></div>
            <div v-if="data?.isReplay" class="live-loop">轮播</div>
        </div>
        <div class="live-card__footer">
            <div class="live-card__name">{{ data?.roomName }}</div>
            <div v-if="slots.dropdown" class="live-card__btns" @click.stop>
                <el-dropdown
                    trigger="click"
                    :persistent="false"
                    placement="bottom-end"
                    @command="onCommand"
                >
                    <el-link>
                        <Icon icon="ele-MoreFilled" rotate="90" />
                    </el-link>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <slot name="dropdown" />
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.live {
    &-avatar {
        width: 32px;
        min-width: 32px;
        height: 32px;
        min-height: 32px;
        border-radius: 50%;
        background-color: #d8d8d8;
    }
    &-owner {
        display: flex;
        align-items: center;
        position: relative;
        z-index: 1;
    }
    &-name {
        margin-left: 6px;
        color: #fff;
        font-size: 14px;
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
    }
    &-type {
        width: 30px;
        height: 30px;
        background-size: cover;
        background-repeat: no-repeat;
        background-position: center center;
        background-color: rgba($color: #000, $alpha: 0.5);
        border-bottom-right-radius: var(--w-border-radius);
        border-top-left-radius: var(--w-border-radius);
        position: absolute;
        top: 0;
        left: 0;
        z-index: 1000;
        @include when(type1) {
            background-image: url("@/assets/live/douyu.png");
        }
        @include when(type2) {
            background-image: url("@/assets/live/huya.png");
        }
        @include when(type3) {
            background-image: url("@/assets/live/douyin.png");
        }
    }
    &-loop {
        height: 22px;
        line-height: 1;
        padding: 0 4px;
        display: flex;
        align-items: center;
        gap: 4px;
        font-size: 12px;
        color: #fff;
        background: rgba(0, 0, 0, 0.6);
        border-radius: var(--w-border-radius);
        position: absolute;
        top: 6px;
        right: 6px;
        &::before {
            content: "";
            display: inline-block;
            width: 16px;
            height: 16px;
            background-image: url("@/assets/live/live-loop.png");
            background-size: 100% 100%;
            background-repeat: no-repeat;
        }
    }
    &-img {
        width: 100%;
        height: 100%;
        background-image: var(--live-url);
        background-size: cover;
        background-repeat: no-repeat;
        position: absolute;
        left: 0;
        top: 0;
        transition: scale 0.2s;
    }
    &-card {
        cursor: pointer;
        position: relative;
        max-height: max-content;
        box-shadow: var(--w-box-bg-shadow);
        border-radius: var(--w-border-radius);
        overflow: hidden;
        &:hover .live-img {
            scale: 1.1;
        }
        &__img {
            aspect-ratio: 300 / 168;
            overflow: hidden;
            display: flex;
            align-items: flex-end;
            padding: 5px;
            position: relative;
            z-index: 0;
            transition: box-shadow 0.2s;
            &:hover {
                box-shadow: var(--el-box-shadow);
                @include when-dark {
                    box-shadow: var(--el-box-shadow-white);
                }
            }
        }
        &__footer {
            display: flex;
            align-items: center;
            justify-content: space-between;
            gap: 10px;
            padding: 8px 2px 8px 8px;
        }
        &__name {
            text-overflow: ellipsis;
            overflow: hidden;
            white-space: nowrap;
            flex-wrap: 1;
        }
        &__btns {
            & > * {
                width: 22px;
                height: 22px;
                display: inline-flex;
                align-items: center;
                justify-content: center;
                transition: background-color 0.3s;
                border-radius: 2px;
                &:hover {
                    background-color: var(--el-fill-color-light);
                }
            }
        }
        @include when-not(live) {
            .live-card__img::before {
                content: "未开播";
                position: absolute;
                left: 0;
                top: 0;
                width: 100%;
                height: 100%;
                z-index: 999;
                display: flex;
                align-items: center;
                justify-content: center;
                background: rgba(0, 0, 0, 0.8);
                color: #fff;
            }
        }
    }
}
</style>