<script setup lang="ts">
import type { CSSProperties } from "vue";

defineOptions({
    name: "WImage",
});

const {
    size = "100% 100%",
    type = "img",
} = defineProps<{
    src?: string;
    type?: "back" | "img";
    size?: CSSProperties["backgroundSize"];
}>();
</script>

<template>
    <img v-if="type === 'img'" class="w-image" :src draggable="false" loading="lazy">
    <div
        v-else
        class="w-image is-back"
        :style="{
            '--w-image-path': `url(${src})`,
            '--w-image-size': size,
        }"
    />
</template>

<style lang="scss" scoped>
.w-image {
    @include when(back) {
        display: inline-flex;
        background-image: var(--w-image-path);
        background-size: var(--w-image-size);
        background-repeat: no-repeat;
    }
}
</style>