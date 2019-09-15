<template>
<div style="margin-top: 40px">
  <Header/>
  <a-row>
    <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8"></a-col>
    <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8">
      <h2 class="title is-3" style="text-align: center">订阅 {{ name }}</h2>
      <a-input type="text"
               class="s-input"
               v-model="email"
               placeholder="请输入您的邮箱地址"
               size="large"/>
      <a-button type="primary"
                size="large"
                @click="handleSubmit"
                class="s-button"
                style="">订阅</a-button>
    </a-col>
    <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8"></a-col>
  </a-row>
  <Footer/>
</div>
</template>

<script>
    import pkg from '../../config'
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
    const url = pkg.api + "/api/v1/subscribe/";

    export default {
        name: "subscribe",
        components: {
            Header,
            Footer,
        },
        head: ()=>({
            title: "订阅 | " + pkg.name
        }),
        data: ()=>({
            email:"",
            name: pkg.name
        }),
        methods: {
            handleSubmit(){
                if (!this.email || !/^\w+((.\w+)|(-\w+))@[A-Za-z0-9]+((.|-)[A-Za-z0-9]+).[A-Za-z0-9]+$/.test(this.email)){
                    this.$message.error("请输入正确的邮箱地址!")
                }else {
                    this.$axios.$post(url, { email: this.email })
                        .then((res)=>{
                            if (res.statusCode === 200) {
                                this.$message.success("订阅成功!");
                                this.email = "";
                            }else if (res.msg.Number === 1062){
                                this.$message.warn("您的邮箱已经订阅过了!☺")
                            }else {
                                this.$message.error("订阅失败!")
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
  .ant-row{
    margin-top: 20%;
    min-height: 600px;
  }
  .s-button{
    width: 100%;
    margin-top: 20px
  }
  @media screen and (max-width: 1024px){
    .ant-row{
      margin-top: 40%;
      min-height: 300px;
      margin-bottom: -20%;
    }
    .s-input{
      width: 90%;
      margin-left: 5%;
      margin-right: 5%;
    }
    .s-button {
      width: 90%;
      margin-left: 5%;
      margin-right: 5%;
    }
  }
</style>
