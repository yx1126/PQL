<script setup lang="ts">
import type { VNode } from "vue";

defineOptions({
    name: "WSiderActions",
});

defineSlots<{
    default?: () => VNode[];
}>();
</script>

<template>
    <div class="w-sider-actions">
        <slot />
    </div>
</template>

<style lang="scss" scoped>
.w-sider-actions {
    min-height: 40px;
    display: flex;
    gap: 8px;
    :slotted(.el-button) {
        height: 35px;
        flex: 1;
        --el-font-size-base: 22px;
        span {
            font-size: calc(var(--el-font-size-base) - 6px);
        }
        & + .el-button {
            margin: 0;
        }
    }
}
</style>