<script setup lang="ts">
defineOptions({
    name: "WBacktop",
});

defineProps<{
    target?: string;
    right?: number;
    bottom?: number;
}>();

const emit = defineEmits<{
    refresh: [];
}>();

function onRefresh() {
    emit("refresh");
}
</script>

<template>
    <div
        class="el-backtop"
        :style="{
            right: right ? right + 'px' : '',
            bottom: bottom ? bottom + 'px' : '',
        }"
        @click="onRefresh"
    >
        <Icon icon="ele-Refresh" />
    </div>
    <el-backtop
        :target
        :right
        :bottom="bottom ? bottom + 50 : undefined"
    />
</template>
