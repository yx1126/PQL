<script setup lang="ts">
import { useTransition, type UseTransitionOptions } from "@vueuse/core";

defineOptions({
    name: "CountTo",
});

const {
    value,
    options,
    duration = 500,
} = defineProps<{
    value: number;
    duration?: number;
    options?: UnRefable<Omit<UseTransitionOptions<number>, "duration">>;
}>();

const source = ref(0);
const precision = ref(0);

const disabled = computed(() => source.value === 0);

const outputValue = useTransition(source, {
    ...options,
    duration,
    disabled,
});

watch(() => value, val => {
    const str = val.toString();
    precision.value = str.includes(".") ? str.split(".").at(-1)!.length : 0;
    nextTick(() => {
        source.value = val;
    });
}, {
    immediate: true,
});
</script>

<template>
    <el-statistic :value="outputValue" :precision />
</template>