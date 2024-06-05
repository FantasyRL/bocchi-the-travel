<script setup>
import axios from "axios";
import Cookies from "js-cookie";
</script>
<script>
export default {
  data() {
    return {
      trnumber: 10,
      access_token: 0 // 假设您将令牌存储在localStorage中
    };
  },
  methods: {
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
    <div v-for="trnumber in trnumber" :key="trnumber">
      <br />
      <el-card style="min-width: 90vw">
        <template #header>
          <div class="card-header">
            <span>Card name {{ trnumber }}</span>
          </div>
        </template>
        <p v-for="o in 4" :key="o" class="text item">{{ "List item " + o }}</p>
        <template #footer>Footer content</template>
      </el-card>
      <br />
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
