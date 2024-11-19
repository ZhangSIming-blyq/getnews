<template>
  <div class="news-board">
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
        <div class="news-table-wrapper">
          <el-table :data="sortedNewsList(source)" style="width: 100%">
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
        </div>
      </el-card>
    </div>

    <div class="controls">
      <!-- 全局显示条数选择 -->
      <div class="limit-select-wrapper">
        <el-select v-model="globalLimit" placeholder="显示条数" @change="fetchAllNews" class="limit-select">
          <el-option label="10条" value="10" />
          <el-option label="20条" value="20" />
          <el-option label="50条" value="50" />
        </el-select>
      </div>
      <!-- 全局刷新按钮 -->
      <el-button
        class="refresh-btn"
        type="primary"
        icon="el-icon-refresh"
        @click="fetchAllNews"
        :loading="isAnyLoading"
      >
        刷新
      </el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'

const sources = ['Weibo', 'Zhihu', 'BiliBili']
const newsData = ref({})
const loading = ref({ Weibo: false, Zhihu: false, BiliBili: false })
const globalLimit = ref('10')
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'
const isAnyLoading = computed(() => Object.values(loading.value).some(status => status))

const sortedNewsList = (source) => {
  const uniqueTitles = new Set()
  return (newsData.value[source] || [])
    .sort((a, b) => a.rank - b.rank)
    .filter((item) => {
      if (uniqueTitles.has(item.title)) return false
      uniqueTitles.add(item.title)
      return true
    })
    .slice(0, parseInt(globalLimit.value))
}

const fetchNews = async (source) => {
  loading.value[source] = true
  try {
    const response = await axios.get(`${API_BASE_URL}/news`, {
      params: { source, limit: globalLimit.value },
      headers: { 'Cache-Control': 'no-cache', 'Pragma': 'no-cache', 'Expires': '0' }
    })
    newsData.value[source] = response.data
  } catch (error) {
    ElMessage.error(`获取${source}新闻失败：` + (error.response?.data?.error || error.message))
  } finally {
    loading.value[source] = false
  }
}

const fetchAllNews = () => {
  sources.forEach(source => fetchNews(source))
}

const rankClass = (rank) => {
  if (rank === 1) return 'rank-red'
  if (rank === 2) return 'rank-orange'
  if (rank === 3) return 'rank-yellow'
  return 'rank-white'
}

onMounted(fetchAllNews)
</script>

<style scoped>
.news-board {
  padding: 60px 20px 20px; /* 增加顶部间距，与页头拉开距离 */
  background-color: transparent;
  display: flex;
  flex-direction: column;
  align-items: flex-start; /* 左对齐 */
}

.board-grid {
  display: flex;
  flex-wrap: wrap; /* 允许换行 */
  gap: 20px;
  justify-content: flex-start; /* 从左开始排列 */
  width: 100%;
}

.board-container {
  width: 20%; /* 保持固定宽度 */
  height: auto;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
}

.news-table-wrapper {
  flex: 1;
  overflow: hidden; /* 移除滚动条 */
}

.refresh-btn {
  font-size: 16px;
  color: white;
  background-color: #409eff;
  border-radius: 4px;
}

.limit-select {
  width: 100px;
}

/* 排名样式 */
.rank-red, .rank-orange, .rank-yellow, .rank-white {
  display: inline-block;
  padding: 4px 8px;
  color: #fff;
  border-radius: 4px;
  font-weight: bold;
}

.rank-red { background-color: #ff4949; }
.rank-orange { background-color: #ffa726; }
.rank-yellow { background-color: #ffeb3b; color: #000; }
.rank-white { background-color: #fff; color: #000; }

.news-link {
  color: #333;
  text-decoration: none;
  transition: color 0.3s;
}

.news-link:hover {
  color: #333; /* 颜色与默认状态保持一致 */
}
</style>
