<script setup lang="ts" generic="T extends object">
import jp, { type ParseOption, type RuleItem } from "@/utils/parse";
import type { PagingType } from "@/components/WUI";

defineOptions({
    name: "VideoListBox",
});

const page = defineModel<number>("page", { default: 0 });
const size = defineModel<number>("size", { default: 0 });

const {
    data = [],
    paging = true,
} = defineProps<{
    data?: T[];
    config?: RuleItem;
    total?: number;
    pagingType?: PagingType;
    maxSize?: ParseOption["defaultMaxSize"];
    getImgUri?: (path?: string | undefined) => string;
    paging?: boolean;
}>();

const emit = defineEmits<{
    "item-click": [item: T];
    "current-change": [page: number];
    "size-change": [size: number];
}>();

function onClick(item: T) {
    emit("item-click", item);
}

function onPageChange(value: number) {
    emit("current-change", value);
}

function onSizeChange(value: number) {
    emit("size-change", value);
}
</script>

<template>
    <div class="vl">
        <div class="vl-list">
            <el-empty v-if="data.length <= 0" class="size-full col-span-6" />
            <template v-for="item, i in data" v-else :key="i">
                <w-video
                    :title="jp.value(item, config?.namePath)"
                    :top="jp.value(item, config?.coverTopPath)"
                    :left="jp.value(item, config?.coverLeftPath)"
                    :right="jp.value(item, config?.coverRightPath)"
                    :src="getImgUri ? getImgUri(jp.value(item, config?.srcPath)) : undefined"
                    @click="onClick(item)"
                />
            </template>
        </div>
        <div v-if="paging" class="vl-paging">
            <w-paging
                v-model:page="page"
                v-model:size="size"
                :total="total"
                :type="pagingType"
                :is-last="maxSize ? data.length < maxSize : false"
                @current-change="onPageChange"
                @size-change="onSizeChange"
            />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.vl {
    &-list {
        display: grid;
        grid-template-columns: repeat(5, minmax(150px, 1fr));
        gap: 30px;
        position: relative;
        @media screen and (width > 1400px) and (width <= 1600px) {
            grid-template-columns: repeat(6, minmax(160px, 1fr));
        }
        @media screen and (width > 1600px) {
            grid-template-columns: repeat(8, minmax(160px, 1fr));
        }
    }
    &-paging {
        margin-top: 15px;
    }
}
</style>