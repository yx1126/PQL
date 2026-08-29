<script lang="tsx">
import { useNamespace } from "element-plus";
import type { CSSProperties, PropType, SlotsType } from "vue";

export default defineComponent({
    name: "ElThing",
    props: {
        contentIndented: Boolean,
        content: String,
        contentClass: String,
        contentStyle: [String, Object] as PropType<string | CSSProperties>,
        description: String,
        descriptionClass: String,
        descriptionStyle: [String, Object] as PropType<string | CSSProperties>,
        titleExtra: String,
        title: String,
    },
    slots: Object as SlotsType<{
        default?: () => VNode[];
        avatar?: () => VNode[];
        header?: () => VNode[];
        "header-extra"?: () => VNode[];
        description?: () => VNode[];
        footer?: () => VNode[];
        action?: () => VNode[];
    }>,
    setup(props, { slots }) {
        const ns = useNamespace("thing");
        return () => {
            return (
                <div class={ns.b()}>
                    {slots.avatar && props.contentIndented ? (<div class={ns.b("avatar")}>{slots.avatar()}</div>) : null}
                    <div class={ns.b("main")}>
                        {!props.contentIndented
                            && (slots.header
                                || props.title
                                || slots["header-extra"]
                                || props.titleExtra
                                || slots.avatar)
                            ? (
                                <div class={ns.b("avatar-header-wrapper")}>
                                    {slots.avatar ? (<div class={ns.b("avatar")}>{slots.avatar()}</div>) : null}
                                    {slots.header
                                        || props.title
                                        || slots["header-extra"]
                                        || props.titleExtra
                                        ? (
                                            <div class={ns.b("header-wrapper")}>
                                                <div class={ns.b("header")}>
                                                    {slots.header || props.title
                                                        ? (
                                                            <div class={ns.be("header", "title")}>
                                                                {slots.header ? slots.header() : props.title}
                                                            </div>
                                                        )
                                                        : null}
                                                    {slots["header-extra"] || props.titleExtra
                                                        ? (
                                                            <div class={ns.be("header", "extra")}>
                                                                {slots["header-extra"] ? slots["header-extra"]() : props.titleExtra}
                                                            </div>
                                                        )
                                                        : null}
                                                </div>
                                                {slots.description || props.description
                                                    ? (
                                                        <div
                                                            class={[
                                                                ns.be("main", "description"),
                                                                props.descriptionClass,
                                                            ]}
                                                            style={props.descriptionStyle}
                                                        >
                                                            {slots.description ? slots.description() : props.description}
                                                        </div>
                                                    )
                                                    : null}
                                            </div>
                                        )
                                        : null}
                                </div>
                            )
                            : (
                                <>
                                    {slots.header
                                        || props.title
                                        || slots["header-extra"]
                                        || props.titleExtra
                                        ? (
                                            <div class={ns.b("header")}>
                                                {slots.header || props.title
                                                    ? (
                                                        <div class={ns.be("header", "title")}>
                                                            {slots.header ? slots.header() : props.title}
                                                        </div>
                                                    )
                                                    : null}
                                                {slots["header-extra"] || props.titleExtra
                                                    ? (
                                                        <div class={ns.be("header", "extra")}>
                                                            {slots["header-extra"] ? slots["header-extra"]() : props.titleExtra}
                                                        </div>
                                                    )
                                                    : null}
                                            </div>
                                        )
                                        : null}
                                    {slots.description || props.description
                                        ? (
                                            <div
                                                class={[
                                                    ns.be("main", "description"),
                                                    props.descriptionClass,
                                                ]}
                                                style={props.descriptionStyle}
                                            >
                                                {slots.description ? slots.description() : props.description}
                                            </div>
                                        )
                                        : null}
                                </>
                            )}
                        {slots.default || props.content
                            ? (
                                <div
                                    class={[
                                        ns.be("main", "content"),
                                        props.contentClass,
                                    ]}
                                    style={props.contentStyle}
                                >
                                    {slots.default ? slots.default() : props.content}
                                </div>
                            )
                            : null}
                        {slots.footer ? (<div class={ns.be("main", "footer")}>{slots.footer()}</div>) : null}
                        {slots.action ? (<div class={ns.be("main", "action")}>{slots.action()}</div>) : null}
                    </div>
                </div>
            );
        };
    },
});
</script>

<style lang="scss" scoped>
.el-thing {
    display: flex;
    &-avatar-header-wrapper {
        display: flex;
        flex-wrap: nowrap;
    }
    &-main {
        flex-grow: 1;
        &__content,
        &__footer,
        &__action {
            &:not(:first-child) {
                margin-top: 12px;
            }
        }
    }
    &-avatar {
        margin-right: 12px;
        margin-top: 2px;
    }
    &-header {
        display: flex;
        margin-bottom: 4px;
        justify-content: space-between;
        align-items: center;
        &-wrapper {
            flex: 1;
        }
        &__title {
            font-size: 16px;
            font-weight: 500;
        }
    }
}
</style>