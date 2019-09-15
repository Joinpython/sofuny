<template>
  <div v-if="loading" class="loading-page">
    <a-spin size="large" tip="Loading..." />
  </div>
</template>

<script>
    export default {
        name: "loading",
        data: ()=>({
            loading:false
        }),
        methods: {
            start () {
                this.loading = true
            },
            finish () {
                this.loading = false
            }
        }
    }
</script>

<style scoped>
  .loading-page {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(255, 255, 255, 0.8);
    text-align: center;
    padding-top: 400px;
    font-size: 30px;
    z-index: 9999;
    font-family: sans-serif;
  }
</style>
