import Douyu from "@/assets/live/douyu.png";
import Huya from "@/assets/live/huya.png";
import Douyin from "@/assets/live/douyin.png";
import type { MenuSiderItem } from "@/components/WUI";

export const liveList = [
    { type: "1", label: "斗鱼", icon: Douyu, path: "/home/live/douyu?type=1" },
    { type: "2", label: "虎牙", icon: Huya, path: "/home/live/huya?type=2" },
    { type: "3", label: "抖音", icon: Douyin, path: "/home/live/douyin?type=3" },
];

export const menuList: MenuSiderItem[] = [
    { id: "0", label: "关注", icon: "heart", size: 18, path: "/home/live" },
    ...liveList.map(v => {
        return {
            id: v.type,
            label: v.label,
            icon: v.icon,
            type: "img",
            size: 18,
            path: v.path,
        } satisfies MenuSiderItem;
    }),
    // { id: 1, label: "赛事", icon: "competition", path: "/home/live/competition" },
    // { id: 1, label: "搜索", icon: "ele-Search", path: "/home/live/search" },
];