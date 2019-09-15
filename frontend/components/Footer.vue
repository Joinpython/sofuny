<template>
  <footer class="footer">
    <a-back-top />
    <div class="content has-text-centered">
      <a-row :gutter="10">
        <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8">
          <a-card :bordered="false" class="page">
            <span slot="title" class="page">网站相关</span>
            <p>暂无数据...</p>
          </a-card>
        </a-col>
        <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8">
          <a-card :bordered="false" class="page">
            <span slot="title" class="page">免责声明</span>
            <P>部分文章及配图来源于互联网,</P>
            <P>大部分已注明来源及其相关信息,</P>
            <P>版权归原创者所有. 如果不小心侵犯了您的权益,</P>
            <p>请及时通知我们, 我们会删除侵权内容, 谢谢合作.</p>
            <p>Email: {{ email }}</p>
          </a-card>
        </a-col>
        <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8">
          <a-card :bordered="false" class="page">
            <span slot="title" class="page">服务</span>
            <a href="https://www.upyun.com" target="_blank" style="display: inline">
              <img :src="staticApi+`/bg/upyun.png`"
                   style="height: 48px;width: 90px;"
                   title="提供云存储服务..."
                   alt="upyun-logo">
            </a>
          </a-card>
        </a-col>
      </a-row>
      <a-divider />
      <div>
        <a-row :gutter="20">
          <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8" class="footer-bottom">
            ©  All 2019
            <nuxt-link to="/">{{ name }}</nuxt-link>
            黔ICP备17011880号-3
          </a-col>
          <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8" class="footer-bottom">
            <p>{{ daySentence }}</p>
          </a-col>
          <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8">
            <span class="icon" title="qq">
              <a-icon type="qq" />
            </span>
            <span class="icon" title="微信">
             <a-icon type="wechat" />
            </span>
            <span class="icon" title="微博">
              <a-icon type="weibo" />
            </span>
            <span class="icon" title="github">
              <a-icon type="github" />
            </span>
          </a-col >
        </a-row>
      </div>
    </div>
  </footer>
</template>

<script>
    import pkg from '../../config'
    export default {
        name: "Footer",
        data: ()=>({
            name: pkg.name,
            email: pkg.email,
            staticApi: "https://upyun.sofuny.cn",
            daySentence: "慢慢来，总有些东西能治愈你所有的悲伤。",
        }),
        mounted() {
            var _hmt = _hmt || [];
            var hm = document.createElement("script");
            hm.src = "https://hm.baidu.com/hm.js?38c5eb77ac9f4290642bee782a43b0c9";
            var s = document.getElementsByTagName("script")[0];
            s.parentNode.insertBefore(hm, s);
        }
    }
</script>

<style scoped>
  .footer {
    padding: 1rem 1rem 1.5rem;
    background-color: #2c2c2c;
    color: #999;
    z-index: 999999;
    margin-top: 40px;
  }
  .disclaimer > p {
    color: #999;
    font-size: 12px;
  }
  .page {
    background-color: #2c2c2c;color: #999;
  }
  .icon {
    font-size: 30px;
    margin-left: 15px;
  }
  @media screen and (max-width: 1024px){
    .footer-bottom{
      margin-top: 10px;
    }
  }
</style>
