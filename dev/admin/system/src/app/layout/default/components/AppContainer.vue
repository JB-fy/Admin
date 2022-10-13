<script setup lang="ts">
import { useUserStore } from '@/stores/user';

const userStore = useUserStore()
</script>

<template>
    <router-view v-slot="{ Component, route }">
        <transition mode="out-in" name="transform">
            <keep-alive :include="userStore.cacheRouteInclude" :exclude="userStore.cacheRoute.exclude"
                :max="userStore.cacheRoute.max">
                <component v-if="userStore.cacheRoute.exclude.indexOf(route.path) === -1" :is="Component"
                    :key="route.path" />
            </keep-alive>
        </transition>
    </router-view>
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