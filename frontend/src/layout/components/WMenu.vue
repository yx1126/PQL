<script setup lang="ts">
defineOptions({
    name: "WindowMenu",
});

const modelValue = defineModel<string>();

defineProps<{
    stretch?: boolean;
    data?: Menu[];
}>();

export interface Menu {
    name: string;
    icon: string;
    size: number;
    path: string;
    keepPath?: string;
}

function onMenuClick(menu: Menu) {
    modelValue.value = menu.path;
}
</script>

<template>
    <div class="w-menu">
        <template v-for="menu, i in data" :key="i">
            <div
                class="w-menu__item"
                :class="{
                    'is-active': modelValue === menu.path
                }"
                @click="onMenuClick(menu)"
            >
                <Icon :icon="menu.icon" :size="menu.size" />
                <span v-if="menu.name">{{ menu.name }}</span>
            </div>
        </template>
    </div>
</template>

<style lang="scss" scoped>
.w-menu {
    --wails-draggable: no-drag;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    gap: 2px;
    &__item {
        height: 100%;
        padding: 0 15px;
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 8px;
        color: var(--w-text-color);
        cursor: pointer;
        border-radius: 2px;
        transition:
            color 0.3s var(--w-trans),
            background 0.3s var(--w-trans);
        position: relative;
        z-index: 1;
        --w-linear-gradient: color-mix(in oklab, var(--el-color-primary) 10%, transparent);
        &:not(.is-active):hover {
            color: var(--el-color-primary);
            background: linear-gradient(180deg, transparent 10%, var(--w-linear-gradient) 100%);
        }
        @include when(active) {
            color: var(--el-color-primary);
            background: linear-gradient(180deg, transparent 10%, var(--w-linear-gradient) 100%);
            &::after {
                content: "";
                width: 100%;
                height: 3px;
                position: absolute;
                bottom: 0;
                left: 0;
                background-color: var(--el-color-primary);
                border-top-left-radius: 10px;
                border-top-right-radius: 10px;
                pointer-events: none;
            }
        }
    }
}
</style>