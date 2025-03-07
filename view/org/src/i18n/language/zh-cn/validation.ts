export default {
    between: {
        string: '必须在 {min} 到 {max} 个字符之间',
        number: '必须在 {min} 到 {max} 之间',
        select: '必须选择 {min} 到 {max} 项',
        upload: '必须上传 {min} 到 {max} 个文件',
        array: '必须在 {min} 到 {max} 个元素之间',
    },
    min: {
        string: '必须大于等于 {min} 个字符',
        number: '必须大于等于 {min}',
        select: '最少选择 {min} 项',
        upload: '最少上传 {min} 个文件',
        array: '必须大于等于 {min} 个元素',
    },
    max: {
        string: '必须小于等于 {max} 个字符',
        number: '必须小于等于 {max}',
        select: '最多选择 {max} 项',
        upload: '最多上传 {max} 个文件',
        array: '必须小于等于 {max} 个元素',
    },
    size: {
        string: '必须等于 {size} 个字符',
        number: '必须等于 {size}',
        select: '必须选 {size} 项',
        upload: '必须上传 {size} 个文件',
        array: '必须等于 {size} 个元素',
    },
    number: '必须是数字',
    array: '必须是数组',
    required: '必须的',
    input: '请输入',
    select: '请选择',
    upload: '请上传',
    account: '非数字开头，只能包含中文/英文，数字和下划线，至少 4 个字符',
    phone: '必须是有效的手机号码',
    email: '必须是有效的EMAIL格式',
    url: '必须是有效的URL格式',
    ip: '必须是有效的IP格式',
    json: '必须是有效的JSON格式',
    regex: '格式是无效的',
    alpha: '只能包含字母',
    alpha_dash: '只能包含字母、数字、中划线或下划线',
    alpha_num: '只能包含字母和数字',
    repeat_password: '两次密码不一致',
    new_password_diff_old_password: '新密码必须与旧密码不同',
}
