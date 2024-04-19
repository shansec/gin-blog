<template>
    <v-card class="mx-auto position" max-width="300">
      <v-img :src="profileInfo.img">
        <v-card-title>
          <v-col align="center">
            <v-avatar size="130" color="grey">
              <img :src="profileInfo.avatar">
            </v-avatar>
            <div class="ma-4 white--text">May</div>
          </v-col>
        </v-card-title>
        <v-divider></v-divider>
      </v-img>
      <v-col>
        <v-card-title>About Me:{{ profileInfo.name }}</v-card-title>
        <v-card-text>{{ profileInfo.desc }}</v-card-text>
      </v-col>
      <v-divider color="indigo"></v-divider>
      <v-list nav dense>
        <v-list-item>
          <v-list-item-icon>
            <v-icon color="blue darken-2">{{ 'mdi-qqchat' }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey--text">{{ profileInfo.qqChat }}</v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-icon>
            <v-icon color="green darken-2">{{ 'mdi-wechat' }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey--text">{{ profileInfo.wechat }}</v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-icon>
            <v-icon color="orange darken-2">{{ 'mdi-sina-weibo' }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey--text">{{ profileInfo.weibo }}</v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-icon>
            <v-icon color="blue darken-2">{{ 'mdi-youtube' }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey--text">{{ profileInfo.bili }}</v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-icon>
            <v-icon color="indigo darken-2">{{ 'mdi-email' }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey--text">{{ profileInfo.email }}</v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card>
</template>

<script>
export default {
  name: "NavView",
  data() {
    return {
      profileInfo: {
        id: 1,
      }
    }

  },
  created() {
    this.getProfileInfo()
  },
  methods: {
    /**
     * 获取个人设置
     * @returns {Promise<void>}
     */
    async getProfileInfo() {
      const {data: res} = await this.$http.get(`profile/get/${this.profileInfo.id}`)

      this.profileInfo = res.data
    },
  }
}
</script>

<style scoped>
/*.position {*/
/*  position: fixed;*/
/*}*/
</style>