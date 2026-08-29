<script setup lang="ts">
import { parseUnit } from "@/utils/unit";
import type { VNode } from "vue";

defineOptions({
    name: "LoadInline",
});

const {
    size = 20,
} = defineProps<{
    loading?: boolean;
    type?: "tag";
    size?: string | number;
}>();

defineSlots<{
    default: () => VNode[];
}>();

const svg = `<path fill="currentColor" d="M512 64a32 32 0 0 1 32 32v192a32 32 0 0 1-64 0V96a32 32 0 0 1 32-32m0 640a32 32 0 0 1 32 32v192a32 32 0 1 1-64 0V736a32 32 0 0 1 32-32m448-192a32 32 0 0 1-32 32H736a32 32 0 1 1 0-64h192a32 32 0 0 1 32 32m-640 0a32 32 0 0 1-32 32H96a32 32 0 0 1 0-64h192a32 32 0 0 1 32 32M195.2 195.2a32 32 0 0 1 45.248 0L376.32 331.008a32 32 0 0 1-45.248 45.248L195.2 240.448a32 32 0 0 1 0-45.248m452.544 452.544a32 32 0 0 1 45.248 0L828.8 783.552a32 32 0 0 1-45.248 45.248L647.744 692.992a32 32 0 0 1 0-45.248M828.8 195.264a32 32 0 0 1 0 45.184L692.992 376.32a32 32 0 0 1-45.248-45.248l135.808-135.808a32 32 0 0 1 45.248 0m-452.544 452.48a32 32 0 0 1 0 45.248L240.448 828.8a32 32 0 0 1-45.248-45.248l135.808-135.808a32 32 0 0 1 45.248 0"></path>`;
</script>

<template>
    <div
        v-loading="loading"
        class="w-load-inline"
        :class="{
            [`is-${type}`]: !!type
        }"
        element-loading-svg-view-box="0 0 1024 1024"
        :element-loading-spinner="svg"
        element-loading-background="var(--el-color-primary-light-9)"
        :style="{
            '--el-loading-spinner-size': parseUnit(size)
        }"
    >
        <slot />
    </div>
</template>

<style lang="scss" scoped>
.w-load-inline {
    display: inline-flex;
    justify-content: center;
    align-items: center;
    @include when(tag) {
        :deep(.el-loading-mask) {
            border-radius: 9999px;
            color: var(--el-color-primary);
        }
    }
}
</style>