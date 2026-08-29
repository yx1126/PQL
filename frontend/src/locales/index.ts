import { createI18n } from "vue-i18n";
import zhCn from "element-plus/es/locale/lang/zh-cn";
import en from "element-plus/es/locale/lang/en";
import type { Lang } from "#/stores";

const i18n = createI18n({
    fallbackLocale: "zh",
    legacy: false,
});

export const LangMap: Record<Lang, any> = {
    "zh-cn": zhCn,
    en: en,
};

export const langList = [
    { label: "简体中文", value: "zh" },
    { label: "English", value: "en" },
];

export async function loadI18n() {
    const zh = await import(`./langs/zh.json`);
    const en = await import(`./langs/en.json`);
    const messages = {
        zh: {
            ...zh.default,
        },
        en: {
            ...en.default,
        },
    };
    Object.entries(messages).forEach(([key, message]) => {
        i18n.global.setLocaleMessage(key, message);
    });
}

export default i18n;