import { createRouter, createWebHistory } from 'vue-router'
import NewsBoard from '../views/NewsBoard.vue'

const routes = [
  {
    path: '/news-board',
    name: 'NewsBoard',
    component: NewsBoard,
    meta: {
      title: '新闻看板'
    }
  }
  // ... 其他路由
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 