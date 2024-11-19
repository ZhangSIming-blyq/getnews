<template>
    <div class="course-detail">
        <h1>{{ course.title }}</h1>
        <p>{{ course.description }}</p>
        <h2>Articles</h2>
        <ul>
            <li v-for="article in course.articles" :key="article.id">
                <router-link :to="`/article/${article.id}`">{{ article.title }}</router-link>
            </li>
        </ul>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useRoute } from 'vue-router'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

export default {
    setup() {
        const course = ref({})
        const route = useRoute()
        const courseId = route.params.id

        const fetchCourse = async () => {
            try {
                const response = await axios.get(`${API_BASE_URL}/courses/${courseId}`)
                course.value = response.data
            } catch (error) {
                console.error('Error fetching course:', error)
            }
        }

        onMounted(fetchCourse)

        return { course }
    }
}
</script>

<style scoped>
.course-detail {
    padding: 20px;
}

.course-detail h1 {
    color: #333;
}

.course-detail ul {
    list-style: none;
    padding: 0;
}

.course-detail li {
    margin: 10px 0;
}

.course-detail a {
    text-decoration: none;
    color: #42b983;
}

.course-detail a:hover {
    text-decoration: underline;
}
</style>