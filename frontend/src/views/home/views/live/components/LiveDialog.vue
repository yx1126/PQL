<script setup lang="ts">
import { LiveService } from "@bind/service";
import Douyu from "@/assets/live/douyu.png";
import Huya from "@/assets/live/huya.png";
import Douyin from "@/assets/live/douyin.png";
import type { CreateLiveVo } from "@bind/vo";
import type { FormRules } from "element-plus";

defineOptions({
    name: "LiveDialog",
});

const emit = defineEmits<{
    success: [];
}>();

const formRef = useTemplateRef("formRef");
const message = useMessage();

const visible = ref(false);
const loading = ref(false);

const form = ref<CreateLiveVo>({
    roomId: "" as `${number}`,
    type: "1",
});

const rules: FormRules<keyof CreateLiveVo> = {
    roomId: [{ required: true, message: "房间号不能为空", trigger: "blur" }],
    type: [{ required: true, message: "类型不能为空", trigger: "change" }],
};

function open() {
    visible.value = true;
}

function onCancel() {
    visible.value = false;
}

function onClosed() {
    formRef.value?.resetFields();
}

function onSubmit() {
    formRef.value?.validate(async valid => {
        if(valid) {
            try {
                loading.value = true;
                await LiveService.CreateLive(form.value);
                emit("success");
                onCancel();
            } catch (error) {
                message.error((error as any)?.message || "未知错误！");
            } finally {
                loading.value = false;
            }
        }
    });
}

defineExpose({
    open,
});
</script>

<template>
    <el-dialog
        v-model="visible"
        title="添加直播间"
        width="500"
        align-center
        append-to-body
        :close-on-click-modal="false"
        :draggable="false"
        destroy-on-close
        @closed="onClosed"
    >
        <template #footer>
            <el-form ref="formRef" size="large" :model="form" :rules="rules" label-width="80px" label-suffix="：">
                <el-form-item prop="roomId" label="房间号">
                    <el-input v-model="form.roomId" placeholder="请输入房间号" />
                </el-form-item>
                <el-form-item prop="type" label="类型">
                    <el-radio-group v-model="form.type" class="is-vertical">
                        <el-radio value="1" border>
                            <w-image class="size-6" :src="Douyu" />
                            <span>斗鱼</span>
                        </el-radio>
                        <el-radio value="2" border>
                            <w-image class="size-6" :src="Huya" />
                            <span>虎牙</span>
                        </el-radio>
                        <el-radio value="3" border>
                            <w-image class="size-6" :src="Douyin" />
                            <span>抖音</span>
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <div class="dialog-footer">
                <el-button @click="onCancel">取消</el-button>
                <el-button type="primary" :loading @click="onSubmit">确认</el-button>
            </div>
        </template>
    </el-dialog>
</template>
