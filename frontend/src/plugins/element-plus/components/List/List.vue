<script setup lang="ts">
import { useNamespace } from "element-plus";
import { ElListContext } from ".";

defineOptions({
    name: "ElList",
});

const props = withDefaults(defineProps<{
    bordered?: boolean;
    hoverable?: boolean;
    clickable?: boolean;
    showDivider?: boolean;
}>(), {
    showDivider: true,
});

const slots = defineSlots<Slotsable<"default" | "header" | "footer">>();

provide(ElListContext, reactive({ ...toRefs(props) }));

const ns = useNamespace("list");
</script>

<template>
    <ul
        :class="[
            ns.b(),
            {
                [ns.m('bordered')]: bordered,
            }
        ]"
    >
        <div v-if="slots.header" :class="ns.e('header')">
            <slot name="header" />
        </div>
        <slot />
        <div v-if="slots.footer" :class="ns.e('footer')">
            <slot name="footer" />
        </div>
    </ul>
</template>

<style lang="scss" scoped>
.el-list {
    font-size: 14px;
    margin: 0;
    padding: 0;
    list-style-type: none;
    background-color: #fff;
    transition: var(--el-transition-duration);
    @include when-dark {
        background-color: var(--el-bg-color-overlay);
    }
    &__header {
        border-bottom: 1px solid var(--el-border-color);
        padding: 12px 0;
    }
    &__footer {
        padding: 12px 0;
        border-top: 1px solid var(--el-border-color);
    }
    &--bordered {
        border: 1px solid var(--el-border-color);
        border-radius: 3px;
        .el-list {
            &__header,
            &__footer {
                padding: 12px 20px;
            }
        }
    }
}
</style>