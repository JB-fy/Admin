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
    }
})
</script>

<template>
    <template v-for="(item, index) in tree" :key="index">
        <ElSubMenu v-if="item.children.length" :index="subMenuIndexPrefix + '/' + index">
            <template #title>
                <MyIconDynamic :icon="item.icon" />
                <span>{{ languageStore.getMenuTitle(item) }}</span>
            </template>
            <LeftMenuItem :tree="item.children" :subMenuIndexPrefix="subMenuIndexPrefix + '/' + index" />
        </ElSubMenu>
        <ElMenuItem v-else :index="item.url">
            <MyIconDynamic :icon="item.icon" />
            <template #title>
                <span>{{ languageStore.getMenuTitle(item) }}</span>
            </template>
        </ElMenuItem>
    </template>
</template>
