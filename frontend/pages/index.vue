<template>
  <div>
    <Header/>
    <section class="hero is-primary is-medium"
             :style="{
             backgroundImage: 'url('+staticApi+'/bg/index.bg.jpg)',
             backgroundSize: '100% 100%'
             }">
      <!-- Hero head: will stick at the top -->
      <div class="hero-head"></div>
      <!-- Hero content: will be in the middle -->
      <div class="hero-body">
        <div class="container has-text-centered">
          <h1 class="name" style="">
            {{ name }}
          </h1>
          <h2 class="subtitle">
            {{ description }}
          </h2>
          <a-button type="primary" @click="()=>{ this.$router.push({ name: 'subscribe' })}">订阅</a-button>
        </div>
      </div>
    </section>

    <!--antd vue-->
    <div style="margin-top: 20px;margin-left: 3%;margin-right: 3%">
<!--      <a-row :gutter="20">-->
<!--        <a-col :md="18" :sm="24" :xs="24" :xl="16">-->
<!--            <a-carousel arrows>-->
<!--              <div-->
<!--                slot="prevArrow" slot-scope="props"-->
<!--                class="custom-slick-arrow"-->
<!--                style="left: 10px;z-index: 1"-->
<!--              >-->
<!--                <a-icon type="left-circle" />-->
<!--              </div>-->
<!--              <div-->
<!--                slot="nextArrow" slot-scope="props"-->
<!--                class="custom-slick-arrow"-->
<!--                style="right: 10px"-->
<!--              >-->
<!--                <a-icon type="right-circle" />-->
<!--              </div>-->
<!--              <div><h3>1</h3></div>-->
<!--              <div><h3>2</h3></div>-->
<!--              <div><h3>3</h3></div>-->
<!--              <div><h3>4</h3></div>-->
<!--            </a-carousel>-->
<!--        </a-col>-->

<!--        <a-col :md="6" :sm="24" :xs="24" :xl="8">-->
<!--          <a-card hoverable style="width: 100%;height: 300px;">-->
<!--            <img-->
<!--              alt="example"-->
<!--              src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"-->
<!--              slot="cover"-->
<!--              style="height: 157px"-->
<!--            />-->
<!--            <template class="ant-card-actions" slot="actions">-->
<!--              <a-icon type="setting" />23-->
<!--              <a-icon type="edit" />232-->
<!--              <a-icon type="ellipsis" />314-->
<!--            </template>-->
<!--            <a-card-meta-->
<!--              title="Card title"-->
<!--              description="This is the description">-->
<!--              <a-avatar slot="avatar" src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />-->
<!--            </a-card-meta>-->
<!--          </a-card>-->
<!--        </a-col>-->
<!--      </a-row>-->

      <a-row :gutter="20">
          <a-col :xs="24" :sm="24" :md="18" :lg="18" :xl="18" :xxl="16">

            <a-row :gutter="20">
              <a-col :xs="24" :sm="24" :md="8" :lg="8" :xl="8" :xxl="6"
                     v-for="(item, index) in articleList"
                     :key="index"
                     style="margin-top: 20px"
              >
                <a-card hoverable style="width: 100%">
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
          </a-col>
          <a-col :xs="24" :sm="24" :md="6" :lg="6" :xl="6" :xxl="8">
            <a-card :bordered="false">
              <span slot="title">搜索</span>
              <a-input-search
                class="search-input"
                size="large"
                placeholder="请输入关键字..."
                @search="onSearch"
              />
            </a-card>

            <a-card title="简介" :bordered="false">
              <div class="subtitle is-6">
                “ <span>{{ abstract }}</span> ”
              </div>
            </a-card>

            <a-card title="文章分类" :bordered="false">
              <div v-for="(item, index) in categorys" :key="index">
                <span class="subtitle is-6" @click="goTo(`cate`, item)">
                        {{ item }}</span>
              </div>
            </a-card>

            <a-card title="热门文章" :bordered="false">
              <div v-if="hotArticleList.length=== 0">
                <span>暂无数据...</span>
              </div>
              <div v-else v-for="(item, index) in hotArticleList" :key="index">
               <span class="subtitle is-6"
                     @click="goTo(`post`, item.uuid)"
                     :title="item.title">
                         {{ item.title }}
                </span>
              </div>
            </a-card>

            <a-card title="友情链接" :bordered="false">
              <span>暂无数据...</span>
            </a-card>
          </a-col>
        </a-row>
    </div>

    <Footer />
  </div>
</template>

<script>
  import pkg from '../../config'
  import Header from '@/components/Header.vue'
  import Footer from '@/components/Footer.vue'
  import moment from 'moment'
  moment.locale('zh-cn');
  const url = pkg.api + "/api/v1/article/"

  export default {
      name: "index",
      components: { Header, Footer },
      head: ()=>({
          title: pkg.name
      }),
      async asyncData ( { app }) {
          let readTime;
          let articleList;
          let hotArticleList;
          const data = await app.$axios.$get(url);
          if (data.statusCode !== 200) { articleList = [] }
          articleList = data.data.articleList;
          articleList.forEach((item)=>{
              item.category = pkg.category[item.category]
              item.timeTitle = moment(item.created_at).format('YYYY-MM-DD HH:mm:ss');
              item.created_at = moment(item.created_at).fromNow();
              readTime = (item.content.trim().replace(/\s+/g,"").length/ 600).toFixed(1)
          })
          hotArticleList = data.data.hotArticleList;
          hotArticleList.forEach((item)=>{
              item.category = pkg.category[item.category]
              item.timeTitle = moment(item.created_at).format('YYYY-MM-DD HH:mm:ss');
              item.created_at = moment(item.created_at).fromNow();
          })
          return { articleList, readTime, hotArticleList}
      },
      data: ()=>({
          currentDate: new Date(),
          api: pkg.api,
          staticApi: "https://upyun.sofuny.cn",
          description: pkg.description,
          name: pkg.name,
          author: pkg.author,
          categorys: pkg.category,
          abstract: pkg.abstract,
      }),
      methods: {
          goTo(value, item) {
              if (value === "post") {
                  this.$router.push({path:"/post/"+item, params: {uuid: item}})
              }else if (value === "cate") {
                  this.$router.push({path:"/category/"+item, params: {cate: item}})
              }
          },
          onSearch(value, event) {
              // console.log(value)
              // console.log(event)
              this.$message.warn("我们正在努力打造中...")
          }
      }
  }
</script>

<style scoped>
  /* For demo */
  .ant-carousel >>> .slick-slide {
    text-align: center;
    height: 300px;
    line-height: 300px;
    background: #364d79;
    overflow: hidden;
  }

  .ant-carousel >>> .custom-slick-arrow {
    width: 25px;
    height: 25px;
    font-size: 25px;
    color: #fff;
    background-color: rgba(31,45,61,.11);
    opacity: 0.3;
  }
  .ant-carousel >>> .custom-slick-arrow:before {
    display: none;
  }
  .ant-carousel >>> .custom-slick-arrow:hover {
    opacity: 0.5;
  }

  .ant-carousel >>> .slick-slide  h3 {
    color: #fff;
  }
  .name {
    color: #fff;
    font-weight: 700;
    margin: 0 0 15px;
    letter-spacing: 0.8px;
    font-size: 3.25rem;
    margin-bottom: 10px;
  }
  @media screen and (max-width: 1024px){
    .container {
      flex-grow: 1;
      margin: 0 auto;
      position: relative;
      width: auto;
      margin-top: 50px;
    }
    .name{
      font-weight: 400;
      font-size: 2.25rem;
    }
  }
</style>
