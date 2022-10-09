import { createStore } from "vuex";

const store = createStore({
    /*--------各模块内必须定义（namespaced: true），否则不同模块有相同命名的会冲突 开始--------*/
    modules: {
        ...getImportComponents(require.context('@/app/basic/store/modules', false, /\.js$/))
    },
    /*--------各模块内必须定义（namespaced: true），否则不同模块有相同命名的会冲突 结束--------*/
    /*--------这里也可以把一些常用的放最外层（因分模块了，故不建议做这一步） 开始--------*/
    /* state:{},
    getters:{},
    mutations:{},
    actions:{} */
    /*--------这里也可以把一些常用的放最外层（因分模块了，故不建议做这一步） 结束--------*/
})

export default store

/*--------使用方式 开始--------*/
/* const store = useStore()
store.state.setting.leftMenuFold
store.getters['setting/leftMenuWidth']
store.commit('setting/leftMenuFold')
store.dispatch('setting/leftMenuFold') */
/*--------使用方式 结束--------*/