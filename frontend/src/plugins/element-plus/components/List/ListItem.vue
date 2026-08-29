<script setup lang="ts">
import { useNamespace } from "element-plus";
import { ElListContext } from ".";

defineOptions({
    name: "ElListItem",
});

const slots = defineSlots<Slotsable<"default" | "prefix" | "suffix">>();

const listContext = inject(ElListContext);

const ns = useNamespace("list-item");
</script>

<template>
    <div
        :class="[
            ns.b(),
            {
                [ns.m('bordered')]: listContext?.bordered,
                [ns.m('hoverable')]: listContext?.hoverable,
                [ns.m('clickable')]: listContext?.clickable,
                [ns.m('divider')]: listContext?.showDivider,
            }
        ]"
    >
        <div v-if="slots.prefix" :class="ns.e('prefix')">
            <slot name="prefix" />
        </div>
        <div :class="ns.e('main')">
            <slot />
        </div>
        <div v-if="slots.suffix" :class="ns.e('suffix')">
            <slot name="suffix" />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.el-list-item {
    position: relative;
    padding: 12px 0;
    display: flex;
    flex-wrap: nowrap;
    align-items: center;
    &__prefix {
        margin-right: 20px;
        flex: 0;
    }
    &__main {
        flex: 1;
    }
    &__suffix {
        margin-left: 20px;
        flex: 0;
    }
    &.el-list-item--bordered {
        padding: 12px 20px;
    }
    &.el-list-item--hoverable:hover {
        background-color: rgba(243, 243, 245);
        border-radius: 3px;
    }
    &.el-list-item--clickable {
        cursor: pointer;
    }
    &.el-list-item--divider {
        & + & {
            border-top: 1px solid var(--el-border-color);
        }
    }
}
</style>