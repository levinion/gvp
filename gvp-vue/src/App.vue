<template>
  <NavigationBar/>
  <router-view/>
</template>
<script>
import NavigationBar from './views/layout/NavigationBar.vue';

export default {
  components: { NavigationBar },
  data() {
    return {};
  },
};
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

</style>
