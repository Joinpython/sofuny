<template>
  <nav class="navbar is-fixed-top" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
      <nuxt-link to="/" class="navbar-item log-image">
        <img :src="staticApi+`/bg/sofuny.logo.png`" width="112" height="28" alt="">
      </nuxt-link>
      <a role="button" id="navbar" @click="showMuse" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>

    <div id="navbarBasic" class="navbar-menu">
      <div class="navbar-start">
<!--        <a class="navbar-item"-->
<!--           v-for="(item, index) in categorys"-->
<!--           @click="goToCate(item)"-->
<!--           :key="index">-->
<!--          {{item}}-->
<!--        </a>-->

        <nuxt-link class="navbar-item" to="/bookmark">
          书签
        </nuxt-link>

<!--        <nuxt-link class="navbar-item" to="/admin">-->
<!--          后台-->
<!--        </nuxt-link>-->
      </div>
      <div class="navbar-end">
        <div class="navbar-item">
          <div class="buttons">
            <a class="button is-primary" @click="()=>{ this.$router.push({name:'account-login'})}">
              <strong>登录</strong>
            </a>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
    import pkg from '../../config'
    export default {
        name: "Header",
        data: ()=>({
            drawer: false,
            categorys: pkg.category,
            staticApi: "https://upyun.sofuny.cn",
        }),
        methods: {
            showMuse () {
                document.getElementById("navbar").classList.toggle('is-active');
                document.getElementById("navbarBasic").classList.toggle('is-active');
            },
            goToCate(value) {
                this.$router.push({
                    path: "/category/"+value,
                    params: {
                        cate: value
                    }
                })
            }
        }
    }
</script>

<style scoped>
  .navbar{
    min-height: 4.25rem;
    /*background-color: darkgray;*/
    border-bottom:1px solid darkgray;
  }
  .navbar-item img {
    max-height: 4.25rem;
  }
  .navbar-brand{
    min-height: 4.25rem;
  }
  .navbar-burger{
    min-height: 4.25rem;
  }
  .navbar-item {
    font-size: 20px;
  }
  @media screen and (max-width: 1087px){
    .navbar-brand, .navbar-tabs {
      align-items: stretch;
      display: flex;
      flex-shrink: 0;
      min-height: 4.25rem;
    }
    .navbar-item {
      display: inline-block;
      flex-grow: 0;
      flex-shrink: 0;
    }
    .log-image{
      margin-top: 1px;
    }
  }
</style>
