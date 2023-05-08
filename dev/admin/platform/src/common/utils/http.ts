import axios from 'axios'

const option = {
    apiSceneName: import.meta.env.VITE_AUTH_SCENE_NAME,
    apiSceneCode: import.meta.env.VITE_AUTH_SCENE_CODE,
    languageName: import.meta.env.VITE_LANGUAGE_NAME,
    accessTokenName: import.meta.env.VITE_ACCESS_TOKEN_NAME,
    baseURL: function () {
        if (import.meta.env.DEV && import.meta.env.VITE_HTTP_HOST.indexOf('http') != 0) {
            return import.meta.env.VITE_DEV_API_PREFIX + import.meta.env.VITE_HTTP_HOST
        }
        return import.meta.env.VITE_HTTP_HOST
    }(),
    timeout: parseInt(import.meta.env.VITE_HTTP_TIMEOUT),
}

const http = axios.create({
    baseURL: option.baseURL,
    timeout: option.timeout
})

http.interceptors.request.use(
    (config: any) => {
        config.headers[option.apiSceneName] = option.apiSceneCode
        config.headers[option.languageName] = getLanguage()
        config.headers[option.accessTokenName] = getAccessToken()
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

http.interceptors.response.use(
    (response) => {
        if (response.data.code === 0) {
            return response.data
        }
        return Promise.reject(new ApiError(JSON.stringify(response.data)))
    },
    (error) => {
        return Promise.reject(error)
    }
)

export default http