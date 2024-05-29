<script setup>
import {RouterLink} from "vue-router";
</script>

<template>
  <div class="create">
    <a-page-header
        style="border: 1px solid rgb(235, 237, 240)"
        title="创建行程"
        @back="() => $router.go(-1)"
    />
  </div>

  <router-link to="/finish/1">行程结束debug</router-link>
  <router-link to="/travels/1">我的行程debug</router-link>

</template>

<style scoped></style>
