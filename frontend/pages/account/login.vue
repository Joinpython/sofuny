<template>
  <div>
    <Header/>
    <a-row :gutter="20">
      <a-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8" :xxl="8"></a-col>
      <a-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8" :xxl="8">
        <a-form-item label="支持以下方式登录">
          <a-icon class="icon" type="github" title="github" @click="getAuth(`github`)"/>
        </a-form-item>
      </a-col>
      <a-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8" :xxl="8"></a-col>
    </a-row>
    <Footer/>
  </div>
</template>

<script>
    import pkg from '../../../config'
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'

    const auth_url = pkg.api +"/api/v1/auth/";

    export default {
        name: "login",
        components: { Footer, Header },
        head: ()=>({
            title: "登录 | " + pkg.name
        }),
        data: ()=>({

        }),
        methods: {
            getAuth(value) {
                if (value === "github") {
                    this.$axios.post(auth_url, { type: value})
                        .then((res)=>{
                            if (res.data.statusCode === 200) {
                                // console.log(res.data)
                                let url = "https://github.com/login/oauth/authorize?client_id=";
                                window.location.href = url+res.data.data.app_id
                            }else{
                                this.$message.error("出了点bug...")
                            }
                        }).catch((error)=>{
                            console.log(error)
                    })
                }else {

                }
            }
        }
    }
</script>

<style scoped>
.icon {
  font-size: 40px;
  text-align: center;
}

.ant-row{
  margin-top: 10%;margin-bottom: 15%
}
@media screen and (max-width: 1024px){
  .ant-row{
    margin-top: 20%;
  }
}
</style>
