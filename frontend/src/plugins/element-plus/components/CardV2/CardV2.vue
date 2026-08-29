<script lang="ts" setup>
import { useNamespace } from "element-plus";
import type { StyleValue } from "vue";

defineOptions({
    name: "ElCardV2",
});

const {
    hover,
    always,
    border = false,
    padding = true,
} = defineProps<{
    header?: string;
    footer?: string;
    border?: boolean | "header" | "footer";
    padding?: boolean;
    bodyStyle?: StyleValue;
    bodyClass?: string;
    hover?: boolean;
    always?: boolean;
}>();

const slots = defineSlots<Slotsable<"default" | "header" | "footer" | "extra">>();

const ns = useNamespace("card");

const shadow = computed(() => always ? "always" : hover ? "hover" : "never");
</script>

<template>
    <div
        :class="[
            ns.b(),
            ns.b('v2'),
            ns.is(`${shadow}-shadow`),
            {
                [ns.is('border-header')]: [true, 'header'].includes(border),
                [ns.is('border-footer')]: [true, 'footer'].includes(border)
            }
        ]"
    >
        <div v-if="slots.header || header || slots.extra" :class="ns.e('header')">
            <div :class="ns.e('header-main')">
                <slot name="header">{{ header }}</slot>
            </div>
            <div v-if="slots.extra" :class="ns.e('extra')">
                <slot name="extra" />
            </div>
        </div>
        <div
            :class="[
                ns.e('body'),
                bodyClass,
                {
                    [ns.is('padding')]: padding,
                    [ns.is('header-padding')]: (slots.header || header || slots.extra) && [false, 'footer'].includes(border),
                    [ns.is('footer-padding')]: (slots.footer || footer) && [false, 'header'].includes(border)
                }
            ]" :style="bodyStyle"
        >
            <slot />
        </div>
        <div v-if="slots.footer || footer" :class="ns.e('footer')">
            <slot name="footer">{{ footer }}</slot>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.el-card {
    display: flex;
    flex-direction: column;
    &__header {
        border-bottom: none;
        display: flex;
        justify-content: space-between;
        align-items: center;
        column-gap: 10px;
        &-main {
            flex: 1;
        }
    }
    &__footer {
        border-top: none;
    }
    @include when(border-header) {
        .el-card__header {
            border-bottom: 1px solid var(--el-card-border-color);
        }
    }
    @include when(border-footer) {
        .el-card__footer {
            border-top: 1px solid var(--el-card-border-color);
        }
    }
    .el-card__body {
        flex: 1;
        padding: 0;
        @include when(padding) {
            padding: var(--el-card-padding);
            @include when(header-padding) {
                padding-top: 0;
            }
            @include when(footer-padding) {
                padding-bottom: 0;
            }
        }
    }
}
</style>