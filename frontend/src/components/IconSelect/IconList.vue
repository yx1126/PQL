<script setup lang="ts">
import Pagination from "./Pagination.vue";

defineOptions({
    name: "IconList",
});

const {
    value,
    options = [],
} = defineProps<{
    value?: string;
    options?: string[];
}>();

const emit = defineEmits<{
    change: [string];
}>();

const page = ref(1);
const size = ref(24);
const iconName = ref("");

const showList = computed(() => {
    const pageNum = (page.value - 1) * size.value;
    return options.slice(pageNum, pageNum + size.value);
});

watch(() => options, () => {
    reset();
}, {
    deep: true,
});

function onClick(icon: string) {
    emit("change", icon);
}

function reset() {
    let current = 1;
    if(value) {
        const index = options.findIndex(v => v === value);
        if(index !== -1) {
            current = Math.max(1, Math.ceil((index + 1) / size.value));
        }
    }
    page.value = current;
}

function onIconMouse(icon: string) {
    iconName.value = icon;
}
defineExpose({
    reset,
});
</script>

<template>
    <div class="icon-list">
        <ul class="icon-select-list">
            <template v-for="icon in showList" :key="icon">
                <li
                    :class="{ 'is-active': icon === value }"
                    @click="onClick(icon)"
                    @mouseenter="onIconMouse(icon)"
                    @mouseleave="onIconMouse('')"
                >
                    <Icon :icon />
                </li>
            </template>
        </ul>
        <div class="paging">
            <div class="icon">{{ iconName }}</div>
            <Pagination
                v-model:page="page"
                :page-size="size"
                :total="options.length || 0"
            />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.icon-list {
    .icon-select-list {
        --n-bezier: cubic-bezier(0.4, 0, 0.2, 1);
        height: 166px;
        margin: 10px 0;
        display: grid;
        grid-template-columns: repeat(6, 50px);
        grid-template-rows: repeat(4, 34px);
        grid-gap: 10px;
        align-content: flex-start;
        & > li {
            border: 1px solid var(--el-border-color);
            border-radius: 3px;
            padding: 0 14px;
            line-height: 34px;
            font-size: 20px;
            cursor: pointer;
            transition:
                color 0.3s var(--n-bezier),
                background-color 0.3s var(--n-bezier),
                opacity 0.3s var(--n-bezier),
                border-color 0.3s var(--n-bezier);
            &:hover,
            &:active,
            &.is-active {
                border-color: var(--el-color-primary);
                color: var(--el-color-primary);
            }
        }
    }
    .paging {
        display: flex;
        justify-content: space-between;
    }
}
</style>