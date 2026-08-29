<script setup lang="ts">
import KeepRouteView from "@/components/KeepRouteView";
import type { MenuSiderItem } from "@/components/WUI";

defineOptions({
    name: "Live",
});

const route = useRoute();

const menuList: MenuSiderItem[] = [
    { id: 0, label: "关注", icon: "heart", size: 18, path: "/home/live" },
    // { id: 1, label: "赛事", icon: "competition", path: "/home/live/competition" },
    // { id: 1, label: "搜索", icon: "ele-Search", path: "/home/live/search" },
];
</script>

<template>
    <w-sider-layout is-wbox>
        <template #sider>
            <w-sider-menu class="flex-1" :default-active="route.path" :data="menuList" router />
        </template>
        <keep-route-view />
    </w-sider-layout>
</template>
