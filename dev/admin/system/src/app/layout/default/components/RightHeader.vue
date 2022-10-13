<script lang="ts">
import { useSettingStore } from '@/stores/setting';
import { useUserStore } from '@/stores/user';

export default {
    setup: () => {
        const settingStore = useSettingStore()
        const userStore = useUserStore()
        const route = useRoute()
        const router = useRouter()
        const state = reactive({
            leftMenuFold: () => {
                settingStore.leftMenuFold = !settingStore.leftMenuFold
            },
            userDropdown: {
                status: false,
                visibleChange: (status) => {
                    state.userDropdown.status = status
                },
            },
            menuTab: {
                refList: {},
                visibleChange: (status, path) => {
                    if (status) {
                        for (let key in state.menuTab.refList) {
                            if (!state.menuTab.refList[key]) {
                                delete state.menuTab.refList[key]
                                continue
                            }
                            if (key !== path) {
                                state.menuTab.refList[key].handleClose()
                            }
                        }
                    }
                },
                change: (path) => {
                    if (path === route.path) {  //左侧菜单点击会触发这个函数，故判断路由是否相同，相同不再跳转
                        return false
                    }
                    router.push(path)
                },
                remove: (path) => {
                    userStore.closeSelfMenuTab(path)
                },
                buttonDropdown: {
                    status: false,
                    visibleChange: (status) => {
                        state.menuTab.buttonDropdown.status = status
                        state.menuTab.visibleChange(status, '')
                    },
                },
            }
        })
        return {
            ...toRefs(state),
            settingStore,
            userStore
        }
    }
}
</script>

<template>
    <el-row>
        <el-col :span="12">
            <el-space :size="20" style="height: 100%; margin-left: 20px;">
                <el-link :underline="false" @click="leftMenuFold">
                    <autoicon-ep-fold :class="{ 'fold-icon': true, 'is-fold': settingStore.leftMenuFold }" />
                </el-link>
                <el-breadcrumb separator=">">
                    <el-breadcrumb-item v-for="(item, key) in $route.meta.pMenuList" :key="key">
                        <el-space :size="0">
                            <icon-dynamic :icon="item.icon" />
                            <span>{{ item.title }}</span>
                        </el-space>
                    </el-breadcrumb-item>
                    <el-breadcrumb-item>
                        <el-space :size="0">
                            <icon-dynamic :icon="$route.meta.icon" />
                            <span>{{ $route.meta.title }}</span>
                        </el-space>
                    </el-breadcrumb-item>
                </el-breadcrumb>
            </el-space>
        </el-col>
        <el-col :span="12" style="text-align: right;">
            <el-space :size="20" style="height: 100%;">
                <el-link :underline="false">
                    <autoicon-ep-lock />
                </el-link>
                <el-link :underline="false">
                    <autoicon-ep-search />
                </el-link>
                <el-link :underline="false">
                    <autoicon-ep-bell />
                </el-link>
                <el-link :underline="false" @click="userStore.refreshMenuTab($route.path)">
                    <autoicon-ep-refresh />
                </el-link>
                <el-dropdown @visible-change="userDropdown.visibleChange">
                    <el-link :underline="false">
                        <el-avatar :src="userStore.info.avatar" :size="40">
                            <autoicon-ep-user-filled />
                        </el-avatar>
                        <span>{{ userStore.info.nickname }}</span>
                        <autoicon-ep-arrow-down
                            :class="{ 'dropdown-icon': true, 'is-dropdown': userDropdown.status }" />
                    </el-link>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item>
                                <router-link to="/profile" :custom="true" v-slot="{ href, navigate, route }">
                                    <el-link :href="href" @click="navigate" :underline="false">
                                        {{ route.meta.title }}
                                    </el-link>
                                </router-link>
                            </el-dropdown-item>
                            <el-dropdown-item @click="userStore.logout()">
                                退出登录
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </el-space>
        </el-col>
        <el-col :span="24">
            <el-tabs class="menu-tabs" type="card" :model-value="$route.path" @tab-change="menuTab.change"
                @tab-remove="menuTab.remove">
                <template v-for="(item, key) in userStore.menuTabList" :key="key">
                    <el-tab-pane :name="item.path" :closable="item.closable">
                        <template #label>
                            <el-dropdown :ref="(el) => { menuTab.refList[item.path] = el }" trigger="contextmenu"
                                @visible-change="(status) => { menuTab.visibleChange(status, item.path) }"
                                style="height: 100%;">
                                <el-space :size="0">
                                    <icon-dynamic :icon="item.icon" />
                                    <span>{{ item.title }}</span>
                                </el-space>
                                <template #dropdown>
                                    <el-dropdown-menu>
                                        <el-dropdown-item @click="userStore.refreshMenuTab(item.path)">
                                            刷新
                                        </el-dropdown-item>
                                        <el-dropdown-item @click="userStore.closeOtherMenuTab(item.path)">
                                            关闭其他
                                        </el-dropdown-item>
                                        <el-dropdown-item @click="userStore.closeLeftMenuTab(item.path)">
                                            关闭左侧
                                        </el-dropdown-item>
                                        <el-dropdown-item @click="userStore.closeRightMenuTab(item.path)">
                                            关闭右侧
                                        </el-dropdown-item>
                                        <el-dropdown-item @click="userStore.closeAllMenuTab()">
                                            关闭全部
                                        </el-dropdown-item>
                                    </el-dropdown-menu>
                                </template>
                            </el-dropdown>
                        </template>
                    </el-tab-pane>
                </template>
            </el-tabs>

            <el-dropdown class="menu-tabs-button" @visible-change="menuTab.buttonDropdown.visibleChange">
                <el-link :underline="false">
                    <autoicon-ep-menu
                        :class="{ 'dropdown-icon': true, 'is-dropdown': menuTab.buttonDropdown.status }" />
                </el-link>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item @click="userStore.refreshMenuTab($route.path)">
                            刷新
                        </el-dropdown-item>
                        <el-dropdown-item @click="userStore.closeOtherMenuTab($route.path)">
                            关闭其他
                        </el-dropdown-item>
                        <el-dropdown-item @click="userStore.closeLeftMenuTab($route.path)">
                            关闭左侧
                        </el-dropdown-item>
                        <el-dropdown-item @click="userStore.closeRightMenuTab($route.path)">
                            关闭右侧
                        </el-dropdown-item>
                        <el-dropdown-item @click="userStore.closeAllMenuTab()">
                            关闭全部
                        </el-dropdown-item>
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
    transition: all 3s, color 0s;
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