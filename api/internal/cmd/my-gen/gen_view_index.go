package my_gen

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// 视图模板Index生成
func genViewIndex(option myGenOption, tpl myGenTpl) {
	tplView := `<script setup lang="tsx">
import List from './List.vue'
import Query from './Query.vue'`
	if option.IsCreate || option.IsUpdate {
		tplView += `
import Save from './Save.vue'`
	}

	tplView += `

const { t } = useI18n()`
	if option.IsAuthAction {
		tplView += `
const adminStore = useAdminStore()`
	}
	tplView += `

const authAction: { [propName: string]: boolean } = {`
	isReadStr := `true`
	if option.IsAuthAction {
		isReadStr = `adminStore.IsAction('` + gstr.CaseCamelLower(tpl.LogicStructName) + `Read')`
	}
	tplView += `
    isRead: ` + isReadStr + `,`
	if option.IsCreate {
		isCreateStr := `true`
		if option.IsAuthAction {
			isCreateStr = `adminStore.IsAction('` + gstr.CaseCamelLower(tpl.LogicStructName) + `Create')`
		}
		tplView += `
    isCreate: ` + isCreateStr + `,`
	}
	if option.IsUpdate {
		isUpdateStr := `true`
		if option.IsAuthAction {
			isUpdateStr = `adminStore.IsAction('` + gstr.CaseCamelLower(tpl.LogicStructName) + `Update')`
		}
		tplView += `
    isUpdate: ` + isUpdateStr + `,`
	}
	if option.IsDelete {
		isDeleteStr := `true`
		if option.IsAuthAction {
			isDeleteStr = `adminStore.IsAction('` + gstr.CaseCamelLower(tpl.LogicStructName) + `Delete')`
		}
		tplView += `
    isDelete: ` + isDeleteStr + `,`
	}
	tplView += `
}
provide('authAction', authAction)

//搜索
const queryCommon = reactive({
    data: {},
})
provide('queryCommon', queryCommon)

//列表
const listCommon = reactive({
    ref: null as any,
})
provide('listCommon', listCommon)`
	if option.IsCreate || option.IsUpdate {
		tplView += `

//保存
const saveCommon = reactive({
    visible: false,
    title: '', //新增|编辑|复制
    data: {},
})
provide('saveCommon', saveCommon)`
	}
	tplView += `
</script>

<template>
    <div v-if="!authAction.isRead" style="text-align: center; font-size: 60px; color: #f56c6c">{{ t('common.tip.notAuthActionRead') }}</div>
    <template v-else>
        <el-container class="main-table-container">
            <el-header>
                <query />
            </el-header>

            <list :ref="(el: any) => listCommon.ref = el" />`
	if option.IsCreate || option.IsUpdate {
		tplView += `

            <!-- 加上v-if每次都重新生成组件。可防止不同操作之间的影响；新增操作数据的默认值也能写在save组件内 -->
            <save v-if="saveCommon.visible" />`
	}
	tplView += `
        </el-container>
    </template>
</template>
`

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Index.vue`
	gfile.PutContents(saveFile, tplView)
}
