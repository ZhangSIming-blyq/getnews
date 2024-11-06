<template>
  <div class="news-board">
    <div class="controls">
      <!-- 全局显示条数选择 -->
      <el-select v-model="globalLimit" placeholder="显示条数" @change="fetchAllNews" class="limit-select">
        <el-option label="10条" value="10" />
        <el-option label="20条" value="20" />
        <el-option label="50条" value="50" />
      </el-select>
      <!-- 全局刷新按钮 -->
      <el-button
        class="refresh-btn"
        type="text"
        icon="el-icon-refresh"
        @click="fetchAllNews"
        :loading="isAnyLoading"
      >
        刷新
      </el-button>
    </div>
    <div class="board-grid">
      <el-card
        v-for="source in sources"
        :key="source"
        class="board-container"
      >
        <!-- 数据源标题 -->
        <template #header>
          <h2>{{ source }}</h2>
        </template>

        <!-- 新闻列表 -->
        <el-table v-loading="loading[source]" :data="sortedNewsList(source)" style="width: 100%">
          <el-table-column prop="rank" label="排名" width="80">
            <template #default="{ row }">
              <span :class="rankClass(row.rank)">
                {{ row.rank }}
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
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

// 数据源列表
const sources = ['Weibo', 'Zhihu', 'BiliBili']
const newsData = ref({}) // 存储每个 Source 的新闻数据
const loading = ref({
  Weibo: false,
  Zhihu: false,
  BiliBili: false
}) // 每个 Source 的加载状态
const globalLimit = ref('10') // 全局显示条数

// API 基础 URL
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

// 检查是否有任一加载状态为 true
const isAnyLoading = computed(() => Object.values(loading.value).some(status => status))

// 排序后的新闻列表（按排名排序，排除重复项）
const sortedNewsList = (source) => {
  const uniqueTitles = new Set()
  return (newsData.value[source] || [])
    .sort((a, b) => a.rank - b.rank) // 按排名升序排序
    .filter((item) => {
      if (uniqueTitles.has(item.title)) return false
      uniqueTitles.add(item.title)
      return true
    })
    .slice(0, parseInt(globalLimit.value)) // 限制显示条数
}

// 获取新闻列表
const fetchNews = async (source) => {
  loading.value[source] = true
  try {
    const response = await axios.get(`${API_BASE_URL}/news`, {
      params: { source, limit: globalLimit.value }
    })
    newsData.value[source] = response.data
  } catch (error) {
    ElMessage.error(`获取${source}新闻失败：` + (error.response?.data?.error || error.message))
  } finally {
    loading.value[source] = false
  }
}

// 全局获取所有新闻
const fetchAllNews = () => {
  sources.forEach(source => fetchNews(source))
}

// 排名颜色框样式
const rankClass = (rank) => {
  if (rank === 1) return 'rank-red'
  if (rank === 2) return 'rank-orange'
  if (rank === 3) return 'rank-yellow'
  return 'rank-white'
}

// 组件挂载时获取数据
onMounted(fetchAllNews)
</script>

<style scoped>
.news-board {
  padding: 20px;
  background-color: #f5f5f5;
}

.controls {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-bottom: 20px;
  gap: 12px;
}

.board-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.board-container {
  padding: 16px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
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
