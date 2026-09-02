<script setup lang="ts">
import { DriveService } from "@bind/service";
import BaiduAuth from "./BaiduAuth.vue";
import type { AuthVo } from "@bind/vo";

defineOptions({
    name: "CloudDrive",
});

const baiduAuthRef = useTemplateRef("baiduAuthRef");

const message = useMessage();
const msgbox = useMessageBox();

const { data, query } = useService({
    request: DriveService.GetAuthList,
    default: [],
    immediate: true,
    isLayoutLoad: true,
});

function onItemClick(item: AuthVo) {
    if(item.type === "baidu") {
        baiduAuthRef.value?.open();
    }
}

function onUnBind(item: AuthVo) {
    msgbox.confirm("确认要解绑吗？").then(async () => {
        await DriveService.UnBindBaidu(item.type);
        message.success("解绑成功！");
        query();
    });
}
</script>

<template>
    <div class="cloud-drive">
        <template v-for="item in data" :key="item.id">
            <div class="cloud-drive-item">
                <div>
                    <w-text :icon="item.icon" size="30">{{ item.name }}</w-text>
                </div>
                <el-tag v-if="item.token && item.refresh_token" class="cursor-pointer" type="primary" size="large" round @click="onUnBind(item)">解绑</el-tag>
                <el-tag v-else class="cursor-pointer" type="danger" size="large" round @click="onItemClick(item)">绑定</el-tag>
            </div>
        </template>
        <baidu-auth ref="baiduAuthRef" @success="query" />
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