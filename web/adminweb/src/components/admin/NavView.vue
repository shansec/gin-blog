<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'My Blog' }}</span>
    </div>
    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index">
        <a-icon type="dashboard"></a-icon>
        <span>仪表盘</span>
      </a-menu-item>
      <a-sub-menu>
        <span slot="title"><a-icon type="file"/><span>文章管理</span></span>
        <a-menu-item key="addArticle">
          <a-icon type="form"/>
          <span>写文章</span>
        </a-menu-item>
        <a-menu-item key="artList">
          <a-icon type="snippets"/>
          <span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="categoryList">
        <a-icon type="book"/>
        <span>分类列表</span></a-menu-item>
      <a-menu-item key="userList">
        <a-icon type="user"/>
        <span>用户列表</span></a-menu-item>
      <a-menu-item key="profile">
        <a-icon type="setting"/>
        <span>个人设置</span></a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  name: "NavView",
  data() {
    return {
      collapsed: false
    }
  },
  methods: {
    goToPage(item) {
      this.$router.push('/' + item.key).catch((err) => err)
    }
  }
}
</script>

<style scoped>
.log {
  height: 32px;
  margin: 16px;
  background-color: #ffffff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 16px;
}
</style>