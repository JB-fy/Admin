<script setup lang="ts">
import { useKeepAliveStore } from '@/stores/keepAlive';

//const { t } = useI18n()
const keepAliveStore = useKeepAliveStore()
</script>

<template>
    <RouterView v-slot="{ Component, route }">
        <Transition mode="out-in" name="transform">
            <KeepAlive :include="keepAliveStore.appContainerInclude" :exclude="keepAliveStore.appContainerExclude"
                :max="keepAliveStore.appContainerMax">
                <component v-if="keepAliveStore.appContainerExclude.indexOf(route.path) === -1" :is="Component"
                    :key="route.path" />
            </KeepAlive>
        </Transition>
    </RouterView>
    <!-- <Suspense>
        <template #default>
            <RouterView v-slot="{ Component, route }">
                <Transition mode="out-in" name="transform">
                    <KeepAlive :include="keepAliveStore.appContainerInclude"
                        :exclude="keepAliveStore.appContainerExclude" :max="keepAliveStore.appContainerMax">
                        <component v-if="keepAliveStore.appContainerExclude.indexOf(route.path) === -1" :is="Component"
                            :key="route.path" />
                    </KeepAlive>
                </Transition>
            </RouterView>
        </template>
        <template #fallback> t('common.loading') </template>
    </Suspense> -->
</template>

<style scoped>
.transform-enter-active,
.transform-leave-active {
    transition: opacity 0.3s ease;
}

/* .transform-enter, */
.transform-enter-from,
.transform-leave-to {
    opacity: 0;
}
</style>