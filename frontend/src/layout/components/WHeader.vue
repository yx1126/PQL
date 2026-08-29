<script setup lang="ts">
import { Window, Events, Application } from "@wailsio/runtime";
import { windowClose } from "@/utils/window";
import WMenu from "./WMenu.vue";
import { SetService } from "@bind/service";

defineOptions({
    name: "WindowHeader",
});

export type EventType = "message" | "setting" | "always-top" | "line" | "maxmin" | "close";

const sourceTitle = import.meta.env.VITE_APP_TITLE_SHORT;
const mainName = import.meta.env.VITE_APP_MAIN_NAME;

const route = useRoute();
const router = useRouter();
const menuStore = useMenuStore();
const set = useSetStore();

const maximiseIcon = ref<"maximization" | "restore">("restore");

const isAlwaysTop = ref(false);

const title = computed(() => {
    const r = route.matched.findLast(v => v.meta.title);
    return r?.meta.title || sourceTitle;
});

const showMenu = computed(() => {
    return !!route.matched.findLast(v => v.meta.showMenu);
});

const tabActive = computed({
    get: () => {
        return route.meta.activePath || route.path;
    },
    set: path => {
        const menu = menuStore.menuList.find(v => v.path === path)!;
        const keepPath = menuStore.keepMap[menu.icon];
        if(keepPath) {
            router.push(keepPath);
            return;
        }
        router.push(path);
    },
});

onBeforeMount(async () => {
    SetService.GetMenuList().then(menus => {
        menuStore.menuList = menus || [];
    });
    Window.IsMaximised().then(isMaximised => {
        maximiseIcon.value = isMaximised ? "restore" : "maximization";
    });
    Events.On(WailsEvents.WindowMaximise, ({ data }) => {
        if(data.name === set.windownName) {
            maximiseIcon.value = "restore";
        }
    });
    Events.On(WailsEvents.WindowUnMaximise, ({ data }) => {
        if(data.name === set.windownName) {
            maximiseIcon.value = "maximization";
        }
    });
    Events.On(WailsEvents.WindowRestore, ({ data }) => {
        if(data.name === set.windownName) {
            maximiseIcon.value = "maximization";
        }
    });
});

onBeforeUnmount(() => {
    Events.Off("window:maximise", "window:restore", "window:unMaximise");
});

async function onClick(type: EventType) {
    switch(type) {
    case "message":
        router.push("/sub/message");
        break;
    case "setting":
        if(route.path === "/sub/setting") break;
        router.push("/sub/setting");
        break;
    case "always-top":
        isAlwaysTop.value = !isAlwaysTop.value;
        Window.SetAlwaysOnTop(isAlwaysTop.value);
        break;
    case "line":
        Window.Minimise();
        Window.Show;
        break;
    case "maxmin":
        Window.ToggleMaximise();
        break;
    case "close":
        {
            if(set.windownName === import.meta.env.VITE_APP_MAIN_NAME) {
                if(set.closeBehavior == 1) {
                    await Application.Quit();
                } else {
                    await Window.Hide();
                }
            } else {
                await windowClose(set.windownName);
            }
        }
        break;
    default:
        break;
    }
}
</script>

<template>
    <div class="w-header">
        <div class="w-header__left">
            <div class="logo">
                <Icon icon="logo" type="primary" />
                <span class="text-[var(--w-title-color)]">{{ title }}</span>
            </div>
            <transition name="slide-top">
                <w-menu v-if="showMenu" v-model="tabActive" class="w-header__menu" :data="menuStore.menuList" />
            </transition>
        </div>
        <div class="w-header__right">
            <template v-if="set.windownName === mainName">
                <div class="btn" @click="onClick('message')">
                    <Icon icon="ele-Bell" />
                </div>
                <div class="btn" @click="onClick('setting')">
                    <Icon icon="ele-Setting" />
                </div>
                <el-divider direction="vertical" />
            </template>
            <div class="btn" @click="onClick('always-top')">
                <Icon v-if="!isAlwaysTop" icon="pin" />
                <Icon v-else icon="pin-fill" rotate="-45" type="primary" />
            </div>
            <div class="btn" @click="onClick('line')">
                <Icon icon="line" />
            </div>
            <div class="btn" @click="onClick('maxmin')">
                <Icon :icon="maximiseIcon" size="16" />
            </div>
            <div class="btn is-danger" @click="onClick('close')">
                <Icon icon="close" />
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.w-header {
    --wails-draggable: drag;
    width: 100%;
    height: var(--w-header-height);
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 15px;
    // border-bottom: 1px solid var(--el-color-primary);
    border-bottom: 1px solid var(--w-border-color);
    & > div {
        height: 100%;
        display: flex;
        align-items: center;
    }
    &__left {
        gap: 50px;
    }
    &__right {
        justify-content: flex-end;
        gap: 15px;
    }
    .logo {
        display: flex;
        align-items: center;
        font-size: 20px;
        gap: 8px;
    }
    .btn {
        --wails-draggable: no-drag;
        cursor: pointer;
        display: inline-flex;
        justify-content: center;
        align-items: center;
        transition: all 150ms var(--w-trans);
        font-size: 20px;
        color: var(--w-text-color);
        position: relative;
        z-index: 99;
        &:hover {
            color: var(--el-color-primary);
        }
        @include when(danger) {
            &:hover {
                color: #e74c3c;
            }
        }
    }
}
</style>