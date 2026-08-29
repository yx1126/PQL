import type { Lang } from "#/stores";
import { useI18n } from "vue-i18n";

function setHtmlLang(value: Lang) {
    document.documentElement.setAttribute("lang", value);
}

interface LocalesOptions {
    immediate?: boolean;
}

export type I18nBack = ReturnType<typeof useI18n>;

interface LocalBack extends I18nBack {
    $t: I18nBack["t"];
    ti: I18nBack["t"];
    ts: I18nBack["t"];
    tv: I18nBack["t"];
    $ti: I18nBack["t"];
    $ts: I18nBack["t"];
    $tv: I18nBack["t"];
}

export function useLocal(...args: Parameters<typeof useI18n>): LocalBack {
    const i18n = useI18n(...args);
    const { t } = useI18n({ useScope: "global" });

    /**
     * 输入框本地翻译 `请输入{0}`
     */
    function ti(...args: Parameters<I18nBack["t"]>) {
        return t("input", [i18n.t(...args)]);
    }
    /**
     * 输入框全局翻译 `请输入{0}`
     */
    function $ti(...args: Parameters<I18nBack["t"]>) {
        return t("input", [t(...args)]);
    }

    /**
     * 选择框本地翻译 `请选择{0}`
     */
    function ts(...args: Parameters<I18nBack["t"]>) {
        return t("select", [i18n.t(...args)]);
    }
    /**
     * 选择框全局翻译 `请选择{0}`
     */
    function $ts(...args: Parameters<I18nBack["t"]>) {
        return t("select", [t(...args)]);
    }

    /**
     * 表单本地翻译 `{0}不能为空`
     */
    function tv(...args: Parameters<I18nBack["t"]>) {
        return t("valid.not-null", [i18n.t(...args)]);
    }
    /**
     * 表单全局翻译 `{0}不能为空`
     */
    function $tv(...args: Parameters<I18nBack["t"]>) {
        return t("valid.not-null", [t(...args)]);
    }

    return {
        ...i18n as any,
        $t: t,
        ti,
        ts,
        tv,
        $ti,
        $ts,
        $tv,
    };
}

export function useLocales(options?: LocalesOptions) {
    const set = useSetStore();

    const { locale } = useI18n({
        useScope: "global",
    });

    const lang = computed<Lang>({
        get() {
            return set.lang as Lang;
        },
        set: setLang,
    });

    onBeforeMount(() => {
        if(options?.immediate) {
            setLang(lang.value);
        }
    });

    function setLang(value: Lang) {
        set.lang = value;
        setHtmlLang(value);
        locale.value = value;
    }

    return {
        lang,
        setLang,
    };
}