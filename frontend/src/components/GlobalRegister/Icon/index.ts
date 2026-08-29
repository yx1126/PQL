import * as Icons from "@element-plus/icons-vue";
import svgIconList from "virtual:svg-icons-list";
import { withInstall } from "@/utils/withInstall";
import Icon from "./Icon.vue";
import SvgIcon from "./SvgIcon.vue";

export {
    Icon,
    SvgIcon,
};
/**
 * element-plus icon names
 */
export const EleIconNames = Object.keys(Icons).map(v => `Ele${v}`);
export const LocalIconNames = svgIconList.map(v => v.name);

export const ElIcon = withInstall(app => {
    for(const [key, component] of Object.entries(Icons)) {
        app.component(`Ele${key}`, component);
    }
});

export default withInstall(Icon);