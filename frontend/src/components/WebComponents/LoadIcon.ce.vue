<script lang="tsx">
export default defineComponent({
    name: "LoadIcon",
    render() {
        return (
            <div class="load-icon">
                <div class="icon-wrapper">
                    <i></i>
                    <i></i>
                    <i></i>
                    <i></i>
                </div>
            </div>
        );
    },
});
</script>

<style lang="scss" scoped>
.load-icon {
    width: 100vw;
    height: 100vh;
    position: fixed;
    left: 0;
    top: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    & > div {
        display: grid;
        grid-template-columns: repeat(2, 15px);
        grid-template-rows: repeat(2, 15px);
        gap: 10px;
        animation: rotate 1s linear infinite;
        & > i {
            background-color: var(--el-color-primary);
            animation: rotate2 1s linear infinite;
            &:nth-child(1) {
                background-color: var(--el-color-primary-light-7);
            }
            &:nth-child(2) {
                background-color: var(--el-color-primary-light-5);
            }
            &:nth-child(4) {
                background-color: var(--el-color-primary-light-3);
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
}
</style>