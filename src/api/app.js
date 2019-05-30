import request from '@/utils/request'

export function sendSlackInvite(data) {
  return request({
    url: '/invite',
    method: 'post',
    data
  })
}

export function getSiteSettings() {
  return request({
    url: '/settings',
    method: 'get'
  })
}
