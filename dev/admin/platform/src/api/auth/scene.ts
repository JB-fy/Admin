export function list(data: any) {
  return http({
    url: '/auth/scene/list',
    method: 'post',
    data: data
  })
}
