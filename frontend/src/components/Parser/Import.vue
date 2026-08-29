<script setup lang="ts">
import { toString } from "@/utils/validata";
import { ParserService } from "@bind/service";
import type { ParseOption } from "@/utils/parse";

defineOptions({
    name: "ImportDialog",
});

const emit = defineEmits<{
    success: [data: ParseOption[]];
}>();

const message = useMessage();
const state = useLoading();

const visible = ref(false);
const inputValue = ref("");

const disabled = computed(() => inputValue.value.trimStart().trimEnd() === "");

function open() {
    visible.value = true;
}

function onCancel() {
    visible.value = false;
}

async function onSubmit() {
    try {
        state.setLoad(true);
        if(inputValue.value.trimStart().trimEnd() === "") {
            return;
        }
        const res = await ParserService.CreateParsers(inputValue.value);

        emit("success", res.data || []);
        message.success(res.message);
        onCancel();
        nextTick(() => {
            inputValue.value = "";
        });
    } catch (error) {
        message.error((error as Error)?.message || toString(error));
    } finally {
        state.setLoad(false);
    }
}

defineExpose({
    open,
});
</script>

<template>
    <el-dialog
        v-model="visible"
        title="导入源"
        width="500"
        align-center
        append-to-body
        :close-on-click-modal="false"
        :draggable="false"
        destroy-on-close
    >
        <el-input v-model="inputValue" class="input" resize="none" type="textarea" :rows="5" />
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="onCancel">取消</el-button>
                <el-button type="primary" :disabled @click="onSubmit">确认</el-button>
            </div>
        </template>
    </el-dialog>
</template>

<style lang="scss" scoped>
.input :deep(.el-textarea__inner) {
    overflow-y: auto;
    height: 250px;
    max-height: 250px;
    @include hidden-scroll;
}
</style>