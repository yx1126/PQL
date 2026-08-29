<script setup lang="ts">
import { decimalDiv, decimalMul } from "@/utils/decimal";
import { mapList, pointIcon, getDistance, getCenterPoint, type Location, type PointType } from "@/utils/maps";
import Img from "@/utils/image";
import Loading from "@/components/Loading";
import pointSvg from "@/assets/icon/map.svg?raw";
import { debounce } from "lodash-es";

defineOptions({
    name: "MapRender",
});

const { map } = defineProps<{
    map?: string;
}>();

export interface GridOptions {
    w: number;
    h: number;
    size: number;
    save?: boolean;
    before?: (i: number) => void;
    after?: (i: number) => void;
    showEdge?: boolean;
}

export type ScaleType = "plus" | "minus";

// 格子大小（你可以改 5 / 10 / 20 都行）
const GRID_SIZE = 10;

const canvasRef = useTemplateRef("canvasRef");
const canvasCtx = shallowRef<CanvasRenderingContext2D | null>(null);

const mapStore = useMapStore();

const loading = ref(false);

const state = reactive({
    step: 0.05,
    scale: 1,
    transformX: 0,
    transformY: 0,
    minScale: 0.05,
    maxScale: 2,
    points: [] as Location[],
});

watch(() => map, async () => {
    try {
        loading.value = true;
        state.points = [];
        onWatcherCleanup(() => {
            canvasRef.value?.removeEventListener("mousedown", onMousedown);
            canvasRef.value?.removeEventListener("wheel", onWheel);
            canvasRef.value?.removeEventListener("contextmenu", onContextmenu);
        });
        const mapItem = mapList.find(v => v.label === map);
        if(mapItem) {
            const pointsMap = mapItem.points || {};
            const types = Object.keys(pointsMap);
            mapStore.pointTypeList = mapStore.pointTypeList.filter(v => types.includes(v.type)).map(v => {
                return {
                    ...v,
                    data: pointsMap[v.type] || [],
                };
            });
        }
        await nextTick();
        await init();
    } catch (e) {
        console.error(e);
    } finally {
        loading.value = false;
    }
}, {
    immediate: true,
});

const onResize = debounce(() => {
    updateCanvas();
    render();
}, 100);

onBeforeMount(() => {
    window.addEventListener("resize", onResize);
});

onBeforeUnmount(() => {
    window.removeEventListener("resize", onResize);
});

function updateCanvas() {
    const wrapper = canvasRef.value?.parentNode as HTMLDivElement;
    canvasRef.value!.width = wrapper.clientWidth;
    canvasRef.value!.height = wrapper.clientHeight;
}

async function init() {
    updateCanvas();
    if(!canvasRef.value) return;
    const ctx = canvasRef.value!.getContext("2d")!;
    canvasCtx.value = ctx;
    canvasRef.value.addEventListener("mousedown", onMousedown);
    canvasRef.value.addEventListener("wheel", onWheel);
    canvasRef.value.addEventListener("contextmenu", onContextmenu);
    await render();
    renderCenter();
}

function clearPoint() {
    if(state.points.length <= 0) return;
    state.points = [];
    render();
}

async function render() {
    clearRect();
    renderBackLine();
    const { scale, transformX: tx, transformY: ty } = state;
    canvasCtx.value?.setTransform(scale, 0, 0, scale, tx, ty);
    await renderMap();
}

function onContextmenu(e: PointerEvent) {
    e.preventDefault();
}

function onMousedown(e: MouseEvent) {
    if(e.button === 0) { // 鼠标左键
        onMouseLeftEvent(e);
    } else if(e.button === 2) { // 鼠标右键
        onMouseRightEvent(e);
    }
}

function onWheel(e: WheelEvent) {
    e.preventDefault();
    onScale(e.deltaY > 0 ? "minus" : "plus", { x: e.offsetX, y: e.offsetY });
}

function onMouseLeftEvent(dEvent: MouseEvent) {
    const point = transformPoint(dEvent.offsetX, dEvent.offsetY);
    if(import.meta.env.DEV) {
        navigator.clipboard.writeText(JSON.stringify(point));
    }
    const canvas = canvasRef.value;
    const ctx = canvasCtx.value;
    if(!canvas || !ctx) return;
    // x 轴距离 y 轴距离
    const { e: offsetX, f: offsetY } = ctx.getTransform();
    const onMousemove = (event: MouseEvent) => {
        state.transformX = offsetX + (event.offsetX - dEvent.offsetX);
        state.transformY = offsetY + (event.offsetY - dEvent.offsetY);
        render();
    };
    const onMouseup = () => {
        canvas.removeEventListener("mousemove", onMousemove);
        canvas.removeEventListener("mouseup", onMouseup);
        canvas.removeEventListener("mouseleave", onMouseup);
    };
    canvas.addEventListener("mousemove", onMousemove);
    canvas.addEventListener("mouseup", onMouseup);
    canvas.addEventListener("mouseleave", onMouseup);
}

function onMouseRightEvent(e: MouseEvent) {
    if(state.points.length >= 2) {
        state.points = [];
    }
    const point = transformPoint(e.offsetX, e.offsetY);
    state.points.push(point);
    render();
}

function clearRect() {
    const ctx = canvasCtx.value;
    if(!canvasRef.value || !ctx) return;
    const { width: w, height: h } = canvasRef.value;
    // 保存当前 transform
    ctx.save();
    // 重置矩阵
    ctx.setTransform(1, 0, 0, 1, 0, 0);
    ctx.clearRect(0, 0, w, h);
    // 恢复 transform
    ctx.restore();
}

async function renderMap() {
    const item = mapList.find(v => v.label === map);
    const ctx = canvasCtx.value;
    if(!ctx || !item) return;
    const mapImage = await Img.loadLocal(item.map);
    ctx.drawImage(mapImage, 0, 0);
    const size = mapImage.width / (item.size === "8x8" ? 80 : 40);
    const { width: w, height: h } = mapImage;
    renderGrid({
        w,
        h,
        size,
        before: i => {
            ctx.strokeStyle = i % 10 === 0 ? "#000" : "#fff";
            ctx.lineWidth = i % 10 === 0 ? 2 : 0.5;
        },
    });
    renderPoints();
}
function renderBackLine() {
    const ctx = canvasCtx.value;
    if(!ctx || !canvasRef.value) return;
    ctx.strokeStyle = "grey";
    ctx.lineWidth = 0.3;
    const { width: w, height: h } = canvasRef.value;
    renderGrid({ w, h, size: GRID_SIZE, save: true });
}

async function renderPoints() {
    const ctx = canvasCtx.value;
    if(!ctx) return;
    const ponit = await Img.loadSvg(pointSvg);
    const size = 30;
    const offsetX = size / 2 - 1;
    ctx.save();
    ctx.setTransform(1, 0, 0, 1, 0, 0);
    const { points, transformX: tx, transformY: ty, scale } = state;
    // 图标
    for(let i = 0; i < mapStore.pointTypeList.length; i++) {
        const item = mapStore.pointTypeList[i];
        const png = pointIcon[item.type].png;
        const icon = await Img.load(png);
        item.data.forEach(({ x, y }) => {
            // 计算正确的屏幕坐标
            const cx = tx + x * scale;
            const cy = ty + y * scale;
            ctx.drawImage(
                icon,
                cx - offsetX,
                cy - size,
                size,
                size,
            );
        });
    }
    points.forEach(({ x, y }, i) => {
        // 计算正确的屏幕坐标
        const cx = tx + x * scale;
        const cy = ty + y * scale;
        ctx.drawImage(
            ponit,
            cx - offsetX,
            cy - size,
            size,
            size,
        );
        if(i != 0) {
            ctx.strokeStyle = "#f4ea2a";
            ctx.lineWidth = 1.5;
            const { x: lx, y: ly } = points[i - 1];
            const lcx = tx + lx * scale;
            const lcy = ty + ly * scale;
            ctx.beginPath();
            ctx.moveTo(lcx, lcy);
            ctx.lineTo(cx, cy);
            ctx.stroke();
            // 真实距离
            const distance = getDistance({ x: x, y: y }, { x: lx, y: ly });
            // 中心点
            const cp = getCenterPoint({ x: cx, y: cy }, { x: lcx, y: lcy });
            ctx.font = "26px sans-serif";
            const { width, actualBoundingBoxAscent, actualBoundingBoxDescent } = ctx.measureText(String(distance));
            const height = actualBoundingBoxAscent + actualBoundingBoxDescent;
            // 文字背景
            ctx.fillStyle = "gray";
            ctx.shadowBlur = 20;
            ctx.shadowColor = "black";
            ctx.shadowBlur = 0;
            const padding = [15, 10];
            ctx.fillRect(cp.x - width / 2 - padding[0] / 2, cp.y - height / 2 - padding[1] / 2, width + padding[0], height + padding[1]);
            // 文字填充
            ctx.textAlign = "center";
            ctx.fillStyle = "#f4ea2a";
            ctx.fillText(String(distance), cp.x, cp.y + height / 2);
        }
    });
    ctx.restore();
}

async function renderPoint(type: PointType) {
    const index = mapStore.pointTypeList.findIndex(v => v.type === type);
    const item = mapList.find(v => v.label === map)?.points || {};
    if(index !== -1) {
        mapStore.pointTypeList.splice(index, 1);
    } else {
        mapStore.pointTypeList.push({
            type,
            data: item[type] || [],
        });
    }
    render();
}

function renderCenter() {
    const item = mapList.find(v => v.label === map);
    if(!item) return;
    const image = Img.get(item.map)!;
    const { width: imageWidth, height: imageHeight } = image;
    const { width: canvasWidth, height: canvasHeight } = canvasRef.value!.getBoundingClientRect();
    const scaleX = canvasWidth / imageWidth;
    const scaleY = canvasHeight / imageHeight;
    const scale = Math.min(scaleX, scaleY);
    if(scale < 1) {
        state.scale = scale;
    }
    const scaledImageWidth = imageWidth * state.scale;
    const scaledImageHeight = imageHeight * state.scale;
    state.transformX = (canvasWidth - scaledImageWidth) / 2;
    state.transformY = (canvasHeight - scaledImageHeight) / 2;
    render();
}

function transformPoint(x: number, y: number) {
    return {
        x: (x - state.transformX) / state.scale,
        y: (y - state.transformY) / state.scale,
    };
}

function renderGrid({ w, h, size, save, before, after, showEdge }: GridOptions) {
    const ctx = canvasCtx.value;
    if(!ctx) return;
    const fn = () => {
        const wi = decimalDiv(w, size);
        // 画竖线
        for(let i = 0; i <= wi; i++) {
            if(!showEdge && (i === 0 || i === wi)) {
                continue;
            }
            before && before(i);
            ctx.beginPath();
            const x = decimalMul(i, size);
            ctx.moveTo(x, 0);
            ctx.lineTo(x, h);
            ctx.stroke();
            after && after(i);
        }
        const hi = decimalDiv(h, size);
        // 画横线
        for(let i = 0; i <= hi; i++) {
            if(!showEdge && (i === 0 || i === hi)) {
                continue;
            }
            before && before(i);
            ctx.beginPath();
            const y = decimalMul(i, size);
            ctx.moveTo(0, y);
            ctx.lineTo(w, y);
            ctx.stroke();
            after && after(i);
        }
    };
    if(save) {
        ctx.save();
        ctx.setTransform(1, 0, 0, 1, 0, 0);
        fn();
        ctx.restore();
    } else {
        fn();
    }
}

function onScaleChange(type: ScaleType) {
    const { clientWidth, clientHeight } = canvasRef.value!;
    onScale(type, { x: clientWidth / 2, y: clientHeight / 2 });
}

function onScale(type: ScaleType, point: Location) {
    const { scale: preScale, minScale, maxScale, step } = state;
    let newScale = preScale;
    switch(type) {
    case "minus":
        newScale = Math.max(minScale, preScale - step);
        break;
    case "plus":
        newScale = Math.min(maxScale, preScale + step);
        break;
    }
    const { e: eOffsetX, f: fOffsetY } = canvasCtx.value!.getTransform();
    state.transformX = point.x - ((point.x - eOffsetX) * newScale) / preScale;
    state.transformY = point.y - ((point.y - fOffsetY) * newScale) / preScale;
    state.scale = newScale;
    render();
}

defineExpose({
    clearPoint,
    renderPoint,
    restore: renderCenter,
    onScaleChange,
});
</script>

<template>
    <div class="map-render">
        <canvas ref="canvasRef" class="map-canvas" />
        <Loading v-if="loading" position="absolute" />
    </div>
</template>

<style lang="scss" scoped>
.map {
    &-render {
        width: 100%;
        height: 100%;
        position: relative;
    }
}
</style>