import { isFn } from "@/utils/validata";
import { tryOnScopeDispose } from "@vueuse/core";
import { ElDialog, ElConfigProvider, type DialogProps } from "element-plus";
import type { AppContext, Component, InjectionKey, Ref, VNode } from "vue";
import { useI18n } from "vue-i18n";

export interface DialogOptions<
    Form extends object = any,
    Comp extends Component = any,
> extends Partial<Omit<Writable<DialogProps>, "modelValue" | "title" | "width">> {
    width?: string | number | (() => string | number);
    component: Comp;
    context?: Nullable<AppContext>;
    title?: string | ((form: Form) => string);
    renderHeader?: () => VNode | VNode[];
    renderFooter?: () => VNode | VNode[];
    showCancel?: boolean;
    cancelText?: string;
    showConfirm?: boolean;
    confirmText?: string;
    showActions?: boolean;
    onCancel?: () => void;
    onSubmit?: () => void;
}

let dialogWrapper: Nullable<HTMLDivElement> = null;
const DialogInstance: VNode[] = [];

function createWrapper() {
    if(!dialogWrapper) {
        dialogWrapper = document.createElement("div");
        dialogWrapper.classList.add("dialog-wrapper");
        document.body.appendChild(dialogWrapper);
    }
}

function destoryWrapper(instance?: Nullable<VNode>) {
    const index = DialogInstance.findIndex(v => v === instance);
    if(index != -1) {
        DialogInstance.splice(index, 1);
    }
    if(DialogInstance.length <= 0) {
        if(dialogWrapper) {
            document.body.removeChild(dialogWrapper);
        }
        dialogWrapper = null;
    }
}

interface DialogSubmitState {
    loading: boolean;
}

interface DialogLifeCycle {
    open: Ref<Array<(form: any) => void>>;
    opened: Ref<Array<(form: any) => void>>;
    close: Ref<Array<() => void>>;
    closed: Ref<Array<() => void>>;
    openAutoFocus: Ref<Array<() => void>>;
    closeAutoFocus: Ref<Array<() => void>>;
    cancel: Ref<Array<() => void>>;
    submit: Ref<Array<(instance: DialogSubmitState, done: () => void) => void>>;
}

const DialogContext = Symbol() as InjectionKey<DialogLifeCycle>;

export function createLifeCycle<T extends keyof DialogLifeCycle>(key: T) {
    return (fn: UnRef<DialogLifeCycle[T]>[number]) => {
        const dialog = inject(DialogContext);
        if(!dialog) return;
        dialog[key].value.push(fn as any);
        tryOnScopeDispose(() => {
            const index = dialog[key].value.findIndex(v => v === fn);
            if(index !== -1) {
                dialog[key].value.splice(index, 1);
            }
        });
    };
}

export const onDialogOpen = createLifeCycle("open");
export const onDialogOpened = createLifeCycle("opened");
export const onDialogClose = createLifeCycle("close");
export const onDialogClosed = createLifeCycle("closed");
export const onDialogOpenAutoFocus = createLifeCycle("openAutoFocus");
export const onDialogCloseAutoFocus = createLifeCycle("closeAutoFocus");
export const onDialogSubmit = createLifeCycle("submit");
export const onDialogCancel = createLifeCycle("cancel");

function createLifeCycleList<T extends keyof DialogLifeCycle>(_key: T) {
    return ref([]) as DialogLifeCycle[T];
}

export function useDialog<Form extends object = any, Comp extends Component = any>(options: DialogOptions<Form, Comp>) {
    const {
        width,
        component,
        title,
        context,
        showActions = true,
        showCancel = true,
        showConfirm = true,
        cancelText,
        confirmText,
        renderHeader,
        renderFooter,
        onCancel,
        onSubmit,
        ...otherDialogProps
    } = Object.assign({}, options);

    const app = useApp();
    const { t: $t } = useI18n({
        useScope: "global",
    });

    let divWrapper: Nullable<HTMLDivElement> = null;
    let instance: Nullable<VNode> = null;

    const state = reactive({
        loading: false,
        visible: false,
    });

    const formData = ref<Nullable<Form>>(null);

    const dialog = defineComponent({
        name: "DialogWrapper",
        inheritAttrs: false,
        emits: ["open", "opened", "close", "closed", "open-auto-focus", "close-auto-focus"],
        setup(_, { emit }) {
            const openLifeCycle = createLifeCycleList("open");
            const openedLifeCycle = createLifeCycleList("opened");
            const closeLifeCycle = createLifeCycleList("close");
            const closedLifeCycle = createLifeCycleList("closed");
            const openAutoFocusLifeCycle = createLifeCycleList("openAutoFocus");
            const closeAutoFocusLifeCycle = createLifeCycleList("closeAutoFocus");
            const cancelLifeCycle = createLifeCycleList("cancel");
            const submitLifeCycle = createLifeCycleList("submit");

            provide(DialogContext, {
                open: openLifeCycle,
                opened: openedLifeCycle,
                close: closeLifeCycle,
                closed: closedLifeCycle,
                openAutoFocus: openAutoFocusLifeCycle,
                closeAutoFocus: closeAutoFocusLifeCycle,
                cancel: cancelLifeCycle,
                submit: submitLifeCycle,
            });

            const EventMap: Record<string, () => void> = {
                onOpen: () => {
                    emit("open");
                    openLifeCycle.value.forEach(fn => fn(formData.value));
                },
                onOpened: () => {
                    emit("opened");
                    openedLifeCycle.value.forEach(fn => fn(formData.value));
                },
                onClose: () => {
                    emit("close");
                    closeLifeCycle.value.forEach(fn => fn());
                },
                onClosed: () => {
                    emit("closed");
                    closedLifeCycle.value.forEach(fn => fn());
                },
                onOpenAutoFocus: () => {
                    emit("open-auto-focus");
                    openAutoFocusLifeCycle.value.forEach(fn => fn());
                },
                onCloseAutoFocus: () => {
                    emit("close-auto-focus");
                    closeAutoFocusLifeCycle.value.forEach(fn => fn());
                },
            };

            function close(fn?: () => void) {
                state.visible = false;
                fn && fn();
            }

            function onClick(action: "cancel" | "submit") {
                if(action === "cancel") {
                    close();
                    onCancel?.();
                    cancelLifeCycle.value.forEach(fn => fn());
                } else if(action === "submit") {
                    submitLifeCycle.value.forEach(fn => {
                        function done() {
                            close(() => onSubmit?.());
                        }
                        fn(state, done);
                    });
                }
            }

            return {
                EventMap,
                onClick,
            };
        },
        render() {
            const { $attrs, EventMap, onClick } = this;
            return (
                <ElConfigProvider>
                    <ElDialog
                        v-model={state.visible}
                        {...$attrs}
                        {...EventMap}
                        title={isFn(title) ? title(formData.value) : title}
                        width={isFn(width) ? width() : width}
                    >
                        {{
                            default: () => h(component),
                            header: renderHeader,
                            footer: () => {
                                if(renderFooter) return renderFooter();
                                if(!showActions) return null;
                                return (
                                    <div class="dialog-footer">
                                        {showCancel ? <el-button onClick={() => onClick("cancel")}>{cancelText || $t("button.cancel")}</el-button> : null}
                                        {showConfirm ? <el-button type="primary" loading={state.loading} onClick={() => onClick("submit")}>{confirmText || $t("button.submit")}</el-button> : null}
                                    </div>
                                );
                            },
                        }}
                    </ElDialog>
                </ElConfigProvider>
            );
        },
    });

    tryOnScopeDispose(destory);

    function open(form?: Partial<Nullable<Form>>) {
        formData.value = form;
        createWrapper();
        if(!instance) {
            divWrapper = document.createElement("div");
            instance = createVNode(dialog, {
                draggable: true,
                destroyOnClose: true,
                ...otherDialogProps,
            });
            instance.appContext = context || app.appContext;
            render(instance, divWrapper);
            dialogWrapper!.appendChild(divWrapper);
            DialogInstance.push(instance);
        }
        nextTick(() => {
            state.visible = true;
        });
    }

    function destory() {
        if(divWrapper) {
            render(null, divWrapper);
            if(dialogWrapper) {
                dialogWrapper.removeChild(divWrapper);
            }
        }
        destoryWrapper(instance);
    }

    return {
        open,
        destory,
    };
}