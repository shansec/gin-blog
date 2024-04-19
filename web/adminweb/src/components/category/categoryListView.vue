<template>
  <div>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="1">
          <a-button type="primary" @click='addCategoryVisible = true'>新增分类</a-button>
        </a-col>
      </a-row>

      <a-table
          rowKey="ID"
          :columns="columns"
          :pagination="pagination"
          :dataSource="categoryList"
          bordered
          @change="handleTableChange"
      >
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button
                type="primary"
                icon="edit"
                style="margin-right: 15px"
                @click="editCategory(data.ID)"
            >编辑
            </a-button>
            <a-button
                type="danger"
                icon="delete"
                style="margin-right: 15px"
                @click="deleteCategory(data.ID)"
            >删除
            </a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增分类区域 -->
    <a-modal
        closable
        title="新增分类"
        :visible="addCategoryVisible"
        width="60%"
        @ok="addCategoryOk"
        @cancel="addCategoryCancel"
        destroyOnClose
    >
      <a-form-model :model="newCategory" :rules="addCategoryRules" ref="addCategoryRef">
        <a-form-model-item label="分类名" prop="name">
          <a-input v-model="newCategory.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑分类区域 -->
    <a-modal
        closable
        destroyOnClose
        title="编辑分类"
        :visible="editCategoryVisible"
        width="60%"
        @ok="editCategoryOk"
        @cancel="editCategoryCancel"
    >
      <a-form-model :model="categoryInfo" :rules="categoryRules" ref="addCategoryRef">
        <a-form-model-item label="分类名" prop="name">
          <a-input v-model="categoryInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>
<script>
import dayjs from "dayjs";

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center'
  },
  {
    title: '注册时间',
    dataIndex: 'CreatedAt',
    width: '20%',
    key: 'CreateAt',
    align: 'center',
    customRender: (val) => {
      return val ? dayjs(val).format('YYYY年MM月DD日') : '暂无'
    }
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    align: 'center',
    scopedSlots: {customRender: 'action'}
  }
]
export default {
  name: "categoryListView",
  data() {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      categoryInfo: {
        name: ''
      },
      categoryList: [],
      newCategory: {
        name: ''
      },
      columns,
      queryParam: {
        name: '',
        pagesize: 5,
        pagenum: 1
      },
      ediVisible: false,
      categoryRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.categoryInfo.name === '') {
                callback(new Error('请输入分类名'))
              }
              if ([...this.categoryInfo.name].length < 4
                  || [...this.categoryInfo.name].length > 12) {
                callback(new Error('分类名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          }
        ]
      },
      addCategoryRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.newCategory.name === '') {
                callback(new Error('请输入分类名'))
              }
              if ([...this.newCategory.name].length < 4
                  || [...this.newCategory.name].length > 12) {
                callback(new Error('用户名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur',
          }
        ]
      },
      editCategoryVisible: false,
      addCategoryVisible: false
    }
  },
  created() {
    this.getCategoryList()
  },
  methods: {
    /**
     * 获取分类列表
     * @returns {Promise<void>}
     */
    async getCategoryList() {
      const {data: res} = await this.$http.get('category/get', {
        params: {
          name: this.queryParam.name,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        this.$message.error(res.message)
      }
      this.categoryList = res.data
      this.pagination.total = res.total
    },
    /**
     * 更改分页
     * @param pagination
     * @param filters
     * @param sorter
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
      this.getCategoryList()
    },
    /**
     * 删除分类
     * @param id
     */
    deleteCategory(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该分类吗？删除无法恢复',
        onOk: async () => {
          const {data: res} = await this.$http.delete(`category/${id}`)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$message.success('删除成功')
          this.getCategoryList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        }
      })
    },
    /**
     * 新增分类
     */
    addCategoryOk() {
      this.$refs.addCategoryRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const {data: res} = await this.$http.post('category/add', {
           name: this.newCategory.name,
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.addCategoryRef.resetFields()
        this.addCategoryVisible = false
        this.$message.success('添加分类成功')
        this.getCategoryList()
      })
    },
    addCategoryCancel() {
      this.$refs.addCategoryRef.resetFields()
      this.addCategoryVisible = false
      this.$message.info('新增分类已取消')
    },
    /**
     * 编辑分类
     * @param id
     * @returns {Promise<void>}
     */
    async editCategory(id) {
      this.editCategoryVisible = true
      const {data: res} = await this.$http.get(`category/getInfo/${id}`)
      this.categoryInfo = res.data
      this.categoryInfo.id = id
    },
    editCategoryOk() {
      this.$refs.addCategoryRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const {data: res} = await this.$http.put(`category/${this.categoryInfo.id}`, {
          name: this.categoryInfo.name,
        })
        if (res.status != 200) return this.$message.error(res.message)
        this.editCategoryVisible = false
        this.$message.success('更新分类信息成功')
        this.getCategoryList()
      })
    },
    editCategoryCancel() {
      this.editCategoryVisible = false
      this.$refs.addCategoryRef.resetFields()
      this.$message.info('编辑已取消')
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>