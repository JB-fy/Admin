<script setup lang="tsx">
const route: any = useRoute()
const router = useRouter()
const { t } = useI18n()
const keepAliveStore = useKeepAliveStore()
const languageStore = useLanguageStore()
const settingStore = useSettingStore()
const adminStore = useAdminStore()

const leftMenuFold = () => {
    settingStore.leftMenuFold = !settingStore.leftMenuFold
}

const userDropdown = reactive({
    status: false,
    visibleChange: (status: boolean) => {
        userDropdown.status = status
    },
})

const menuTab = reactive({
    refList: {} as { [propName: string]: any },
    visibleChange: (status: boolean, fullPath: string) => {
        if (status) {
            for (let key in menuTab.refList) {
                if (!menuTab.refList[key]) {
                    delete menuTab.refList[key]
                    continue
                }
                if (key !== fullPath) {
                    menuTab.refList[key].handleClose()
                }
            }
        }
    },
    change: (fullPath: string) => {
        if (fullPath === route.fullPath) {
            //左侧菜单点击会触发这个函数，故判断路由是否相同，相同不再跳转
            return false
        }
        router.push(fullPath)
    },
    remove: (fullPath: string) => {
        adminStore.closeSelfMenuTab(fullPath)
    },
    buttonDropdown: {
        status: false,
        visibleChange: (status: boolean) => {
            menuTab.buttonDropdown.status = status
            menuTab.visibleChange(status, '')
        },
    },
})
</script>

<template>
    <el-row>
        <el-col :span="12">
            <el-space :size="20" style="height: 100%; margin-left: 20px">
                <el-link :underline="false" @click="leftMenuFold">
                    <autoicon-ep-fold :class="{ 'fold-icon': true, 'is-fold': settingStore.leftMenuFold }" />
                </el-link>
                <el-breadcrumb separator=">">
                    <el-breadcrumb-item v-for="(item, index) in adminStore.getCurrentMenuChain" :key="index">
                        <el-space :size="0">
                            <my-icon-dynamic :icon="item.icon" />
                            <span>{{ item.title }}</span>
                        </el-space>
                    </el-breadcrumb-item>
                </el-breadcrumb>
            </el-space>
        </el-col>
        <el-col :span="12" style="text-align: right">
            <el-space :size="20" style="height: 100%; margin-right: 20px">
                <el-link :underline="false" @click="keepAliveStore.refreshMenuTab(route.meta.componentName)">
                    <autoicon-ep-refresh />
                </el-link>

                <el-dropdown>
                    <el-link :underline="false">
                        <autoicon-ep-place />
                    </el-link>
                    <template #dropdown>
                        <el-dropdown-menu v-for="(item, index) in settingStore.language" :key="index">
                            <el-dropdown-item @click="languageStore.changeLanguage(index)">
                                {{ item }}
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>

                <el-link :underline="false">
                    <full-screen-icon />
                </el-link>

                <el-dropdown @visible-change="userDropdown.visibleChange">
                    <el-link :underline="false">
                        <el-avatar :src="adminStore.info.avatar" :size="40">
                            <autoicon-ep-avatar />
                        </el-avatar>
                        <span>{{ adminStore.info.nickname }}</span>
                        <autoicon-ep-arrow-down :class="{ 'dropdown-icon': true, 'is-dropdown': userDropdown.status }" />
                    </el-link>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item>
                                <router-link to="/profile" :custom="true" v-slot="{ href, navigate, route }">
                                    <el-link :href="href" @click="navigate" :underline="false">
                                        {{ languageStore.getMenuTitle(route?.meta?.menu) }}
                                    </el-link>
                                </router-link>
                            </el-dropdown-item>
                            <el-dropdown-item @click="adminStore.logout()">
                                {{ t('common.logout') }}
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </el-space>
        </el-col>
        <el-col :span="24">
            <el-tabs class="menu-tabs" type="card" :model-value="route.fullPath" @tab-change="menuTab.change" @tab-remove="menuTab.remove">
                <template v-for="(item, index) in adminStore.getMenuTabList" :key="index">
                    <el-tab-pane :name="item.url" :closable="item.closable">
                        <template #label>
                            <el-dropdown
                                :ref="(el: any) => menuTab.refList[item.url] = el"
                                trigger="contextmenu"
                                @visible-change="
                                    (status: boolean) => {
                                        menuTab.visibleChange(status, item.url)
                                    }
                                "
                                style="height: 100%"
                                :key="item.url"
                            >
                                <el-space :size="0">
                                    <my-icon-dynamic :icon="item.icon" />
                                    <span>{{ item.title }}</span>
                                </el-space>
                                <template #dropdown>
                                    <el-dropdown-menu>
                                        <el-dropdown-item @click="keepAliveStore.refreshMenuTab(item.componentName)"> <autoicon-ep-refresh />{{ t('common.refresh') }} </el-dropdown-item>
                                        <el-dropdown-item @click="adminStore.closeSelfMenuTab(item.url)"> <autoicon-ep-close />{{ t('common.close') }} </el-dropdown-item>
                                        <el-dropdown-item @click="adminStore.closeLeftMenuTab(item.url)"> <autoicon-ep-back />{{ t('common.closeLeft') }} </el-dropdown-item>
                                        <el-dropdown-item @click="adminStore.closeRightMenuTab(item.url)"> <autoicon-ep-right />{{ t('common.closeRight') }} </el-dropdown-item>
                                        <el-dropdown-item @click="adminStore.closeOtherMenuTab(item.url)"> <autoicon-ep-switch />{{ t('common.closeOther') }} </el-dropdown-item>
                                        <el-dropdown-item @click="adminStore.closeAllMenuTab()"> <autoicon-ep-circle-close />{{ t('common.closeAll') }} </el-dropdown-item>
                                    </el-dropdown-menu>
                                </template>
                            </el-dropdown>
                        </template>
                    </el-tab-pane>
                </template>
            </el-tabs>

            <el-dropdown class="menu-tabs-button" @visible-change="menuTab.buttonDropdown.visibleChange">
                <el-link :underline="false">
                    <autoicon-ep-menu :class="{ 'dropdown-icon': true, 'is-dropdown': menuTab.buttonDropdown.status }" />
                </el-link>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item @click="keepAliveStore.refreshMenuTab(route.meta.componentName)"> <autoicon-ep-refresh />{{ t('common.refresh') }} </el-dropdown-item>
                        <el-dropdown-item @click="adminStore.closeSelfMenuTab(route.fullPath)"> <autoicon-ep-close />{{ t('common.close') }} </el-dropdown-item>
                        <el-dropdown-item @click="adminStore.closeLeftMenuTab(route.fullPath)"> <autoicon-ep-back />{{ t('common.closeLeft') }} </el-dropdown-item>
                        <el-dropdown-item @click="adminStore.closeRightMenuTab(route.fullPath)"> <autoicon-ep-right />{{ t('common.closeRight') }} </el-dropdown-item>
                        <el-dropdown-item @click="adminStore.closeOtherMenuTab(route.fullPath)"> <autoicon-ep-switch />{{ t('common.closeOther') }} </el-dropdown-item>
                        <el-dropdown-item @click="adminStore.closeAllMenuTab()"> <autoicon-ep-circle-close />{{ t('common.closeAll') }} </el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </el-col>
    </el-row>
</template>

<style scoped>
.el-row .el-col {
    height: 50px;
}

.fold-icon {
    font-size: 20px;
    transition:
        all 1s,
        color 0s;
}

.fold-icon.is-fold {
    transform: rotateY(180deg);
}

.dropdown-icon {
    font-size: 12px;
    transition-duration: 0.5s;
}

.dropdown-icon.is-dropdown {
    color: var(--el-link-hover-text-color);
    transform: rotate(180deg);
}

.el-row .el-col:nth-child(3) {
    padding-top: 9px;
    border-top: 1px solid var(--el-border-color);
    display: flex;
    align-content: center;
    align-items: center;
}

.menu-tabs {
    width: calc(100% - 80px);
    margin: 0 20px;
}

.menu-tabs-button {
    margin: -9px 20px 0 0;
}

.menu-tabs-button .el-link {
    width: 20px;
    height: 20px;
}

.menu-tabs-button .el-link .dropdown-icon {
    font-size: 20px;
}

.menu-tabs :deep(.el-tabs__header) {
    margin-bottom: 0;
    border: none;
}

.menu-tabs :deep(.el-tabs__nav) {
    border: none;
}

.menu-tabs :deep(.el-tabs__item:last-child) {
    margin-right: 0;
}

.menu-tabs :deep(.el-tabs__nav-prev),
.menu-tabs :deep(.el-tabs__nav-next),
.menu-tabs :deep(.el-tabs__item) {
    height: 40px;
    line-height: 40px;
}

.menu-tabs :deep(.el-tabs__item) {
    margin-right: -15px;
    padding: 0 20px 0 20px !important;
    transition: 0.3s;
    border: none;
}

.menu-tabs :deep(.el-tabs__item.is-active) {
    color: var(--el-color-primary);
    background: var(--el-color-primary-light-9);
    -webkit-mask: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANoAAAAkBAMAAAAdqzmBAAAAMFBMVEVHcEwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlTPQ5AAAAD3RSTlMAr3DvEM8wgCBA379gj5//tJBPAAAAnUlEQVRIx2NgAAM27fj/tAO/xBsYkIHyf9qCT8iWMf6nNQhAsk2f5rYheY7Dnua2/U+A28ZEe8v+F9Ax2v7/F4DbxkUH2wzgtvHTwbYPo7aN2jZq26hto7aN2jZq25Cy7Qvctnw62PYNbls9HWz7S8/G6//PsI6H4396gAUQy1je08W2jxDbpv6nD4gB2uWp+J9eYPsEhv/0BPS1DQBvoBLVZ3BppgAAAABJRU5ErkJggg==);
    mask: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANoAAAAkBAMAAAAdqzmBAAAAMFBMVEVHcEwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlTPQ5AAAAD3RSTlMAr3DvEM8wgCBA379gj5//tJBPAAAAnUlEQVRIx2NgAAM27fj/tAO/xBsYkIHyf9qCT8iWMf6nNQhAsk2f5rYheY7Dnua2/U+A28ZEe8v+F9Ax2v7/F4DbxkUH2wzgtvHTwbYPo7aN2jZq26hto7aN2jZq25Cy7Qvctnw62PYNbls9HWz7S8/G6//PsI6H4396gAUQy1je08W2jxDbpv6nD4gB2uWp+J9eYPsEhv/0BPS1DQBvoBLVZ3BppgAAAABJRU5ErkJggg==);
    -webkit-mask-size: 100% 100%;
    mask-size: 100% 100%;
}

.menu-tabs :deep(.el-tabs__item:not(.is-active):hover) {
    color: var(--el-color-black);
    background: #dcdfe6;
    -webkit-mask: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANoAAAAkBAMAAAAdqzmBAAAAMFBMVEVHcEwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlTPQ5AAAAD3RSTlMAr3DvEM8wgCBA379gj5//tJBPAAAAnUlEQVRIx2NgAAM27fj/tAO/xBsYkIHyf9qCT8iWMf6nNQhAsk2f5rYheY7Dnua2/U+A28ZEe8v+F9Ax2v7/F4DbxkUH2wzgtvHTwbYPo7aN2jZq26hto7aN2jZq25Cy7Qvctnw62PYNbls9HWz7S8/G6//PsI6H4396gAUQy1je08W2jxDbpv6nD4gB2uWp+J9eYPsEhv/0BPS1DQBvoBLVZ3BppgAAAABJRU5ErkJggg==);
    mask: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAANoAAAAkBAMAAAAdqzmBAAAAMFBMVEVHcEwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlTPQ5AAAAD3RSTlMAr3DvEM8wgCBA379gj5//tJBPAAAAnUlEQVRIx2NgAAM27fj/tAO/xBsYkIHyf9qCT8iWMf6nNQhAsk2f5rYheY7Dnua2/U+A28ZEe8v+F9Ax2v7/F4DbxkUH2wzgtvHTwbYPo7aN2jZq26hto7aN2jZq25Cy7Qvctnw62PYNbls9HWz7S8/G6//PsI6H4396gAUQy1je08W2jxDbpv6nD4gB2uWp+J9eYPsEhv/0BPS1DQBvoBLVZ3BppgAAAABJRU5ErkJggg==);
    -webkit-mask-size: 100% 100%;
    mask-size: 100% 100%;
}

.menu-tabs :deep(.el-tabs__item:not(.is-active) .is-icon-close) {
    margin-left: 0;
}

.menu-tabs :deep(.el-tabs__item:not(.is-active):hover .is-icon-close) {
    margin-left: 5px;
}
</style>
