<script setup lang="tsx">
const languageStore = useLanguageStore()
defineProps({
    tree: {
        //type: Array as PropType<{ title: string, url: string, icon: string, children: [] }[]>,
        type: Array as any,
        required: true,
    },
    subMenuIndexPrefix: {
        type: String,
        default: '',
    },
})
</script>

<template>
    <template v-for="(item, index) in tree" :key="index">
        <el-sub-menu v-if="item.children.length" :index="subMenuIndexPrefix + '/' + index">
            <template #title>
                <my-icon-dynamic :icon="item.icon" />
                <span>{{ languageStore.getMenuTitle(item) }}</span>
            </template>
            <left-menu-item :tree="item.children" :subMenuIndexPrefix="subMenuIndexPrefix + '/' + index" />
        </el-sub-menu>
        <el-menu-item v-else :index="item.url">
            <my-icon-dynamic :icon="item.icon" />
            <template #title>
                <span>{{ languageStore.getMenuTitle(item) }}</span>
            </template>
        </el-menu-item>
    </template>
</template>
