import { withInstall } from "@/utils/withInstall";

export interface ElListContextType {
    bordered?: boolean;
    hoverable?: boolean;
    clickable?: boolean;
    showDivider?: boolean;
}

export const ElListContext = Symbol() as InjectionKey<ElListContextType>;

/**
 * @see https://www.naiveui.com/zh-CN/light/components/list
 */

export default withInstall(app => {
    app.component("ElList", defineAsyncComponent(() => import("./List.vue")));
    app.component("ElListItem", defineAsyncComponent(() => import("./ListItem.vue")));
});