// 当配置文件用（不用翻译。方便在HTML代码中使用，HTML不支持直接用import.meta.env.XXXX）
export default {
    webTitle: '平台后台',
    VITE_HTTP_API_PREFIX: import.meta.env.VITE_HTTP_API_PREFIX, //不用翻译
    const: {
        tagType: ['', 'success', 'danger', 'info', 'warning'], //不用翻译
    },
}
