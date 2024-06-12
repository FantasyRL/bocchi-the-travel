<script setup>
import { h } from "vue";
import axios from "axios";
import Cookies from "js-cookie";
import { CheckOutlined, TeamOutlined } from "@ant-design/icons-vue";
</script>
<script>
export default {
  data() {
    return {
      trnumber: 10,
      access_token: 0,
      items: [{}]
    };
  },
  methods: {
    tohave() {
      this.$router.push(`/alltravelshave/`);
    },
    tomember(i) {
      this.$router.push(`/member/${i}`);
    },
    ToEnd(i) {
      axios.get("https://api.xiey.work/bocchi/party/status?party_id=" + i + "&action_type=1", {
        headers: { "access-token": this.access_token }
      });
      this.$router.push(`/finish/${i}`);
    },
    applylist(pagenum) {
      axios
        .get(
          "https://api.xiey.work/bocchi/party/party/my?page_num=" + pagenum,

          {
            headers: {
              "access-token": this.access_token
            }
          }
        )
        .then((response) => {
          console.log(response.data);
          this.items = response.data.party_list;
          console.log(this.items);
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
    this.applylist(1);
  },
  computed: {
    sortedItems() {
      return this.items.sort((a, b) => b.id - a.id);
    }
  }
};
</script>

<template>
  <div class="travels">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="所有行程"
      @back="() => $router.go(-1)"
    />
  </div>
  <div class="card-container" @click="tohave">
    <a-card
      title="查看已完成的行程"
      style="width: 90%; margin: 5px 5px 5px; justify-content: center; text-align: center"
    >
    </a-card>
  </div>
  <div class="center">
    <div v-for="item in sortedItems" :key="item.id" class="item">
      <div v-if="!item.status">
        <a-divider orientation="left" class="separate">活动名: {{ item.title }}</a-divider>
        <div style="padding: 20px" class="relative">
          <a-row>
            <a-card class="relative-card" title="活动详情" :bordered="false">
              <template #extra
                ><router-link :to="`/partys/${item.id}`"> 查看详细 </router-link></template
              >
              <router-link :to="`/partys/${item.id}`">
                <a-row>
                  <!-- <a-col class="inner-img"><img :src="getImgUrl(1)" alt="" /></a-col> -->
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
                <a-button
                  style="margin-right: 5px"
                  :icon="h(TeamOutlined)"
                  @click="tomember(item.id)"
                  >查看成员</a-button
                >
                <a-button
                  type="primary"
                  @click="ToEnd(item.id)"
                  style="margin-right: 5px"
                  :icon="h(CheckOutlined)"
                  >结束行程</a-button
                >
              </div>
            </a-card>
          </a-row>
        </div>
      </div>
    </div>
  </div>
  <br />
  <br />
  <br />
  <br />
  <br />
  <br />
  <br />
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
.buttongroup {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 10px;
  margin-bottom: 10px;
}
.card-container {
  display: flex;
  justify-content: center;
  cursor: pointer;
  margin-top: 10px;
}
</style>
