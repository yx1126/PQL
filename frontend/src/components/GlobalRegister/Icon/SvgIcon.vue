<script lang="ts" setup>
defineOptions({
    name: "SvgIcon",
});

const {
    prefix = "local",
    icon,
} = defineProps<{
    prefix?: string;
    icon?: string;
}>();

const symbolId = computed(() => `#${prefix}-${icon}`);
</script>

<template>
    <svg class="svg-icon" aria-hidden="true">
        <use :href="symbolId" :xlink:href="symbolId" />
    </svg>
</template>

<style>
.svg-icon {
    width: 1em;
    height: 1em;
    fill: currentColor;
    overflow: hidden;
}
</style>