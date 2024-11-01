<template>
  <el-container>
    <!-- 你的侧边栏导航 -->
    <el-aside width="200px">
      <el-menu router>
        <el-menu-item index="/news-board">
          <el-icon><Document /></el-icon>
          <span>新闻看板</span>
        </el-menu-item>
        <!-- 其他菜单项 -->
      </el-menu>
    </el-aside>

    <!-- 主内容区 -->
    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template> 