<script setup lang="ts">
import WImage from "./WImage.vue";
import { toString } from "@/utils/validata";

defineOptions({
    name: "WVideo",
});

defineProps<{
    title?: string | number;
    src?: string;
    top?: string | number;
    left?: string | number;
    right?: string | number;
}>();
</script>

<template>
    <div class="w-video" :title="title ? toString(title) : undefined">
        <div class="w-video-img">
            <w-image class="video" :src="src" />
            <div v-if="top" class="w-video-top">{{ top }}</div>
            <div v-if="left || right" class="w-video-mask">
                <span v-if="left" :class="{'is-full': !right}">{{ left }}</span>
                <span v-if="right" :class="{'is-full': !left}">{{ right }}</span>
            </div>
        </div>
        <div class="w-video-title">{{ title }}</div>
    </div>
</template>

<style lang="scss" scoped>
.w-video {
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 5px;
    cursor: pointer;
    line-height: 1;
    @media screen and (width > 1600px) {
        .w-video-top,
        .w-video-mask,
        .w-video-title {
            font-size: 18px;
        }
        .w-video-mask {
            padding: 8px;
        }
    }
    &-img {
        width: 100%;
        border-radius: 5px;
        aspect-ratio: 2 / 3;
        overflow: hidden;
        position: relative;
        .video {
            width: 100%;
            height: 100%;
            &:hover {
                position: relative;
                transform: scale(1.2);
                transition: transform 0.3s;
            }
        }
    }
    &-top {
        position: absolute;
        top: 5px;
        right: 5px;
        padding: 2px 10px;
        border-radius: var(--w-border-radius);
        color: var(--el-color-primary);
        background-color: color-mix(in oklab, #000 80%, transparent);
        font-size: 14px;
    }
    &-mask {
        position: absolute;
        left: 0;
        right: 0;
        bottom: 0;
        padding: 5px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        color: var(--el-color-primary);
        background-color: color-mix(in oklab, #000 80%, transparent);
        font-size: 14px;
        & > span {
            display: inline-block;
            @include when(full) {
                width: 100% !important;
            }
            &:first-child {
                width: 80%;
                overflow: hidden;
                text-overflow: ellipsis;
                white-space: nowrap;
            }
            &:last-child {
                align-self: flex-end;
                text-align: right;
            }
        }
    }
    &-title {
        width: 100%;
        text-align: center;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }
}
</style>