<script setup lang="ts">
import type { Component } from "vue";

defineOptions({
    name: "Setting",
});

const Base = defineAsyncComponent(() => import("./components/Base.vue"));
const Page = defineAsyncComponent(() => import("./components/Page.vue"));
const About = defineAsyncComponent(() => import("./components/About.vue"));
const Video = defineAsyncComponent(() => import("./components/Video.vue"));
const CloudDrive = defineAsyncComponent(() => import("./components/CloudDrive.vue"));

interface SetMenu {
    name: string;
    icon: string;
    type: string;
    component?: Component;
}

const route = useRoute();
const router = useRouter();

const menuList: SetMenu[] = [
    { name: "通用设置", type: "base", icon: "ele-Setting", component: Base },
    { name: "界面设置", type: "ui", icon: "ele-Monitor", component: Page },
    { name: "视频设置", type: "video", icon: "ele-VideoCamera", component: Video },
    { name: "网盘设置", type: "drive", icon: "cloud-drive", component: CloudDrive },
    { name: "下载设置", type: "download", icon: "ele-Download" },
    { name: "关于PQL", type: "about", icon: "ele-Warning", component: About },
];

const tabActive = ref((route.query.type as string) || menuList.at(0)?.type);

const current = computed(() => menuList.find(v => v.type === tabActive.value));

function onTabClick(item: SetMenu) {
    tabActive.value = item.type;
    router.replace(`${route.path}?type=${item.type}`);
}
</script>

<template>
    <w-sider-layout is-wbox :padding="false">
        <template #sider>
            <div class="menu-list">
                <template v-for="menu, i in menuList" :key="i">
                    <div
                        class="menu-item"
                        :class="{'is-active': menu.type === tabActive}"
                        @click="onTabClick(menu)"
                    >
                        <Icon :icon="menu.icon" />
                        <span>{{ menu.name }}</span>
                    </div>
                </template>
            </div>
        </template>
        <div class="setting-title">{{ current?.name }}</div>
        <div
            :class="{
                'setting-box w-box': !['drive'].includes(current?.type || '')
            }"
        >
            <keep-alive>
                <component
                    :is="current?.component"
                    v-if="current?.component"
                    :key="current.type"
                />
            </keep-alive>
        </div>
    </w-sider-layout>
</template>

<style lang="scss" scoped>
.setting {
    &-title {
        margin-bottom: 10px;
        font-size: 20px;
        color: var(--w-text-color);
    }
    &-box {
        padding: var(--w-layout-space-large);
    }
}

.menu {
    &-list {
        width: var(--w-sider-min-width);
        height: 100%;
        display: flex;
        flex-direction: column;
        gap: 2px;
        padding: var(--w-layout-space);
    }
    &-item {
        height: 34px;
        display: flex;
        align-items: center;
        padding: 0 12px;
        line-height: 1;
        gap: var(--w-layout-space);
        cursor: pointer;
        border-radius: var(--w-border-radius);
        transition: background-color 0.3s var(--w-trans);
        @include when-hover(active) {
            background-color: var(--el-color-primary-light-7);
        }
    }
}
</style>