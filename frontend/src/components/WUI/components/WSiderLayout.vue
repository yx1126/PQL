<script setup lang="ts">
import type { VNode } from "vue";

defineOptions({
    name: "WSiderLayout",
});

const {
    padding = true,
} = defineProps<{
    isWbox?: boolean;
    padding?: boolean;
}>();

defineSlots<{
    sider?: () => VNode[];
    default?: () => VNode[];
}>();
</script>

<template>
    <div class="w-sider-layout" :class="{'is-padding': padding}">
        <div class="w-sider-layout__sider" :class="{'is-wbox': isWbox}">
            <slot name="sider" />
        </div>
        <div class="w-sider-layout__main">
            <slot name="default" />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.w-sider-layout {
    width: 100%;
    height: 100%;
    display: flex;
    gap: var(--w-layout-space);
    @include when(padding) {
        padding: var(--w-layout-space);
    }
    &__sider {
        width: var(--w-sider-min-width);
        height: 100%;
        display: flex;
        flex-direction: column;
        gap: var(--w-layout-space);
        animation-name: slide-left;
        animation-duration: 0.3s;
        animation-fill-mode: forwards;
        @include when(wbox) {
            & > * {
                background-color: var(--w-box-bg);
                border-radius: var(--w-border-radius);
                border: 1px solid var(--w-border-color);
                padding: var(--w-layout-space);
            }
        }
    }
    &__main {
        width: calc(100% - var(--w-layout-space) - var(--w-sider-min-width));
        height: 100%;
        display: flex;
        flex-direction: column;
        animation-name: slide-right;
        animation-duration: 0.3s;
        animation-fill-mode: forwards;
    }
}
</style>