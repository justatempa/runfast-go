import request from '@/utils/request'

// API接口类型定义
export interface UserToken {
  id: number
  token_name: string
  token: string
  created_at: string
  updated_at: string
}

// 生成用户Token
export function generateUserToken(tokenName: string, adminToken: string) {
  return request.get('/admin/token/generate', {
    params: { token_name: tokenName },
    headers: { Authorization: `Bearer ${adminToken}` }
  })
}

// 获取Token列表
export function getTokenList(adminToken: string) {
  return request.get('/admin/token/list', {
    headers: { Authorization: `Bearer ${adminToken}` }
  })
}

// 删除Token
export function deleteToken(token: string, adminToken: string) {
  return request.get('/admin/token/remove', {
    params: { token },
    headers: { Authorization: `Bearer ${adminToken}` }
  })
}