<script setup lang="ts">
import { keepNames } from "@/router/staticRoutes";

defineOptions({
    name: "KeepRouteView",
});

defineProps<{
    isFull?: boolean;
}>();
</script>

<template>
    <!-- <component :is="Component" :key="route.fullPath" /> -->
    <router-view #default="{ Component, route }">
        <KeepAlive :include="keepNames">
            <component :is="Component" :key="isFull ? route.fullPath : undefined" />
        </KeepAlive>
    </router-view>
</template>
