<template>
<a-row>
  <a-col>
    <a-row :gutter="20">
      <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8">
        <a-form-item label="网站名">
          <a-input v-model="setting.name"/>
        </a-form-item>
      </a-col>
      <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8">
        <a-form-item label="作者">
          <a-input v-model="setting.author"/>
        </a-form-item>
      </a-col>
      <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8">
        <a-form-item label="Api">
          <a-input v-model="setting.api"/>
        </a-form-item>
      </a-col>
    </a-row>

    <a-form-item label="分类">
      <a-input/>
    </a-form-item>

    <a-row :gutter="20">
      <a-col :xs="24" :sm="24" :md="12" :xl="12" :xxl="12">
        <a-form-item label="后台登录图片">
          <img :src="staticApi+`/bg/login.bg.jpg`"
               style="width: 100%;height: 200px;"
               alt="">
        </a-form-item>
      </a-col>
      <a-col :xs="24" :sm="24" :md="12" :xl="12" :xxl="12">
        <a-form-item label="首页图片">
          <img :src="staticApi+`/bg/index.bg.jpg`"
               style="width: 100%;height: 200px;"
               alt="">
        </a-form-item>

      </a-col>
    </a-row>

    <a-form-item>
      <div style="text-align: center;">
        <a-button type="primary">修改</a-button>
      </div>
    </a-form-item>
  </a-col>
</a-row>
</template>

<script>
    import pkg from '../../../config'
    const url = pkg.api + "/api/v1/setting/";
    export default {
        layout: "admin",
        name: "setting",
        middleware: "auth",
        head: ()=>({
            title: "基本设置 | " + pkg.name
        }),
        async asyncData({ app, store }) {
            let setting;
            const data = await app.$axios.$get(url);
            if (data.statusCode !== 200) {
                setting = {}
            }
            setting = data.data;
            return { setting }
        },
        data: ()=>({
              staticApi: "https://upyun.sofuny.cn"
        }),
        methods: {

        }
    }
</script>

<style scoped>

</style>
