<template>
<div>
  <Header/>
  <a-row style="width: 94%;margin-top: 80px;">
    <a-col :span="24" :offset="1">
      <a-tabs @change="callback"
              size="large"
              tabPosition="top"
      >
        <a-tab-pane :tab="item.name" :key="index+1" v-for="(item, index) in big_cates">

          <a-row :gutter="20">
            <a-col :xs="24" :sm="24" :md="8" :xl="8" :xxl="8"
                   v-for="(ca, index) in cates"
                   :key="index"
                   v-if="ca.big_id === item.id"
                   style="margin-top: 20px">
              <a-card class="box-card">
                <div slot="title" class="clearfix">
                  <span>{{ ca.name }}</span>
                </div>
                <div v-for="(book, index) in bookmarks" :key="index" class="text item">
                  <a-button type="link"
                            :title="book.origin"
                            v-if="book.cate_id === ca.id"
                            @click="goToUrl(book.url)">{{ book.name }}</a-button>
                </div>
              </a-card>
            </a-col>
          </a-row>
        </a-tab-pane >
        <a-button slot="tabBarExtraContent" @click="()=>{ this.$router.push({ name:'bookmark-add' })}">
          <a-icon type="plus"/>
        </a-button>
      </a-tabs>
    </a-col>
  </a-row>
<Footer/>
</div>
</template>

<script>
  // TODO 标签页面有待完善　时间仓促　简单上线
    import pkg from '../../../config'
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'

    const allDataUrl = pkg.api +"/api/v1/bookmark/allData/";

    export default {
        name: "bookmark",
        components: {
          Header,
          Footer,
        },
        head: ()=>({
            title: "书签 | " + pkg.name
        }),
        async asyncData( {app, error}) {
            let big_cates;
            let cates;
            let bookmarks;
            const data = await app.$axios.$get(allDataUrl);
            if ( data.statusCode !== 200) {
                error({ statusCode: 500, message: JSON.stringify(data.msg) })
            }
            big_cates = data.data.big_cates;
            cates = data.data.cates;
            bookmarks = data.data.bookmarks;
            return { big_cates, cates, bookmarks }
        },
        methods: {
            callback(key) {

            },
            goToUrl(url){
                window.location.href = url
            }
        }
    }
</script>

<style scoped>

</style>
