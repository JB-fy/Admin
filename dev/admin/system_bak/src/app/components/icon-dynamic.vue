<template>
    <el-icon v-if="icon">
        <van-icon v-if="prefix === 'vant'" :name="iconCode" />
        <component v-else :is="iconCode" />
        <!-- <component v-else-if="icon.indexOf('ep-') === 0" :is="iconCode" />
        <component v-else-if="icon.indexOf('autoicon-') === 0" :is="iconCode" /> -->
    </el-icon>
</template>

<script>
export default {
    props: {
        icon: {
            type: String,
            required: true,
            default: ''
        }
    },
    setup: (props) => {
        const state = reactive({
            prefix: computed(() => {
                return props.icon.slice(0, props.icon.indexOf('-'))
            }),
            iconCode: computed(() => {
                if (state.prefix === 'vant') {
                    return props.icon.slice(props.icon.indexOf('-') + 1)
                }
                return props.icon
            }),
        })
        return {
            ...toRefs(state),
        }
    }
}
</script>