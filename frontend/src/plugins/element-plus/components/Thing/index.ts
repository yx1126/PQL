import { withInstall } from "@/utils/withInstall";

/**
 * @see https://www.naiveui.com/zh-CN/light/components/Thing
 */

export default withInstall(app => {
    app.component("ElThing", defineAsyncComponent(() => import("./Thing.vue")));
});