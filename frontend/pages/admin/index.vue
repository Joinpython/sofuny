<template>
  <h2>首页</h2>
</template>

<script>
    import pkg from '../../../config'
    export default {
        layout: "admin",
        name: "index",
        middleware: 'auth',
        head: ()=>({
           title: "后台管理 | " + pkg.name
        }),
        data: ()=>({

        })
    }
</script>

<style scoped>

</style>
