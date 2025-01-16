<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const table = reactive({
    columns: [
        {
            dataKey: 'id',
            title: t('common.name.id'),
            key: 'id',
            align: 'center',
            width: 200,
            fixed: 'left',
            sortable: true,
            headerCellRenderer: () => {
                const allChecked = table.data.every((item: any) => item.checked)
                const someChecked = table.data.some((item: any) => item.checked)
                return [
                    //阻止冒泡
                    <div class="id-checkbox" onClick={(event: any) => event.stopPropagation()}>
                        <el-checkbox model-value={table.data.length ? allChecked : false} indeterminate={someChecked && !allChecked} onChange={(val: boolean) => table.data.forEach((item: any) => (item.checked = val))} />
                    </div>,
                    <div>{t('common.name.id')}</div>,
                ]
            },
            cellRenderer: (props: any): any => {
                return [<el-checkbox class="id-checkbox" model-value={props.rowData.checked} onChange={(val: boolean) => (props.rowData.checked = val)} />, <div>{props.rowData.id}</div>]
            },
        },
        {
            dataKey: 'nickname',
            title: t('users.users.name.nickname'),
            key: 'nickname',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'avatar',
            title: t('users.users.name.avatar'),
            key: 'avatar',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                if (!props.rowData.avatar) {
                    return
                }
                const imageList = [props.rowData.avatar]
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {imageList.map((item) => {
                                return <el-image src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            },
        },
        {
            dataKey: 'gender',
            title: t('users.users.name.gender'),
            key: 'gender',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let statusList = tm('users.users.status.gender') as { value: any; label: string }[]
                let statusIndex = statusList.findIndex((item) => item.value == props.rowData.gender)
                return <el-tag type={tagType[statusIndex % tagType.length]}>{statusList[statusIndex]?.label}</el-tag>
            },
        },
        {
            dataKey: 'birthday',
            title: t('users.users.name.birthday'),
            key: 'birthday',
            align: 'center',
            width: 100,
            sortable: true,
        },
        {
            dataKey: 'address',
            title: t('users.users.name.address'),
            key: 'address',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'phone',
            title: t('users.users.name.phone'),
            key: 'phone',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'email',
            title: t('users.users.name.email'),
            key: 'email',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'account',
            title: t('users.users.name.account'),
            key: 'account',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'wx_openid',
            title: t('users.users.name.wx_openid'),
            key: 'wx_openid',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'wx_unionid',
            title: t('users.users.name.wx_unionid'),
            key: 'wx_unionid',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'id_card_no',
            title: t('users.users.name.id_card_no'),
            key: 'id_card_no',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'id_card_name',
            title: t('users.users.name.id_card_name'),
            key: 'id_card_name',
            align: 'center',
            width: 150,
            /* cellRenderer: (props: any): any => {
                if (!authAction.isUpdate) {
                    return [<div class="el-table-v2__cell-text">{props.rowData.id_card_name}</div>]
                }
                if (!props.rowData?.editIdCardName?.isEdit) {
                    return [
                        <div class="el-table-v2__cell-text inline-edit" onClick={() => (props.rowData.editIdCardName = { isEdit: true, oldValue: props.rowData.id_card_name })}>
                            {props.rowData.id_card_name}
                        </div>,
                    ]
                }
                let currentRef: any
                return [
                    <el-input
                        ref={(el: any) => (el?.focus(), (currentRef = el))}
                        v-model={props.rowData.id_card_name}
                        placeholder={t('users.users.name.id_card_name')}
                        maxlength={30}
                        show-word-limit={true}
                        onBlur={() => {
                            props.rowData.editIdCardName.isEdit = false
                            if (props.rowData.id_card_name == props.rowData.editIdCardName.oldValue) {
                                return
                            }
                            if (!props.rowData.id_card_name) {
                                props.rowData.id_card_name = props.rowData.editIdCardName.oldValue
                                return
                            }
                            handleUpdate(props.rowData.id, { id_card_name: props.rowData.id_card_name }).catch(() => (props.rowData.id_card_name = props.rowData.editIdCardName.oldValue))
                        }}
                        onKeydown={(event: any) => [13].includes(event.keyCode) && currentRef?.blur()} //13：Enter键 27：Esc键 32：空格键
                    />,
                ]
            }, */
        },
        {
            dataKey: 'id_card_gender',
            title: t('users.users.name.id_card_gender'),
            key: 'id_card_gender',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let statusList = tm('users.users.status.id_card_gender') as { value: any; label: string }[]
                let statusIndex = statusList.findIndex((item) => item.value == props.rowData.id_card_gender)
                return <el-tag type={tagType[statusIndex % tagType.length]}>{statusList[statusIndex]?.label}</el-tag>
            },
        },
        {
            dataKey: 'id_card_birthday',
            title: t('users.users.name.id_card_birthday'),
            key: 'id_card_birthday',
            align: 'center',
            width: 100,
            sortable: true,
        },
        {
            dataKey: 'id_card_address',
            title: t('users.users.name.id_card_address'),
            key: 'id_card_address',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'is_stop',
            title: t('users.users.name.is_stop'),
            key: 'is_stop',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                let statusList = tm('common.status.whether') as { value: any; label: string }[]
                return [
                    <el-switch
                        model-value={props.rowData.is_stop}
                        active-value={statusList[1].value}
                        inactive-value={statusList[0].value}
                        active-text={statusList[1].label}
                        inactive-text={statusList[0].label}
                        inline-prompt={true}
                        disabled={!authAction.isUpdate}
                        onChange={(val: any) => handleUpdate(props.rowData.id, { is_stop: val }).then(() => (props.rowData.is_stop = val))}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                    />,
                ]
            },
        },
        {
            dataKey: 'updated_at',
            title: t('common.name.updatedAt'),
            key: 'updated_at',
            align: 'center',
            width: 150,
            sortable: true,
        },
        {
            dataKey: 'created_at',
            title: t('common.name.createdAt'),
            key: 'created_at',
            align: 'center',
            width: 150,
            sortable: true,
        },
        /* {
            title: t('common.name.action'),
            key: 'action',
            align: 'center',
            width: 80 * ((authAction.isCreate ? 1 : 0) + (authAction.isUpdate ? 1 : 0) + (authAction.isDelete ? 1 : 0)),
            fixed: 'right',
            hidden: !(authAction.isCreate || authAction.isUpdate || authAction.isDelete),
            cellRenderer: (props: any): any => {
                let vNode: any = []
                if (authAction.isUpdate) {
                    vNode.push(
                        <el-button type="primary" size="small" onClick={() => handleEditCopy(props.rowData.id)}>
                            <autoicon-ep-edit />
                            {t('common.edit')}
                        </el-button>
                    )
                }
                return vNode
            },
        }, */
    ] as any,
    data: [],
    loading: false,
    sort: { key: 'id', order: 'desc' } as any,
    handleSort: (sort: any) => {
        table.sort.key = sort.key
        table.sort.order = sort.order
        getList()
    },
})

/* const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
//编辑|复制
const handleEditCopy = (id: number, type: string = 'edit') => {
    request(t('config.VITE_HTTP_API_PREFIX') + '/users/users/info', { id: id }).then((res) => {
        saveCommon.data = { ...res.data.info }
        saveCommon.title = t('common.' + type)
        if (type == 'copy') {
            delete saveCommon.data.id
        }
        saveCommon.visible = true
    })
} */
//更新
const handleUpdate = async (id: number | number[], param: { [propName: string]: any }) => {
    param[Array.isArray(id) ? 'id_arr' : 'id'] = id
    await request(t('config.VITE_HTTP_API_PREFIX') + '/users/users/update', param, true)
}

//分页
const settingStore = useSettingStore()
const pagination = reactive({
    total: 0,
    page: 1,
    size: settingStore.pagination.size,
    sizeList: settingStore.pagination.sizeList,
    layout: settingStore.pagination.layout,
    sizeChange: () => getList(),
    pageChange: () => getList(),
})

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
//列表
const getList = async (resetPage: boolean = false) => {
    resetPage && (pagination.page = 1)
    const param = {
        field: [],
        filter: removeEmptyOfObj(queryCommon.data, true, true),
        sort: table.sort.key + ' ' + table.sort.order,
        page: pagination.page,
        limit: pagination.size,
    }
    table.loading = true
    try {
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/users/users/list', param)
        table.data = res.data.list?.length ? res.data.list : []
        pagination.total = res.data.count
    } finally {
        table.loading = false
    }
}
getList()

//暴露组件接口给父组件
defineExpose({
    getList,
})
</script>

<template>
    <el-row class="main-table-tool">
        <el-col :span="16">
            <el-space :size="10" style="height: 100%; margin-left: 10px"></el-space>
        </el-col>
        <el-col :span="8" style="text-align: right">
            <el-space :size="10" style="height: 100%; margin-right: 10px">
                <my-export-button i18nPrefix="users.users" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/users/users/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
                <el-dropdown max-height="300" :hide-on-click="false">
                    <el-button type="info" :circle="true"><autoicon-ep-hide /></el-button>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item v-for="(item, index) in table.columns" :key="index">
                                <el-checkbox v-model="item.hidden">{{ item.title }}</el-checkbox>
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </el-space>
        </el-col>
    </el-row>

    <el-main>
        <el-auto-resizer>
            <template #default="{ height, width }">
                <el-table-v2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort" @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="50">
                    <template v-if="table.loading" #overlay>
                        <el-icon class="is-loading" color="var(--el-color-primary)" :size="25"><autoicon-ep-loading /></el-icon>
                    </template>
                </el-table-v2>
            </template>
        </el-auto-resizer>
    </el-main>

    <el-row class="main-table-pagination">
        <el-col :span="24">
            <el-pagination
                :total="pagination.total"
                v-model:currentPage="pagination.page"
                v-model:page-size="pagination.size"
                @size-change="pagination.sizeChange"
                @current-change="pagination.pageChange"
                :page-sizes="pagination.sizeList"
                :layout="pagination.layout"
                :background="true"
            />
        </el-col>
    </el-row>
</template>
