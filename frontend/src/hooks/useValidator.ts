import type { FormItemRule } from "element-plus";

export type Triggers = "blur" | "change";

export function Require(message: FormItemRule["message"], trigger: Triggers = "blur", options?: FormItemRule): FormItemRule {
    const value: FormItemRule = {
        ...options,
        required: true,
        message,
        trigger,
    };
    return value;
}

export function RequiredValidator(validator: FormItemRule["validator"], trigger: Triggers = "blur", options?: FormItemRule): FormItemRule {
    return {
        ...options,
        required: true,
        validator,
        trigger,
    };
}

export function Validator(validator: FormItemRule["validator"], trigger: Triggers = "blur", options?: FormItemRule): FormItemRule {
    return {
        ...options,
        validator,
        trigger,
    };
}