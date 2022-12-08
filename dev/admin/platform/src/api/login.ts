export function login(data: any) {
  return http({
    url: '/login',
    method: 'post',
    data: data
  })
}