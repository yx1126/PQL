<script lang="ts" setup>
import Artplayer, { type Option } from "artplayer";
import type { Selector, Control } from ".";
import poster from "@/assets/video/poster.png";
import Flv from "flv.js";
import Hls from "hls.js";
import QualitySvg from "@/assets/icon/quality.svg?raw";
import CdnsSvg from "@/assets/icon/cdns.svg?raw";
import { hotKeyPlugin } from "./plugins";

const {
    src,
    isLive,
    type,
    cdns = [],
    quality = [],
} = defineProps<{
    src?: string;
    isLive?: boolean;
    type?: string;
    cdns?: Selector[];
    quality?: Selector[];
}>();

const emit = defineEmits<{
    "quality-change": [value: Selector["value"]];
    "cdn-change": [value: Selector["value"]];
}>();

const playerRef = useTemplateRef("playerRef");
const artPlayer = shallowRef<Nullable<Artplayer>>(null);

watch(() => src, () => {
    if(artPlayer.value && src) {
        artPlayer.value.switchUrl(src);
    }
});

watch(() => type, () => {
    destroy();
    init();
});

watch(() => cdns, v => {
    if(!artPlayer.value) return;
    artPlayer.value.controls.remove("cnds");
    if(v.length > 0) {
        artPlayer.value.controls.add(createCndsControl());
    }
}, {
    deep: true,
});

watch(() => quality, v => {
    if(!artPlayer.value) return;
    artPlayer.value.controls.remove("quality");
    if(v.length > 0) {
        artPlayer.value.controls.add(createQualityControl());
    }
}, {
    deep: true,
});

onMounted(init);

onBeforeUnmount(destroy);

function init() {
    const options: Option = {
        container: playerRef.value!,
        url: src || "",
        theme: "var(--el-color-primary)",
        poster,
        autoSize: false,
        autoMini: true,
        playbackRate: !isLive,
        aspectRatio: true,
        setting: true,
        pip: true,
        fullscreen: true,
        fullscreenWeb: true,
        miniProgressBar: true,
        lang: "zh-cn",
        lock: true,
        isLive: isLive,
        flip: true,
        autoplay: isLive,
        hotkey: !isLive,
        customType: {
            flv: playFlv,
            m3u8: playM3u8,
        },
        controls: [],
    };
    if(type) {
        options.type = type;
    }
    if(quality.length) {
        options.controls?.push(createQualityControl());
    }
    if(cdns.length) {
        options.controls?.push(createCndsControl());
    }
    const art = new Artplayer(options);
    art.contextmenu.remove("version");
    if(isLive) {
        art.plugins.add(hotKeyPlugin);
    }
    // 等待播放器DOM渲染完成
    art.on("ready", () => {

    });
    artPlayer.value = art;
}

function createCndsControl(): Control {
    const html = `<i class="art-icon art-icon-cdns">${CdnsSvg}</i>`;
    return {
        name: "cnds",
        index: 12,
        html,
        position: "right",
        selector: cdns,
        onSelect: item => {
            emit("cdn-change", item.value);
            return html;
        },
    };
}

function createQualityControl(): Control {
    const html = `<i class="art-icon art-icon-qualitys">${QualitySvg}</i>`;
    return {
        name: "quality",
        index: 13,
        html,
        position: "right",
        selector: quality,
        onSelect: item => {
            emit("quality-change", item.value);
            return html;
        },
    };
}

function playFlv(video: HTMLMediaElement, url: string, art: Artplayer) {
    if(Flv.isSupported()) {
        if(art.flv) {
            (art.flv as Flv.Player)?.destroy();
        }
        const flv = Flv.createPlayer({ type: "flv", url });
        flv.attachMediaElement(video);
        flv.load();
        art.flv = flv;
        art.on("destroy", () => {
            flv.unload();
            flv.destroy();
            art.flv = null;
        });
    } else {
        art.notice.show = "Unsupported playback format: flv";
    }
}

function playM3u8(video: HTMLMediaElement, url: string, art: Artplayer) {
    if(Hls.isSupported()) {
        if(art.hls) {
            (art.flv as Flv.Player).destroy();
        }
        const hls = new Hls();
        hls.loadSource(url);
        hls.attachMedia(video);
        art.hls = hls;
        art.on("destroy", () => {
            hls.destroy();
            art.hls = null;
        });
    } else if(video.canPlayType("application/vnd.apple.mpegurl")) {
        video.src = url;
    } else {
        art.notice.show = "Unsupported playback format: m3u8";
    }
}

function destroy() {
    artPlayer.value?.destroy(true);
}

defineExpose({
    destroy,
});
</script>

<template>
    <div class="player-wrapper">
        <div ref="playerRef" class="player" />
    </div>
</template>

<style lang="scss" scoped>
.player {
    width: 100%;
    height: 100%;
    &-wrapper {
        width: 100%;
        height: 100%;
        overflow: hidden;
    }
}
</style>