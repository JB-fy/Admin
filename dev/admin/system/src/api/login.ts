export function getEncryptStr(data: any) {
  return http({
    url: '/login/getEncryptStr',
    method: 'post',
    data: data
  })
}

export function login(data: any) {
  return http({
    url: '/login',
    method: 'post',
    data: data
  })
}

export function getInfo() {
  return http({
    url: '/login/getInfo',
    method: 'post'
  })
}

export function updateInfo(data: any) {
  return http({
    url: '/login/updateInfo',
    method: 'post',
    data: data
  })
}

export function getMenuTree() {
  return http({
    url: '/login/getMenuTree',
    method: 'post'
  })
}
