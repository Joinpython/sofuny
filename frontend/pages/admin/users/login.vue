<template>
<a-row :gutter="40">
  <a-col :xs="24" :sm="24" :md="8" :xl="18" :xxl="8">
    <img class="image"
         :src="staticApi+`/bg/login.bg.jpg`" alt="">
  </a-col>
  <a-col :xs="24" :sm="24" :md="8" :xl="6" :xxl="8">
      <a-card :bordered="false"
              class="child">
        <h3 class="title" slot="title">
          <span @click="()=>{ this.$router.push({ name: 'index' })}"> {{ webName }}</span> |
          登录
        </h3>
        <a-form-item label="用户名">
          <a-input v-model="user.name"/>
        </a-form-item>
        <a-form-item label="密码">
          <a-input type="password" v-model="user.password"/>
        </a-form-item>
        <a-form-item>
          <div style="text-align: center">
            <a-button type="primary" @click="login">登录</a-button>
          </div>
        </a-form-item>
      </a-card>
  </a-col>
</a-row>
</template>

<script>
    import pkg from '../../../../config'
    const url = pkg.api + "/api/v1/user/login/";
    export default {
        name: "login",
        head: ()=>({
            title: "后台登录 | " + pkg.name
        }),
        data: ()=>({
            webName: pkg.name,
            api: pkg.api,
            staticApi: "https://upyun.sofuny.cn",
            user: {
                name:"",
                password:""
            }
        }),
        methods: {
            login() {
                if (!this.user.name || !this.user.password){
                    this.$message.error("登录名或者密码不允许为空!")
                }else {
                    this.$axios.$post(url, this.user)
                        .then((res)=>{
                            if (res.statusCode === 200) {
                                this.$store.dispatch("acLogin", res.data);
                                this.$message.success(res.msg);
                                setTimeout(()=>{
                                    this.$router.push({ name: "admin"})
                                }, 1000)
                            }else {
                                this.$message.error(res.msg)
                            }
                        }).catch((error)=>{
                            console.log(error)
                    })
                }
            }
        }
    }
</script>

<style scoped>
  .image {
    width:100%;			/* ... */
    height:100%;		/* ... */
    display:block;	/* ... */

  }
  .child {
    margin-top: 60%;
  }
  @media screen and (max-width: 1024px){
    .child {
      margin-top: -10%;
    }
  }
</style>
