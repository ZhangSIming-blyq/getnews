<template>
    <div class="course-list">
        <h1>All Courses</h1>
        <ul>
            <li v-for="course in courses" :key="course.id">
                <router-link :to="`/course/${course.id}`">{{ course.title }}</router-link>
            </li>
        </ul>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

export default {
    setup() {
        const courses = ref([])

        const fetchCourses = async () => {
            try {
                const response = await axios.get(`${API_BASE_URL}/courses`)
                courses.value = response.data
            } catch (error) {
                console.error('Error fetching courses:', error)
            }
        }

        onMounted(fetchCourses)

        return { courses }
    }
}
</script>

<style scoped>
.course-list {
    padding: 20px;
}
.course-list ul {
    list-style: none;
    padding: 0;
}
.course-list li {
    margin: 10px 0;
}
.course-list a {
    text-decoration: none;
    color: #42b983;
    font-size: 1.2em;
}
.course-list a:hover {
    text-decoration: underline;
}
</style>