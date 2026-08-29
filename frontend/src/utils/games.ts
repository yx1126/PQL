import Graph from "@/assets/games/graph.png";
import Sevendays from "@/assets/games/7days.ico";
import Pubg from "@/assets/games/pubg.ico";

export const ICON_MAO: Record<string, string> = {
    DEFAULT: Graph,
    "7days": Sevendays,
    pubg: Pubg,
};

export function getGameIcon(type: string) {
    return ICON_MAO[type] || ICON_MAO.DEFAULT;
}