export function getLoginToken(data) {
  return http({
    url: '/admin/login/getLoginToken',
    method: 'post',
    data: data
  })
}

export function login(data) {
  return http({
    url: '/admin/login/index',
    method: 'post',
    data: data
  })
}

export function getInfo() {
  return http({
    url: '/admin/login/getInfo',
    method: 'post'
  })
}

export function updateInfo(data) {
  return http({
    url: '/admin/login/updateInfo',
    method: 'post',
    data: data
  })
}

export function getMenuTree() {
  return http({
    url: '/admin/authMenu/getSelfTree',
    method: 'post'
  })
}
