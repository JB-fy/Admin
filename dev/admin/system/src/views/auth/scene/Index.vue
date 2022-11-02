<script setup lang="ts">
onMounted(async () => {
    request('index.index')
})
let input = ref('')
let aaaa = ref(false)
</script>

<template>
    <div>
        <div @click="aaaa = true">点击</div>
        <el-input v-model="input" placeholder="请输入" />
        <!-- <el-dialog v-model="aaaa" title="Shipping address">
            111111
        </el-dialog> -->
        <el-drawer v-model="aaaa" title="新增" size="50%" :show-close="true">
            <div style="text-align: center; font-size: 300px; color: #409EFF;">场景列表</div>
            <template #footer>
                <div>
                    <el-button>cancel</el-button>
                    <el-button type="primary">confirm</el-button>
                </div>
            </template>
        </el-drawer>
        <div style="text-align: center; font-size: 300px; color: #409EFF;">场景列表</div>
    </div>
</template>