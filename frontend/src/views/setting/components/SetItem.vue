<script setup lang="ts">
import type { VNode } from "vue";

defineOptions({
    name: "SetItem",
});

defineProps<{
    title?: string;
    desc?: string;
}>();

const slots = defineSlots<{
    default?: () => VNode[];
    extra?: () => VNode[];
    header?: () => VNode[];
}>();
</script>

<template>
    <div class="set-item">
        <div class="set-item__header">
            <div>
                <div class="set-item__title">{{ title }}</div>
                <slot name="header" />
            </div>
            <div class="set-item__extra">
                <slot />
            </div>
        </div>
        <div v-if="desc || slots.extra" class="set-item__body">
            <slot name="extra">{{ desc }}</slot>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.set-item {
    width: 100%;
    &__header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        & > * {
            height: 100%;
            display: flex;
            align-items: center;
            gap: var(--w-layout-space);
            line-height: 1;
        }
    }
    &__title {
        font-size: 16px;
    }
    &__body {
        margin-top: 5px;
        color: #626b79;
        font-size: 14px;
    }
}
</style>