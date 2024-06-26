<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
              v-model="queryParam.title"
              placeholder="输入文章名查找"
              enter-button
              allowClear
              @search="getArticleList"
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="$router.push('addArticle')">新增</a-button>
        </a-col>

        <a-col :span="3">
          <a-select placeholder="请选择分类" style="width: 200px" @change="CateChange">
            <a-select-option v-for="item in categoryList" :key="item.ID" :value="item.ID">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-col>
        <a-col :span="4">
          <a-button type="info" @click="getArticleList">显示全部</a-button>
        </a-col>
      </a-row>

      <a-table
          rowKey="ID"
          :columns="columns"
          :pagination="pagination"
          :dataSource="articleList"
          bordered
          @change="handleTableChange"
      >
        <span class="ArtImg" slot="img" slot-scope="img">
          <img :src="img"/>
        </span>
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
                size="small"
                type="primary"
                icon="edit"
                style="margin-right: 15px"
                @click="$router.push(`addArticle/${data.ID}`)"
            >编辑
            </a-button>
            <a-button
                size="small"
                type="danger"
                icon="delete"
                style="margin-right: 15px"
                @click="deleteArticle(data.ID)"
            >删除
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
import day from 'dayjs'

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id',
    align: 'center',
  },
  {
    title: '分类',
    dataIndex: 'Category.name',
    width: '10%',
    key: 'name',
    align: 'center',
  },
  {
    title: '更新日期',
    dataIndex: 'UpdatedAt',
    width: '15%',
    key: 'UpdatedAt',
    align: 'center',
    customRender: (val) => {
      return val ? day(val).format('YYYY年MM月DD日') : '暂无'
    },
  },
  {
    title: '文章标题',
    dataIndex: 'title',
    width: '15%',
    key: 'title',
    align: 'center',
  },
  {
    title: '文章描述',
    dataIndex: 'desc',
    width: '20%',
    key: 'desc',
    align: 'center',
  },
  {
    title: '缩略图',
    dataIndex: 'img',
    width: '20%',
    key: 'img',
    align: 'center',
    scopedSlots: {customRender: 'img'},
  },
  {
    title: '操作',
    width: '15%',
    key: 'action',
    align: 'center',
    scopedSlots: {customRender: 'action'},
  },
]

export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['4', '8', '12'],
        pageSize: 4,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      articleList: [],
      categoryList: [],
      columns,
      queryParam: {
        title: '',
        pagesize: 5,
        pagenum: 1,
      },
    }
  },
  created() {
    this.getArticleList()
    this.getCategoryList()
  },
  methods: {
    /**
     * 获取文章列表
     * @returns {Promise<void>}
     */
    async getArticleList() {
      const {data: res} = await this.$http.get('article/get', {
        params: {
          title: this.queryParam.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      })
      if (res.status !== 200) {
        // eslint-disable-next-line no-constant-condition
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }

      this.articleList = res.data
      this.pagination.total = res.total
    },
    /**
     * 获取分类
     * @returns {Promise<MessageType>}
     */
    async getCategoryList() {
      const {data: res} = await this.$http.get('category/get')
      if (res.status !== 200) return this.$message.error(res.message)
      this.categoryList = res.data
      // this.pagination.total = res.total
    },
    /**
     * 更改分页
     * @param pagination
     */
    handleTableChange(pagination) {
      var pager = {...this.pagination}
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getArticleList()
    },
    /**
     * 删除文章
     * @param id
     */
    deleteArticle(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该文章吗？一旦删除，无法恢复',
        onOk: async () => {
          const {data: res} = await this.$http.delete(`article/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getArticleList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
    /**
     * 查询分类下的文章
     * @param value
     * @constructor
     */
    CateChange(value) {
      this.getCategoryArt(value)
    },
    async getCategoryArt(id) {
      const {data: res} = await this.$http.get(`article/cateArt/${id}`, {
        params: {pagesize: this.queryParam.pagesize, pagenum: this.queryParam.pagenum},
      })
      if (res.status !== 200) return this.$message.error(res.message)
      this.articleList = res.data
      this.pagination.total = res.total
    },
  },
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}

.ArtImg {
  height: 100%;
  width: 100%;
}

.ArtImg img {
  width: 100px;
  height: 80px;
}
</style>
