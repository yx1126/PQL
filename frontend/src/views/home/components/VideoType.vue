<script setup lang="ts" generic="T extends object">
import jp, { type Parse, type RuleItem } from "@/utils/parse";

defineOptions({
    name: "VideoType",
});

const {
    form = {},
} = defineProps<{
    data?: T[];
    config?: RuleItem;
    form?: Parse.Data;
}>();

const emit = defineEmits<{
    "item-click": [item: unknown, primary?: string];
}>();

function onClick(item: unknown, primary?: string) {
    emit("item-click", item, primary);
}
</script>

<template>
    <div class="flex flex-col gap-2">
        <template
            v-for="item, index in data"
            :key="jp.value(item, config?.primaryPath) || index"
        >
            <el-divider v-if="index !== 0" class="!m-0" />
            <div class="flex flex-wrap gap-[8px_2px] leading-none">
                <template
                    v-for="sub, i in jp.value(item, config?.childrenPath)"
                    :key="jp.value(sub, config?.childPrimaryPath) || i"
                >
                    <w-button
                        plain
                        :type="form[jp.value(item, config?.primaryPath)] === jp.value(sub, config?.childPrimaryPath) ? 'primary' : ''"
                        @click="onClick(sub, jp.value(item, config?.primaryPath))"
                    >
                        {{ jp.value(sub, config?.childNamePath) }}
                    </w-button>
                </template>
            </div>
        </template>
    </div>
</template>
