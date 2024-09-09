<script setup>
import loginpage from "../module/app/login.vue";
import homebar from "../module/app/homebar.vue";
import { RouterView } from "vue-router";
import "mdui/components/icon.js";
</script>

<template>
  <loginpage></loginpage>
  <!-- 登录,注册界面 -->

  <div style="width: 100%; height: 32px"></div>
  <!-- 这是给手机端预留的状态栏 -->

  <router-view></router-view>
  <!-- 路由视图 -->

  <homebar></homebar>
  <!-- 底栏 -->
</template>

<style></style>
