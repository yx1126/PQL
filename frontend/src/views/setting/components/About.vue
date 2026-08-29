<script setup lang="ts">
defineOptions({
    name: "SettingAbout",
});

const title = import.meta.env.VITE_APP_TITLE;
const sortTitle = import.meta.env.VITE_APP_TITLE_SHORT;
</script>

<template>
    <div class="set-about">
        <div class="logo-box"></div>
        <div class="text-6">{{ sortTitle }}</div>
        <div>{{ title }}</div>
        <el-button type="primary">检查更新</el-button>
    </div>
</template>

<style lang="scss" scoped>
.set-about {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 15px;
    .logo-box {
        width: 100px;
        height: 100px;
        background-image: url("/PQL.png");
        background-repeat: no-repeat;
        background-size: 100% 100%;
    }
}
</style>