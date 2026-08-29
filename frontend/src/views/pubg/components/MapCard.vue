<script setup lang="ts">
const {
    src,
} = defineProps<{
    mapName: string;
    src: string;
}>();
</script>

<template>
    <div class="card">
        <div class="img__wrapper">
            <w-image :src="`/local/pubg/${src}`" :alt="mapName" class="card__img" size="cover" />
        </div>
        <div class="text_wrapper">
            <p>{{ mapName }}</p>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.card {
    display: flex;
    flex-direction: column;
    background-color: white;
    border-radius: var(--w-border-radius);
    cursor: pointer;
    transition-duration: 150ms;
    color: #333;
    &:hover {
        background-color: var(--el-color-primary);
        transform: scale(1.1);
        color: #fff;
    }
    .img__wrapper {
        padding: 3px;
    }
    &__img {
        width: 100%;
        min-width: 200px;
        height: 200px;
    }
    .text_wrapper {
        margin-bottom: 5px;
        text-align: center;
        font-size: x-large;
    }
}
</style>
