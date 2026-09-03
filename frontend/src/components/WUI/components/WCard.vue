<script setup lang="ts">
defineOptions({
    name: "WCard",
});

const {
    showBody = true,
} = defineProps<{
    title?: string;
    description?: string;
    isList?: boolean;
    hover?: boolean;
    border?: boolean;
    showBody?: boolean;
}>();

const slots = defineSlots<{
    default?: () => VNode[];
    header?: () => VNode[];
    extra?: () => VNode[];
}>();
</script>

<template>
    <div
        class="w-card"
        :class="{
            'is-hover': hover,
        }"
    >
        <div v-if="title || slots.header || slots.extra" class="w-card__header">
            <div>
                <slot name="header">{{ title }}</slot>
            </div>
            <div class="w-card__extra">
                <slot name="extra" />
            </div>
        </div>
        <template v-if="showBody">
            <el-divider v-if="border" />
            <div :class="`w-card__${isList ? 'list' : 'desc'}`">
                <slot>{{ description }}</slot>
            </div>
        </template>
    </div>
</template>

<style lang="scss" scoped>
@use "@/styles/w-lib" as *;
.w-card {
    padding: var(--w-layout-space-large);
    @extend .w-box;
    &__header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        & > div {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: var(--w-layout-space);
        }
    }
    &__list {
        display: flex;
        flex-wrap: wrap;
        gap: var(--w-layout-space-large);
        margin-top: 10px;
    }
    &__desc {
        margin-top: 10px;
        width: 100%;
        min-width: 0;
        font-size: 14px;
        color: #999;
        /* 关键换行 */
        word-break: break-all;
        overflow-wrap: break-word;
        white-space: normal;
    }
    @include when(hover) {
        .w-card__desc:hover {
            color: var(--el-color-primary-light-3);
        }
    }
}
</style>