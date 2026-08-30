<script setup lang="ts">
import KeepRouteView from "@/components/KeepRouteView";

defineOptions({
    name: "SubLayout",
});

const route = useRoute();
const router = useRouter();

const subTitle = computed(() => route.matched.findLast(v => v.meta.subTitle)?.meta.subTitle || "");

function onBack() {
    router.back();
}
</script>

<template>
    <div class="sublayout">
        <div class="sublayout__header">
            <div class="sublayout__header-main">
                <el-link icon="ele-Back" @click="onBack" />
                <span class="title">{{ subTitle }}</span>
            </div>
            <div class="sublayout__header-extra">
                <portal-target name="layout-extra" />
            </div>
        </div>
        <div class="sublayout__main">
            <keep-route-view />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.sublayout {
    width: 100%;
    height: 100%;
    padding: var(--w-layout-space-large);
    &__header {
        height: 60px;
        background-color: var(--w-box-bg);
        border: 1px solid var(--w-border-color);
        border-radius: var(--w-border-radius);
        display: flex;
        align-items: center;
        jubstify-content: space-between;
        padding: 0 15px;
        animation-name: slide-top;
        animation-duration: 0.3s;
        animation-fill-mode: forwards;
        &-main {
            flex: 1;
            display: flex;
            align-items: center;
            gap: var(--w-layout-space);
            --el-font-size-base: 18px;
            .title {
                // font-size: 18px;
                line-height: 1;
            }
        }
        &-extra {
            display: flex;
            align-items: center;
            gap: var(--w-layout-space);
        }
    }
    &__main {
        width: 100%;
        margin-top: var(--w-layout-space);
        height: calc(100% - var(--w-layout-space) - 60px);
        min-height: 0;
        display: flex;
        flex-direction: column;
        animation-name: slide-bottom;
        animation-duration: 0.3s;
        animation-fill-mode: forwards;
    }
}
</style>