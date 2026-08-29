<script setup lang="ts">
import VideoType from "../../components/VideoType.vue";
import VideoList from "../../components/VideoList.vue";
import { useVideoList } from "../../hooks/useVideoList.ts";

defineOptions({
    name: "AnimeList",
});

const {
    typeList,
    dataList,
    form,
    total,
    types,
    datas,
    maxSize,
    onRefresh,
    getDataList,
    onTypeClick,
    onVideoClick,
} = useVideoList("anime");
</script>

<template>
    <div id="videoListTarget" class="videolist">
        <video-type
            v-if="types.isShow"
            class="videolist-item"
            :data="typeList"
            :form
            :config="types"
            @item-click="onTypeClick"
        />
        <video-list
            v-model:page="form.page"
            v-model:size="form.size"
            class="videolist-item"
            :total="total"
            :max-size="maxSize"
            :paging-type="datas.pagingType"
            :data="dataList"
            :config="datas"
            :get-img-uri="datas.getImgUri"
            @current-change="getDataList"
            @size-change="getDataList"
            @item-click="onVideoClick"
        />
        <w-backtop
            target="#videoListTarget"
            :right="20"
            :bottom="20"
            @refresh="onRefresh"
        />
    </div>
</template>

<style lang="scss" scoped>
.videolist {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: var(--w-layout-space);
    overflow-x: hidden;
    overflow-y: auto;
    @include hidden-scroll;
    &-item {
        background-color: var(--w-box-bg);
        border-radius: var(--w-border-radius);
        border: 1px solid var(--w-border-color);
        padding: var(--w-layout-space-large);
    }
}
</style>