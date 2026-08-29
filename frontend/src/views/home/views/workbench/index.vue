<script setup lang="ts">
defineOptions({
    name: "Workbench",
});
const cardList = [
    { title: "CPU", icon: "win-cpu" },
    { title: "GPU", icon: "win-gpu" },
    { title: "内存", icon: "win-memory" },
    { title: "硬盘", icon: "win-hard-disk" },
    { title: "网络", icon: "win-internet" },
];
</script>

<template>
    <div class="workbench">
        <div>
            <div class="grid grid-cols-5 gap-[var(--w-layout-space)]">
                <template v-for="card, i in cardList" :key="i">
                    <div class="yj-card">
                        <div class="yj-card-title">
                            <Icon :icon="card.icon" />
                            <span>{{ card.title }}</span>
                        </div>
                    </div>
                </template>
            </div>
        </div>
        <w-card title="系统快捷工具">
            系统快捷工具
        </w-card>
    </div>
</template>

<style lang="scss" scoped>
@use "@/styles/w-lib" as *;
.workbench {
    width: 100%;
    height: 100%;
    display: grid;
    grid-template-columns: 2fr 1fr;
    padding: var(--w-layout-space);
    gap: var(--w-layout-space);
    .yj-card {
        @extend .w-box;
        padding: var(--w-layout-space-large);
        line-height: 1;
        &-title {
            display: flex;
            align-items: center;
            gap: var(--w-layout-space);
        }
    }
}
</style>