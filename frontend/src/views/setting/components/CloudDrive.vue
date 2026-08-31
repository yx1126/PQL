<script setup lang="ts">
defineOptions({
    name: "CloudDrive",
});
</script>

<template>
    <div class="cloud-drive">
        <div class="cloud-drive-item">
            <div>
                <w-text icon="baidu" size="30">百度网盘</w-text>
            </div>
            <el-tag class="cursor-pointer" type="danger" size="large" round>绑定</el-tag>
        </div>
    </div>
</template>

<style lang="scss" scoped>
@use "@/styles/w-lib" as *;
.cloud-drive {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: var(--w-layout-space-large);
    &-item {
        @extend .w-box;
        padding: var(--w-layout-space-large);
        line-height: 1;
        display: flex;
        align-items: center;
        justify-content: space-between;
        & > * {
            display: flex;
            align-items: center;
            gap: var(--w-layout-space);
        }
    }
}
</style>