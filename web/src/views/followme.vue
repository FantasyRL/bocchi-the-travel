<script setup>
import Cookies from "js-cookie";
import axios from "axios";
</script>
<script>
export default {
  data() {
    return { items: [], token: null, access_token: Cookies.get("access_token") };
  },
  methods: {
    applyuser(i) {
      const token = this.access_token;
      axios
        .post(
          "https://api.xiey.work/bocchi/trust/action?object_uid=" + i + "&action_type=1",
          {},
          {
            headers: {
              "access-token": token
            }
          }
        )
        .then((res) => {
          console.log(res.data.msg);
          alert(res.data.base.msg);
          this.init();
        });
    },
    toinfo(i) {
      this.$router.push(`/about/${i}`);
    },
    init() {
      axios
        .get(
          "https://api.xiey.work/bocchi/trust/follower?page_num=1&user_id=" +
            Number(this.$route.params.id)
        )
        .then((res) => {
          console.log(res);
          console.log(res.data.base.msg);
          this.items = res.data.follower_list;
        });
    }
  },
  mounted() {
    this.init();
  }
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
    <div v-for="item in items" :key="item" style="display: grid; justify-content: center">
      <a-card hoverable style="width: 300px">
        <template #actions>
          <!-- <setting-outlined key="setting" />
        <edit-outlined key="edit" /> -->
          <div @click="applyuser(item.id)">回关</div>
        </template>
        <a-card-meta :title="item.name" :description="item.signature" @click="toinfo(item.id)">
          <template #avatar>
            <a-avatar :src="item.avatar" />
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
