import type { Menu } from "@/layout/components/WMenu.vue";
import type { PointType, Location } from "@/utils/maps";
import type { ParserVo, UpdateSettingVo } from "@bind/vo";

export type RouterTrans = "fade" | "scale" | "scale-slide" | "null";

export type Lang = "zh-cn" | "en";

export interface MenuState {
    menuList: Menu[];
    keepMap: Record<string, string>;
}

export interface ParserState {
    parserList: ParserVo[];
}

export interface SetState extends PartialId<UpdateSettingVo> {
    windownName: string;
}

export interface MapState {
    pointTypeList: PointTypeItem[];
}

// eslint-disable-next-line @typescript-eslint/no-empty-object-type
export interface StatesState {
}

export interface PointTypeItem {
    type: PointType;
    data: Location[];
}