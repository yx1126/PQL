<script lang="tsx">
import WImage from "./WImage.vue";
import { parseUnit } from "@/utils/unit";
import type { PropType, SlotsType, VNode } from "vue";

export default defineComponent({
    name: "WText",
    props: {
        type: { type: String as PropType<"icon" | "img">, default: "icon" },
        icon: { type: String },
        size: { type: [String, Number] as PropType<Unit>, default: 16 },
        gap: { type: [String, Number] as PropType<Unit> },
        reverse: { type: Boolean },
    },
    slots: Object as SlotsType<{
        default?: () => VNode[];
    }>,
    setup(props, { slots }) {
        function renderIcon() {
            const { type, icon: ic } = props;
            if(!ic) return null;
            switch(type) {
            case "icon":
                return <icon class="w-text__icon" icon={ic} />;
            case "img":
                return <WImage class="w-text__icon" src={ic} />;
            default:
                return null;
            }
        }

        return () => {
            const { size, gap, reverse } = props;
            const text = <span>{renderSlot(slots, "default")}</span>;
            return (
                <div
                    class="w-text"
                    style={{
                        "--w-text-size": parseUnit(size),
                        "--w-text-gap": parseUnit(gap),
                    }}
                >
                    {reverse ? text : null}
                    {renderIcon()}
                    {!reverse ? text : null}
                </div>
            );
        };
    },
});
</script>

<style lang="scss" scoped>
.w-text {
    display: inline-flex;
    align-items: center;
    line-height: 1;
    gap: var(--w-text-gap, var(--w-layout-space));
    &__icon {
        font-size: var(--w-text-size);
    }
    &__img {
        width: var(--w-text-size);
        min-width: var(--w-text-size);
        height: var(--w-text-size);
        min-height: var(--w-text-size);
    }
}
</style>