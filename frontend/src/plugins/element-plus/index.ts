import { withInstall } from "@/utils/withInstall";
import { makeInstaller, ElButtonGroup, type ConfigProviderContext as $ConfigProviderContext } from "element-plus";
// @ts-ignore
import BaseComponents from "element-plus/es/component";
// @ts-ignore
import Plugin from "element-plus/es/plugin";
import ElButton from "./components/ElButton";
import ElLink from "./components/ElLink";
import ElTag from "./components/ElTag";
import ElList from "./components/List";
import ElCardV2 from "./components/CardV2";
import ElThing from "./components/Thing";
import { DelayKey } from "./keys";
import type { App } from "vue";

const Exclude: Record<`El${string}`, Component> = {
    ElButton,
    ElLink,
    ElButtonGroup,
    ElTag,
};

const Components = BaseComponents.filter((comp: any) => !Object.keys(Exclude).includes((comp).name));

const ExtraComponents = withInstall(app => {
    // rewrite
    for(const name in Exclude) {
        app.component(name, Exclude[name as keyof typeof Exclude]);
    }
    // custom
    app.use(ElList);
    app.use(ElThing);
    app.use(ElCardV2);
});

export interface ConfigProviderContext extends $ConfigProviderContext {
    delay?: number;
}

export default {
    install: (app: App, options?: ConfigProviderContext) => {
        const ElementPlus = makeInstaller([...Components, ExtraComponents, ...Plugin]);
        // ElButton、ElLink `click` 事件节流配置
        app.provide(DelayKey, options?.delay || 500);
        app.use(ElementPlus, options);
    },
};