<template>
  <div>
    <v-app-bar app color="indigo darken-2" flat>
      <v-avatar size="35" color="teal"/>
      <v-container class="py-0 fill-height justify-center">
        <v-btn small style="margin-right: 10px" text color="white">首页</v-btn>
        <v-btn small style="margin-right: 10px" v-for="item in categoryList" :key="item.ID" text color="white">
          {{ item.name }}
        </v-btn>
      </v-container>
      <v-spacer></v-spacer>
<!--      <v-responsive max-width="260" color="white">-->
<!--        <v-text-field dense flat hide-details solo-inverted rounded dark></v-text-field>-->
<!--      </v-responsive>-->
    </v-app-bar>
  </div>

</template>

<script>
export default {
  name: "HeaderView",
  data() {
    return {
      categoryList: []
    }
  },
  created() {
    this.getCategoryList()
  },
  methods: {
    /**
     * 获取全部分类
     * @returns {Promise<void>}
     */
    async getCategoryList() {
      const {data: res} = await this.$http.get("category/get")
      this.categoryList = res.data
    }
  },
}
</script>

<style scoped>

</style>