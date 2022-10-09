<template>
    <template v-for="(item, key) in tree" :key="key">
        <el-sub-menu v-if="item.children.length" :index="subMenuIndexPrefix + '/' + key">
            <template #title>
                <icon-dynamic :icon="item.icon" />
                <span>{{ item.title }}</span>
            </template>
            <left-menu-item :tree="item.children" :subMenuIndexPrefix="subMenuIndexPrefix + '/' + key" />
        </el-sub-menu>
        <el-menu-item v-else :index="item.path">
            <icon-dynamic :icon="item.icon" />
            <template #title>
                <span>{{ item.title }}</span>
            </template>
        </el-menu-item>
    </template>
</template>

<script>
export default {
    name: 'left-menu-item',
    props: {
        tree: {
            type: Array,
            required: true
        },
        subMenuIndexPrefix: {
            type: String,
            default: ''
        },
    }
}
</script>