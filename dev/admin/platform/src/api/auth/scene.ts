import axios from 'axios'
const CancelToken = axios.CancelToken
let cancel: Function | null

export function list(data: any) {
  if (cancel) {
    cancel()
  }
  return http({
    url: '/auth/scene/list',
    method: 'post',
    data: data,
    cancelToken: new CancelToken(function executor(c) {
      cancel = c
    })
  })
}

export function info(data: any) {
  return http({
    url: '/auth/scene/info',
    method: 'post',
    data: data
  })
}

export function save(data: any) {
  if (data?.id > 0) {
    return http({
      url: '/auth/scene/update',
      method: 'post',
      data: data
    })
  } else {
    return http({
      url: '/auth/scene/create',
      method: 'post',
      data: data
    })
  }
}

export function del(data: any) {
  return http({
    url: '/auth/scene/delete',
    method: 'post',
    data: data
  })
}