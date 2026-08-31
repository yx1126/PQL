<script setup lang="ts">
import WImage from "./WImage.vue";
import { parseUnit } from "@/utils/unit";
import type { VNode } from "vue";

defineOptions({
    name: "WText",
});

const {
    type = "icon",
    size = 16,
} = defineProps<{
    type?: "icon" | "img";
    icon?: string;
    size?: Unit;
    gap?: Unit;
    reverse?: boolean;
}>();

defineSlots<{
    default?: () => VNode[];
}>();
</script>

<template>
    <div
        class="w-text"
        :style="{
            '--w-text-size': parseUnit(size),
            '--w-text-gap': parseUnit(gap),
        }"
    >
        <span v-if="reverse"><slot /></span>
        <template v-if="icon">
            <Icon v-if="type === 'icon'" class="w-text__icon" :icon />
            <w-image v-if="type === 'img'" class="w-text__img" :src="icon" />
        </template>
        <span v-if="!reverse"><slot /></span>
    </div>
</template>

<style lang="scss" scoped>
.w-text {
    display: inline-flex;
    align-items: center;
    line-height: 1;
    gap: var(--w-text-gap, var(--w-layout-space));
    &__icon {
        font-size: var(--w-text-size);
    }
    &__img {
        width: var(--w-text-size);
        min-width: var(--w-text-size);
        height: var(--w-text-size);
        min-height: var(--w-text-size);
    }
}
</style>