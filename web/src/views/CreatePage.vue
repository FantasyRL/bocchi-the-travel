<script setup>
import axios from "axios";
import Cookies from "js-cookie";
import { ref } from "vue";
import cityOptions from "@/city";
import { useCounterStore } from "@/stores/api";
const title = ref();
const type = ref();
const content = ref();
const partytime = ref();
const city = ref();
const options = ref([{ value: "放松" }, { value: "深度" }, { value: "人文" }]);
</script>

<script>
const setapiurl = useCounterStore();
export default {
  data() {
    return {
      trnumber: 10,
      access_token: "",
      refresh_token: "",
      fieldNames: { label: "label", value: "label", children: "children" },
      cityOptions: cityOptions
    };
  },
  methods: {
    partycreate(title, type, content, city, time) {
      const dateTimeStr = time[0];
      const date = new Date(dateTimeStr);
      const oldday = date.toISOString().slice(0, 50);

      const dateTimeStr2 = time[1];
      const date2 = new Date(dateTimeStr2);
      const newday = date2.toISOString().slice(0, 50);
      console.log(time);
      if (city["2"] === undefined) {
        console.log("未定义");
        city["2"] = "";
      }
      axios
        .post(
          setapiurl.apiurl +
            "/party/create?title=" +
            title +
            "&content=" +
            content +
            "&type=" +
            type +
            "&province=" +
            city["0"] +
            "&city=" +
            city["1"] +
            city["2"] +
            "&start_time=" +
            oldday +
            "&end_time=" +
            newday,
          {},
          {
            headers: {
              "access-token": this.access_token
            }
          }
        )
        .then((response) => {
          console.log(response.data);
          if (response.data.base.code == 10000) {
            console.log("创建成功");
            this.$router.push("/partys/" + response.data.party.id);
          }
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
  },
  watch: {}
};
</script>

<template>
  <div class="create">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="创建活动"
      @back="() => $router.go(-1)"
    />
  </div>
  <div class="create">
    <a-divider orientation="left" class="separate">名称</a-divider>
    <div class="input-box">
      <a-input v-model:value="title" :bordered="false" size="large" placeholder="标题" />
    </div>

    <a-divider orientation="left" class="separate">城市</a-divider>
    <a-form-item class="input-box">
      <a-cascader
        :options="cityOptions"
        placeholder="请选择地区"
        :fieldNames="fieldNames"
        v-model:value="city"
        size="large"
        :bordered="false"
      />
    </a-form-item>

    <a-divider orientation="left" class="separate">旅行心情</a-divider>
    <div class="input-box">
      <a-select
        v-model:value="type"
        mode="multiple"
        size="large"
        style="width: 100%"
        :bordered="false"
        placeholder="请选择你的旅行心情"
        :options="options"
      />
    </div>

    <a-divider orientation="left" class="separate">简介</a-divider>
    <div class="input-box">
      <a-textarea
        v-model:value="content"
        placeholder="请输入活动简介"
        :rows="4"
        size="large"
        :bordered="false"
      />
    </div>

    <a-divider orientation="left" class="separate">时间</a-divider>
    <div class="time-box">
      <a-range-picker v-model:value="partytime" size="large" :bordered="false" />
    </div>

    <br /><br />

    <div class="start-button">
      <button class="pushable" @click="partycreate(title, type, content, city, partytime)">
        <span class="shadow"></span>
        <span class="edge"></span>
        <span class="front"> 开启旅途 </span>
      </button>
    </div>
  </div>
  <br /><br /><br /><br /><br />
</template>

<style scoped>
.create {
  position: relative;
  max-width: 100vw;
}
.ant-cascader-menus {
  max-width: 20px;
}
.travels-card {
  display: flex;
  justify-content: center;
}

.separate {
  margin-top: 20px;
  font-weight: bold;
  font-size: 24px;
  margin-bottom: 0;
}

.pushable {
  position: relative;
  background: transparent;
  padding: 0;
  border: none;
  cursor: pointer;
  outline-offset: 4px;
  transition: filter 250ms;
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
  bottom: 15%;
}

.shadow {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background: hsl(226, 25%, 69%);
  border-radius: 8px;
  filter: blur(2px);
  will-change: transform;
  transform: translateY(2px);
  transition: transform 600ms cubic-bezier(0.3, 0.7, 0.4, 1);
}

.input-box {
  font-size: 24px;
  border-radius: 8px;
  border: 3px solid #d9d9d9;
  margin: 5%;
}

.time-box {
  display: flex;
  font-size: 24px;
  border-radius: 8px;
  border: 3px solid #d9d9d9;
  justify-content: center;
  width: fit-content;
  margin: 5% auto auto;
}

.start-button {
  display: flex;
  justify-content: center;
}

.edge {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  border-radius: 8px;
  background: linear-gradient(
    to right,
    hsl(248, 39%, 39%) 0%,
    hsl(248, 39%, 49%) 8%,
    hsl(248, 39%, 39%) 92%,
    hsl(248, 39%, 29%) 100%
  );
}

.front {
  display: block;
  position: relative;
  border-radius: 8px;
  background: hsl(248, 53%, 58%);
  padding: 16px 32px;
  color: white;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell,
    "Open Sans", "Helvetica Neue", sans-serif;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 1.5px;
  font-size: 1rem;
  transform: translateY(-4px);
  transition: transform 600ms cubic-bezier(0.3, 0.7, 0.4, 1);
}

.pushable:hover {
  filter: brightness(110%);
}

.pushable:hover .front {
  transform: translateY(-6px);
  transition: transform 250ms cubic-bezier(0.3, 0.7, 0.4, 1.5);
}

.pushable:active .front {
  transform: translateY(-2px);
  transition: transform 34ms;
}

.pushable:hover .shadow {
  transform: translateY(4px);
  transition: transform 250ms cubic-bezier(0.3, 0.7, 0.4, 1.5);
}

.pushable:active .shadow {
  transform: translateY(1px);
  transition: transform 34ms;
}

.pushable:focus:not(:focus-visible) {
  outline: none;
}
</style>
