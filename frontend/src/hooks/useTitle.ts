import { isStr } from "@/utils/validata";
import type { MaybeRef } from "vue";

export function useTitle(title?: MaybeRef<string>, template = "%s") {
    const value = ref(title ?? document.title ?? null);
    watch(value, (v, ov) => {
        if(isStr(v) && v !== ov && document) {
            document.title = template.replaceAll("%s", v).replace(/^[-]+|[-]+$/g, "");
        }
    }, {
        immediate: true,
    });
    return value;
}