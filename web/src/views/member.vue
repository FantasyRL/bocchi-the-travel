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
      id: this.$route.params.id,
      members: {}, // 用于存储活动成员信息
      applylist: {} // 用于存储申请列表信息
    };
  },
  methods: {
    applyuser(userid) {
      axios
        .get(
          setapiurl.apiurl + "/party/apply/permit?party_id=" + this.id + "&member_id=" + userid,
          {
            headers: { "access-token": Cookies.get("access_token") }
          }
        )
        .then((response) => {
          console.log(response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    get_party_members(party_id, page_num) {
      axios
        .get(setapiurl.apiurl + "/party/members?party_id=" + party_id + "&page_num=" + page_num, {
          headers: {}
        })
        .then((response) => {
          console.log(response);
          this.members = response.data.member_list;
        })
        .catch((error) => {
          console.log(error);
        });
    },
    apply_list(party_id, page_num) {
      axios
        .get(
          setapiurl.apiurl + "/party/apply/list?party_id=" + party_id + "&page_num=" + page_num,
          {
            headers: {
              "access-token": Cookies.get("access_token")
            }
          }
        )
        .then((response) => {
          console.log(response.data); // 打印返回的数据
          this.applylist = response.data.applicant_list;
        })
        .catch((error) => {
          console.log(error);
        });
    }
  },
  mounted() {
    this.get_party_members(this.id, 1);
    this.apply_list(this.id, 1);
  }
};
</script>
<script setup></script>
<template>
  <div class="travels">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="活动成员"
      @back="() => $router.go(-1)"
    />
  </div>

  <div class="main">
    <div class="member">
      <div style="text-align: center; margin-bottom: 5vw">团队成员</div>
      <div v-for="item in members" :key="item" style="display: grid; justify-content: center">
        <a-card hoverable style="width: 300px; justify-content: center" v-if="item.id">
          <template #actions>
            <!-- <div @click="applyuser()">同意加入</div> -->
          </template>
          <a-card-meta
            :title="item.name"
            :description="item.signature"
            @click="this.$router.push(`/about/${item.id}`)"
          >
            <template #avatar>
              <a-avatar :src="item.avatar" />
            </template>
          </a-card-meta>
        </a-card>
      </div>
    </div>
    <div class="applylist">
      <div style="text-align: center; margin-bottom: 5vw">申请列表:</div>
      <div v-for="good in applylist" :key="good" style="display: grid; justify-content: center">
        <a-card hoverable style="width: 300px" v-if="good.id">
          <template #actions>
            <!-- <setting-outlined key="setting" />
            <edit-outlined key="edit" /> -->
            <div @click="applyuser(good.id)">同意加入</div>
          </template>
          <a-card-meta
            :title="good.name"
            :description="good.signature"
            @click="this.$router.push(`/about/${good.id}`)"
          >
            <template #avatar>
              <a-avatar :src="good.avatar" />
            </template>
          </a-card-meta>
        </a-card>
      </div>
    </div>
  </div>
</template>

<style scoped>
.main {
  display: grid;
  justify-content: center;
  grid-template-columns: 1fr; /* 设置两个列，各占1份空间 */
  margin-top: 5vw;
}
.member {
  justify-content: center;
}
.applylist {
  justify-content: center;
  margin-top: 5vw;
}
</style>
