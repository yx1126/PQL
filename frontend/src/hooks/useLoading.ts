const loadCount = ref(0);

export function useLoading() {
    const loading = computed(() => loadCount.value > 0);

    onBeforeUnmount(() => {
        loadCount.value = 0;
    });

    function setLoad(value = true) {
        if(value) {
            loadCount.value++;
        } else {
            loadCount.value = Math.max(loadCount.value - 1, 0);
        }
    }
    return {
        loading,
        loadCount,
        setLoad,
    };
}