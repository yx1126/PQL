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
const scrollHeight = ref(0);

let resizeObserver: ResizeObserver | null = null;

onMounted(() => {
    onCheckOverflow();
    let lastWidth = 0;
    if(contentRef.value) {
        resizeObserver = new ResizeObserver(entries => {
            const width = entries[0]?.contentRect.width ?? 0;
            // 只处理宽度变化
            if(Math.abs(width - lastWidth) < 1) return;
            lastWidth = width;
            onCheckOverflow();
        });
        resizeObserver.observe(contentRef.value);
    }
});

onBeforeUnmount(() => {
    resizeObserver?.disconnect();
});

function onShowMore() {
    isExpanded.value = !isExpanded.value;
}

async function onCheckOverflow() {
    if(isExpanded.value) return;
    await nextTick();

    const el = contentRef.value;
    if(!el) return;
    const height = parseFloat(getComputedStyle(el).height);
    scrollHeight.value = el.scrollHeight;
    overflow.value = el.scrollHeight > height + 2;
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
            '--video-item-scroll-height': scrollHeight + 'px',
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
        transition: height 0.2s var(--w-trans);
    }
    &__more {
        align-self: flex-start;
    }
    @include when(expand) {
        .video-item__body {
            height: var(--video-item-scroll-height);
        }
    }
}
</style>