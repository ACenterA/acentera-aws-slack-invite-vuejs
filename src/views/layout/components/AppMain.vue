<template>
  <transition name="fade-transform" mode="out-in">
    <keep-alive>
      <router-view :key="key"/>
    </keep-alive>
  </transition>
</template>

<script>
export default {
  name: 'AppMain',
  computed: {
    key() {
      return this.$route.fullPath
    }
  }
}
</script>

<style scoped>
.app-main {
  /*84 = navbar + tags-view = 50 +34 */
  min-height: calc(100vh - 84px);
  width: 100%;
  position: relative;
  overflow: hidden;
}
</style>

<style rel="stylesheet/scss" lang="scss">
.base-grey {
  background-color: rgb(240, 242, 245);
}
</style>
