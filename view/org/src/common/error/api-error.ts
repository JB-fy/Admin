//自定义错误类
export class ApiError extends Error {
    constructor(message: string) {
        super(message)
        this.name = 'ApiError'
    }
}
