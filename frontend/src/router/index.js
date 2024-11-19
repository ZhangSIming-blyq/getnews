import { createRouter, createWebHistory } from 'vue-router'
import News from '../views/News.vue'
import LandingPage from '../views/LandingPage.vue'
import Course from '../views/Course.vue'
import Article from '../views/Article.vue'
import CourseDetail from '../views/CourseDetail.vue'

const routes = [
  {
    path: '/news',
    name: 'News',
    component: News,
  },
  {
    path: '/',
    name: 'LandingPage',
    component: LandingPage,
  },
  {
    path: '/courses',
    name: 'Course',
    component: Course,
  },
  {
    path: '/article/:id',
    name: 'Article',
    component: Article,
    props: true
  },
  {
    path: '/course/:id',
    name: 'CourseDetail',
    component: CourseDetail,
    props: true
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 