<template>
  <a-layout style="min-height: 100vh">

    <a-layout-sider
      collapsible
      v-model="collapsed"
      :class="{'layout-admin-sider':true}"
    >
      <div class="logo" @click="()=>{this.$router.push({path:'/admin'})}"/>
      <a-menu theme="dark"
              :defaultSelectedKeys="['1']" mode="inline">

        <a-menu-item key="1" @click="()=>{this.$router.push({path:'/'})}">
          <a-icon type="pie-chart" />
          <span>首页</span>
        </a-menu-item>

        <a-menu-item key="2" @click="()=>{this.$router.push({name:'admin-blogs-create'}) }">
          <a-icon type="plus" />
          <span>发布文章</span>
        </a-menu-item>

        <a-menu-item key="22" @click="()=>{this.$router.push({name:'admin-blogs-list'}) }">
          <a-icon type="bars" />
          <span>文章列表</span>
        </a-menu-item>

        <a-menu-item key="23" @click="()=>{this.$router.push({name:'admin-setting'}) }">
          <a-icon type="setting" />
          <span>基本设置</span>
        </a-menu-item>

        <a-menu-item key="24" @click="()=>{this.$router.push({name:'admin-bookmark-list'}) }">
          <a-icon type="message" />
          <span>书签管理</span>
        </a-menu-item>

        <a-sub-menu key="sub1">
          <span slot="title">
            <a-icon type="user" /><span>用户管理</span></span>
          <a-menu-item key="3">Tom</a-menu-item>
          <a-menu-item key="4">Bill</a-menu-item>
          <a-menu-item key="5">Alex</a-menu-item>
        </a-sub-menu>

        <a-sub-menu
          key="sub2"
        >
          <span slot="title">
            <a-icon type="inbox" /><span>博客管理</span>
          </span>
          <a-menu-item key="80" @click="()=>{this.$router.push({path:'/admin/blogs/create/'})}">发布文章</a-menu-item>
          <a-menu-item key="81" @click="()=>{this.$router.push({path:'/admin/blogs/list/'})}">文章列表</a-menu-item>
        </a-sub-menu>

        <a-menu-item key="9">
          <a-icon type="file" />
          <span>File</span>
        </a-menu-item>

      </a-menu>
    </a-layout-sider>

    <a-layout>

      <a-layout-header
        :class="{'layout-admin-header':true, offsetLeft:offsetLeft}">

        <a-icon
          class="trigger"
          :type="collapsed ? 'menu-unfold' : 'menu-fold'"
          @click="showSider"
        />
      </a-layout-header>

      <a-layout-content :class="{'layout-admin-content':true, offsetLeftContent:offsetLeft}">
        <div :style="{ padding: '5px', background: '#fff', minHeight: '600px' }">
          <nuxt />
        </div>
      </a-layout-content>

      <a-layout-footer style="text-align: center">
        <a-back-top />
        {{ siteName }} ©2018 Created by Ant UED
      </a-layout-footer>

    </a-layout>

  </a-layout>
</template>

<script>
    import pkg from '../../config'
    export default {
        data: ()=>({
            collapsed: true,
            offsetLeft: false,
            siteName: pkg.name
        }),
        methods: {
            showSider () {
                this.collapsed = !this.collapsed
            }
        },
        watch: {
            collapsed: function (now, old) {
                this.offsetLeft = !now;
            }
        }
    }
</script>

<style scoped>

  .logo {
    height: 32px;
    background: rgba(255,255,255,.2);
    margin: 16px;
  }

  .trigger {
    font-size: 18px;
    line-height: 64px;
    padding: 0 24px;
    cursor: pointer;
    transition: color .3s;
  }

  .trigger:hover {
    color: #1890ff;
  }

  .layout-admin-sider{
    overflow: auto;
    z-index: 99;
    position: fixed;
    left: 0;
    top: 0;
    min-height: 100vh;
  }

  .layout-admin-header{
    background: #fff;
    padding: 0;
    position: fixed;
    width: 100%;
    z-index: 99;
    left: 85px;
  }

  .offsetLeft {
    left: 208px;
  }
  .offsetLeftContent {
    margin-left: 208px!important;
  }
  .layout-admin-content{
    margin-top: 80px;
    margin-left: 85px;
  }

  @media screen and (max-width: 1087px){

  }

  /*  .layout-admin-header{*/
  /*    background: #fff;*/
  /*    padding: 0;*/
  /*    position: fixed;*/
  /*    width: 100%;*/
  /*    z-index: 99;*/
  /*    left:0px!important;*/
  /*  }*/
  /*  .layout-admin-content{*/
  /*    margin-top: 80px;*/
  /*    margin-left: 0px!important;*/
  /*  }*/
  /*}*/
</style>
