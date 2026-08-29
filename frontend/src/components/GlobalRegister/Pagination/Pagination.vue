<script setup lang="ts">
defineOptions({
    name: "Pagination",
});

const {
    layout = ["total", "->", "prev", "pager", "next", "sizes", "jumper"].join(", "),
} = defineProps<{
    layout?: string;
}>();
</script>

<template>
    <el-pagination :layout />
</template>
