<script setup lang="ts">
import type { ButtonProps, PaginationProps } from "element-plus";

defineOptions({
    name: "WPaging",
});

const page = defineModel<number>("page", { default: 0 });
const size = defineModel<number>("size", { default: 0 });

const {
    type = "default",
    isLast,
    buttonSize = "large",
    layout = "total, prev, pager, next, jumper",
} = defineProps<{
    type?: PagingType;
    total?: number;
    isLast?: boolean;
    buttonSize?: ButtonProps["size"];
    layout?: PaginationProps["layout"];
}>();

const emit = defineEmits<{
    "current-change": [page: number];
    "size-change": [size: number];
}>();

export type PagingType = "default" | "paging";

function onChange(type: "home" | "up" | "down") {
    if(type === "home") {
        page.value = 1;
        onPageChange(1);
    } else if(type === "up") {
        page.value = Math.max(page.value - 1, 1);
        onSizeChange(page.value);
    } else if(type === "down") {
        // 15条/每页
        if(!isLast) {
            page.value = page.value + 1;
            onPageChange(page.value);
        }
    }
}

function onPageChange(value: number) {
    emit("current-change", value);
}

function onSizeChange(value: number) {
    emit("size-change", value);
}
</script>

<template>
    <div class="paging">
        <template v-if="type === 'default'">
            <el-button-group type="primary" :size="buttonSize">
                <el-button icon="ele-ArrowLeft" :disabled="page <= 1" @click="onChange('up')" />
                <el-button v-if="page > 1" icon="ele-House" @click="onChange('home')" />
                <el-button icon="ele-ArrowRight" :disabled="isLast" @click="onChange('down')" />
            </el-button-group>
        </template>
        <template v-if="type === 'paging'">
            <el-pagination
                v-model:current-page="page"
                class="pagination"
                :page-size="size"
                :size="buttonSize"
                background
                :layout
                :total="total"
                @current-change="onPageChange"
                @size-change="onSizeChange"
            />
        </template>
    </div>
</template>

<style lang="scss" scoped>
.paging {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    .pagination {
        --el-pagination-font-size: 16px;
    }
}
</style>