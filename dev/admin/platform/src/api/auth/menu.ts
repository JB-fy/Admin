export function list(data: any) {
  return http({
    url: '/auth/menu/list',
    method: 'post',
    data: data
  })
}

export function info(data: any) {
  return http({
    url: '/auth/menu/info',
    method: 'post',
    data: data
  })
}

export function save(data: any) {
  if (data?.id > 0) {
    return http({
      url: '/auth/menu/update',
      method: 'post',
      data: data
    })
  } else {
    return http({
      url: '/auth/menu/create',
      method: 'post',
      data: data
    })
  }
}

export function del(data: any) {
  return http({
    url: '/auth/menu/delete',
    method: 'post',
    data: data
  })
}