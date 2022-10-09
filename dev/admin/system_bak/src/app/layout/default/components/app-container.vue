<template>
    <router-view v-slot="{ Component, route }">
        <transition mode="out-in" name="transform">
            <keep-alive :include="$store.getters['user/cacheRouteInclude']"
                :exclude="$store.state.user.cacheRoute.exclude" :max="$store.state.user.cacheRoute.max">
                <component v-if="$store.state.user.cacheRoute.exclude.indexOf(route.path) === -1" :is="Component"
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