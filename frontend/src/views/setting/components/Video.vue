<script setup lang="ts">
import { AppService } from "@bind/service/index.ts";
import SetItem from "./SetItem.vue";
import { LoadInline } from "@/components/Loading";

defineOptions({
    name: "SettingVideo",
});

const { data: isInstall, loading } = useService(AppService.HasHEVCExtension, false);
</script>

<template>
    <div class="set-video">
        <set-item title="Hevc视频扩展">
            <load-inline :loading type="tag">
                <el-tag v-if="isInstall" type="primary" size="large" round>已安装</el-tag>
                <el-tag v-else type="danger" size="large" round>未安装</el-tag>
            </load-inline>
        </set-item>
    </div>
</template>