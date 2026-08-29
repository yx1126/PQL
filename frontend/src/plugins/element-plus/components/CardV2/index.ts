import { withInstall } from "@/utils/withInstall";

export default withInstall(app => {
    app.component("ElCardV2", defineAsyncComponent(() => import("./CardV2.vue")));
});