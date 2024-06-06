<script setup>
import { h } from "vue";
import axios from "axios";
import Cookies from "js-cookie";
import { CheckOutlined, ReconciliationOutlined, TeamOutlined } from "@ant-design/icons-vue";
const baseUrl = "https://severj.top/img/";
const getImgUrl = (i) => {
  return `${baseUrl}background${i}.webp`;
};
</script>
<script>
export default {
  data() {
    return {
      trnumber: 10,
      access_token: 0, // 假设您将令牌存储在localStorage中
      items: [{}]
    };
  },
  methods: {
    ToEnd(i) {
      this.$router.push(`/finish/${i}`);
    },
    applylist(pagenum) {
      axios
        .get(
          "/bocchi/party/party/my?page_num=" + pagenum,

          {
            headers: {
              "access-token": this.access_token
            }
          }
        )
        .then((response) => {
          console.log(response.data);
          this.items = response.data.party_list; // 假设响应数据中有一个名为"list"的数组包含所有行程对象
        })
        .catch((error) => {
          console.log(error);
        });
    }
  },
  mounted() {
    this.id = Cookies.get("id");
    this.access_token = Cookies.get("access_token");
    this.refresh_token = Cookies.get("refresh_token");
    this.applylist(1); // 调用函数
  }
};
</script>
<!-- 我的行程页面  -->
<template>
  <div class="travels">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="所有行程"
      @back="() => $router.go(-1)"
    />
  </div>

  <div class="center">
    <div v-for="item in items" :key="item.id" class="item">
      <a-divider orientation="left" class="separate">{{ item.title }}</a-divider>
      <div style="padding: 20px" class="relative">
        <a-row>
          <a-card class="relative-card" title="活动详情" :bordered="false">
            <router-link :to="`/partys/${item.id}`">
              <a-row>
                <a-col class="inner-img"><img :src="getImgUrl(1)" alt="" /></a-col>
                <a-col class="inner-word">
                  <b>行程ID：</b>{{ item.id }}
                  <br />
                  <b>创建者ID：</b>{{ item.founder_id }}
                  <br />
                  <b>简介：</b> {{ item.content }}
                  <br />
                  <b>类型：</b> {{ item.type }}
                  <br />
                  <b>城市：</b>{{ item.province }}，{{ item.city }}
                  <br />
                  <b>成员人数：</b> {{ item.member_count }}
                  <br />
                </a-col>
              </a-row>
            </router-link>
            <a-divider></a-divider>
            <div class="buttongroup">
              <a-button style="margin-right: 5%" :icon="h(ReconciliationOutlined)"
                >查看申请</a-button
              >
              <a-button style="margin-right: 5%" :icon="h(TeamOutlined)">查看成员</a-button>
              <a-button @click="ToEnd(item.id)" style="margin-right: 5%" :icon="h(CheckOutlined)"
                >结束行程</a-button
              >
            </div>
          </a-card>
        </a-row>
      </div>
    </div>

    <br /><br /><br /><br /><br />
    <br />
  </div>
</template>

<style scoped>
.separate {
  margin-top: 22px;
  font-weight: bold;
  font-size: 20px;
  margin-bottom: 0;
}
.relative-card {
  width: 100%;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}
.inner-img img {
  width: 150px;
  height: 150px;
  object-fit: cover;
  object-position: center;
  border-radius: 8px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.3);
}
.inner-word {
  margin-left: 20px;
  font-size: 15px;
}
</style>
