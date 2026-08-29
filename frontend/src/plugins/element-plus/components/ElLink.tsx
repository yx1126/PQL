import { renderIcon } from "../renderIcon";
import type { Component, PropType } from "vue";
import { ElLink as ELink } from "element-plus";
import { throttle } from "lodash-es";
import { DelayKey } from "../keys";

export default defineComponent({
    name: "ElLink",
    props: {
        icon: [String, Object] as PropType<string | Component>,
        delay: { type: Number },
    },
    emits: ["click"],
    setup(props, { slots, emit }) {
        const delay = inject(DelayKey, 500);

        const onClick = throttle(() => {
            emit("click");
        }, props.delay ?? delay, { trailing: false });
        return () => {
            const { icon } = props;
            return (
                <ELink icon={renderIcon(icon)} onClick={onClick}>{{ ...slots }}</ELink>
            );
        };
    },
});