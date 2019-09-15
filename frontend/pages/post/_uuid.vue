<template>
    <div>
      <Header/>
      <div class="article">
        <a-row :gutter="20">
          <a-col :xs="24" :sm="24" :md="18" :xl="18" :xxl="18">
            <img :src="article.mediaUrl"
                 alt="文章配图"
                 :title="article.title"
                 class="article-image">
            <a-card :bordered="false">
              <div slot="title">
                <h2 class="title is-3">{{ article.title }}</h2>
                <div style="line-height: 1.5;color: rgba(0, 0, 0, 0.65);" class="clearfix">
                  <div class="article-left">
                    <small>
                      <a-icon type="user" />
                      {{ author }}
                    </small>
                    <small>
                      <a-icon type="calendar" />
                      {{ article.created_at }}
                    </small>
                    <small>
                      <a-icon type="calculator" />
                      {{ wordCount }} 字</small>
                    <small>
                      <a-icon type="read" />
                      {{ readTime }} 分钟</small>
                  </div>
                  <div class="article-right">
                    <small>
                      <a-icon type="like" />({{ article.likeCounts}})
                    </small>
                    <a-divider type="vertical" />
                    <small>
                      <a-icon type="message" />({{ article.commentCounts}})
                    </small>
                  </div>
                </div>
              </div>
              <!--TODO 文章h系列标签未渲染-->
              <vue-markdown class="article-context"
                            v-highlight
                            :source="article.content" />

              <a-divider>-- END --</a-divider>

              <div style="text-align: center">
                <a-icon type="like" style="font-size: 40px" @click="addArticleLike(article.id)"/>({{ article.likeCounts}})
              </div>

              <a-card title="相关推荐" :bordered="false">
                <a-row :gutter="20">
                  <a-col v-if="relateArticle.length === 0">
                    <span>暂无数据...</span>
                  </a-col>
                  <a-col v-else :xs="24" :sm="24" :md="8" :xl="8" :xxl="8"
                         v-for="(item, index) in relateArticle"
                         :key="index"
                         style="margin-top: 10px"
                  >
                    <a-card :hoverable="true"
                            style="width: 100%">
                      <img
                        alt="文章配图"
                        :src="item.mediaUrl"
                        slot="cover"
                        :title="item.title"
                        @click="goTo(`post`, item.uuid)"
                        style="min-height: 160px;max-height: 160px;width: 100%"
                      />
                      <template class="ant-card-actions" slot="actions">
                        <a-icon type="eye" />{{ item.viewCounts }}
                        <a-icon type="message" />{{ item.commentCounts }}
                        <a-icon type="like" />{{ item.likeCounts }}
                      </template>
                      <a-card-meta>
                        <h2 slot="title" class="title is-5">
                      <span @click="goTo(`post`, item.uuid)" :title="item.title">
                         {{ item.title }}
                      </span>
                        </h2>
                        <div slot="description">
                      <span class="subtitle is-6" @click="goTo(`cate`, item.category)">
                        {{ item.category }}</span>
                          <a-divider type="vertical" />
                          <span :title="item.timeTitle">
                          <a-icon type="calendar" />
                          {{ item.created_at }}
                        </span>
                          <!--                  <a-divider type="vertical" />-->
                          <!--                  <span>read {{ readTime }} 分钟</span>-->
                        </div>
                        <a-avatar slot="avatar"
                                  :title="author"
                                  src="/sofuny.png" />
                      </a-card-meta>
                    </a-card>
                  </a-col>
                </a-row>
              </a-card>

              <!--gitalk 评论挂载元素-->
              <div id="gitalk-container"></div>

<!--              gitment-->
              <div id="comments"></div>
            </a-card>
          </a-col>
          <!--下一篇-->
          <a-col :xs="24" :sm="24" :md="6" :xl="6" :xxl="6">
            <a-affix :offsetTop="120" style="z-index: 99">
              <a-card title="下一篇"
                      :hoverable="true"
                      v-if="nextArticle.title">
                <img
                  alt="文章配图"
                  :src="nextArticle.mediaUrl"
                  slot="cover"
                  :title="nextArticle.title"
                  @click="goTo(`post`, nextArticle.uuid)"
                  style="height: 150px"
                />
                <template class="ant-card-actions" slot="actions">
                  <a-icon type="eye" />{{ nextArticle.viewCounts }}
                  <a-icon type="message" />{{ nextArticle.commentCounts }}
                  <a-icon type="like" />{{ nextArticle.likeCounts }}
                </template>
                <a-card-meta>
                  <h2 slot="title">
                    <span @click="goTo(`post`, nextArticle.uuid)" :title="nextArticle.title">
                          {{ nextArticle.title }}
                    </span>
                  </h2>
                  <p slot="description">
                    <span @click="goTo(`cate`, nextArticle.category)">{{ nextArticle.category }}</span>
                    <a-divider type="vertical" />
                    <span :title="nextArticle.timeTitle">
                      <a-icon type="calendar" />
                      {{ nextArticle.created_at }}
                  </span>
                  </p>
                  <a-avatar slot="avatar" src="/sofuny.png" />
                </a-card-meta>
              </a-card>
              <a-card title="下一篇" v-else>
                <div>
                  <p style="text-align: center;margin-bottom: 10px">正在创造中...</p>
                  <img :src="staticApi+`/bg/nextArticle.gif`"
                       alt="下一篇">
                </div>
              </a-card>
            </a-affix>
          </a-col>
        </a-row>
      </div>
      <Footer/>
    </div>
</template>

<script>
    import pkg from '../../../config'
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
    import VueMarkdown from 'vue-markdown'
    const url = pkg.api + "/api/v1/article/";
    const article_like_url = url+"like/";
    import moment from 'moment'
    moment.locale('zh-cn');
    import 'gitalk/dist/gitalk.css'
    import Gitalk from 'gitalk'
    export default {
        name: "uuid",
        components: { Header, Footer, VueMarkdown },
        head () {
          return { title: this.article.title +" | " +pkg.name }
        },
        async asyncData ( { app, params }) {
            let article;
            let wordCount;
            let readTime;
            let nextArticle;
            let relateArticle;
            const data = await app.$axios.$get(url, { params: { uuid: params.uuid } });
            if (data.statusCode !== 200) { article = {} }
            article = data.data.article;
            article.category = pkg.category[article.category];
            nextArticle = data.data.nextArticle.id === 0 ? {}:data.data.nextArticle;
            nextArticle.category = pkg.category[nextArticle.category];
            nextArticle.timeTitle = moment(nextArticle.created_at).format('YYYY-MM-DD HH:mm:ss');
            nextArticle.created_at = moment(nextArticle.created_at).fromNow();
            wordCount = article.content.trim().replace(/\s+/g,"").length;
            readTime = (wordCount / 600).toFixed(1);
            article.timeTitle = moment(article.created_at).format('YYYY-MM-DD HH:mm:ss');
            article.created_at = moment(article.created_at).format('YYYY-MM-DD HH:mm');
            relateArticle = data.data.relateArticle;
            relateArticle.forEach((item)=>{
                item.category = pkg.category[item.category];
                item.timeTitle = moment(item.created_at).format('YYYY-MM-DD HH:mm:ss');
                item.created_at = moment(item.created_at).fromNow();
            });
            return { article, wordCount, readTime, nextArticle, relateArticle }
        },
        data: ()=>({
            moment,
            loading: false,
            author: pkg.author,
            staticApi: "https://upyun.sofuny.cn",
            siteName: pkg.name
        }),
        methods:{
            goTo(value, item) {
                if (value === "post") {
                    this.$router.push({path:"/post/"+item, params: {uuid: item}})
                }else if (value === "cate") {
                    this.$router.push({path:"/category/"+item, params: {cate: item}})
                }
            },
            addArticleLike (id) {
                this.$axios.$post(article_like_url, { id: id })
                    .then((res)=>{
                        if (res.statusCode === 200) {
                            this._data.article.likeCounts = res.data.newLikeCounts;
                            this.$message.success("点赞成功!")
                        }else {
                            this.$message.error("点赞失败!")
                        }
                    }).catch((error)=>{
                    console.log(error)
                })
            }
        },
        mounted() {
            this.$nextTick(()=>{
                var gitalk = new Gitalk({
                    clientID: pkg.auth.github.app_id,
                    clientSecret: pkg.auth.github.app_key,
                    repo: 'blogs',
                    owner: 'tqblogs',
                    admin: ['tqblogs'],
                    id: location.pathname,      // Ensure uniqueness and length less than 50
                    distractionFreeMode: false  // Facebook-like distraction free mode
                });
                gitalk.render('gitalk-container');
            })
        }
    }
</script>

<style scoped>
  h1, h2, h3, h4, h5, h6 {
    margin-top: 0;
    margin-bottom: 0.5em;
    color: rgba(0, 0, 0, 0.85);
    font-weight: 500;
  }
  .article-context * {
    box-sizing: border-box;
  }
  .article-context{
    min-height: 600px;
    font-size: 16px;
    line-height: 1.5rem;
  }
  .article-context h2 {
    padding-bottom: .3em;
    border-bottom: 1px solid #eaecef;
    margin-top: 24px;
    margin-bottom: 16px;
    font-weight: 600;
    line-height: 1.25;
    font-size: 20px;
  }
  .article-context ol {
    padding-left: 2em;
    margin-top: 0;
    margin-bottom: 16px;
  }

  .article-context > blockquote {
    margin: 0px;
    color: rgb(106, 115, 125);
    padding: 0px 1em;
    border-left: 0.25em solid rgb(223, 226, 229);
    margin-bottom: 16px;
  }
  .article{
    margin-top: 80px;
    margin-left: 1%;
    margin-right: 1%;
    min-height: 600px;
  }
  pre {
    -webkit-overflow-scrolling: touch;
    background-color: whitesmoke;
    color: #4a4a4a;
    font-size: 0.875em;
    overflow-x: auto;
    padding: 0.5rem 0.875rem;
    white-space: pre;
    word-wrap: normal;
  }
  .article-left{
    float: left;
  }
  .article-right {
    float: right;
  }
  .article-image{
    max-height: 300px;
    width: 100%
  }
  @media screen and (max-width: 1024px){
    .article{
      margin-top: 80px;
      margin-left: 0;
      margin-right: 0;
      min-height: 600px;
    }
    .context{
      font-size: 14px;
    }
    .article-right {
      display: none;
    }
    .article-image{
      max-height: 160px;
      width: 100%
    }
  .title.is-3 {
    font-size: 1.5rem;
  }
}
</style>
