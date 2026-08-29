import { SvgIcon } from "@/components/GlobalRegister/Icon";

export function renderIcon(icon?: string, props?: Record<string, any>) {
    if(!icon) return;
    return defineComponent(_props => {
        if(icon?.toLowerCase().startsWith("ele")) {
            return () => h(resolveComponent(icon), { ...props, ..._props, icon });
        }
        return () => h(SvgIcon, { ...props, ..._props, icon });
    });
}