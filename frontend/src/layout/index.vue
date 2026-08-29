<script setup lang="ts">
import WHeader from "./components/WHeader.vue";
import KeepRouteView from "@/components/KeepRouteView";
import Loading from "@/components/Loading";

defineOptions({
    name: "Layout",
});

const state = useStatesStore();
</script>

<template>
    <div class="w-layout h-full">
        <w-header />
        <div class="w-layout__main">
            <Loading v-if="state.loading" position="absolute" mask />
            <keep-route-view />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.w-layout {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    &__main {
        width: 100%;
        height: calc(100% - var(--w-header-height));
        overflow: hidden;
        display: flex;
        flex-direction: column;
        position: relative;
    }
}
</style>