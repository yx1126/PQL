<script setup lang="ts">
import LazyInput from "@/components/LazyInput";
import type { PopoverProps } from "./";
import IconList from "./IconList.vue";
import { Search as ElSearchIcon, CircleClose } from "@element-plus/icons-vue";
import { EleIconNames, LocalIconNames } from "@/components/GlobalRegister/Icon";

defineOptions({
    name: "IconSelect",
    inheritAttrs: false,
});

const modelValue = defineModel<string>();

const {
    clearable,
    placement = "prefix",
    popoverProps,
} = defineProps<{
    placeholder?: string;
    closeable?: boolean;
    placement?: "prefix" | "suffix";
    disabled?: boolean;
    clearable?: boolean;
    popoverProps?: PopoverProps;
}>();

type IconType = "element" | "local";

const { disabled: formDisabled, validate } = useElForm();

const popoverRef = useTemplateRef("popoverRef");
const eleIconRef = useTemplateRef("eleIconRef");
const localIconRef = useTemplateRef("localIconRef");
const { t } = useLocal();

const hovering = ref(false);
const searchKey = ref("");
const iconType = ref<IconType>("element");

const options = computed(() => {
    let iconType: IconType = "element";
    const eleIcons = search(searchKey.value, EleIconNames);
    const localIcons = search(searchKey.value, LocalIconNames);
    if(eleIcons.length > 0) {
        iconType = "element";
    } else if(localIcons.length > 0) {
        iconType = "local";
    }
    return {
        iconType,
        eleIcons,
        localIcons,
    };
});

const showClear = computed(() => clearable && modelValue.value && !formDisabled.value && hovering.value);

const popverAttrs = computed<PopoverProps>(() => {
    return {
        trigger: "click",
        placement: "bottom-start",
        showArrow: false,
        offset: 6,
        hideAfter: 0,
        ...popoverProps,
        width: "378px",
    };
});

watch(modelValue, (val, oldVal) => {
    if(val !== oldVal) {
        validate();
    }
});

watch(() => options.value.iconType, value => {
    iconType.value = value;
});

function search(value?: string, list?: string[]) {
    return list?.filter(n => value ? n?.toLowerCase()?.includes(value?.toLowerCase()) : n) || [];
}

function onChange(icon: string) {
    modelValue.value = icon;
    popoverRef.value?.hide();
}

function onMouseenter() {
    hovering.value = true;
}
function onMouseleave() {
    hovering.value = false;
}
function onClear() {
    modelValue.value = undefined;
}
function onBeforeEnter() {
    nextTick(() => {
        iconType.value = getIconType();
        eleIconRef.value?.reset();
        localIconRef.value?.reset();
    });
}
function onAfterLeave() {
    searchKey.value = "";
}

const ICON_MAP: ([IconType, string[]])[] = [
    ["element", EleIconNames],
    ["local", LocalIconNames],
];

function getIconType(): IconType {
    const icon = modelValue.value || "";
    for(let i = 0; i < ICON_MAP.length; i++) {
        const [type, list] = ICON_MAP[i];
        if(list.includes(icon)) {
            return type;
        }
    }
    return "element";
}
</script>

<template>
    <el-popover
        ref="popoverRef"
        v-bind="popverAttrs"
        :persistent="false"
        @before-enter="onBeforeEnter"
        @after-leave="onAfterLeave"
    >
        <template #reference>
            <el-input
                v-model="modelValue"
                class="icon-select-input"
                readonly
                :placeholder="placeholder || t('placeholder')"
                v-bind="$attrs"
                @mouseenter="onMouseenter"
                @mouseleave="onMouseleave"
            >
                <template v-if="placement === 'prefix'" #prepend>
                    <Icon :icon="modelValue" size="18" />
                </template>
                <template #suffix>
                    <Icon v-if="showClear" :icon="CircleClose" cursor @click.stop="onClear" />
                </template>
                <template v-if="placement === 'suffix'" #append>
                    <Icon :icon="modelValue" size="18" />
                </template>
            </el-input>
        </template>
        <div class="icon-select-wrapper">
            <div class="search">
                <lazy-input v-model.lazy="searchKey" :prefix-icon="ElSearchIcon" :placeholder="t('input')" clearable />
            </div>
            <div class="tabs">
                <el-tabs v-model="iconType">
                    <el-tab-pane label="element" name="element" lazy>
                        <icon-list ref="eleIconRef" :options="options.eleIcons" :value="modelValue" @change="onChange" />
                    </el-tab-pane>
                    <el-tab-pane label="local" name="local" lazy>
                        <icon-list ref="localIconRef" :options="options.localIcons" :value="modelValue" @change="onChange" />
                    </el-tab-pane>
                </el-tabs>
            </div>
        </div>
    </el-popover>
</template>

<style lang="scss" scoped>
.icon-select-input {
    :deep(.el-input-group__prepend) {
        padding: 0 12px;
    }
    :deep(.el-input-group__append) {
        padding: 0 12px;
    }
}
.icon-select-wrapper {
    width: 100%;
}
</style>

<i18n lang="yaml">
zh:
  placeholder: 请选择图标
  input: 请输入图标名称
en:
  placeholder: Please select icon
  input: Please enter icon name
</i18n>