<script setup lang="ts">
import img404 from "@/assets/image/404.jpg";

const { t } = useI18n()
</script>

<template>
    <div style="text-align: center; margin: 20px auto 0 auto;">
        <ElRow style="display: flex; align-items: center;">
            <ElCol :span="10">
                <ElImage :src="img404" />
                <!-- <img src="@/assets/image/404.jpg"> -->
            </ElCol>
            <ElCol :span="14" style="text-align: left;">
                <div>
                    <div style="font-size: 4rem; color: #1989fa;">404</div>
                    <div style="font-size: 1.25rem; color: #1482f0; font-weight: bold;">{{ t('view.404.oops') }}</div>
                    <div style="font-size: 1rem; font-weight: bold;">{{ t('view.404.title') }}</div>
                    <div style="font-size: 0.65rem; color: grey;">{{ t('view.404.tip') }}</div>
                    <RouterLink to="/" :replace="true" :custom="true" v-slot="{ href, navigate, route }">
                        <ElLink :href="href" :underline="false" @click="navigate">
                            <ElButton type="primary" :round="true" size="small">{{ t('view.404.backHome') }}</ElButton>
                        </ElLink>
                    </RouterLink>
                </div>
            </ElCol>
        </ElRow>
    </div>
</template>