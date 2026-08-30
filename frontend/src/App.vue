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
    onWatchSystemTheme();
    onWatch();
});

onBeforeUnmount(() => {
    Events.OffAll();
});

function onWatch() {
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
    // 设置变化
    if(set.windownName !== import.meta.env.VITE_APP_MAIN_NAME) {
        Events.On(WailsEvents.AppSetChange, () => set.load());
    }
    // 系统主题变化
    Events.On(WailsEvents.AppTheme, async ({ data }) => {
        switch(data.type) {
        case "system":
            onWatchSystemTheme();
            break;
        default:
            // TODO BUG
            set.theme = data.theme;
            onWatchSystemTheme(() => {
                setTheme(data.theme);
            });
        }
    });
}

async function onWatchSystemTheme(fn?: () => void) {
    if(set.theme === 2) {
        const dark = await AppService.GetDarkMode();
        setTheme(dark);
        set.updateTheme(dark);
        return;
    };
    fn && fn();
}
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