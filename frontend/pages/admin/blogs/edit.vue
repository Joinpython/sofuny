<template>
  <a-form>
    <a-form-item label="标题:">
      <a-input v-model="article.title" placeholder="请输入文章标题"/>
    </a-form-item>

    <a-form-item label="内容:">
      <client-only>
        <mavon-editor
          v-model="article.content"
          :toolbars="markdownOption"
          :subfield="true"
          :shortCut="false"
          style="min-height: 800px;z-index: 0"
          ref="mavon"
          @imgAdd="uploadImage"
          @imgDel="delImage"
          codeStyle="atom-one-dark"
          placeholder="开始创作吧..."
        />
      </client-only>
    </a-form-item>

    <a-card title="其他选项">

      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item label="分类">
            <a-radio-group v-model="article.category">
              <a-radio :value="index" v-for="(item, index) in categorys" :key="index">{{ item }}</a-radio>
            </a-radio-group>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item label="是否发布">
            <a-radio-group v-model="article.status">
              <a-radio :value="index"
                       v-for="(item, index) in statusList"
                       :key="index">{{ item }}</a-radio>
            </a-radio-group>
          </a-form-item>
        </a-col>
      </a-row>

      <a-form-item label="文章配图">

        <img v-if="article.mediaUrl"
             :src="article.mediaUrl"
             alt=""
             style="width: 100%;height: 300px">
        <span v-else style="text-align: center">暂无配图...</span>

        <a-divider />

        <a-upload-dragger name="file"
                          accept="video/*,image/*"
                          listType="picture-card"
                          :multiple="false"
                          :supportServerRender="true"
                          :action="actionUrl"
                          :remove="handleRemove"
                          @change="handleChange"
                          @preview="handlePreview"
        >
          <p class="ant-upload-drag-icon">
            <a-icon type="inbox" />
          </p>
          <p class="ant-upload-text">单击或拖动文件到此区域进行上传</p>
          <a-modal :visible="previewVisible" :footer="null" @cancel="handleCancel">
            <img alt="example" style="width: 100%" :src="previewImage" />
          </a-modal>
        </a-upload-dragger>
      </a-form-item>

    </a-card>
    <a-form-item>
      <div style="text-align: center;margin-top: 30px">
        <a-button :loading="loading" type="primary" @click="handleSubmit">修改</a-button>
      </div>
    </a-form-item>
  </a-form>

</template>

<script>
    import pkg from '../../../../config'
    const fileUrl = pkg.api + "/api/v1/articleFiles/"
    const article_url = pkg.api + "/api/v1/article/";
    export default {
        layout: "admin",
        name: "edit",
        head: ()=>({
            title: "编辑文章 | " + pkg.name
        }),
        async asyncData ( { app, query }) {
            let article;
            const data = await app.$axios.$get(article_url, { params: { uuid: query.uuid } });
            if (data.statusCode !== 200) { article = {} }
            article = data.data.article;
            article.status = article.status ? 1 : 0
            return { article }
        },
        data: ()=>({
            loading: false,
            markdownOption: {
                bold: true, // 粗体
                italic: true, // 斜体
                header: true, // 标题
                underline: true, // 下划线
                strikethrough: true, // 中划线
                mark: true, // 标记
                superscript: true, // 上角标
                subscript: true, // 下角标
                quote: true, // 引用
                ol: true, // 有序列表
                ul: true, // 无序列表
                link: true, // 链接
                imagelink: true, // 图片链接
                code: true, // code
                table: true, // 表格
                fullscreen: true, // 全屏编辑
                readmodel: true, // 沉浸式阅读
                htmlcode: true, // 展示html源码
                help: true, // 帮助
                /* 1.3.5 */
                undo: true, // 上一步
                redo: true, // 下一步
                trash: true, // 清空
                save: true, // 保存（触发events中的save事件）
                /* 1.4.2 */
                navigation: true, // 导航目录
                /* 2.1.8 */
                alignleft: true, // 左对齐
                aligncenter: true, // 居中
                alignright: true, // 右对齐
                /* 2.2.1 */
                subfield: true, // 单双栏模式
                preview: true, // 预览
            },
            categorys: pkg.category,
            statusList:["草稿", "发布"],
            images:[],
            actionUrl: fileUrl,
            previewVisible: false,
            previewImage: '',
        }),
        methods: {
            // 编辑器内图片 上传
            uploadImage (pos, file) {
                let formdata = new FormData();
                formdata.append('file', file);
                this.$axios.$post(fileUrl, formdata)
                    .then((res) => {
                        if (res.statusCode === 200) {
                            this.$message.success(res.msg);
                            this.$refs.mavon.$img2Url(pos, res.data)
                            this.images.push(res.data)
                        }else {
                            this.$message.warning(res.msg)
                        }
                    }).catch((error)=>{
                    console.log(error)
                })
            },
            delImage (file) {
                this.$axios.$delete(fileUrl, { data: { imgUrl: file[0] }} )
                    .then((res) => {
                        if (res.statusCode === 200) {
                            this.$message.success("移除图片成功!")
                            this.images.splice(this.images.findIndex(item => item === res.data), 1);
                        }else {
                            this.$message.warning(res.msg)
                        }
                    }).catch((error)=>{
                    console.log(error)
                })
            },
            handleSubmit(){
                this.loading = true;
                if (!this.article.title || !this.article.content){
                    this.$message.error("标题或内容不允许为空")
                    this.loading = false;
                }
                this.$axios.$put(article_url, {
                      id: this.article.id,
                      title: this.article.title,
                      content: this.article.content,
                      category: this.article.category,
                      mediaUrl: this.article.mediaUrl,
                    }
                ).then((res)=>{
                        if (res.statusCode === 200){
                            this.$message.success(res.msg)
                            this.loading = false;
                            setTimeout(()=>{
                                this.$router.push({name:"admin-blogs-list"})
                            }, 1000)
                        }else {
                            this.loading = false;
                            this.$message.error("修改失败!");
                        }
                    }).catch((error)=>{
                      this.loading = false;
                      console.log(error)
                })
            },
            handleChange(info) {
                const status = info.file.status;
                if (status !== 'uploading') {
                    // console.log(info.file.response, info.fileList);
                }
                if (status === 'done') {
                    const res = info.file.response;
                    if (res.statusCode === 200) {
                        this.article.mediaUrl = res.data;
                        this.$message.success(`${info.file.name} file uploaded successfully.`);
                    }else {
                        this.$message.error("上传图片失败!")
                    }
                } else if (status === 'error') {
                    this.$message.error(`${info.file.name} file upload failed.`);
                }
            },
            handleRemove(file) {
                this.$axios.$delete(fileUrl, { data: { imgUrl: file.response.data }} )
                    .then((res) => {
                        if (res.statusCode === 200) {
                            this.article.mediaUrl = "";
                            this.$message.success("移除图片成功!")
                        }else {
                            this.$message.warning(res.msg)
                        }
                    }).catch((error)=>{
                    console.log(error)
                })
                return true
            },
            handlePreview (file) {
                this.previewImage = file.url || file.thumbUrl
                this.previewVisible = true
            },
            handleCancel () {
                this.previewVisible = false
            },
        }
    }
</script>

<style scoped>
  /* you can make up upload button and sample style by using stylesheets */
  .ant-upload-select-picture-card i {
    font-size: 32px;
    color: #999;
  }

  .ant-upload-select-picture-card .ant-upload-text {
    margin-top: 8px;
    color: #666;
  }
</style>
