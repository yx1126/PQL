// eslint-disable-next-line @typescript-eslint/no-unused-vars
import type DPlayer from "dplayer";
import type { DPlayerVideo, DPlayerDanmaku } from "dplayer";

// 模块扩充，重写类方法签名
declare module "dplayer" {
    export default interface DPlayer {
    /** 重写switchVideo，弹幕参数可选 */
        switchVideo(video: DPlayerVideo, danmaku?: DPlayerDanmaku): void;
    }
}
