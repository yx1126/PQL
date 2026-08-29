import { ElTag as ElTag } from "element-plus";

export default defineComponent({
    name: "ElTag",
    props: {
        disableTransitions: { type: Boolean, default: true },
    },
    emits: ["click"],
    setup(props, { emit, slots }) {
        return () => {
            const { disableTransitions } = props;
            return (
                <ElTag disableTransitions={disableTransitions} onClick={() => emit("click")}>{{ ...slots }}</ElTag>
            );
        };
    },
});