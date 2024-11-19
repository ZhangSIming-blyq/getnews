<template>
  <el-container direction="vertical" class="app-container">
    <!-- 页眉 -->
    <el-header class="app-header">
      <Header />
    </el-header>

    <!-- 主内容区 -->
    <el-main class="app-main">
      <router-view />
    </el-main>

    <!-- 页脚 -->
    <el-footer class="app-footer">
      <Footer />
    </el-footer>
  </el-container>
</template>

<script>
import Header from '@/components/Header.vue';
import Footer from '@/components/Footer.vue';

export default {
  name: 'App',
  components: {
    Header,
    Footer
  }
}
</script>

<style>
body,
#app {
  margin: 0;
  padding: 0;
  height: 100%;
  background-color: #f0f4f8;
  /* 设置整体背景色为浅色调 */
  font-family: 'Roboto', sans-serif;
}

.app-container {
  min-height: 100vh;
  background-color: transparent;
  /* 使容器背景透明，显示body的背景色 */
}

.app-header,
.app-footer {
  background: transparent;
  /* 页眉和页脚的背景由各自组件控制 */
}

.app-main {
  background-color: #f0f4f8;
  /* 主内容区与body保持一致的背景色 */
}
</style>