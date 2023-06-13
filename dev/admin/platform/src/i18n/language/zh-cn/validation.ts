export default {
    between: {
        string: '必须在 {min} 到 {max} 个字符之间',
        number: '必须在 {min} 到 {max} 之间'
    },
    min: {
        string: '必须大于等于 {min} 个字符',
        number: '必须大于等于 {min}',
        upload: '最少上传 {min} 个文件'
    },
    max: {
        string: '必须小于等于 {max} 个字符',
        number: '必须小于等于 {max}',
        upload: '最多上传 {max} 个文件'
    },
    regex: '格式是无效的',
    json: '必须是有效的JSON格式',
    url: '必须是有效的URL格式',
    email: '必须是有效的EMAIL格式',
    select: '请选择',
    upload: '请上传',
    //account: '必须包含以下有效字符 (中文/英文，数字)',
    account: '不能是纯数字',
    phone: '必须是有效的手机号',
    alpha: '只能包含字母',
    alpha_dash: '只能包含字母、数字、中划线或下划线',
    alpha_num: '只能包含字母和数字',
    repeatPassword: '两次密码不一致',
    newPasswordDiffOldPassword: '新密码必须与旧密码不同',
}