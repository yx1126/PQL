<script setup lang="ts">
import type { CSSProperties } from "vue";

defineOptions({
    name: "WImage",
});

const {
    size = "100% 100%",
} = defineProps<{
    src?: string;
    type?: "default" | "img";
    size?: CSSProperties["backgroundSize"];
}>();
</script>

<template>
    <img v-if="type === 'img'" :src>
    <div
        v-else
        class="w-image"
        :style="{
            '--w-image-path': `url(${src})`,
            '--w-image-size': size,
        }"
    />
</template>

<style lang="scss" scoped>
.w-image {
    display: inline-flex;
    background-image: var(--w-image-path);
    background-size: var(--w-image-size);
    background-repeat: no-repeat;
}
</style>