<script setup lang="ts">
import "mui-player/dist/mui-player.min.css";
// @ts-ignore
import MuiPlayer from "mui-player";
// @ts-ignore
import MuiPlayerDesktopPlugin from "mui-player-desktop-plugin";

defineOptions({
    name: "Player",
});

const {
    src,
} = defineProps<{
    src: string;
}>();

const playerRef = useTemplateRef("playerRef");

const set = useSetStore();

const muiPlayer = shallowRef<MuiPlayer>();

watch(() => src, () => {
    muiPlayer.value?.reloadUrl(src);
});

onMounted(async () => {
    await nextTick();
    init();
});

onBeforeUnmount(destroy);

function destroy() {
    muiPlayer.value?.destroy();
}

function init() {
    if(!playerRef.value) return;
    muiPlayer.value = new MuiPlayer({
        container: playerRef.value,
        src,
        width: "100%",
        height: "100%",
        lang: "zh-cn",
        themeColor: set.primaryColor,
        custom: {
            headControls: [],
            rightSidebar: [
                {
                    slot: "HD", // 对应定义的 slot 值
                    width: "200px", // 侧栏宽度，string | number
                },
            ],
        },
        plugins: [
            new MuiPlayerDesktopPlugin({
                leaveHiddenControls: true,
            }),
        ],
    });
}

defineExpose({
    destroy,
});
</script>

<template>
    <div class="player-wrapper">
        <div ref="playerRef" class="player">
        </div>
    </div>
</template>

<style lang="scss" scoped>
.player {
    width: 100%;
    height: 100% !important;
    &-wrapper {
        width: 100%;
        height: 100%;
    }
}
</style>