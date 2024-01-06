<script setup lang="tsx">
import img404 from '@/assets/image/404.jpg'

const { t } = useI18n()
</script>

<template>
    <div style="text-align: center; margin: 20px auto 0 auto">
        <el-row style="display: flex; align-items: center">
            <el-col :span="10">
                <el-image :src="img404" />
                <!-- <img src="@/assets/image/404.jpg"> -->
            </el-col>
            <el-col :span="14" style="text-align: left">
                <div>
                    <div style="font-size: 4rem; color: #1989fa">404</div>
                    <div style="font-size: 1.25rem; color: #1482f0; font-weight: bold">{{ t('404.oops') }}</div>
                    <div style="font-size: 1rem; font-weight: bold">{{ t('404.title') }}</div>
                    <div style="font-size: 0.65rem; color: grey">{{ t('404.tip') }}</div>
                    <router-link to="/" :replace="true" :custom="true" v-slot="{ href, navigate, route }">
                        <el-link :href="href" :underline="false" @click="navigate">
                            <el-button type="primary" :round="true" size="small">{{ t('404.backHome') }}</el-button>
                        </el-link>
                    </router-link>
                </div>
            </el-col>
        </el-row>
    </div>
</template>
