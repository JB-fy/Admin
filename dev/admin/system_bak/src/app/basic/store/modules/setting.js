export default {
    namespaced: true,
    state: {
        leftMenuFold: false,
    },
    mutations: {
        leftMenuFold: (state) => {
            state.leftMenuFold = !state.leftMenuFold
        }
    },
    actions: {
        /**
         * 折叠左侧菜单
         */
        leftMenuFold: ({ commit }) => {
            commit('leftMenuFold')
        }
    }
}