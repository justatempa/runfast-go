<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const goToTokenManager = () => {
  router.push('/token-manager')
}
</script>

<template>
  <div class="home-container">
    <h1>欢迎使用 RunFast Go 管理系统</h1>
    <p>这是一个基于 Go 和 Vue.js 构建的管理系统。</p>
    
    <div class="quick-links">
      <h2>快速导航</h2>
      <el-button type="primary" @click="goToTokenManager">Token 管理</el-button>
    </div>
  </div>
</template>

<style scoped>
.home-container {
  padding: 20px;
}

.quick-links {
  margin-top: 30px;
}

.quick-links h2 {
  margin-bottom: 15px;
}
</style>
