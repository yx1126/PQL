<script setup lang="ts">
import { parseUnit } from "@/utils/unit";

defineOptions({
    name: "VideoItem",
});

defineProps<{
    height: Unit;
}>();

const contentRef = useTemplateRef("contentRef");

const overflow = ref(false);
const isExpanded = ref(false);

useWindowResize(onCheckOverflow);

onMounted(onCheckOverflow);

onActivated(onCheckOverflow);

function onShowMore() {
    isExpanded.value = !isExpanded.value;
}

async function onCheckOverflow() {
    await nextTick();
    const el = contentRef.value;
    if(!el) return;
    let sum = 0;
    let isShow = false;
    const w = parseFloat(getComputedStyle(el).width);
    for(let index = 0; index < el.childNodes.length; index++) {
        const child = el.childNodes[index];
        if(child instanceof Element) {
            sum += parseFloat(getComputedStyle(child)?.width);
            if(sum > w) {
                isShow = true;
                break;
            }
        } else {
            continue;
        }
    }
    overflow.value = isShow;
}
</script>

<template>
    <div
        class="video-item"
        :class="{
            'is-expand': isExpanded
        }"
        :style="{
            '--video-item-height': parseUnit(height),
        }"
    >
        <div ref="contentRef" class="video-item__body">
            <slot />
        </div>
        <el-link
            v-if="overflow"
            class="video-item__more"
            type="primary"
            :delay="0"
            @click="onShowMore"
        >
            {{ isExpanded ? "收起" : "展开" }}
            <Icon :icon="isExpanded ? 'ele-ArrowUp' : 'ele-ArrowDown'" />
        </el-link>
    </div>
</template>

<style lang="scss" scoped>
.video-item {
    width: 100%;
    display: flex;
    gap: var(--w-layout-space);
    &__body {
        flex: 1;
        min-width: 0;
        display: flex;
        flex-wrap: wrap;
        height: var(--video-item-height);
        gap: 8px 2px;
        overflow: hidden;
    }
    &__more {
        align-self: flex-start;
    }
    @include when(expand) {
        .video-item__body {
            height: max-content;
        }
    }
}
</style>