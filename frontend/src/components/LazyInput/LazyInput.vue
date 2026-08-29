<script setup lang="ts">
import { debounce, type DebouncedFunc } from "lodash-es";

defineOptions({
    name: "LazyInput",
});

const modelValue = defineModel<string>();

const { delay = 500, modelModifiers = { lazy: false } } = defineProps<{
    modelModifiers?: { lazy?: boolean };
    delay?: number;
}>();

const inputRef = useTemplateRef("inputRef");

const inputValue = ref<Noable<string>>(modelValue.value);

let updateValue: Nullable<DebouncedFunc<(value?: string) => void>> = null;

watch(() => delay, value => {
    updateValue = debounce((value?: string) => {
        modelValue.value = value;
    }, modelModifiers?.lazy ? value : 0);
    onWatcherCleanup(() => {
        updateValue?.cancel();
        updateValue = null;
    });
}, {
    immediate: true,
});

watch(inputValue, value => {
    if(updateValue) {
        updateValue(value);
    }
});

watch(modelValue, value => {
    inputValue.value = value;
});

defineExpose(new Proxy({}, {
    get(_, key, receiver) {
        if(!inputRef.value) return;
        return Reflect.get(inputRef.value, key, receiver);
    },
    has(_, key) {
        if(!inputRef.value) return false;
        return Reflect.has(inputRef.value, key);
    },
}));
</script>

<template>
    <el-input ref="inputRef" v-model="inputValue" v-bind="$attrs">
        <template v-for="slot in Object.keys($slots)" #[slot] :key="slot">
            <slot :name="slot" />
        </template>
    </el-input>
</template>