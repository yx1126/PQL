<script setup lang="ts">
import { langList } from "@/locales";
import type { VNode } from "vue";

const { type = "select" } = defineProps<{
    type?: "select" | "dropdown";
}>();

defineSlots<{
    default?: () => VNode[];
}>();
const { lang, setLang } = useLocales();
</script>

<template>
    <el-dropdown v-if="type === 'dropdown'" trigger="click" :persistent="false" @command="setLang">
        <slot />
        <template #dropdown>
            <el-dropdown-menu>
                <template v-for="item in langList" :key="item.value">
                    <el-dropdown-item :command="item.value">{{ item.label }}</el-dropdown-item>
                </template>
            </el-dropdown-menu>
        </template>
    </el-dropdown>
    <el-select v-else-if="type === 'select'" v-model="lang">
        <template v-for="item in langList" :key="item.value">
            <el-option :label="item.label" :value="item.value" />
        </template>
    </el-select>
</template>