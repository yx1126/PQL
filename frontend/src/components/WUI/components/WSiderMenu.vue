<script setup lang="ts">
defineOptions({
    name: "WSiderMenu",
});

const {
    router: isRouter,
} = defineProps<{
    data?: MenuSiderItem[];
    router?: boolean;
    defaultActive?: string;
}>();

const router = useRouter();

export interface MenuSiderItem {
    id?: string | number;
    label?: string;
    icon?: string;
    path?: string;
    size?: Unit;
    disabled?: boolean;
    router?: boolean;
    isActive?: boolean;
}

function onMenuClick(menu: MenuSiderItem) {
    if(menu.disabled) return;
    if(isRouter && (menu.router ?? true) && menu.path) {
        router.push(menu.path);
    }
}
</script>

<template>
    <div class="w-sider-menu">
        <template v-for="menu in data" :key="menu.id">
            <div
                class="w-sider-menu__item"
                :class="{
                    'is-active': (menu.isActive ?? true) ? defaultActive === menu.path : false,
                    'is-disabled': menu.disabled
                }"
                @click="onMenuClick(menu)"
            >
                <Icon v-if="menu.icon" :icon="menu.icon" :size="menu.size" />
                <span>{{ menu.label }}</span>
            </div>
        </template>
    </div>
</template>

<style lang="scss" scoped>
.w-sider-menu {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 5px;
    &__item {
        display: flex;
        align-items: center;
        padding: 0 15px;
        gap: 10px;
        height: 40px;
        line-height: 1;
        border: 1px solid transparent;
        border-radius: var(--w-border-radius);
        cursor: pointer;
        transition:
            border-color 0.2s var(--w-trans),
            background-color 0.2s var(--w-trans);
        &:hover:not(.is-disabled) {
            background-color: var(--el-color-primary-light-9);
            color: var(--el-color-primary);
        }
        @include when(active) {
            &:not(.is-disabled) {
                // border-color: var(--el-color-primary-light-3);
                color: var(--el-color-primary);
                background-color: var(--el-color-primary-light-9);
            }
        }
        @include when(disabled) {
            cursor: not-allowed;
            opacity: 0.25;
        }
    }
}
</style>