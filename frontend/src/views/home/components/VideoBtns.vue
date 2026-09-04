<script setup lang="ts">
import { Clipboard } from "@wailsio/runtime";
import type { ParseOption } from "@/utils/parse";

defineOptions({
    name: "VideoButtons",
});

const {
    type,
} = defineProps<{
    type?: string;
}>();

const emit = defineEmits<{
    import: [data: ParseOption[]];
}>();

const ImportDialog = defineAsyncComponent(() => import("@/components/Parser"));

const importDialogRef = useTemplateRef("importDialogRef");

const message = useMessage();
const msgbox = useMessageBox();
const router = useRouter();
const store = useParserStore();

function onSourceClick() {
    router.push(`/sub/source?type=${type}`);
}

function onSuccess(data: ParseOption[]) {
    emit("import", data);
}

function onExport() {
    msgbox.confirm(`确认要导出 "json" 到剪切板吗？`).then(() => {
        let str = "";
        if(type === "video") {
            str = store.videoSource?.source || "";
        } else if(type === "anime") {
            str = store.animeSource?.source || "";
        }
        if(str) {
            Clipboard.SetText(str);
            message.success("导出成功！");
        }
    });
}
</script>

<template>
    <w-sider-actions>
        <el-button type="primary" icon="source-manage" title="源" @click="onSourceClick" />
        <el-button type="primary" icon="source-import" title="导入" @click="importDialogRef?.open()" />
        <el-button type="primary" icon="source-export" title="导出" @click="onExport" />
    </w-sider-actions>
    <import-dialog ref="importDialogRef" @success="onSuccess" />
</template>