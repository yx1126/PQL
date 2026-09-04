<script setup lang="ts">
import { parseUnit } from "@/utils/unit";
import type { VNode } from "vue";

defineOptions({
    name: "WPlayerLayout",
});

const {
    siderWidth = 400,
} = defineProps<{
    title?: string;
    subTitle?: string;
    siderWidth?: Unit;
}>();

defineSlots<{
    default?: () => VNode[];
    sider?: () => VNode[];
}>();
</script>

<template>
    <div
        class="w-player-layout"
        :style="{
            '--w-player-layout-sider-width': parseUnit(siderWidth)
        }"
    >
        <div class="w-player-title">
            <span>{{ title }}</span>
            <span v-if="subTitle">（{{ subTitle }}）</span>
        </div>
        <div class="w-player-main">
            <div class="w-player__left">
                <slot />
            </div>
            <div class="w-player__right w-box">
                <slot name="sider" />
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.w-player {
    &-layout {
        width: 100%;
        height: 100%;
        padding: var(--w-layout-space-large);
        display: flex;
        flex-direction: column;
        gap: var(--w-layout-space-large);
        position: relative;
    }
    &-title {
        font-size: 24px;
        padding: 5px 0;
    }
    &-main {
        min-height: 0;
        flex: 1;
        display: flex;
        gap: var(--w-layout-space-large);
    }
    &__left {
        flex: 1;
        display: flex;
        flex-direction: column;
        gap: var(--w-layout-space-large);
        overflow-x: hidden;
        overflow-y: auto;
        position: relative;
        @include hidden-scroll;
    }
    &__right {
        width: var(--w-player-layout-sider-width);
        height: 100%;
        gap: var(--w-layout-space-large);
        padding: var(--w-layout-space-large);
    }
}
</style>