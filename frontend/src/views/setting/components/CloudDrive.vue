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

function getBaiduVip(value: number) {
    switch(value) {
    case 0:
        return "普通用户";
    case 1:
        return "会员VIP用户";
    case 2:
        return "超级会员SVIP用户";
    default:
        return "";
    }
}

function isAuth(v: AuthVo) {
    return !!v.token && !!v.refresh_token;
}

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
    <div class="w-drive">
        <template v-for="item in data" :key="item.id">
            <w-card border :show-body="isAuth(item)">
                <template #header>
                    <w-text :icon="item.icon" size="30">{{ item.name }}</w-text>
                </template>
                <template #extra>
                    <el-button v-if="isAuth(item)" type="danger" round plain @click="onUnBind(item)">解绑</el-button>
                    <el-button v-else type="primary" round plain @click="onItemClick(item)">绑定</el-button>
                </template>
                <div v-if="isAuth(item)" class="flex items-center justify-between">
                    <div class="flex items-center gap-[var(--w-layout-space)]">
                        <w-image class="size-[50px] rounded-[50%]" :src="item.avatar" />
                        <span>{{ item.nickname }}（{{ item.username }}）</span>
                    </div>
                    <span>{{ getBaiduVip(item.vip_type) }}</span>
                </div>
            </w-card>
        </template>
        <baidu-auth ref="baiduAuthRef" @success="query" />
    </div>
</template>

<style lang="scss" scoped>
@use "@/styles/w-lib" as *;
.w-drive {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: var(--w-layout-space-large);
}
</style>