// import Player from "./Player.vue";
import Player from "./Artplayer.vue";
import type { Option } from "artplayer";

export type Selector = ArrayGet<Required<ArrayGet<Required<Option>["controls"]>>["selector"]>;

export type Control = ArrayGet<Required<Option>["controls"]>;

export default Player;