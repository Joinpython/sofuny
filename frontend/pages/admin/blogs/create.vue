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

      <a-form-item label="分类">
        <a-radio-group v-model="article.category">
          <a-radio :value="index" v-for="(item, index) in categorys" :key="index">{{ item }}</a-radio>
        </a-radio-group>
      </a-form-item>

      <a-form-item label="文章配图">
        <a-upload-dragger name="file"
                          accept="video/*,image/*"
                          listType="picture-card"
                          :multiple="false"
                          :supportServerRender="true"
                          :action="actionUrl"
                          :headers="uploadHeaders"
                          :remove="handleRemove"
                          @change="handleChange"
                          @preview="handlePreview"
        >
          <p class="ant-upload-drag-icon">
            <a-icon type="inbox" />
          </p>
          <p class="ant-upload-text">单击或拖动文件到此区域进行上传</p>
          <a-modal :visible="previewVisible"
                   :footer="null"
                   @cancel="handleCancel">
            <img alt="文章配图"
                 style="width: 100%"
                 :src="previewImage" />
          </a-modal>
        </a-upload-dragger>
      </a-form-item>

    </a-card>
    <a-form-item>
      <div style="text-align: center;margin-top: 30px">
        <a-button type="primary" :loading="loading" @click="handleSubmit(true)">发布</a-button>
        <a-divider type="vertical" />
        <a-button type="danger" :loading="loading" @click="handleSubmit(false)">草稿</a-button>
      </div>
    </a-form-item>
  </a-form>
</template>

<script>
    import pkg from '../../../../config'
    const fileUrl = pkg.api + "/api/v1/articleFiles/"
    const article_url = pkg.api + "/api/v1/article/"
    export default {
        layout: "admin",
        name: "create",
        head: ()=>({
            title: "发布文章 | " + pkg.name
        }),
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
            article: {
                title:"",
                content:"",
                images:"",
                status: true,
                category: 0,
                mediaUrl:""
            },
            images:[],
            actionUrl: fileUrl,
            previewVisible: false,
            previewImage: '',
            uploadHeaders: {
                "Authorization":'Bearer '+pkg.basicAuth.key
            }
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
            // 编辑器内图片 删除
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
            handleSubmit(value){
                this.loading = true;
                this.article.status = value
                this.article.images = this.images.join(",")
                if (!this.article.title || !this.article.content || !this.article.mediaUrl) {
                    this.$message.error("文章标题或者内容或者媒体地址为空")
                    this.loading = false;
                }else {
                    this.$axios.$post(article_url, this.article)
                        .then((res)=>{
                            if (res.statusCode === 200) {
                                this.$message.success("创建文章成功!");
                                this.loading = false;
                                setTimeout(()=>{
                                    this.$router.push({ name:"admin-blogs-list"})
                                }, 1000)
                            }else {
                                this.loading = false;
                                this.$message.error(res.msg.Message)
                            }
                        }).catch((error)=>{
                            this.loading = false;
                            console.log(error)
                    })
                }
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
