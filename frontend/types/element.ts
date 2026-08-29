import { dropdownItemProps } from "element-plus";
import type { ExtractPublicPropTypes } from "vue";

export interface SelectOption {
    label?: string | number;
    value: string | number | boolean | Record<string, any>;
    disabled?: boolean;
    [key: string]: any;
}

export type DropdownItemOptions = Omit<Writable<ExtractPublicPropTypes<typeof dropdownItemProps>>, "textValue">;