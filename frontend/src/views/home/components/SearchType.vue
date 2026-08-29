<script setup lang="ts" generic="T extends object">
import jp, { type Parse, type RuleItem } from "@/utils/parse";

defineOptions({
    name: "SearchType",
});

const {
    form = {},
} = defineProps<{
    data?: T[];
    config?: RuleItem;
    form?: Parse.Data;
}>();

const emit = defineEmits<{
    "item-click": [item: T];
}>();

function onClick(item: T) {
    emit("item-click", item);
}
</script>

<template>
    <div class="flex gap-[8px_2px] leading-none">
        <template
            v-for="item, i in data"
            :key="jp.value(item, config?.primaryPath) || i"
        >
            <w-button
                :type="form.id === jp.value(item, config?.primaryPath) ? 'primary' : ''"
                @click="onClick(item)"
            >
                {{ jp.value(item, config?.namePath) }}
            </w-button>
        </template>
    </div>
</template>
