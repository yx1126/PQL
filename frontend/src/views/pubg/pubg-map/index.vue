<script setup lang="ts">
import PubgMap from "../components/PubgMap.vue";
import MapHeader from "../components/MapHeader.vue";

defineOptions({
    name: "PubgMap",
});

const route = useRoute();
const router = useRouter();

const mapName = computed({
    get: () => route.params.name as string,
    set: name => {
        router.replace(`/pubg/map/${name}`);
    },
});

function onBack() {
    router.back();
}

function onCardItem(name: string) {
    mapName.value = name;
}
</script>

<template>
    <div class="size-full">
        <map-header v-model="mapName" @back="onBack" @map-item="onCardItem" />
        <div class="h-[calc(100%-50px)]">
            <pubg-map :map-name />
        </div>
    </div>
</template>