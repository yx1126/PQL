<script lang="ts" setup>
import * as Echarts from "echarts";
import { useResizeObserver } from "@vueuse/core";
import type { EChartsOption, ECharts as EchartsInstance } from "echarts";

const {
    options,
    renderer = "canvas",
    dark,
    beforeRender,
    afterRender,
} = defineProps<{
    options: EChartsOption;
    renderer?: "svg" | "canvas";
    dark?: boolean;
    beforeRender?: (echarts: typeof Echarts) => void;
    afterRender?: (instance: EchartsInstance, echarts: typeof Echarts) => void;
}>();

const echartsRef = useTemplateRef("echartsRef");
const echarts = shallowRef<Nullable<EchartsInstance>>(null);

useResizeObserver(echartsRef, onResize);

watch(() => options, value => {
    setOption(value);
}, {
    deep: true,
    flush: "post",
});

watch(() => renderer, init);

watch(() => [dark], () => {
    echarts.value?.setTheme(dark ? "dark" : "default");
});

onMounted(init);

onScopeDispose(() => {
    dispose();
});

async function init() {
    dispose();
    await nextTick();
    if(!echartsRef.value?.clientWidth || !echartsRef.value.clientHeight) return;
    if(beforeRender) beforeRender(Echarts);
    echarts.value = Echarts.init(echartsRef.value!, dark ? "dark" : "default", {
        renderer: renderer,
    });
    setOption(options);
    if(afterRender) {
        afterRender(echarts.value!, Echarts);
    }
}

function setOption(options?: EChartsOption) {
    echarts.value?.setOption({
        backgroundColor: dark ? undefined : "transparent",
        ...options,
    });
}

function onRefresh() {
    onClear();
    onResize();
    setOption(options);
}

function onResize() {
    echarts.value?.resize();
}

function onClear() {
    echarts.value?.clear();
}

function dispose() {
    if(echarts.value) {
        echarts.value.dispose();
    }
}

defineExpose({
    onRefresh,
    onResize,
    onClear,
    setOption,
    dispose,
    instance: echarts,
});
</script>

<template>
    <div class="w-[100%] h-[100%]">
        <div ref="echartsRef" class="w-[100%] h-[100%]" />
    </div>
</template>