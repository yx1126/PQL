<script setup lang="ts">
import { LangMap } from "@/locales";
import { Events } from "@wailsio/runtime";
import Layout from "@/layout/index.vue";
import { setTheme } from "@/utils/theme";
import { AppService } from "@bind/service";

const route = useRoute();
const router = useRouter();
const set = useSetStore();

const { lang } = useLocales({
    immediate: true,
});

const local = computed(() => {
    return LangMap[lang.value];
});

onBeforeMount(() => {
    // 页面跳转
    if(set.windownName === import.meta.env.VITE_APP_MAIN_NAME) {
        Events.On(WailsEvents.PageChange, ({ data }) => {
            if(route.fullPath === data.path) return;
            switch(data.type) {
            case "push":
                router.push(data.path);
                break;
            case "replace":
                router.replace(data.path);
                break;
            default:
                break;
            }
        });
    }
    // 系统主题变化
    Events.On(WailsEvents.AppTheme, async ({ data }) => {
        set.theme = data.theme;
        switch(data.type) {
        case "system":
            if(set.theme === 2) {
                const dark = await AppService.GetDarkMode();
                setTheme(dark);
            };
            break;
        default:
            setTheme(data.theme);
            break;
        }
    });
});

onBeforeUnmount(() => {
    Events.OffAll();
});
</script>

<template>
    <el-config-provider
        :locale="local"
        :link="{
            underline: 'never'
        }"
        :dialog="{
            draggable: true
        }"
        :message="{
            grouping: true
        }"
    >
        <Layout />
    </el-config-provider>
</template>