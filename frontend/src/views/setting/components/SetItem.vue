<script setup lang="ts">
defineOptions({
    name: "SetItem",
});

defineProps<{
    title?: string;
    desc?: string;
}>();
</script>

<template>
    <div class="set-item">
        <div class="set-item__body">
            <div>
                <div class="set-item__title">{{ title }}</div>
            </div>
            <div class="set-item__extra">
                <slot />
            </div>
        </div>
        <div v-if="desc" class="set-item__desc">{{ desc }}</div>
    </div>
</template>

<style lang="scss" scoped>
.set-item {
    width: 100%;
    &__body {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    &__title {
        font-size: 16px;
    }
    &__desc {
        margin-top: 5px;
        color: #626b79;
        font-size: 14px;
    }
}
</style>