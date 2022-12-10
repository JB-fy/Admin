<script setup lang="ts">
const languageStore = useLanguageStore()
defineProps({
    tree: {
        //type: Array as PropType<{ title: string, url: string, icon: string, children: [] }[]>,
        type: Array as any,
        required: true
    },
    subMenuIndexPrefix: {
        type: String,
        default: ''
    },
})
</script>

<template>
    <template v-for="(item, key) in tree" :key="key">
        <ElSubMenu v-if="item.children.length" :index="subMenuIndexPrefix + '/' + key">
            <template #title>
                <MyIconDynamic :icon="item.icon" />
                <span>{{ languageStore.getMenuTitle(item) }}</span>
            </template>
            <LeftMenuItem :tree="item.children" :subMenuIndexPrefix="subMenuIndexPrefix + '/' + key" />
        </ElSubMenu>
        <ElMenuItem v-else :index="item.url">
            <MyIconDynamic :icon="item.icon" />
            <template #title>
                <span>{{ languageStore.getMenuTitle(item) }}</span>
            </template>
        </ElMenuItem>
    </template>
</template>