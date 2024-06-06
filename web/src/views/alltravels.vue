<script setup>
import axios from "axios";
import Cookies from "js-cookie";
</script>
<script>
export default {
  data() {
    return {
      trnumber: 10,
      access_token: 0, // 假设您将令牌存储在localStorage中
      items: [
        {
          id: 1,
          founder_id: 1,
          title: "标题",
          content: "内容",
          type: "类型",
          province: "省份",
          city: "城市",
          start_time: "2006-01-02",
          end_time: "2006-01-03",
          member_count: 0,
          status: 0,
          rectangle: ""
        }
      ]
    };
  },
  methods: {
    nono() {
      alert("我还没写");
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
      title="所有活动"
      @back="() => $router.go(-1)"
    />
  </div>
  <div v-for="item in items" :key="item.id" class="item">
    <br />
    <el-card style="min-width: 90vw">
      <template #header>
        <div class="card-header">
          <span>
            Title: {{ item.title }}
            <br />
            ID: {{ item.id }}
            <br />

            founder_id:{{ item.founder_id }}
            <br />
          </span>
        </div>
      </template>
      Content: {{ item.content }}
      <br />
      city: {{ item.city }}
      <br />
      province: {{ item.province }}
      <br />
      Content: {{ item.content }}

      <br />
      Start Time: {{ item.start_time }}
      <br />
      End Time: {{ item.end_time }}
      <br />
      member_count: {{ item.member_count }}
      <br />
      Type: {{ item.type }}
      <br />
      Rectangle: {{ item.rectangle }}
      <br />
      <el-button @click="nono">查看申请</el-button>
      <el-button @click="nono">查看成员</el-button>
      <el-button @click="nono(item.id)">删除这个活动</el-button>
      <template #footer>Status: {{ item.status }}</template>
    </el-card>
    <br />
  </div>
  <div class="center">
    <div v-for="item in items" :key="item.id" class="item">
      <br />
      <el-card style="min-width: 90vw">
        <template #header>
          <div class="card-header">
            <span>
              Title: {{ item.title }}
              <br />
              ID: {{ item.id }}
              <br />

              founder_id:{{ item.founder_id }}
              <br />
            </span>
          </div>
        </template>
        Content: {{ item.content }}
        <br />
        city: {{ item.city }}
        <br />
        province: {{ item.province }}
        <br />
        Content: {{ item.content }}

        <br />
        Start Time: {{ item.start_time }}
        <br />
        End Time: {{ item.end_time }}
        <br />
        member_count: {{ item.member_count }}
        <br />
        Type: {{ item.type }}
        <br />
        Rectangle: {{ item.rectangle }}
        <br />
        <el-button @click="nono">查看申请</el-button>
        <el-button @click="nono">查看成员</el-button>
        <el-button @click="nono(item.id)">删除这个活动</el-button>
        <template #footer>Status: {{ item.status }}</template>
      </el-card>
      <br />
    </div>

    <div>
      <!-- 循环项的数组 -->
    </div>
    <router-link to="/finish/1">行程结束debug</router-link>

    <br />
    <div class="null"></div>
    <br />
  </div>
</template>

<style scoped>
.center {
  display: grid;
  justify-items: center;
}
.null {
  height: 100px;
}
</style>
