<template>
  <a-card title="文章列表" :bordered="false" >
    <a-table :columns="columns"
             :dataSource="articleList"
             :scroll="{ y: 600 }"
             bordered>

      <template slot="operation" slot-scope="text, record, index">
        <div class='editable-row-operations'>
          <a @click="() => handle(`edit`, record)">编辑</a>
          <a @click="() => handle(`delete`, record)">删除</a>
        </div>
      </template>

    </a-table>
  </a-card>
</template>

<script>
    import pkg from '../../../../config'
    import dayjs from 'dayjs'
    const article_url = pkg.api + "/api/v1/article/";
    const columns = [
        {
          title: 'ID',
          dataIndex: 'id',
          width: '5%',
          scopedSlots: { customRender: 'id' },
        },
        {
          title: '状态',
          dataIndex: 'status',
          width: '8%',
          scopedSlots: { customRender: 'status' },
      },
        {
        title: '创建时间',
        dataIndex: 'created_at',
        width: '14%',
        scopedSlots: { customRender: 'created_at' },
    },
        {
            title: '标题',
            dataIndex: 'title',
            width: '20%',
            scopedSlots: { customRender: 'title' },
        },
        {
            title: '分类',
            dataIndex: 'category',
            width: '8%',
            scopedSlots: { customRender: 'category' },
        },
        {
            title: '浏览量',
            dataIndex: 'viewCounts',
            width: '8%',
            scopedSlots: { customRender: 'viewCounts' },
        },
        {
            title: '评论数',
            dataIndex: 'commentCounts',
            width: '8%',
            scopedSlots: { customRender: 'commentCounts' },
        },
        {
            title: '点赞数',
            dataIndex: 'likeCounts',
            width: '8%',
            scopedSlots: { customRender: 'likeCounts' },
        },
        {
            title: '操作',
            dataIndex: 'operation',
            scopedSlots: { customRender: 'operation' },
        }
    ];
    export default {
        layout: "admin",
        name: "list",
        head: ()=>({
           title: "文章管理 | " + pkg.name
        }),
        async asyncData( { app }){
            let articleList;
            const data = await app.$axios.$get(article_url);
            if (data.statusCode === 200) {
                articleList = [...data.data.articleList];
            }else {
                articleList = []
            }
            if (articleList.length === 0){

            }
            articleList.forEach((item)=>{
                item.created_at = dayjs(item.created_at).format("YYYY-MM-DD HH:mm:ss");
                item.status = item.status ? "发布" : "草稿";
                item.category = pkg.category[item.category];
                item.key = item.uuid;
            });
            return { articleList }
        },
        data: ()=>({
            columns,
        }),
        methods: {
            handle (key, row) {
                if (key === "edit") {
                   this.$router.push({
                       name: 'admin-blogs-edit',
                       query: {uuid: row.uuid}
                   })
                }else {
                    this.$axios.$delete(article_url, { data: { id: row.id } })
                        .then((res)=>{
                            if (res.statusCode === 200) {
                                let arr = this._data.articleList
                                arr.splice(arr.findIndex(item => item.id === res.data.id), 1)
                                this.$message.success("删除成功!")
                            }else {
                                this.$message.error("删除失败!")
                            }
                        }).catch((error)=>{
                            console.log(error)
                        })
                }
            },
        },
    }
</script>

<style scoped>

</style>
