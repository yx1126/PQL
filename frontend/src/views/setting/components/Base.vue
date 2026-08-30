<script setup lang="ts">
import { AppService } from "@bind/service";
import SetItem from "./SetItem.vue";
import { Events } from "@wailsio/runtime";

defineOptions({
    name: "SettingBase",
});

const predefineColors = [
    "#FF9F43",
    "#4D96FF",
    "#22C7D6",
    "#36C98F",
    "#8B6CF6",
    "#F472B6",
    "#F05D5E",
    "#667EEA",
    "#E6B84A",
    "#4FD1C5",
];

const set = useSetStore();

const loading = ref(false);
const isStartUp = ref(false);

const theme = computed({
    get: () => set.theme,
    set: v => {
        Events.Emit(WailsEvents.AppTheme, {
            type: "*",
            theme: v,
        });
    },
});

onBeforeMount(() => {
    getStartStatus();
});

async function onBeforeChange() {
    loading.value = true;
    try {
        if(isStartUp.value) {
            await AppService.DisableOnWindow();
        } else {
            await AppService.StartOnWindow();
        }
        return true;
    } catch (error) {
        console.error(error);
        return false;
    } finally {
        loading.value = false;
    }
}

async function getStartStatus() {
    const res = await AppService.AutoStartStatus();
    isStartUp.value = res.Enabled;
}
</script>

<template>
    <div class="">
        <set-item title="开机启动">
            <el-switch v-model="isStartUp" :loading :before-change="onBeforeChange" />
        </set-item>
        <el-divider />
        <set-item title="主题模式">
            <el-select v-model="theme" class="w-[150px]" :persistent="false">
                <el-option label="深色模式" :value="0" />
                <el-option label="浅色模式" :value="1" />
                <el-option label="跟随系统" :value="2" />
            </el-select>
        </set-item>
        <el-divider />
        <set-item title="主题颜色">
            <el-color-picker
                v-model="set.primaryColor"
                class="w-[150px]"
                show-alpha
                :predefine="predefineColors"
                :persistent="false"
            />
        </set-item>
        <el-divider />
        <set-item title="关闭按钮行为">
            <el-select v-model="set.closeBehavior" class="w-[150px]" :persistent="false">
                <el-option label="最小化到托盘" :value="0" />
                <el-option label="退出程序" :value="1" />
            </el-select>
        </set-item>
        <el-divider />
        <set-item title="默认语言">
            <el-select v-model="set.lang" class="w-[150px]" :persistent="false">
                <el-option label="中文" value="zh-cn" />
            </el-select>
        </set-item>
    </div>
</template>
