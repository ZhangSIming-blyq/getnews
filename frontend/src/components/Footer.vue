<template>
    <footer class="footer">
      <div class="footer-content">
        <div class="footer-logo">
          <img src="@/assets/logo.png" alt="Logo" />
        </div>
        <div class="footer-links">
          <ul>
            <li><a href="#about">关于我们</a></li>
            <li><a href="#courses">课程</a></li>
            <li><a href="#news">新闻</a></li>
            <li><a href="#contact">联系我们</a></li>
          </ul>
        </div>
        <div class="footer-info">
          <p>&copy; {{ currentYear }} Your Company Name. 版权所有.</p>
        </div>
      </div>
    </footer>
  </template>
  
  <script>
  export default {
    name: 'Footer',
    data() {
      return {
        currentYear: new Date().getFullYear()
      }
    }
  }
  </script>
  
  <style scoped>
  .footer {
    width: 100%;
    height: 10vh;
    background: linear-gradient(135deg, rgba(0, 0, 0, 0.7), rgba(50, 50, 50, 0.7));
    border-top-left-radius: 30px;
    border-top-right-radius: 30px;
    box-shadow: 0 -8px 16px rgba(0, 0, 0, 0.3);
    backdrop-filter: blur(10px);
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    overflow: hidden;
  }
  
  .footer::before {
    content: '';
    position: absolute;
    top: -50%;
    left: -50%;
    width: 200%;
    height: 200%;
    background: radial-gradient(circle, rgba(255, 255, 255, 0.1), transparent 70%);
    transform: rotate(45deg);
  }
  
  .footer-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 90%;
    max-width: 1200px;
    z-index: 1;
  }
  
  .footer-logo img {
    width: 60px;
    border-radius: 10px;
  }
  
  .footer-links ul {
    list-style: none;
    display: flex;
    gap: 20px;
    margin: 0;
    padding: 0;
  }
  
  .footer-links a {
    color: #ffffff;
    text-decoration: none;
    font-weight: 500;
    position: relative;
    transition: color 0.3s;
  }
  
  .footer-links a::after {
    content: '';
    display: block;
    width: 0%;
    height: 2px;
    background: #00d1b2;
    transition: width 0.3s;
    position: absolute;
    bottom: -5px;
    left: 0;
  }
  
  .footer-links a:hover {
    color: #00d1b2;
  }
  
  .footer-links a:hover::after {
    width: 100%;
  }
  
  .footer-info p {
    color: #cccccc;
    font-size: 14px;
  }
  
  @media (max-width: 1024px) {
    .footer-content {
      flex-direction: column;
      gap: 20px;
    }
  
    .footer-links ul {
      gap: 15px;
    }
  
    .footer-info p {
      font-size: 13px;
    }
  }
  
  @media (max-width: 768px) {
    .footer {
      height: auto;
      padding: 20px 30px;
    }
  
    .footer-content {
      flex-direction: column;
      align-items: center;
      width: 100%;
      gap: 15px;
    }
  
    .footer-links ul {
      flex-direction: column;
      gap: 10px;
    }
  
    .footer-logo img {
      width: 50px;
    }
  
    .footer-info p {
      font-size: 12px;
    }
  }
  </style>
  