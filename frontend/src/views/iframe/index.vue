<script setup lang="ts">
defineOptions({
    name: "IFramePage",
});

const route = useRoute();

const show = ref(false);
const loading = ref(false);

const src = ref<string>("");

watch(() => route.fullPath, () => {
    show.value = false;
    loading.value = true;
    src.value = getLink();
    nextTick(() => {
        show.value = true;
    });
}, {
    immediate: true,
});

function getLink() {
    const link = route.params.link || route.meta.link;
    return decodeURIComponent(link as string);
}

function onLoad() {
    loading.value = false;
}

function onError() {
    loading.value = false;
}
</script>

<template>
    <div v-loading="loading" class="iframe">
        <el-card class="full-page" shadow="never">
            <iframe v-if="show" class="iframe-main" :src frameborder="0" @load="onLoad" @error="onError"></iframe>
        </el-card>
    </div>
</template>

<style lang="scss" scoped>
.iframe-main {
    width: 100%;
    height: 100%;
}
</style>