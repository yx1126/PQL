<script setup lang="ts">
import SetItem from "./SetItem.vue";

defineOptions({
    name: "SettingPage",
});

const set = useSetStore();
</script>

<template>
    <div class="">
        <set-item title="顶部主题切换">
            <el-switch v-model="set.headerShowTheme" active-value="1" inactive-value="0" />
        </set-item>
    </div>
</template>
