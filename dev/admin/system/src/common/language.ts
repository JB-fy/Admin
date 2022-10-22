const storage = import.meta.env.VITE_LANGUAGE_STORAGE === 'localStorage' ? localStorage : sessionStorage
const languageNAME = import.meta.env.VITE_LANGUAGE_NAME

//获取语言
export const getLanguage = (): string => {
    const language = storage.getItem(languageNAME)
    if (language) {
        return language
    }
    //不存在，则根据当前浏览器环境设置相应语言或默认简体中文
    return (navigator.language || 'zh-cn').toLowerCase()
}

//设置语言
export const setLanguage = (language: string) => {
    storage.setItem(languageNAME, language)
}