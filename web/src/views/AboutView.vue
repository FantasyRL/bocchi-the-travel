<script setup>
import { UserOutlined } from "@ant-design/icons-vue";
/* import { useCounterStore } from '@/stores/counter' */
/* const counter = useCounterStore() */
</script>

<template>
  <a-page-header
    style="border: 1px solid rgb(235, 237, 240)"
    title="个人界面"
    sub-title="副标题"
    @back="() => $router.go(-1)"
  />

  <!-- info -->
  <div class="info">
    <!-- 头像 start -->
    <div class="touxiang">
      <div class="testdiv6">
        <img src="//xiey.work/640.jpg" class="touxiangimg rounded-image" alt="头像" />
        <div><text class="name">用户名1111111111111111111</text></div>
      </div>
    </div>
    <!-- 头像 end -->
    <!-- 社交账号 -->
    <div class="shejiao">
      <text>社交账号</text>
      <text>社交账号</text>
      <text>社交账号</text>
      <text>社交账号</text>
    </div>
    <!-- 社交账号 end -->

    <!-- 个人信息 start -->
    <!--   <div class="ab2">
      <a-empty />
    </div> -->
    <!-- 个人信息 end -->
  </div>
  <!-- info end -->
</template>

<style scoped></style>
