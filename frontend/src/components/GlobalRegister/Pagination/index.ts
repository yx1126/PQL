import { withInstall } from "@/utils/withInstall";

export default withInstall(app => {
    app.component("Pagination", defineAsyncComponent(() => import("./Pagination.vue")));
});