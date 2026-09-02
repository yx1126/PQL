<script setup lang="ts">
import { DriveService } from "@bind/service";
import { Browser, Clipboard } from "@wailsio/runtime";

defineOptions({
    name: "ImportDialog",
});

const emit = defineEmits<{
    success: [];
}>();

const message = useMessage();

const { data: deviceData, query } = useService({
    request: DriveService.StartBaiduAuth,
    default: {
        device_code: "",
        user_code: "",
        verification_url: "",
        qrcode_url: "",
        expires_in: 0,
        interval: 0,
    },
    immediate: false,
    isLayoutLoad: true,
});

const tabActive = ref("code");
const visible = ref(false);

const isFailure = computed(() => deviceData.value.expires_in <= 0);

const validtoken = useInterval(async () => {
    await DriveService.GetBaiduToken(deviceData.value.device_code);
    message.success("授权成功！");
    visible.value = false;
    emit("success");
});

const countdown = useInterval(() => {
    deviceData.value.expires_in--;
    if(deviceData.value.expires_in <= 0) {
        onClosed();
    }
});

async function open() {
    tabActive.value = "code";
    visible.value = true;
    await query();
    onStartTimer();
}

function onStartTimer() {
    countdown.start(1000);
    validtoken.start((deviceData.value.interval + 1) * 1000);
}

function onClosed() {
    countdown.stop();
    validtoken.stop();
}

function onOpenUrl() {
    Browser.OpenURL(deviceData.value.verification_url);
}

async function onCopy() {
    await Clipboard.SetText(deviceData.value.user_code);
    message.success("复制成功！");
}

async function onRefresh() {
    await query();
    onStartTimer();
}

function onCancel() {
    visible.value = false;
}

defineExpose({
    open,
});
</script>

<template>
    <el-dialog
        v-model="visible"
        width="500"
        align-center
        append-to-body
        :close-on-click-modal="false"
        :draggable="false"
        destroy-on-close
        :show-close="false"
        header-class="hidden"
        @closed="onClosed"
    >
        <el-tabs v-model="tabActive" type="card" stretch>
            <el-tab-pane label="授权码登录" name="code">
                <div class="auth-content">
                    <el-form class="w-full" :model="deviceData" label-position="top" label-suffix="：">
                        <el-form-item label="授权地址">
                            <el-input :model-value="deviceData.verification_url" size="large" readonly>
                                <template #append>
                                    <el-button @click="onOpenUrl">打开</el-button>
                                </template>
                            </el-input>
                        </el-form-item>
                        <el-form-item label="用户Code">
                            <el-input :model-value="deviceData.user_code" size="large" readonly>
                                <template #append>
                                    <el-button @click="onCopy">复制</el-button>
                                </template>
                            </el-input>
                        </el-form-item>
                    </el-form>
                    <div v-if="isFailure" class="is-failure">已失效</div>
                </div>
            </el-tab-pane>
            <el-tab-pane label="二维码登录" name="qrcode">
                <div class="auth-content">
                    <w-image v-if="deviceData.qrcode_url" class="size-[300px]" :src="deviceData.qrcode_url" />
                    <div v-if="isFailure" class="is-failure">已失效</div>
                </div>
            </el-tab-pane>
        </el-tabs>
        <template #footer>
            <div class="flex justify-between items-center">
                <el-link type="primary">{{ deviceData.expires_in }}s</el-link>
                <div>
                    <el-button type="primary" :disabled="!isFailure" @click="onRefresh">刷新</el-button>
                    <el-button @click="onCancel">取消</el-button>
                </div>
            </div>
        </template>
    </el-dialog>
</template>

<style lang="scss" scoped>
.auth-content {
    width: 100%;
    height: 330px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: var(--w-layout-space-large);
    padding: var(--w-layout-space-large);
    position: relative;
    .is-failure {
        width: 100%;
        height: 100%;
        position: absolute;
        z-index: 2;
        background-color: rgba($color: #000000, $alpha: 0.5);
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--el-color-primary);
        font-size: 22px;
    }
}
</style>