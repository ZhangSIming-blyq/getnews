<template>
  <div class="news-board">
    <el-card class="board-container">
      <!-- 头部显示数据源标题 -->
      <template #header>
        <div class="header">
          <h2>微博热搜</h2>
          <div class="controls">
            <el-select v-model="limit" placeholder="显示条数" @change="fetchNews" class="limit-select">
              <el-option label="10条" value="10" />
              <el-option label="20条" value="20" />
              <el-option label="50条" value="50" />
            </el-select>
            <el-button
              class="refresh-btn"
              type="text"
              icon="el-icon-refresh"
              @click="refreshNews"
              :loading="loading"
            >
              刷新
            </el-button>
          </div>
        </div>
      </template>

      <!-- 新闻列表 -->
      <el-table v-loading="loading" :data="sortedNewsList" style="width: 100%">
        <el-table-column prop="rank" label="排名" width="80">
          <template #default="{ row }">
            <span :class="rankClass(row.rank)">
              {{ row.rank + 1 }}
            </span>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="标题">
          <template #default="{ row }">
            <a :href="row.link" target="_blank" class="news-link">{{ row.title }}</a>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

// 状态定义
const limit = ref('10')
const newsList = ref([])
const loading = ref(false)

// API 基础URL
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

// 排序后的新闻列表（按排名排序，按需要排除重复项）
const sortedNewsList = computed(() => {
  const uniqueTitles = new Set()
  return newsList.value
    .sort((a, b) => a.rank - b.rank)
    .filter((item) => {
      if (uniqueTitles.has(item.title)) return false
      uniqueTitles.add(item.title)
      return true
    })
    .slice(0, parseInt(limit.value))
})

// 获取新闻列表
const fetchNews = async () => {
  loading.value = true
  try {
    const response = await axios.get(`${API_BASE_URL}/news`, {
      params: { source: 'weibo', limit: limit.value }
    })
    newsList.value = response.data
  } catch (error) {
    ElMessage.error('获取新闻失败：' + (error.response?.data?.error || error.message))
  } finally {
    loading.value = false
  }
}

// 手动刷新新闻
const refreshNews = async () => {
  loading.value = true
  try {
    await fetchNews()
    ElMessage.success('刷新成功')
  } catch (error) {
    ElMessage.error('刷新失败：' + (error.response?.data?.error || error.message))
  } finally {
    loading.value = false
  }
}

// 排名颜色框样式
const rankClass = (rank) => {
  if (rank === 0) return 'rank-red'
  if (rank === 1) return 'rank-orange'
  if (rank === 2) return 'rank-yellow'
  return 'rank-white'
}

// 组件挂载时获取数据
onMounted(fetchNews)
</script>

<style scoped>
.news-board {
  padding: 20px;
  background-color: #f5f5f5;
}

.board-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 16px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.controls {
  display: flex;
  align-items: center;
  gap: 12px;
}

.news-link {
  color: #333;
  text-decoration: none;
  transition: color 0.3s;
}

.news-link:hover {
  color: #409EFF;
}

.refresh-btn {
  font-size: 16px;
  color: #409eff;
}

.limit-select {
  width: 100px;
}

.rank-red {
  display: inline-block;
  padding: 4px 8px;
  background-color: #ff4949;
  color: #fff;
  border-radius: 4px;
  font-weight: bold;
}

.rank-orange {
  display: inline-block;
  padding: 4px 8px;
  background-color: #ffa726;
  color: #fff;
  border-radius: 4px;
  font-weight: bold;
}

.rank-yellow {
  display: inline-block;
  padding: 4px 8px;
  background-color: #ffeb3b;
  color: #000;
  border-radius: 4px;
  font-weight: bold;
}

.rank-white {
  display: inline-block;
  padding: 4px 8px;
  background-color: #fff;
  color: #000;
  border-radius: 4px;
  font-weight: bold;
}
</style>
