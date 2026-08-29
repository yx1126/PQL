<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
    name: "Redirect",
    setup() {
        const route = useRoute();
        const router = useRouter();

        onBeforeMount(() => {
            const path = route.params.path as string[];
            if(!path?.length) {
                router.replace("/");
                return;
            }
            router.replace({ path: `/${path.join("/")}`, query: route.query });
        });
    },
    render() {
        return h("span", { style: "display: none;" });
    },
});
</script>