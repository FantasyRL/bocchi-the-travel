<script setup>
import Cookies from "js-cookie";
import axios from "axios";
</script>
<script>
export default {
  data() {
    return {};
  },
  methods: {},
  mounted() {}
};
</script>
<template>
  <div class="travels">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="我的粉丝"
      @back="() => $router.go(-1)"
    />
  </div>
  <div class="follow">
    <div v-for="item in 2" :key="item" style="display: grid; justify-content: center">
      <a-card hoverable style="width: 300px" @click="this.$router.push(`/about/${item}`)">
        <template #actions>
          <!-- <setting-outlined key="setting" />
        <edit-outlined key="edit" /> -->
          <div @click="applyuser(item)">取消关注</div>
        </template>
        <a-card-meta :title="1" :description="1">
          <template #avatar>
            <a-avatar :src="1" />
          </template>
        </a-card-meta>
      </a-card>
    </div>
  </div>
</template>

<style scoped>
.follow {
  justify-content: center;
  display: grid;
  grid-template-columns: repeat(auto-fill, 300px);
  gap: 30px 10px;
  padding: 10px 10px 10px 10px;
  margin-top: 5vw;
}
</style>
