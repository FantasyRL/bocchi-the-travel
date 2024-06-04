<script setup>
import { RouterLink } from "vue-router";
</script>

<template>
  <div class="create">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="创建行程"
      @back="() => $router.go(-1)"
    />
  </div>
  <div class="travelscard">
    <a-card title="已创建的行程" style="width: 90%">
      <template #extra><router-link to="/travels/1">查看详细</router-link></template>
      <p>这里塞一个显示已经创建的行程信息</p>
    </a-card>
  </div>
  <router-link to="/finish/1">行程结束debug</router-link>
</template>

<style scoped>
.travelscard {
  position: relative;
  margin-left: 5%;
}
</style>
