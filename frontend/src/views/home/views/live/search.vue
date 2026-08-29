<script setup lang="ts">
import Douyu from "@/assets/live/douyu.png";
import Huya from "@/assets/live/huya.png";
import SearchItem from "./components/SearchItem.vue";

defineOptions({
    name: "LiveSearch",
});

const douyuRef = useTemplateRef("douyuRef");
const huyaRef = useTemplateRef("huyaRef");

const keyword = ref("");

function search() {
    douyuRef.value?.search();
    huyaRef.value?.search();
}

function query() {
    douyuRef.value?.refresh();
    huyaRef.value?.refresh();
}
</script>

<template>
    <div id="liveSearchTarget" class="livesearch">
        <div class="w-box p-[var(--w-layout-space-large)] pos-sticky top-0 z-1">
            <div class="livesearch-search">
                <el-input
                    v-model="keyword"
                    class="input"
                    size="large"
                    clearable
                    placeholder="请输入关键词"
                    @keydown.enter="search"
                >
                    <template #append>
                        <el-button icon="ele-Search" @click="search" />
                    </template>
                </el-input>
            </div>
        </div>
        <el-tabs class="w-tabs" type="border-card">
            <el-tab-pane>
                <template #label>
                    <w-image class="size-[20px]" :src="Douyu" />
                    <span>斗鱼</span>
                </template>
                <search-item ref="douyuRef" type="1" :keyword />
            </el-tab-pane>
            <el-tab-pane label="Config">
                <template #label>
                    <w-image class="size-[20px]" :src="Huya" />
                    <span>虎牙</span>
                </template>
                <search-item ref="huyaRef" type="2" :keyword />
            </el-tab-pane>
        </el-tabs>
        <w-backtop
            target="#liveSearchTarget"
            :right="20"
            :bottom="20"
            @refresh="query"
        />
    </div>
</template>

<style lang="scss" scoped>
.livesearch {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: var(--w-layout-space);
    overflow-x: hidden;
    overflow-y: auto;
    @include hidden-scroll;
    &-search {
        display: flex;
        align-items: center;
        justify-content: center;
        .input {
            width: 500px;
        }
    }
}
</style>