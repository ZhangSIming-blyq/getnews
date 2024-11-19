<template>
    <div class="page-container">
      <header class="header">
        <!-- 页眉内容 -->
        <div class="header-content">
          <div class="logo">
            <img src="@/assets/logo.png" alt="Logo" />
          </div>
          <nav class="nav-links">
            <ul>
              <li><router-link to="/">首页</router-link></li>
              <li><router-link to="/courses">课程</router-link></li>
              <li><router-link to="/news">新闻</router-link></li>
            </ul>
          </nav>
        </div>
      </header>
  
      <div class="hero-carousel">
        <el-carousel :interval="5000" arrow="always" type="card">
          <el-carousel-item v-for="(slide, index) in slides" :key="index">
            <div class="slide">
              <img :src="slide.image" :alt="slide.alt" />
              <div class="overlay">
                <h2>{{ slide.title }}</h2>
                <p>{{ slide.description }}</p>
              </div>
            </div>
          </el-carousel-item>
        </el-carousel>
      </div>
  
      <footer class="footer">
        <div class="footer-content">
          <p>&copy; {{ new Date().getFullYear() }} Your Company Name. 版权所有.</p>
        </div>
      </footer>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Hero',
    data() {
      return {
        slides: [
          { image: '@/assets/slide1.jpg', alt: 'Slide 1', title: '欢迎来到我们的平台', description: '探索最新的课程和资源，提升您的技能。' },
          { image: '@/assets/slide2.jpg', alt: 'Slide 2', title: '互动学习体验', description: '通过互动课程，与专家和同学一起学习。' },
          { image: '@/assets/slide3.jpg', alt: 'Slide 3', title: '个性化课程推荐', description: '根据您的兴趣和需求，推荐最适合您的课程。' },
          { image: '@/assets/slide4.jpg', alt: 'Slide 4', title: '持续更新内容', description: '我们不断更新内容，确保您获取最新的信息。' },
        ],
      };
    },
  };
  </script>
  
  <style scoped>
  .page-container {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }
  
  .header {
    height: 10vh; /* 页眉高度 */
    background: rgba(0, 0, 0, 0.7);
  }
  
  .hero-carousel {
    flex: 1; /* 自动填满中间区域 */
    display: flex;
  }
  
  .el-carousel {
    width: 100%;
    height: 100%;
  }
  
  .slide {
    position: relative;
    width: 100%;
    height: 100%;
  }
  
  .slide img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  
  .overlay {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    color: #ffffff;
    text-align: center;
    background: rgba(0, 0, 0, 0.4);
    padding: 20px 40px;
    border-radius: 10px;
  }
  
  .overlay h2 {
    font-size: 2.5rem;
    margin-bottom: 10px;
  }
  
  .overlay p {
    font-size: 1.2rem;
  }
  
  .footer {
    height: 10vh; /* 页脚高度 */
    background: rgba(0, 0, 0, 0.7);
    color: white;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  .footer-content p {
    font-size: 0.9rem;
  }
  
  /* 响应式设计 */
  @media (max-width: 768px) {
    .overlay h2 {
      font-size: 1.8rem;
    }
  
    .overlay p {
      font-size: 1rem;
    }
  }
  </style>
  