<script setup lang="ts">
import { LiveService } from "@bind/service";
import { liveList } from "./data";
import type { CreateLiveVo } from "@bind/vo";
import type { FormRules } from "element-plus";

defineOptions({
    name: "LiveDialog",
});

const emit = defineEmits<{
    success: [data: CreateLiveVo];
}>();

const formRef = useTemplateRef("formRef");
const message = useMessage();

const visible = ref(false);
const loading = ref(false);

const form = ref<CreateLiveVo>({
    roomId: "" as `${number}`,
    type: "1",
    isSpecial: 0,
});

const rules: FormRules<keyof CreateLiveVo> = {
    roomId: [{ required: true, message: "房间号不能为空", trigger: "blur" }],
    type: [{ required: true, message: "类型不能为空", trigger: "change" }],
};

function open(type = "1") {
    form.value.isSpecial = 0;
    form.value.type = type;
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
                emit("success", form.value);
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
        title="添加房间"
        width="500"
        align-center
        append-to-body
        :close-on-click-modal="false"
        :draggable="false"
        destroy-on-close
        @closed="onClosed"
    >
        <template #footer>
            <el-form ref="formRef" size="large" :model="form" :rules="rules" label-position="top" label-width="80px" label-suffix="：">
                <el-form-item prop="roomId" label="房间号">
                    <el-input v-model="form.roomId" type="textarea" :rows="3" resize="none" placeholder="请输入房间号" />
                </el-form-item>
                <el-form-item prop="type" label="类型">
                    <el-radio-group v-model="form.type" class="is-vertical">
                        <template v-for="item in liveList" :key="item.type">
                            <el-radio :value="item.type" border>
                                <w-text type="img" :icon="item.icon" size="24" gap="5">{{ item.label }}</w-text>
                            </el-radio>
                        </template>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <div class="flex justify-between">
                <el-checkbox v-model="form.isSpecial" :true-value="1" :false-value="0">是否关注</el-checkbox>
                <span>
                    <el-button @click="onCancel">取消</el-button>
                    <el-button type="primary" :loading @click="onSubmit">确认</el-button>
                </span>
            </div>
        </template>
    </el-dialog>
</template>
