<script setup lang="ts">
import { mapList, pointIcon, type PointType, type PointIcon } from "@/utils/maps.ts";
import MapRender, { type ScaleType } from "./MapRender.vue";

defineOptions({
    name: "MapV2",
});

const {
    mapName,
} = defineProps<{
    mapName?: string;
}>();

interface IconItem extends PointIcon {
    label: PointType;
    isActive?: boolean;
}

const mapRef = useTemplateRef("mapRef");

const map = useMapStore();
const sourceList = ref<IconItem[]>([]);

watch(() => mapName, () => {
    const points = mapList.find(v => v.label === mapName)?.points || {};
    sourceList.value = (Object.keys(points) as PointType[]).map(item => {
        return {
            label: item,
            ...pointIcon[item],
            isActive: false,
        };
    });
}, {
    immediate: true,
});

const iconList = computed(() => {
    return sourceList.value.map(v => {
        return {
            ...v,
            isActive: map.pointTypeList.findIndex(p => p.type === v.label) !== -1,
        };
    });
});

function onIconClick(item: IconItem) {
    item.isActive = !item.isActive;
    mapRef.value?.renderPoint(item.label);
}

function onClearPoint() {
    mapRef.value?.clearPoint();
}

function onRestore() {
    mapRef.value?.restore();
}

function onScaleChange(type: ScaleType) {
    mapRef.value?.onScaleChange(type);
}
</script>

<template>
    <div class="map-wrapper">
        <map-render ref="mapRef" :map="mapName" />
        <div class="map-tools">
            <template v-for="item in iconList" :key="item.label">
                <el-tooltip :content="item.name" placement="left">
                    <div
                        class="map-tools-item"
                        :class="{
                            'is-active': item.isActive
                        }"
                        @click="onIconClick(item)"
                    >
                        <Icon :icon="item.icon" />
                    </div>
                </el-tooltip>
            </template>
            <el-tooltip content="清空点线" placement="left">
                <div class="map-tools-item" @click="onClearPoint">
                    <Icon icon="point-remove" />
                </div>
            </el-tooltip>
            <el-tooltip content="复位" placement="left">
                <div class="map-tools-item" @click="onRestore">
                    <Icon icon="point-restore" />
                </div>
            </el-tooltip>
            <el-tooltip content="放大" placement="left">
                <div class="map-tools-item" @click="onScaleChange('plus')">
                    <Icon icon="ele-Plus" size="18" />
                </div>
            </el-tooltip>
            <el-tooltip content="缩小" placement="left">
                <div class="map-tools-item" @click="onScaleChange('minus')">
                    <Icon icon="ele-minus" size="18" />
                </div>
            </el-tooltip>
        </div>
    </div>
</template>

<style lang="scss" scoped>
@use "@/styles/global/mixins" as *;
.map {
    &-wrapper {
        height: 100%;
        overflow: hidden;
        position: relative;
    }
    &-tools {
        position: absolute;
        right: 10px;
        top: 50px;
        display: flex;
        flex-direction: column;
        gap: 10px;
        &-item {
            width: 40px;
            height: 40px;
            cursor: pointer;
            display: flex;
            justify-content: center;
            align-items: center;
            background-color: var(--w-box-bg);
            border-radius: var(--w-border-radius);
            transition-duration: 150ms;
            box-shadow: var(--el-box-shadow-dark);
            :deep(.el-color-picker) {
                width: 100%;
                height: 100%;
            }
            &:hover {
                color: var(--el-color-primary);
                background-color: var(--el-color-primary-light-9);
            }
            & > i {
                font-size: 28px;
            }
            @include when(active) {
                background-color: var(--el-color-primary-light-3);
            }
            @include when-dark {
                box-shadow: var(--el-box-shadow-dark-white);
            }
        }
    }
}
</style>