<script lang="tsx">
import type { CSSProperties, PropType } from "vue";
export default defineComponent({
    name: "LoadIcon",
    props: {
        position: { type: String as PropType<CSSProperties["position"]>, default: "fixed" },
        color: { type: String as PropType<CSSProperties["backgroundColor"]>, default: "" },
        mask: { type: Boolean, default: false },
    },
    render() {
        const { position, color, mask } = this;
        return (
            <div
                class="load-icon"
                style={{
                    position,
                    "--bg-color": color ? color : mask ? "var(--el-overlay-color-lighter)" : "",
                }}
            >
                <div class="icon-wrapper">
                    <icon icon="logo" />
                    <icon icon="logo" />
                    <icon icon="logo" />
                    <icon icon="logo" />
                </div>
            </div>
        );
    },
});
</script>

<style lang="scss" scoped>
.load-icon {
    width: 100%;
    height: 100%;
    left: 0;
    top: 0;
    z-index: 99999;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: var(--bg-color);
    .icon-wrapper {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        grid-template-rows: repeat(2, 1fr);
        gap: 10px;
        animation: rotate 1s linear infinite;
        position: relative;
        z-index: 2;
        will-change: transform;
        & > i {
            color: var(--el-color-primary);
            animation: rotate2 1s linear infinite;
            font-size: 20px;
            will-change: transform;
            &:nth-child(1) {
                color: color-mix(in oklab, var(--el-color-primary) 20%, transparent);
            }
            &:nth-child(2) {
                color: color-mix(in oklab, var(--el-color-primary) 60%, transparent);
            }
            &:nth-child(4) {
                color: color-mix(in oklab, var(--el-color-primary) 80%, transparent);
            }
        }
    }

    @keyframes rotate {
        0% {
            transform: rotate(0);
        }
        50% {
            transform: rotate(180deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }

    @keyframes rotate2 {
        0% {
            transform: rotate(0);
        }
        50% {
            transform: rotate(-180deg);
        }
        100% {
            transform: rotate(-360deg);
        }
    }

    @keyframes l13 {
        100% {
            transform: rotate(1turn);
        }
    }
}
</style>