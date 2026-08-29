import { useFormDisabled, useFormItem } from "element-plus";

export default function useElForm() {
    const disabled = useFormDisabled();
    const { formItem } = useFormItem();

    function validate(type: "change" | "blur" = "change") {
        formItem?.validate(type).catch(error => {
            console.warn(error);
        });
    }

    return {
        disabled,
        formItem,
        validate,
    };
}