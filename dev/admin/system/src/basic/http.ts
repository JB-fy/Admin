import axios from 'axios'
import { getLanguage } from '@/i18n'

const option = {
    apiSceneName: import.meta.env.VITE_API_SCENE_NAME,
    apiSceneCode: import.meta.env.VITE_API_SCENE_CODE,
    languageName: import.meta.env.VITE_HTTP_LANGUAGE_NAME,
    accessTokenName: import.meta.env.VITE_ACCESS_TOKEN_NAME,
    baseURL: import.meta.env.VITE_HTTP_HOST,
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
        error.message = JSON.stringify({ code: '999999', msg: error.message, data: {} })
        return Promise.reject(error)
    }
)

http.interceptors.response.use(
    (response) => {
        if (response.data.code === '000000') {
            return response.data
        }
        return Promise.reject(new Error(JSON.stringify(response.data)))
    },
    (error) => {
        error.message = JSON.stringify({ code: '999999', msg: error.message, data: {} })
        return Promise.reject(error)
    }
)

export default http