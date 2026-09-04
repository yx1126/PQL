<script setup lang="ts">
import { keepNames } from "@/router/staticRoutes";
import { isArray } from "@/utils/validata";

defineOptions({
    name: "KeepRouteView",
});

defineProps<{
    isFull?: boolean;
}>();

const keepalive = useMitt("keepalive");

const exclude = ref<string[]>([]);

onBeforeMount(() => {
    keepalive.on(async key => {
        if(!key) return;
        const keys = isArray(key) ? key : [key];
        if(keys.length <= 0) return;
        exclude.value.push(...keys);
        await nextTick();
        exclude.value = [];
    });
});
</script>

<template>
    <!-- <component :is="Component" :key="route.fullPath" /> -->
    <router-view #default="{ Component, route }">
        <KeepAlive :include="keepNames" :exclude>
            <component :is="Component" :key="isFull ? route.fullPath : undefined" />
        </KeepAlive>
    </router-view>
</template>
