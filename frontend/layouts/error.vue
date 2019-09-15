<template>
  <div class="context bruce flex-ct-x">
    <div v-if="error.statusCode === 404">
      <h1 class="fault-text" :data-text="error.statusCode">{{ error.statusCode }}</h1>
      <p>页面不存在</p>
    </div>
    <div v-if="error.statusCode === 500">
      <h1>{{ error.statusCode }}</h1>
      <p>应用发生错误异常</p>
    </div>
    <nuxt-link to="/"> 返回首页 </nuxt-link>
  </div>

</template>

<script>
    import pkg from '../../config'
    export default {
        name: "error",
        props: ['error'],
        head () {
            return {
                title: this.title + " | " + pkg.name
            }
        },
        data: ()=>({
            title: ""
        }),
        mounted() {
            this.title = this.error.statusCode;
        }

    }
</script>

<style scoped>
.context{
  text-align: center;
  margin-top: 15%;
}
.bruce {
  /*background-color: #000;*/
}
</style>
