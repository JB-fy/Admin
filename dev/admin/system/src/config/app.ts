export default readonly({
    webTitle: import.meta.env.VITE_WEB_TITLE,   //网站标题
    apiScene: {
        name: import.meta.env.VITE_API_SCENE_NAME,  //场景名称。作用：设置http的请求头
        code: import.meta.env.VITE_API_SCENE_CODE,  //场景标识。与后台代码配合使用，多后台情况下，复制一份前端代码，更改这个参数，可方便开发。
    },
    accessToken: {
        storage: import.meta.env.VITE_ACCESS_TOKEN_STORAGE === 'localStorage' ? localStorage : sessionStorage, //存储方式。直接填写localStorage对象和sessionStorage对象
        name: import.meta.env.VITE_ACCESS_TOKEN_NAME,   //accessToken名称。作用：设置http的请求头；在storage中存储的键名
        activeTimeName: import.meta.env.VITE_ACCESS_TOKEN_ACTIVE_TIME_NAME, //活跃时间名称。作用：在storage中存储的键名
        activeTimeout: parseInt(import.meta.env.VITE_ACCESS_TOKEN_ACTIVE_TIMEOUT),  //失活时间，大于0生效，即验证是否失活（单位：毫秒时间戳。当前时间与活跃时间相差超过该值，判定失活，删除accessToken）
    },
    http: {
        host: import.meta.env.VITE_HTTP_HOST,    //前后端域名一致时可设置为空，这样上线后就不用改
        timeout: parseInt(import.meta.env.VITE_HTTP_TIMEOUT)   //超时时间。0不限制
    }
})