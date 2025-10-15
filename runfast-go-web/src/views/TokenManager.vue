<template>
  <div class="token-manager">
    <el-row :gutter="20">
      <el-col :span="24">
        <el-card class="admin-token-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span><i class="el-icon-key"></i> 管理员认证</span>
            </div>
          </template>
          <el-form :inline="true" :model="adminTokenForm" class="admin-token-form">
            <el-form-item label="Admin Token">
              <el-input 
                v-model="adminTokenForm.token" 
                placeholder="请输入Admin Token" 
                show-password
                style="width: 300px"
              />
            </el-form-item>
            <el-form-item>
              <el-button 
                type="primary" 
                @click="loadTokenList"
                :loading="loading"
              >
                <i class="el-icon-refresh"></i> 加载Token列表
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="content-row">
      <el-col :span="24">
        <el-card class="generate-token-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span><i class="el-icon-plus"></i> 生成新Token</span>
            </div>
          </template>
          <el-form :inline="true" :model="newTokenForm" class="generate-token-form">
            <el-form-item label="Token名称">
              <el-input 
                v-model="newTokenForm.name" 
                placeholder="请输入Token名称" 
                style="width: 200px"
              />
            </el-form-item>
            <el-form-item>
              <el-button 
                type="success" 
                @click="generateToken"
                :disabled="!adminTokenForm.token"
              >
                <i class="el-icon-circle-plus"></i> 生成Token
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="content-row">
      <el-col :span="24">
        <el-card class="token-list-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span><i class="el-icon-tickets"></i> Token列表</span>
              <el-badge :value="tokenList.length" type="primary" v-if="tokenList.length > 0">
                总数
              </el-badge>
            </div>
          </template>
          <el-table 
            :data="tokenList" 
            style="width: 100%" 
            v-loading="loading"
            stripe
            border
          >
            <el-table-column prop="id" label="ID" width="80" align="center" />
            <el-table-column prop="token_name" label="Token名称" width="200" />
            <el-table-column prop="token" label="Token值">
              <template #default="scope">
                <el-tooltip :content="scope.row.token" placement="top">
                  <span class="token-value">{{ scope.row.token.substring(0, 30) }}...</span>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column prop="created_at" label="创建时间" width="200" align="center">
              <template #default="scope">
                {{ formatDate(scope.row.created_at) }}
              </template>
            </el-table-column>
            <el-table-column label="操作" width="150" align="center" fixed="right">
              <template #default="scope">
                <el-button
                  type="primary"
                  size="small"
                  @click="copyToken(scope.row.token)"
                  plain
                >
                  <i class="el-icon-document-copy"></i> 复制
                </el-button>
                <el-button
                  type="danger"
                  size="small"
                  @click="deleteToken(scope.row.token)"
                  plain
                >
                  <i class="el-icon-delete"></i> 删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
          <div class="empty-table" v-if="tokenList.length === 0 && !loading">
            <el-empty description="暂无Token数据，请先生成Token"></el-empty>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { generateUserToken, getTokenList, deleteToken } from '@/api/auth'
import type { UserToken } from '@/api/auth'

// Admin Token表单
const adminTokenForm = reactive({
  token: ''
})

// 新Token表单
const newTokenForm = reactive({
  name: ''
})

// Token列表
const tokenList = ref<UserToken[]>([])

// 加载状态
const loading = ref(false)

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString('zh-CN')
}

// 复制Token到剪贴板
const copyToken = (token: string) => {
  navigator.clipboard.writeText(token).then(() => {
    ElMessage.success('Token已复制到剪贴板')
  }).catch(() => {
    ElMessage.error('复制失败')
  })
}

// 加载Token列表
const loadTokenList = async () => {
  if (!adminTokenForm.token) {
    ElMessage.error('请输入Admin Token')
    return
  }

  try {
    loading.value = true
    const response: any = await getTokenList(adminTokenForm.token)
    
    if (response.code === 0) {
      tokenList.value = response.data || []
      ElMessage.success('Token列表加载成功')
    } else {
      ElMessage.error(response.msg || '加载Token列表失败')
    }
  } catch (error) {
    ElMessage.error('加载Token列表失败')
  } finally {
    loading.value = false
  }
}

// 生成新Token
const generateToken = async () => {
  if (!adminTokenForm.token) {
    ElMessage.error('请输入Admin Token')
    return
  }

  if (!newTokenForm.name) {
    ElMessage.error('请输入Token名称')
    return
  }

  try {
    const response: any = await generateUserToken(newTokenForm.name, adminTokenForm.token)
    
    if (response.code === 0) {
      ElMessage.success('Token生成成功')
      newTokenForm.name = ''
      // 重新加载Token列表
      loadTokenList()
    } else {
      ElMessage.error(response.msg || '生成Token失败')
    }
  } catch (error) {
    ElMessage.error('生成Token失败')
  }
}

// 删除Token
const handleDeleteToken = async (token: string) => {
  try {
    const response: any = await deleteToken(token, adminTokenForm.token)
    
    if (response.code === 0) {
      ElMessage.success('Token删除成功')
      // 重新加载Token列表
      loadTokenList()
    } else {
      ElMessage.error(response.msg || '删除Token失败')
    }
  } catch (error) {
    ElMessage.error('删除Token失败')
  }
}

// 删除Token确认
const deleteToken = (token: string) => {
  ElMessageBox.confirm(
    '确定要删除这个Token吗？此操作不可恢复！',
    '确认删除',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }
  ).then(() => {
    handleDeleteToken(token)
  }).catch(() => {
    // 取消删除
  })
}
</script>

<style scoped>
.token-manager {
  padding: 20px;
  background-color: #f5f7fa;
  min-height: 100%;
}

.content-row {
  margin-top: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-weight: bold;
}

.admin-token-card {
  border-left: 4px solid #409eff;
}

.generate-token-card {
  border-left: 4px solid #67c23a;
}

.token-list-card {
  border-left: 4px solid #e6a23c;
}

.admin-token-form,
.generate-token-form {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
}

.token-value {
  font-family: monospace;
  font-size: 12px;
}

.empty-table {
  padding: 40px 0;
}
</style>