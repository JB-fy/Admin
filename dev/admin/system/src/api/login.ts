export function encryptStr(data: any) {
  return http({
    url: '/login/encryptStr',
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

export function info() {
  return http({
    url: '/login/info',
    method: 'post'
  })
}

export function update(data: any) {
  return http({
    url: '/login/update',
    method: 'post',
    data: data
  })
}

export function menuTree() {
  return http({
    url: '/login/menuTree',
    method: 'post'
  })
}
