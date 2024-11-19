<template>
    <header class="header">
        <div class="header-content">
            <div class="logo">
                <img src="@/assets/logo.png" alt="Logo" />
            </div>
            <nav class="nav-links">
                <ul>
                    <li><router-link to="/">首页</router-link></li>
                    <li><router-link to="/courses">课程</router-link></li>
                    <li><router-link to="/news">新闻</router-link></li>
                    <!-- 添加更多导航项如果需要 -->
                </ul>
            </nav>
        </div>
    </header>
</template>

<script>
export default {
    name: 'Header',
};
</script>

<style scoped>
.header {
    width: 100%;
    height: 10vh; 
    background: rgba(0, 0, 0, 0.7);
    background: linear-gradient(135deg, rgba(0, 0, 0, 0.7), rgba(50, 50, 50, 0.7));
    border-bottom-left-radius: 30px;
    border-bottom-right-radius: 30px;
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.3);
    backdrop-filter: blur(10px);
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    overflow: hidden;
}

.header::before {
    content: '';
    position: absolute;
    top: -50%;
    left: -50%;
    width: 200%;
    height: 200%;
    background: radial-gradient(circle, rgba(255, 255, 255, 0.1), transparent 70%);
    transform: rotate(45deg);
}

.header-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 90%;
    max-width: 1200px;
    z-index: 1;
}

.logo {
    margin-right: auto; /* 将 logo 放在最左侧 */
}

.logo img {
    height: 80px;
    border-radius: 20px;
}

.nav-links {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-grow: 1;
}

.nav-links ul {
    list-style: none;
    display: flex;
    gap: 40px;
    margin: 0;
    padding: 0;
}

.nav-links a {
    color: #ffffff;
    text-decoration: none;
    font-weight: 600;
    position: relative;
    transition: color 0.3s;
    font-size: 20px;
}

.nav-links a::after {
    content: '';
    display: block;
    width: 0%;
    height: 3px;
    background: #00d1b2;
    transition: width 0.3s;
    position: absolute;
    bottom: -5px;
    left: 0;
}

.nav-links a:hover {
    color: #00d1b2;
}

.nav-links a:hover::after {
    width: 100%;
}

/* 响应式设计 */
@media (max-width: 1024px) {
    .header-content {
        flex-direction: column;
        gap: 20px;
    }

    .nav-links ul {
        gap: 30px;
    }

    .nav-links a {
        font-size: 18px;
    }
}

@media (max-width: 768px) {
    .header {
        height: auto; /* 让高度自适应内容 */
        padding: 20px 30px;
        flex-direction: column;
        align-items: flex-start;
    }

    .header-content {
        flex-direction: column;
        align-items: center;
        width: 100%;
        height: auto; /* 确保高度自适应 */
    }

    .logo {
        margin-bottom: 10px;
    }

    .nav-links ul {
        flex-direction: column;
        gap: 20px;
        align-items: center;
    }

    .nav-links a {
        font-size: 16px;
    }
}
</style>