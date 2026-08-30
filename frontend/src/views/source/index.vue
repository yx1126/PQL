<script setup lang="ts">
import { ParserService } from "@bind/service";
import { Browser } from "@wailsio/runtime";

defineOptions({
    name: "Source",
});

const route = useRoute();
const router = useRouter();
const msgbox = useMessageBox();
const store = useParserStore();

const source = ref(route.query.type as string);

const data = computed(() => {
    if(!route.query.type) return store.parserList;
    return store.parserList.filter(v => v.type === route.query.type);
});

function onDelete(row: any) {
    msgbox.confirm("确认要删除源吗？").then(async () => {
        await ParserService.DeleteParser([row.id as unknown as number]);
        store.load();
    });
}

function onLink(link: string) {
    if(!link.startsWith("http")) {
        link = `https://${link}`;
    }
    Browser.OpenURL(link);
}

function onChange() {
    router.replace("/sub/source?type=" + source.value || "");
}
</script>

<template>
    <div class="source">
        <portal to="layout-extra">
            <el-select
                v-model="source"
                class="w-[150px]"
                placeholder="请选择源类型"
                clearable
                :empty-values="['','']"
                value-on-clear=""
                @change="onChange"
            >
                <el-option label="视频" value="video" />
                <el-option label="动漫" value="anime" />
            </el-select>
        </portal>
        <div class="w-box source-table">
            <el-table :data="data" height="100%" size="large">
                <el-table-column type="index" label="序号" width="80" align="center" />
                <el-table-column label="图标" prop="icon" width="100" align="center">
                    <template #default="{row}">
                        <w-image v-if="row.icon" class="size-[50px]" :src="row.icon" />
                    </template>
                </el-table-column>
                <el-table-column label="类型" prop="type" width="80" align="center" />
                <el-table-column label="子类型" prop="subType" width="100" align="center" />
                <el-table-column label="首页" prop="homePage" align="center">
                    <template #default="{row}">
                        <el-link v-if="row.homePage" underline="hover" @click="onLink(row.homePage)">{{ row.homePage }}</el-link>
                    </template>
                </el-table-column>
                <el-table-column label="版本" prop="version" width="80" align="center" />
                <el-table-column label="作者" prop="author" align="center" />
                <el-table-column label="描述" prop="description" align="center" show-overflow-tooltip />
                <el-table-column label="操作" align="center" width="80">
                    <template #default="{row}">
                        <div class="w-table-actions">
                            <!-- <el-link type="primary" icon="ele-View" @click="onDelete(row)" /> -->
                            <el-link type="danger" icon="ele-Delete" @click="onDelete(row)" />
                        </div>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.source {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    gap: 15px;
    &-search {
        padding: 15px;
    }
    &-table {
        flex: 1;
        display: flex;
        flex-direction: column;
        overflow: hidden;
    }
    .code {
        width: 100%;
        max-height: 200px;
        overflow: auto;
        white-space: pre-wrap;
        @include hidden-scroll;
    }
}
</style>