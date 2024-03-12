package my_gen

import (
	"context"

	"github.com/gogf/gf/v2/os/gfile"
)

// 视图模板Index生成
func genViewIndex(ctx context.Context, option myGenOption, tpl myGenTpl) {
	tplView := `<script setup lang="tsx">
import List from './List.vue'
import Query from './Query.vue'`
	if option.IsCreate || option.IsUpdate {
		tplView += `
import Save from './Save.vue'`
	}
	tplView += `

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
`

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Index.vue`
	gfile.PutContents(saveFile, tplView)
}
