import { createApp } from "vue";
import App from "./App.vue";
import i18n, { loadI18n } from "@/locales";
import router from "@/router";
import pinia from "@/stores";
import Directive from "@/directive";
import ElementPlus from "@/plugins/element-plus";
import GlobalRegister from "@/components/GlobalRegister";
import "@/components/WebComponents";
import WUI from "@/components/WUI";
import "@/styles/index.scss";
import "virtual:svg-icons-load";
import "virtual:uno.css";

const app = createApp(App);

app.use(pinia)
    .use(router)
    .use(ElementPlus)
    .use(Directive)
    .use(i18n)
    .use(GlobalRegister)
    .use(WUI);

router.isReady().then(async () => {
    const set = useSetStore();
    await loadI18n();
    await useParserStore().load();
    await set.load();
    app.mount("#app");
});