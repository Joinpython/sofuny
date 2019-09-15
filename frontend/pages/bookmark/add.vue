<template>
  <div>
    <Header/>
    <a-row :gutter="20" style="margin-top: 80px">
      <a-col :span="12" :offset="6">
          <a-row :gutter="20">
            <a-col :span="12">
              <a-form-item label="分类" :required="true">
                <a-select defaultValue="选择分类" style="width: 200px" @change="handleChange">
                  <a-select-opt-group v-for="(item, index) in big_cates" :key="index">
                    <span slot="label">
                      {{ item.name}}
                    </span>
                    <a-select-option :value="ch.id"
                                     v-for="(ch, i) in cates"
                                     :key="i">{{ ch.name }}</a-select-option>
                  </a-select-opt-group>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item label="来源:">
                <a-input placeholder="请输入来源..." v-model="bookmark.origin"/>
              </a-form-item>
            </a-col>
          </a-row>

          <a-form-item label="标题" :required="true">
            <a-input v-model="bookmark.name" placeholder="请输入标题..."/>
          </a-form-item>
          <a-form-item label="链接" :required="true">
            <a-input v-model="bookmark.url" placeholder="请输入链接..."/>
          </a-form-item>

          <a-form-item label="简介(选填)">
            <a-textarea v-model="bookmark.abstract" placeholder="请输入简介..." :rows="4"/>
          </a-form-item>

          <a-form-item>
            <div style="text-align: center">
              <a-button type="primary" @click="submitBookmark">提交</a-button>
            </div>
          </a-form-item>
      </a-col>
    </a-row>
    <Footer/>
  </div>
</template>

<script>
    import pkg from '../../../config'
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
    const allDataUrl = pkg.api +"/api/v1/bookmark/allData/";
    const bookmarkUrl = pkg.api + "/api/v1/bookmark/"
    export default {
        name: "add",
        components: { Footer, Header },
        head: ()=>({
            title: "添加书签 | " + pkg.name
        }),
        async asyncData ({ app }){
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
        data: ()=>({
            categorys: pkg.category,
            bookmark: {
                name:"",
                url:"",
                category:"",
                origin:"",
                abstract:"暂无简介...",
                cate_id: ""
            },
        }),
        methods: {
            submitBookmark() {
               if (!this.bookmark.name || !this.bookmark.url || !this.bookmark.category) {
                   this.$message.warn("请输入标题、链接、分类!")
               }else {
                   this.$axios.$post(bookmarkUrl, this.bookmark)
                       .then((res)=>{
                           if (res.statusCode === 200) {
                               this.$message.success("创建成功！")
                           }else {
                               this.$message.error(JSON.stringify(res.msg))
                           }
                       }).catch((error)=>{
                           console.log(error)
                   })
               }
            },
            handleChange(value) {
                this.bookmark.cate_id = value;
                this.bookmark.category = this.cates.filter(item => item.id === value)[0].name
            }
        }
    }
</script>

<style scoped>

</style>
