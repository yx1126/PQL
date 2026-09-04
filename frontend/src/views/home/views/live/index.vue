<script setup lang="ts">
import KeepRouteView from "@/components/KeepRouteView";
import { menuList } from "./components/data";
import type { CreateLiveVo } from "@bind/vo/models";

defineOptions({
    name: "Live",
});

const LiveDialog = defineAsyncComponent(() => import("./components/LiveDialog.vue"));

const liveDialogRef = useTemplateRef("liveDialogRef");

const keepalive = useMitt("keepalive");
const livelistener = useMitt("live:data:refresh");
const route = useRoute();

function onSuccess(data: CreateLiveVo) {
    const n = route.name as string;
    const rtype = route.query.type || "";
    const keys: string[] = [];
    // 非关注页面添加直播间选中关注 移除缓存
    if(["LiveMine"].includes(n)) {
        if(data.isSpecial === 1) livelistener.emit("LiveMine");
    } else if(["LiveDouyu", "LiveHuya", "LiveDouyin"].includes(n)) {
        if(rtype === data.type) livelistener.emit(n);
        if(data.isSpecial === 1) keys.push("LiveMine");
        // 平台页面
        switch(data.type) {
        case "1":
            rtype != "1" && keys.push("LiveDouyu");
            break;
        case "2":
            rtype != "2" && keys.push("LiveHuya");
            break;
        case "3":
            rtype != "3" && keys.push("LiveDouyin");
            break;
        }
    }
    keepalive.emit(keys);
}
</script>

<template>
    <w-sider-layout is-wbox>
        <template #sider>
            <w-sider-menu class="flex-1" :default-active="route.fullPath" :data="menuList" router />
            <w-sider-actions>
                <el-button
                    type="primary"
                    icon="ele-Plus"
                    title="添加房间"
                    @click="liveDialogRef?.open(route.query.type as string)"
                >
                    添加房间
                </el-button>
            </w-sider-actions>
        </template>
        <keep-route-view />
    </w-sider-layout>
    <live-dialog ref="liveDialogRef" @success="onSuccess" />
</template>
