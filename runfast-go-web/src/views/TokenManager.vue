<template>
  <div class="token-manager">
    <h1>Token 管理</h1>
    
    <!-- Admin Token 输入框 -->
    <div class="admin-token-section">
      <el-form :inline="true" :model="adminTokenForm" class="demo-form-inline">
        <el-form-item label="Admin Token">
          <el-input 
            v-model="adminTokenForm.token" 
            placeholder="请输入Admin Token" 
            style="width: 300px"
            show-password
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadTokenList">加载Token列表</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 生成新Token -->
    <div class="generate-token-section">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>生成新Token</span>
          </div>
        </template>
        <el-form :inline="true" :model="newTokenForm" class="demo-form-inline">
          <el-form-item label="Token名称">
            <el-input 
              v-model="newTokenForm.name" 
              placeholder="请输入Token名称" 
              style="width: 200px"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="success" @click="generateToken">生成Token</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>

    <!-- Token列表 -->
    <div class="token-list-section">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>Token列表</span>
          </div>
        </template>
        <el-table :data="tokenList" style="width: 100%" v-loading="loading">
          <el-table-column prop="id" label="ID" width="80" />
          <el-table-column prop="token_name" label="Token名称" width="200" />
          <el-table-column prop="token" label="Token值" />
          <el-table-column prop="created_at" label="创建时间" width="200" />
          <el-table-column label="操作" width="120">
            <template #default="scope">
              <el-button
                type="danger"
                size="small"
                @click="deleteToken(scope.row.token)"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>
    </div>
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
  ElMessageBox.confirm('确定要删除这个Token吗？', '确认删除', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    handleDeleteToken(token)
  }).catch(() => {
    // 取消删除
  })
}
</script>

<style scoped>
.token-manager {
  padding: 20px;
}

.admin-token-section {
  margin-bottom: 20px;
}

.generate-token-section {
  margin-bottom: 20px;
}

.token-list-section {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>