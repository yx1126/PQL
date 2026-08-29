<script setup lang="ts" generic="T">
import { parseTemp } from "@/utils/parse";
import { isArray, isObject, isStr, toString } from "@/utils/validata";
import { Clipboard } from "@wailsio/runtime";

defineOptions({
    name: "LinkList",
});

const {
    label,
    options = [],
    separator = "/",
    disabled,
    template,
    templateData = {},
} = defineProps<{
    options?: T | T[];
    separator?: string;
    templateData?: Record<string, any>;
    label?: string;
    disabled?: boolean;
    template?: string;
}>();

const dataList = computed<any[]>(() => {
    return isArray<T[]>(options)
        ? options
        : [{ label: options }].map(v => {
            return isObject(v) ? v : { label: v };
        });
});

function getKey(item: T): string {
    const value = isStr(item) ? item : toString(item[(label || "label") as keyof T]);
    if(template) {
        const MAP: Record<string, any> = {
            ...templateData,
            value,
        };
        return parseTemp(MAP, template);
    }
    return value;
}

async function onClick(item: T) {
    if(disabled) return;
    const text = getKey(item);
    if(text) {
        Clipboard.SetText(text);
    }
}
</script>

<template>
    <div class="link-list">
        <template v-for="item, i in dataList" :key="i">
            <span v-if="disabled" @click="onClick(item)">{{ getKey(item) }}</span>
            <el-link v-else @click="onClick(item)">{{ getKey(item) }}</el-link>
            <span v-if="i !== dataList.length - 1" class="link-list__separator">{{ separator || "/" }}</span>
        </template>
    </div>
</template>

<style lang="scss" scoped>
.link-list {
    width: 100%;
    height: 100%;
    display: inline-flex;
    flex-wrap: wrap;
    gap: 5px;
}
</style>