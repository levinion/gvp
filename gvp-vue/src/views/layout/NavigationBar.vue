<template>
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
  >
    <el-menu-item index="0" @click="$router.replace({name:'home'})">GVP</el-menu-item>
    <el-menu-item index="1" @click="$router.replace({name:'about'})">about</el-menu-item>
    <div class="flex-grow"/>
    <el-sub-menu index="2" v-if="userInfo">
      <template #title>{{userInfo.name}}</template>
      <el-menu-item index="2-1" @click="$router.push({name:'profile'})">主页</el-menu-item>
      <el-menu-item index="2-2" @click="logout">登出</el-menu-item>
    </el-sub-menu>
    <el-sub-menu index="3" v-if="!userInfo">
        <template #title>
          menu
        </template>
        <el-menu-item index="3-1" v-if="$route.name!='register'"
        @click="$router.replace({name:'register'})">register</el-menu-item>
        <el-menu-item index="3-2" v-if="$route.name!='login'"
        @click="$router.replace({name:'login'})">login</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script>
import { mapActions, mapState } from 'vuex';

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: mapActions('userModule', ['logout']),
};
</script>
<style>
.flex-grow {
  flex-grow: 1;
}
</style>
