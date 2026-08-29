<script setup lang="ts">
import KeepRouteView from "@/components/KeepRouteView";

defineOptions({
    name: "Home",
});

const route = useRoute();
const menuStore = useMenuStore();

watch(() => route.fullPath, () => {
    onSavePath();
});

// 解决切换时子页面重定向问题
function onSavePath() {
    if(route.meta.keepType) {
        const menu = menuStore.menuList.find(v => v.icon === route.meta.keepType);
        if(menu) {
            menuStore.keepMap[menu.icon] = route.fullPath;
        }
    }
}
</script>

<template>
    <div class="home">
        <keep-route-view />
    </div>
</template>

<style lang="scss" scoped>
.home {
    width: 100%;
    height: 100%;
}
</style>
