<script setup lang="ts">
import { isStr } from "@/utils/validata";
import SvgIcon from "./SvgIcon.vue";
import type { CSSProperties, Component } from "vue";
import { parseUnit } from "@/utils/unit";
import type { LinkProps } from "element-plus";

defineOptions({
    name: "Icon",
});

const { rotate = 0, cursor = false, size, type } = defineProps<{
    icon?: string | Component;
    rotate?: string | number;
    cursor?: boolean | CSSProperties["cursor"];
    size?: number | string;
    type?: LinkProps["type"];
}>();

const style = computed(() => {
    const styles: CSSProperties = { fontSize: parseUnit(size) };
    if(rotate) {
        styles.transform = `rotate(${rotate}deg)`;
    }
    if(cursor) {
        styles.cursor = cursor ? isStr(cursor) ? cursor : "pointer" : undefined;
    }
    if(type) {
        styles.color = `var(--el-color-${type})`;
    }
    return styles;
});
</script>

<template>
    <el-icon :style="style">
        <slot>
            <svg-icon v-if="isStr(icon) && !(icon.startsWith('ele-') || icon.startsWith('Ele'))" :icon="icon" />
            <component :is="icon" v-else />
        </slot>
    </el-icon>
</template>