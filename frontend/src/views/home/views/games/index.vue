<script setup lang="ts">
import { GameService } from "@bind/service";
import { windowOpen } from "@/utils/window";
import { getGameIcon } from "@/utils/games";
import type { GameVo } from "@bind/vo";

defineOptions({
    name: "Games",
});

interface ToolItem {
    name: string;
    path: string;
    icon: string;
}

const router = useRouter();

const { data: gameList } = useService(GameService.GetGameList, []);

const tabActive = ref(0);
const toolList: ToolItem[] = [
    { name: "下载管理", icon: "ele-Download", path: "/sub/download" },
];

async function onClick(item: GameVo, i: number) {
    if(item.isSupportOpenWindow === 1 && item.path) {
        windowOpen({
            title: item.name,
            name: item.name,
            path: item.path,
        });
    } else {
        tabActive.value = i;
    }
}

async function onToolClick(item: ToolItem) {
    router.push(item.path);
}
</script>

<template>
    <div class="games">
        <div class="games__list">
            <div class="games__list-body">
                <el-scrollbar>
                    <template v-for="item, i in gameList" :key="i">
                        <div
                            class="games__list-item"
                            :class="{ 'is-active': i === tabActive }"
                            @click="onClick(item, i)"
                        >
                            <div class="games__list-icon">
                                <w-image class="size-full" :src="getGameIcon(item.type)" alt="" />
                            </div>
                            <span>{{ item.name }}</span>
                        </div>
                    </template>
                </el-scrollbar>
            </div>
            <div class="games__list-footer">
                <template v-for="tool in toolList" :key="tool.path">
                    <div class="tools-item" @click="onToolClick(tool)">
                        <Icon :icon="tool.icon" />
                        <span>{{ tool.name }}</span>
                    </div>
                </template>
            </div>
        </div>
        <div class="games__main">
            <div class="games__main-body">
                <el-empty class="size-full" />
            </div>
            <div class="games__main-footer">
                <el-button class="start" type="primary">开始游戏</el-button>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.games {
    width: 100%;
    height: 100%;
    color: #fff;
    display: flex;
    &__list {
        width: var(--w-sider-width);
        height: 100%;
        border-right: 1px solid var(--w-border-color);
        padding: 15px 10px 15px 15px;
        display: flex;
        gap: 15px;
        flex-direction: column;
        background-color: #fcfcfd;
        animation-name: slide-left;
        animation-duration: 0.3s;
        animation-fill-mode: forwards;
        @include when-dark {
            background-color: transparent;
        }
        &-icon {
            width: 50px;
            min-width: 50px;
            height: 50px;
            padding: 2px;
            background-color: var(--w-box-bg);
            border: 1px solid var(--w-border-color);
            transition: border-color 0.3s var(--w-trans);
            border-radius: 8px;
            overflow: hidden;
        }
        &-item {
            width: 100%;
            height: 70px;
            padding: 10px;
            display: flex;
            align-items: center;
            gap: 12px;
            background-color: var(--w-box-bg);
            border: 1px solid var(--w-border-color);
            border-radius: var(--w-border-radius);
            color: var(--w-text-color);
            position: relative;
            transition:
                border-color 0.3s var(--w-trans),
                color 0.3s var(--w-trans);
            cursor: pointer;
            overflow: hidden;
            @include when-hover(active) {
                border-color: var(--el-color-primary);
                @include when-dark {
                    border-color: var(--el-color-primary-light-5);
                }
                color: var(--el-color-primary);
                .games__list-icon {
                    border-color: var(--el-color-primary);
                    @include when-dark {
                        border-color: var(--el-color-primary-light-5);
                    }
                }
                &::before {
                    transition: background-color 0.3s var(--w-trans);
                    background-color: var(--el-color-primary);
                }
                &::after {
                    background:
                        linear-gradient(
                            to right,
                            color-mix(in oklab, var(--el-color-primary) 10%, transparent),
                            transparent 10px
                        ),
                        linear-gradient(
                            to left,
                            color-mix(in oklab, var(--el-color-primary) 10%, transparent),
                            transparent 10px
                        ),
                        linear-gradient(
                            to bottom,
                            color-mix(in oklab, var(--el-color-primary) 10%, transparent),
                            transparent 10px
                        ),
                        linear-gradient(
                            to top,
                            color-mix(in oklab, var(--el-color-primary) 10%, transparent),
                            transparent 10px
                        );
                }
            }
            &::before {
                content: "";
                position: absolute;
                top: 50%;
                left: 0;
                transform: translateY(-50%);
                width: 2px;
                height: 15px;
                background: transparent;
            }
            &::after {
                content: "";
                width: 100%;
                height: 100%;
                position: absolute;
                top: 0;
                left: 0;
                pointer-events: none;
                background: transparent;
            }
            & + & {
                margin-top: 8px;
            }
        }
        &-body {
            flex: 1;
        }
        &-footer {
            width: 100%;
            background-color: var(--w-box-bg);
            border-radius: var(--w-border-radius);
            overflow: hidden;
            box-shadow: var(--el-box-shadow-lighter);
            @include when-dark {
                box-shadow: none;
            }
            .tools-item {
                display: flex;
                align-items: center;
                gap: 10px;
                padding: 10px 15px;
                cursor: pointer;
                line-height: 1;
                color: var(--w-text-color);
                .el-icon {
                    font-size: 18px;
                }
                &:hover {
                    color: var(--el-color-primary);
                    background-color: var(--el-color-primary-light-7);
                }
            }
        }
    }
    &__main {
        width: calc(100% - var(--w-sider-width));
        height: 100%;
        display: flex;
        flex-direction: column;
        animation-name: slide-right;
        animation-duration: 0.3s;
        animation-fill-mode: forwards;
        &-body {
            width: 100%;
            flex: 1;
        }
        &-footer {
            width: 100%;
            height: 80px;
            padding: 10px;
            display: flex;
            justify-content: flex-end;
            background-color: var(--w-box-bg);
            .start {
                width: 200px;
                height: 60px;
                --el-font-size-base: 20px;
            }
        }
    }
}
</style>