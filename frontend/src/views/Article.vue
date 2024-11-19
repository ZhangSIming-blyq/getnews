<template>
    <div class="article">
        <h2>{{ article.title }}</h2>
        <p>{{ article.content }}</p>
    </div>
</template>

<script>
export default {
    props: ['article']
}
</script>

<style scoped>
/* Add styles for the article page */
</style>