import { withInstall } from "@/utils/withInstall";
import Icon, { ElIcon } from "./Icon";
import Pagination from "./Pagination";

/**
 * 全局注册
 * 类型文件 types/vue.ts
 */
export default withInstall(app => {
    app.use(Icon);
    app.use(ElIcon);
    app.use(Pagination);
});