<script setup lang="ts">
defineOptions({
    name: "WCard",
});

defineProps<{
    title?: string;
    description?: string;
    isList?: boolean;
    hover?: boolean;
}>();

defineSlots<{
    default?: () => VNode[];
    extra?: () => VNode[];
}>();
</script>

<template>
    <div class="w-card" :class="{ 'is-hover': hover }">
        <div v-if="title" class="w-card__title">
            <div>{{ title }}</div>
            <div class="w-card__extra">
                <slot name="extra" />
            </div>
        </div>
        <div :class="`w-card__${isList ? 'list' : 'desc'}`">
            <slot>{{ description }}</slot>
        </div>
    </div>
</template>

<style lang="scss" scoped>
@use "@/styles/w-lib" as *;
.w-card {
    padding: var(--w-layout-space-large);
    @extend .w-box;
    &__title {
        margin-bottom: 10px;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    &__extra {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: var(--w-layout-space);
    }
    &__list {
        display: flex;
        flex-wrap: wrap;
        gap: var(--w-layout-space-large);
    }
    &__desc {
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