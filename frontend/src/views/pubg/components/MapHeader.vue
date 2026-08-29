<script lang="ts" setup>
import { mapList } from "@/utils/maps";

const modelValue = defineModel<string>();

const emit = defineEmits<{
    "map-item": [name: string];
    back: [];
}>();

function onClick(name: string) {
    emit("map-item", name);
}
function onBack() {
    emit("back");
}
</script>

<template>
    <div class="navbar">
        <div class="navbar__menu">
            <div class="navbar__logo" @click="onBack">[PUBG] RangeFinder</div>
            <template
                v-for="map, i in mapList"
                :key="i"
            >
                <div
                    class="navbar__btn"
                    :class="{
                        'is-active': modelValue == map.label
                    }"

                    @click="onClick(map.label)"
                >
                    <span class="navbar_text">{{ map.name }}</span>
                </div>
            </template>
        </div>
        <div class="navbar__right">
        </div>
    </div>
</template>

<style lang="scss" scoped>
.navbar {
    height: 50px;
    background-color: var(--w-box-bg);
    display: flex;
    align-items: center;
    gap: 20px;
    justify-content: space-between;
    padding: 0 15px;
    box-shadow: var(--el-box-shadow);
    @include when-dark {
        box-shadow: 0px 4px 8px 0px rgb(0 0 0 / 65%);
    }
    z-index: 1;
    transition-duration: 250ms;
    &__menu {
        display: flex;
        gap: 30px;
    }
    &__logo {
        margin-right: 20px;
        cursor: pointer;
        font-size: large;
        font-weight: 600;
        display: inline-block;
        transition-duration: 250ms;
        &:hover {
            color: var(--el-color-primary);
            transform: scale(1.1);
        }
    }
    &_text {
        text-decoration: inherit;
        font-weight: inherit;
        display: inline-block;
    }
    &__btn {
        cursor: pointer;
        font-weight: 600;
        font-size: large;
        transition-duration: 150ms;
        &:hover {
            color: var(--el-color-primary);
            transform: scale(1.2);
        }
        @include when(active) {
            transform: scale(1.2);
            color: var(--el-color-primary);
        }
    }
    &__right {
        display: flex;
        align-items: center;
        a {
            margin-left: 40px;
            color: initial;
            background-color: initial;
            text-decoration: none;
        }
    }
}
</style>
