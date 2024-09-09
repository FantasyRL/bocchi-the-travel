<script setup>
import Cookies from "js-cookie";
import axios from "axios";
import { useCounterStore } from "@/stores/api";
</script>
<script>
const setapiurl = useCounterStore();
export default {
  data() {
    return {
      items: [],
      token: null,
      access_token: Cookies.get("access_token")
    };
  },
  methods: {
    applyuser(i) {
      const token = this.access_token;
      axios
        .post(
          setapiurl.apiurl + "/trust/action?object_uid=" + i + "&action_type=0",
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
          this.initt();
        });
    },
    toinfo(i) {
      this.$router.push(`/about/${i}`);
    },
    initt() {
      axios
        .get(
          setapiurl.apiurl + "/trust/following?page_num=1&user_id=" + Number(this.$route.params.id)
        )
        .then((res) => {
          console.log(res);
          this.items = res.data.following_list;
        });
    }
  },
  mounted() {
    this.token = Cookies.get("access-token");
    this.initt();
  }
};
</script>

<template>
  <div class="travels">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="我的关注"
      sub-title="This is a subtitle"
      @back="() => $router.go(-1)"
    />
  </div>
  <div class="follow">
    <div v-for="item in items" :key="item" style="display: grid; justify-content: center">
      <a-card hoverable style="width: 300px">
        <template #actions>
          <div @click="applyuser(item.id)">取消关注</div>
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
