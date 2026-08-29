<script setup lang="ts">
import SearchType from "../../components/SearchType.vue";
import VideoList from "../../components/VideoList.vue";
import { useVideoSearch } from "../../hooks/useVideoSearch.ts";

defineOptions({
    name: "VideoSearch",
});

const {
    inputValue,
    typeList,
    dataList,
    form,
    total,
    sTypes,
    sDatas,
    maxSize,
    onRefresh,
    onSearch,
    onGetData,
    onTypeClick,
    onVideoClick,
} = useVideoSearch("video");
</script>

<template>
    <div id="videoSearchTarget" class="videosearch">
        <div class="w-box pos-sticky top-0 z-1">
            <div class="videosearch-search">
                <el-input v-model="inputValue" class="input" size="large" clearable placeholder="请输入关键词" @keydown.enter="onSearch">
                    <template #append>
                        <el-button icon="ele-Search" @click="onSearch" />
                    </template>
                </el-input>
            </div>
            <search-type
                v-if="sTypes.isShow"
                class="mt-[10px]"
                :data="typeList"
                :form
                :config="sTypes"
                @item-click="onTypeClick"
            />
        </div>
        <video-list
            v-model:page="form.page"
            v-model:size="form.size"
            class="w-box"
            :total="total"
            :max-size="maxSize"
            :paging-type="sDatas.pagingType"
            :data="dataList"
            :config="sDatas"
            :get-img-uri="sDatas.getImgUri"
            @current-change="onGetData"
            @size-change="onGetData"
            @item-click="onVideoClick"
        />
        <w-backtop
            target="#videoSearchTarget"
            :right="20"
            :bottom="20"
            @refresh="onRefresh"
        />
    </div>
</template>

<style lang="scss" scoped>
.videosearch {
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
    & > * {
        padding: var(--w-layout-space-large);
    }
}
</style>