<template>
  <div>
    <a-button type="danger" @click="logout">退出</a-button>
  </div>
</template>

<script>
export default {
  name: "HeaderView",
  methods: {
    logout() {
      window.sessionStorage.clear()
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>

</style>