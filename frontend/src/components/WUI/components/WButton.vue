<script setup lang="ts">
import type { VNode } from "vue";

defineOptions({
    name: "WButton",
});

const {
    type = "",
} = defineProps<{
    type?: "" | "default" | "primary";
    size?: "" | "default" | "medium" | "large";
    plain?: boolean;
    border?: boolean;
}>();

defineSlots<{
    default?: () => VNode[];
}>();
</script>

<template>
    <div
        class="w-button"
        :class="{
            [`w-button--${type}`]: !!type,
            [`w-button--${size}`]: !!size,
            'w-button--bordered': border,
            'is-plain': plain
        }"
    >
        <slot />
    </div>
</template>

<style lang="scss" scoped>
%hover-active {
    color: #fff;
    background-color: var(--el-color-primary);
    @include when-dark {
        @include when(plain) {
            color: var(--w-text-color);
            background-color: var(--el-color-primary-light-7);
        }
    }
}

.w-button {
    min-height: 0;
    min-width: 0;
    line-height: 1;
    cursor: pointer;
    padding: 5px 10px;
    border-radius: 4px;
    text-align: center;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    transition:
        color 0.3s var(--w-trans),
        background-color 0.3s var(--w-trans);
    &--primary {
        @extend %hover-active;
    }
    &--medium {
        padding: 10px;
    }
    &--large {
        padding: 15px;
    }
    &--bordered {
        border: 1px solid var(--w-border-color);
    }
    &:hover {
        @extend %hover-active;
    }
}
</style>