import { renderIcon } from "../renderIcon";
import type { PropType } from "vue";
import { ElButton as EButton, type ButtonProps } from "element-plus";
import { throttle } from "lodash-es";
import { DelayKey } from "../keys";

export default defineComponent({
    name: "ElButton",
    props: {
        icon: [String, Object] as PropType<ButtonProps["icon"]>,
        delay: { type: Number },
    },
    emits: ["click"],
    setup(props, { slots, emit }) {
        const delay = inject(DelayKey, 500);

        const onClick = throttle(() => {
            emit("click");
        }, props.delay ?? delay, { trailing: false });
        return () => {
            return (
                <EButton icon={renderIcon(props.icon)} onClick={onClick}>{{ ...slots }}</EButton>
            );
        };
    },
});