<script setup lang="tsx">
const keepAliveStore = useKeepAliveStore()
const languageStore = useLanguageStore()
const settingStore = useSettingStore()
const adminStore = useAdminStore()

const { t, tm } = useI18n()

const route: any = useRoute()
const router = useRouter()

const leftMenuFold = () => {
    settingStore.leftMenuFold = !settingStore.leftMenuFold
}

const userDropdown = reactive({
    status: false,
    visibleChange: (status: boolean) => {
        userDropdown.status = status
    }
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
        }
    }
})
</script>

<template>
    <ElRow>
        <ElCol :span="12">
            <ElSpace :size="20" style="height: 100%; margin-left: 20px">
                <ElLink :underline="false" @click="leftMenuFold">
                    <AutoiconEpFold :class="{ 'fold-icon': true, 'is-fold': settingStore.leftMenuFold }" />
                </ElLink>
                <ElBreadcrumb separator=">">
                    <ElBreadcrumbItem v-for="(item, index) in adminStore.getCurrentMenuChain" :key="index">
                        <ElSpace :size="0">
                            <MyIconDynamic :icon="item.icon" />
                            <span>{{ item.title }}</span>
                        </ElSpace>
                    </ElBreadcrumbItem>
                </ElBreadcrumb>
            </ElSpace>
        </ElCol>
        <ElCol :span="12" style="text-align: right">
            <ElSpace :size="20" style="height: 100%">
                <ElLink :underline="false" @click="keepAliveStore.refreshMenuTab(route.meta.componentName)">
                    <AutoiconEpRefresh />
                </ElLink>

                <ElDropdown>
                    <ElLink :underline="false">
                        <AutoiconEpPlace />
                    </ElLink>
                    <template #dropdown>
                        <ElDropdownMenu v-for="(item, index) in settingStore.language" :key="index">
                            <ElDropdownItem @click="languageStore.changeLanguage(index)">
                                {{ item }}
                            </ElDropdownItem>
                        </ElDropdownMenu>
                    </template>
                </ElDropdown>

                <ElLink :underline="false">
                    <FullScreenIcon />
                </ElLink>

                <ElDropdown @visible-change="userDropdown.visibleChange">
                    <ElLink :underline="false">
                        <ElAvatar :src="adminStore.info.avatar" :size="40">
                            <AutoiconEpAvatar />
                        </ElAvatar>
                        <span>{{ adminStore.info.nickname }}</span>
                        <AutoiconEpArrowDown :class="{ 'dropdown-icon': true, 'is-dropdown': userDropdown.status }" />
                    </ElLink>
                    <template #dropdown>
                        <ElDropdownMenu>
                            <ElDropdownItem>
                                <RouterLink to="/profile" :custom="true" v-slot="{ href, navigate, route }">
                                    <ElLink :href="href" @click="navigate" :underline="false">
                                        {{ languageStore.getMenuTitle(route?.meta?.menu) }}
                                    </ElLink>
                                </RouterLink>
                            </ElDropdownItem>
                            <ElDropdownItem @click="adminStore.logout()">
                                {{ t('common.logout') }}
                            </ElDropdownItem>
                        </ElDropdownMenu>
                    </template>
                </ElDropdown>
            </ElSpace>
        </ElCol>
        <ElCol :span="24">
            <ElTabs class="menu-tabs" type="card" :model-value="route.fullPath" @tab-change="menuTab.change" @tab-remove="menuTab.remove">
                <template v-for="(item, index) in adminStore.getMenuTabList" :key="index">
                    <ElTabPane :name="item.url" :closable="item.closable">
                        <template #label>
                            <ElDropdown
                                :ref="(el: any) => (menuTab.refList[item.url] = el)"
                                trigger="contextmenu"
                                @visible-change="
                                    (status: boolean) => {
                                        menuTab.visibleChange(status, item.url)
                                    }
                                "
                                style="height: 100%"
                                :key="item.url"
                            >
                                <ElSpace :size="0">
                                    <MyIconDynamic :icon="item.icon" />
                                    <span>{{ item.title }}</span>
                                </ElSpace>
                                <template #dropdown>
                                    <ElDropdownMenu>
                                        <ElDropdownItem @click="keepAliveStore.refreshMenuTab(item.componentName)"> <AutoiconEpRefresh />{{ t('common.refresh') }} </ElDropdownItem>
                                        <ElDropdownItem @click="adminStore.closeSelfMenuTab(item.url)"> <AutoiconEpClose />{{ t('common.close') }} </ElDropdownItem>
                                        <ElDropdownItem @click="adminStore.closeLeftMenuTab(item.url)"> <AutoiconEpBack />{{ t('common.closeLeft') }} </ElDropdownItem>
                                        <ElDropdownItem @click="adminStore.closeRightMenuTab(item.url)"> <AutoiconEpRight />{{ t('common.closeRight') }} </ElDropdownItem>
                                        <ElDropdownItem @click="adminStore.closeOtherMenuTab(item.url)"> <AutoiconEpSwitch />{{ t('common.closeOther') }} </ElDropdownItem>
                                        <ElDropdownItem @click="adminStore.closeAllMenuTab()"> <AutoiconEpCircleClose />{{ t('common.closeAll') }} </ElDropdownItem>
                                    </ElDropdownMenu>
                                </template>
                            </ElDropdown>
                        </template>
                    </ElTabPane>
                </template>
            </ElTabs>

            <ElDropdown class="menu-tabs-button" @visible-change="menuTab.buttonDropdown.visibleChange">
                <ElLink :underline="false">
                    <AutoiconEpMenu :class="{ 'dropdown-icon': true, 'is-dropdown': menuTab.buttonDropdown.status }" />
                </ElLink>
                <template #dropdown>
                    <ElDropdownMenu>
                        <ElDropdownItem @click="keepAliveStore.refreshMenuTab(route.meta.componentName)"> <AutoiconEpRefresh />{{ t('common.refresh') }} </ElDropdownItem>
                        <ElDropdownItem @click="adminStore.closeSelfMenuTab(route.fullPath)"> <AutoiconEpClose />{{ t('common.close') }} </ElDropdownItem>
                        <ElDropdownItem @click="adminStore.closeLeftMenuTab(route.fullPath)"> <AutoiconEpBack />{{ t('common.closeLeft') }} </ElDropdownItem>
                        <ElDropdownItem @click="adminStore.closeRightMenuTab(route.fullPath)"> <AutoiconEpRight />{{ t('common.closeRight') }} </ElDropdownItem>
                        <ElDropdownItem @click="adminStore.closeOtherMenuTab(route.fullPath)"> <AutoiconEpSwitch />{{ t('common.closeOther') }} </ElDropdownItem>
                        <ElDropdownItem @click="adminStore.closeAllMenuTab()"> <AutoiconEpCircleClose />{{ t('common.closeAll') }} </ElDropdownItem>
                    </ElDropdownMenu>
                </template>
            </ElDropdown>
        </ElCol>
    </ElRow>
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
