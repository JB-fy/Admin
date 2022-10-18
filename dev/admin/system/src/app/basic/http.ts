import axios from 'axios'
import { getLanguage } from '@/i18n'

const defaultOption = {
    apiSceneName: config('app.apiScene.name'),
    languageName: import.meta.env.VITE_HTTP_LANGUAGE_NAME,
    apiSceneCode: config('app.apiScene.code'),
    accessTokenName: config('app.accessToken.name'),
    baseURL: config('app.http.host'),
    timeout: config('app.http.timeout'),
}

export const getHttp = (option = {}) => {
    option = Object.assign({}, defaultOption, option)
    const http = axios.create({
        baseURL: option.baseURL,
        timeout: option.timeout
    })

    http.interceptors.request.use(
        (config) => {
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

    return http
}

const http = getHttp()

export default http