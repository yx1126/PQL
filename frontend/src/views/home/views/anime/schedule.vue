<script setup lang="ts">
import { createParse } from "@/utils/parse";
import { useState } from "../../hooks/useState.ts";
import VideoList from "../../components/VideoList.vue";
import dayjs from "dayjs";
import isoWeek from "dayjs/plugin/isoWeek";

defineOptions({
    name: "AnimeSchedule",
});

dayjs.extend(isoWeek);

interface Weekly {
    id: number;
    num: number;
    date: string;
    name: string;
    upper: string;
    lower: string;
    upperShort: string;
    lowerShort: string;
    japan: string;
    day: number;
}

const Week = {
    name: ["周一", "周二", "周三", "周四", "周五", "周六", "周日"],
    upper: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
    lower: ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
    upperShort: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
    lowerShort: ["mon", "tue", "wed", "thu", "fri", "sat", "sun"],
    japan: ["月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日", "日曜日"],
};

const today = dayjs().format("MM-DD");
const weekList = onGetWeekList();

const segmentedList = [
    { label: "周", value: "cn" },
    { label: "曜", value: "japan" },
    { label: "英", value: "en" },
];

const state = useStatesStore();
const set = useSetStore();
const { dataList, store, open } = useState("anime");

const tabActive = ref<number>(dayjs().day());

const schedule = computed(() => {
    const { config, getImgUri, createRequest } = createParse(store.value, "scheduleApi");
    return {
        ...config,
        getImgUri,
        request: createRequest(),
    };
});

onBeforeMount(() => {
    onGetDataList();
});

function onGetWeekList() {
    const weekStart = dayjs().startOf("isoWeek");
    const weekList: Weekly[] = [];
    for(let i = 0; i < 7; i++) {
        const current = weekStart.add(i, "day");
        weekList.push({
            id: i,
            num: i + 1,
            day: current.day(),
            date: current.format("MM-DD"),
            japan: Week.japan[i],
            name: Week.name[i],
            upper: Week.upper[i],
            lower: Week.lower[i],
            upperShort: Week.upperShort[i],
            lowerShort: Week.lowerShort[i],
        });
    }
    return weekList;
}

async function onGetDataList() {
    try {
        state.setLoad();
        const item = weekList.find(v => v.day === tabActive.value);
        if(!item) return;
        const res = await schedule.value.request(item);
        dataList.value = res.data;
    } finally {
        state.setLoad(false);
    }
}

function onWeekClick(item: Weekly) {
    tabActive.value = item.day;
    onGetDataList();
}

function onVideoClick(item: any) {
    const { namePath, next } = schedule.value;
    open(item, namePath, next);
}
</script>

<template>
    <div class="schedule">
        <div class="week w-box p-[15px]">
            <div class="week-item week-tools">
                <div></div>
                <el-segmented v-model="set.animeWeeklyType" :options="segmentedList" />
            </div>
            <template v-for="w in weekList" :key="w.id">
                <div
                    class="week-item"
                    :class="{
                        'is-active': tabActive === w.day,
                        'is-today': today === w.date
                    }"
                    @click="onWeekClick(w)"
                >
                    <div v-if="set.animeWeeklyType === 'japan'">{{ w.japan }}</div>
                    <div v-else-if="set.animeWeeklyType === 'en'">{{ w.upperShort }}</div>
                    <div v-else>{{ w.name }}</div>
                    <div class="text-[14px]">{{ w.date }}</div>
                </div>
            </template>
        </div>
        <div id="scheduleListTarget" class="schedule-list w-box">
            <video-list
                :paging="false"
                :data="dataList"
                :config="schedule"
                :get-img-uri="schedule.getImgUri"
                @item-click="onVideoClick"
            />
            <w-backtop
                target="#scheduleListTarget"
                :right="20"
                :bottom="20"
                @refresh="onGetDataList"
            />
        </div>
    </div>
</template>

<style lang="scss" scoped>
.schedule {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 15px;
    .week {
        display: grid;
        justify-content: space-between;
        grid-template-columns: repeat(7, 1fr);
        gap: 10px;
        &-tools {
            grid-column: span 7 / span 7;
            display: flex;
            align-items: center;
            justify-content: space-between;
        }
        &-item:not(.week-tools) {
            flex: 1;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            line-height: 1;
            gap: 10px;
            border: 1px solid var(--w-border-color);
            padding: 25px 10px;
            border-radius: var(--w-border-radius);
            cursor: pointer;
            position: relative;
            transition:
                color 0.3s,
                border-color 0.3s,
                background-color 0.3s;
            &::after {
                content: "";
                display: inline-block;
                width: 60%;
                height: 4px;
                background-color: transparent;
                position: absolute;
                bottom: 0;
                left: 50%;
                transform: translateX(-50%);
                transition: background-color 0.3s;
                border-top-left-radius: var(--w-border-radius);
                border-top-right-radius: var(--w-border-radius);
            }
            @include when-hover(active) {
                color: var(--el-color-primary);
                border-color: var(--el-color-primary-light-7);
                background-color: var(--el-color-primary-light-9);
                &::after {
                    background-color: var(--el-color-primary);
                }
            }
            @include when(today) {
                &::before {
                    content: "Today";
                    display: inline-block;
                    position: absolute;
                    top: 0;
                    right: 5%;
                    font-size: 12px;
                    color: #fff;
                    background-color: var(--el-color-primary);
                    letter-spacing: 0.3px;
                    padding: 5px;
                    border-bottom-left-radius: var(--w-border-radius);
                    border-bottom-right-radius: var(--w-border-radius);
                    box-shadow:
                        0 6px 18px rgba(255, 138, 0, 0.28),
                        inset 0 1px 0 rgba(255, 255, 255, 0.25);
                }
            }
        }
    }
    &-list {
        flex: 1;
        padding: 15px;
        overflow-x: hidden;
        overflow-y: auto;
        @include hidden-scroll;
    }
}
</style>