<template>
  <div>
    <Header/>
    <a-card :title="category"
            :bordered="false"
            class="category">
      <a-row :gutter="20">
        <a-col v-if="cateList.length=== 0">
          <div>
            暂无数据...
          </div>
        </a-col>
        <a-col v-else :xs="24" :sm="24" :md="6" :lg="6" :xl="6" :xxl="8" v-for="(item, index) in cateList" :key="index">
          <a-card hoverable style="width: 300px;">
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
                        src="/sofuny.png"/>
            </a-card-meta>
          </a-card>
        </a-col>
      </a-row>
    </a-card>
    <Footer/>
  </div>
</template>

<script>
    import pkg from '../../../config'
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
    import moment from 'moment'
    moment.locale('zh-cn');
    const url = pkg.api + "/api/v1/category/";
    export default {
        name: "cate",
        components: {
          Header,
          Footer
        },
        head () { return { title: this.category+" | "+pkg.name } },
        async asyncData( { app, params}) {
            let cate;
            let cateList;
            let category = params.cate;
            pkg.category.forEach((item, index)=>{ if (item === params.cate) { cate = index } });
            const data = await app.$axios.$get(url, { params: { cate: cate } });
            if (data.statusCode !== 200) { cateList = [] }
            cateList = data.data;
            cateList.forEach((item)=>{
                item.category = pkg.category[item.category];
                item.timeTitle = moment(item.created_at).format('YYYY-MM-DD HH:mm:ss');
                item.created_at = moment(item.created_at).fromNow();
            });
            return { category, cateList }
        },
        methods: {
            goTo(value, item) {
                if (value === "post") {
                    this.$router.push({path:"/post/"+item, params: {uuid: item}})
                }else if (value === "cate") {
                    this.$router.push({path:"/category/"+item, params: {cate: item}})
                }
            }
        }
    }
</script>

<style scoped>
  .category{
    margin-top: 80px;
    min-height: 500px;
    margin-left: 1%;
    margin-right: 1%
  }
  @media screen and (max-width: 1024px){
    .ant-col-xs-24 {
      display: block;
      box-sizing: border-box;
      width: 100%;
      margin-top: 20px;
      margin-left: 5px;
    }
    .ant-col-xs-24:first-child {
      margin-top: 0px;
    }
    .category {
      margin-left: 1%;
    }
  }
</style>
