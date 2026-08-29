<script setup lang="ts">
import MapCard from "../components/MapCard.vue";
import { mapList } from "@/utils/maps.ts";

defineOptions({
    name: "PubgHome",
});

const router = useRouter();

function onClick(name: string) {
    router.push(`/pubg/map/${name}`);
}
</script>

<template>
    <div class="pubg-main">
        <div class="pubg-main_cards">
            <template v-for="map,i in mapList" :key="i">
                <map-card
                    :src="map.thumbnail"
                    :map-name="map.name"
                    @click="onClick(map.label)"
                />
            </template>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.pubg-main {
    width: 100%;
    height: 100%;
    background-color: black;
    background-image: url("@/assets/pubg/pubg-bg.jpg");
    background-size: cover;
    background-position: center;
    filter: brightness(0.9);
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    align-items: center;
    padding: 100px 0;
    &_cards {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 20px;
        flex-wrap: wrap;
    }
}
</style>
