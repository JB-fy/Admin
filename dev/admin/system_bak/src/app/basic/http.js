import axios from 'axios'

const defaultOption = {
    apiSceneName: config('app.apiScene.name'),
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

            let accessToken = getAccessToken()
            if (accessToken) {
                config.headers[option.accessTokenName] = accessToken
            }
            return config
        },
        (error) => {
            error.message = JSON.stringify({ code: 9999, msg: error.message, data: {} })
            return Promise.reject(error)
        }
    )

    http.interceptors.response.use(
        (response) => {
            if (response.data.code === 0) {
                return response.data
            }
            return Promise.reject(new Error(JSON.stringify(response.data)))
        },
        (error) => {
            error.message = JSON.stringify({ code: 9999, msg: error.message, data: {} })
            return Promise.reject(error)
        }
    )

    return http
}

const http = getHttp()

export default http