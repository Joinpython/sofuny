<template>
  <div>
    <a-card title="大分类" :bordered="false" >
      <span slot="extra">
        <a-button type="primary" @click="addHandle('add', '大分类')">添加</a-button>
      </span>
      <a-table :columns="BigCateColumns"
               :dataSource="big_cates"
               :scroll="{ y: 600 }"
               bordered>

        <template slot="operation" slot-scope="text, record, index">
          <div class='editable-row-operations'>
            <a @click="BigHandle('edit', record)">编辑</a>
            <a @click="BigHandle('delete', record)">删除</a>
          </div>
        </template>
      </a-table>
    </a-card>

    <a-card title="小分类" :bordered="false" >
      <span slot="extra">
        <a-button type="primary" @click="addHandle('add', '小分类')">添加</a-button>
      </span>
      <a-table :columns="CateColumns"
               :dataSource="cates"
               :scroll="{ y: 600 }"
               bordered>
        <template slot="operation" slot-scope="text, record, index">
          <div class='editable-row-operations'>
            <a @click="() => CateHandle(`edit`, record)">编辑</a>
            <a @click="() => CateHandle(`delete`, record)">删除</a>
          </div>
        </template>

      </a-table>
    </a-card>

    <a-modal
      :title="action+` | `+actionName"
      :visible="visible"
      @ok="handleOk(action, actionName)"
      :confirmLoading="loading"
      @cancel="handleCancel"
    >
      <a-form-item label="所属大分类:" v-if="actionName === `小分类`">
        <a-select :defaultValue="cateDefault" style="width: 200px" @change="changeCateHandle">
          <a-select-option :value="item.id"
                           :key="index"
                           v-for="(item, index) in big_cates">{{ item.name }}</a-select-option>
        </a-select>
      </a-form-item>

      <a-form-item label="名称:">
        <a-input v-model="name"/>
      </a-form-item>
      <a-form-item label="状态:">
        <a-switch :defaultChecked="true"
                  checkedChildren="激活"
                  unCheckedChildren="关闭"
                  @change="changeHandle"/>
      </a-form-item>
    </a-modal>

    <a-card title="书签列表" :bordered="false" >
      <span slot="extra">
        <a-button type="primary" @click="addHandle('add', '书签')">添加</a-button>
      </span>
      <a-table :columns="BigCateColumns"
               :dataSource="bookmarks"
               :scroll="{ y: 600 }"
               bordered>

        <template slot="operation" slot-scope="text, record, index">
          <div class='editable-row-operations'>
            <a @click="() => BookmarkHandle(`edit`, record)">编辑</a>
            <a @click="() => BookmarkHandle(`delete`, record)">删除</a>
          </div>
        </template>

      </a-table>
    </a-card>
  </div>
</template>

<script>
    import pkg from '../../../../config'
    const allDataUrl = pkg.api +"/api/v1/bookmark/allData/";
    const bigCateUrl = pkg.api +"/api/v1/bookmark/big/";
    const cateUrl = pkg.api + "/api/v1/bookmark/cate/";
    import moment from 'moment'
    moment.locale('zh-cn');
    const BigCateColumns = [
        {
            title: 'ID',
            dataIndex: 'id',
            width: '5%',
        },
        {
            title: '状态',
            dataIndex: 'status',
            width: '8%',
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            width: '14%',
        },
        {
            title: '更新时间',
            dataIndex: 'updated_at',
            width: '14%',
        },
        {
            title: '名称',
            dataIndex: 'name',
            width: '10%',
        },
        {
            title: '操作',
            dataIndex: 'operation',
            scopedSlots: { customRender: 'operation' },
        }
    ];
    const CateColumns = [
        {
            title: 'ID',
            dataIndex: 'id',
            width: '5%',
        },
        {
            title: '状态',
            dataIndex: 'status',
            width: '8%',
        },
        {
            title: '创建时间',
            dataIndex: 'created_at',
            width: '14%',
        },
        {
            title: '更新时间',
            dataIndex: 'updated_at',
            width: '14%',
        },
        {
            title: '名称',
            dataIndex: 'name',
            width: '10%',
        },
        {
            title: '所属大分类',
            dataIndex: 'bigCate',
            width: '10%',
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
            title: "书签管理 | " + pkg.name
        }),
        async asyncData({ app, error }) {
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
            big_cates.forEach((item)=>{
                item.created_at = moment(item.created_at).format("YYYY-MM-DD HH:mm:ss");
                item.updated_at = moment(item.updated_at).format("YYYY-MM-DD HH:mm:ss");
                item.status = item.status ? "使用中" : "关闭";
                item.key = item.uuid;
            });
            cates.forEach((item)=>{
                item.created_at = moment(item.created_at).format("YYYY-MM-DD HH:mm:ss");
                item.updated_at = moment(item.updated_at).format("YYYY-MM-DD HH:mm:ss");
                item.status = item.status ? "使用中" : "关闭";
                item.bigCate = big_cates.filter(it => it.id === item.big_id)[0].name;
                item.key = item.uuid;
            });
            bookmarks.forEach((item)=>{
                item.created_at = moment(item.created_at).format("YYYY-MM-DD HH:mm:ss");
                item.status = item.status ? "发布" : "草稿";
                item.category = pkg.category[item.category];
                item.key = item.uuid;
            });
            return { big_cates, cates, bookmarks }
        },
        data: ()=>({
            BigCateColumns,
            CateColumns,
            visible: false,
            loading: false,
            action: "",
            actionName:"",
            name: "",
            id: "",
            status: true,
            cate: {
                big_id: "",
                name:"",
                status:true
            },
            cateDefault: "选择大分类",
        }),
        methods: {
            addHandle (key, name) {
                this.visible = true;
                this.action = key;
                this.actionName = name;
            },
            BigHandle(key, row) {
                 if (key === "edit") {
                    this.visible = true;
                    this.action = key;
                    this.actionName = "大分类";
                    this.name = row.name;
                    this.id = row.id;
                }else {
                    this.$axios.$delete(bigCateUrl, { data: { id: row.id } })
                        .then((res)=>{
                            if (res.statusCode === 200) {
                                let arr = this._data.big_cates;
                                let name = arr.filter(item => item.id === res.data.id)[0].name;
                                arr.splice(arr.findIndex(item => item.id === res.data.id), 1);
                                this.$message.success("删除 "+name+ " 成功!")
                            }else {
                                this.$message.error(JSON.stringify(res.msg))
                            }
                        }).catch((error)=>{
                            console.log(error)
                    })
                }
            },
            CateHandle(key, row) {
                if (key === "edit") {
                    this.visible = true;
                    this.action = key;
                    this.actionName = "小分类";
                    this.name = row.name;
                    this.id = row.id;
                    this.cateDefault = row.bigCate;
                    this.cate.big_id = row.big_id;
                }else {
                    this.$axios.$delete(cateUrl, { data: { id: row.id } })
                        .then((res)=>{
                            if (res.statusCode === 200) {
                                let arr = this._data.cates;
                                arr.splice(arr.findIndex(item => item.id === res.data.id), 1);
                                this.$message.success("删除成功!")
                            }else {
                                this.$message.error(JSON.stringify(res.msg))
                            }
                        }).catch((error)=>{
                        console.log(error)
                    })
                }
            },
            BookmarkHandle(key, row) {
                 if (key === "edit") {

                }else {

                }
            },
            handleOk(action, name) {
                if (!this.name) {
                    this.$message.error("请填写数据!")
                }
                // 添加　大分类
                if (action === "add" && name === "大分类" && this.name !== "") {
                    this.$axios.$post(bigCateUrl, {
                        name: this.name,
                        status: this.status,
                    }).then((res)=>{
                        if (res.statusCode === 200) {
                            res.data.created_at = moment(res.data.created_at).format("YYYY-MM-DD HH:mm:ss");
                            res.data.updated_at = moment(res.data.updated_at).format("YYYY-MM-DD HH:mm:ss");
                            res.data.status = res.data.status ? "使用中" : "关闭";
                            res.data.key = res.data.uuid;
                            this._data.big_cates.push(res.data);
                            this.name = "";
                            this.status = true;
                            this.visible = false;
                            this.$message.success("添加 "+res.data.name+" "+res.msg)
                        }else {
                            this.name = "";
                            this.status = true;
                            this.visible = false;
                            this.$message.error(JSON.stringify(res.msg))
                        }}).catch((error)=>{
                          this.name = "";
                          this.status = true;
                          this.visible = false;
                          console.log(error)
                    })
                }
                // 编辑　大分类
                if (action === "edit" && name === "大分类" && this.name !== "") {
                    this.$axios.$put(bigCateUrl, {
                        id: this.id,
                        name: this.name,
                        status: this.status,
                    }).then((res)=>{
                        if (res.statusCode === 200) {
                            this.id = "";
                            this.name = "";
                            this.visible = false;
                            this.status = true;
                            let arr = this._data.big_cates;
                            let name = arr.filter(item => item.id === res.data.id)[0].name;
                            arr.filter(item => item.id === res.data.id)[0].name = res.data.name;
                            this.$message.success("编辑 "+name+" "+res.msg);
                        }else {
                            this.id = "";
                            this.name = "";
                            this.visible = false;
                            this.status = true;
                            this.$message.error(JSON.stringify(res.msg))
                        }
                    }).catch((error)=>{
                        this.id = "";
                        this.name = "";
                        this.visible = false;
                        this.status = true;
                        console.log(error)
                    })
                }

                // 添加　小分类
                if (action === "add" && name === "小分类" && this.name !== "") {
                    this.cate.name = this.name
                    this.cate.status = this.status
                    this.$axios.$post(cateUrl, this.cate)
                        .then((res)=>{
                            if (res.statusCode === 200) {
                                this.name = "";
                                res.data.created_at = moment(res.data.created_at).format("YYYY-MM-DD HH:mm:ss");
                                res.data.updated_at = moment(res.data.updated_at).format("YYYY-MM-DD HH:mm:ss");
                                res.data.status = res.data.status ? "使用中" : "关闭";
                                res.data.key = res.data.uuid;
                                res.data.bigCate = this._data.big_cates.filter(it => it.id === res.data.big_id)[0].name;
                                this._data.cates.push(res.data);
                                this.visible = false;
                                this.cateDefault = "选择大分类"
                                this.$message.success(res.msg)
                            }else {
                                this.name = "";
                                this.visible = false;
                                this.$message.error(JSON.stringify(res.msg))
                            }
                        }).catch((error)=>{
                        console.log(error)
                    })
                }
                // 编辑　小分类
                if (action === "edit" && name === "小分类" && this.name !== "") {
                    console.log(this.cate)
                    this.$axios.$put(cateUrl, {
                        id: this.id,
                        name: this.name,
                        status: this.status,
                        big_id: this.cate.big_id,
                    }).then((res)=>{
                        if (res.statusCode === 200) {
                            this.id = "";
                            this.name = "";
                            this.visible = false;
                            this.status = true;
                            let arr = this._data.cates;
                            let name = arr.filter(item => item.id === res.data.id)[0].name;
                            arr.filter(item => item.id === res.data.id)[0].name = res.data.name;
                            this.$message.success("编辑 "+name+" "+res.msg);
                        }else {
                            this.id = "";
                            this.name = "";
                            this.visible = false;
                            this.status = true;
                            this.$message.error(JSON.stringify(res.msg))
                        }
                    }).catch((error)=>{
                        this.id = "";
                        this.name = "";
                        this.visible = false;
                        this.status = true;
                        console.log(error)
                    })
                }

                // 添加书签
                // 编辑书签

            },
            handleCancel() {
                this.name = "";
                this.id = "";
                this.visible = false;
            },
            changeHandle(value) {
                console.log(value)
                this.status = value
            },
            changeCateHandle (value) {
                console.log(value)
                this.cate.big_id = value
            }
        }
    }
</script>

<style scoped>

</style>
