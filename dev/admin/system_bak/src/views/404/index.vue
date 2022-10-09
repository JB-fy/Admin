<template>
    <div style="text-align: center; margin: 20px auto 0 auto;">
        <el-row>
            <el-col :span="10" style="text-align: right;">
                <!-- <el-image :src="require('@/assets/404.jpg')" /> -->
                <img src="@/assets/404.jpg">
            </el-col>
            <el-col :span="14" style="text-align: left; margin-top: 20px; display: flex; align-items: center;">
                <div>
                    <div style="font-size: 80px; color: #1989fa;">404</div>
                    <div style="font-size: 22px; color: #1482f0; font-weight: bold; margin-bottom: 5px;">哎呀！</div>
                    <div style="font-size: 15px; font-weight: bold; margin-bottom: 5px;">网页跟小鸡鸡一起不见了...</div>
                    <div style="font-size: 12px; color: grey; margin-bottom: 10px;">请检查您输入的网址是否正确，或点击下面按钮返回主页。</div>
                    <router-link to="/" :replace="true" :custom="true" v-slot="{ href, navigate, route }">
                        <el-link :href="href" :underline="false" @click="navigate">
                            <el-button type="primary" :round="true" size="small">返回主页</el-button>
                        </el-link>
                    </router-link>
                </div>
            </el-col>
        </el-row>
    </div>
</template>