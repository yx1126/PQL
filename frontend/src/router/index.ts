import { createRouter } from "vue-router";
import { createWebHashHistory } from "vue-router";
import { staticRoutes } from "./staticRoutes";

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    strict: true,
    routes: [
        ...staticRoutes,
    ],
});

router.beforeEach(() => {
});

router.afterEach(() => {
});

router.onError(error => {
    console.error("路由错误", error);
});

export default router;