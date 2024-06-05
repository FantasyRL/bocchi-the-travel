<script setup>
import axios from "axios";
import Cookies from "js-cookie";
import { watch, ref } from "vue";
import cityOptions from "@/city";

const title = ref("");
const value1 = ref("");
const type = ref("");
const partytime = ref("");
const city = ref("");
const oldday = ref("");

const couter = ref(0);

/* watch(partytime, () => {
  const dateTimeStr = partytime.value[0];
  const date = new Date(dateTimeStr);
  const oldday = date.toISOString().slice(0, 50);
  console.log(oldday);
  const dateTimeStr2 = partytime.value[1];
  const date2 = new Date(dateTimeStr2);
  const newday = date2.toISOString().slice(0, 50);
  console.log(newday);
  Cookies.set("oldday", oldday);
  Cookies.set("newday", newday);
}); */
</script>
<script>
export default {
  data() {
    return {
      trnumber: 10,
      access_token: "", // 假设您将令牌存储在localStorage中
      refresh_token: "",
      fieldNames: { label: "label", value: "label", children: "children" }, //重置默认字段
      cityOptions: cityOptions //数据
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
      if (city["2"] == undefined) {
        console.log("未定义");
        city["2"] = "";
      }
      axios
        .post(
          "/bocchi/party/create?title=" +
            title +
            "&content=" +
            content +
            "&type=" +
            type +
            "&province=" +
            city["0"] +
            "&city=" +
            city["0"] +
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
          Cookies.remove("oldday");
          Cookies.remove("newday");
        })
        .catch((error) => {
          console.log(error);
          Cookies.remove("oldday");
          Cookies.remove("newday");
        });
    }
  },
  mounted() {
    this.id = Cookies.get("id");
    this.access_token = Cookies.get("access_token");
    this.refresh_token = Cookies.get("refresh_token");
  },
  watch: {
    //该方法会在数据变化，触发执行
    couter(newValue, oldValue) {
      console.log(`couter changed from ${oldValue} to ${newValue}`); // 打印变化后的值和变化前的值
    },
    partytime(newValue, oldValue) {
      console.log(`couter changed from ${oldValue} to ${newValue}`);
    }
  }
};
</script>
<template>
  <div class="create">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="创建party活动"
      @back="() => $router.go(-1)"
    />
  </div>
  <div class="travelscard">
    <a-card title="已创建的party活动" style="width: 90%">
      <template #extra><router-link to="/alltravels/">查看详细</router-link></template>
      <p>这里塞一个最新创建活动</p>
      <!-- <p>
        <br />
        {{ access_token }}
        <br />
        couter:
        {{ couter }}
        <br />
        title:{{ title }}
        <br />
        value1:{{ value1 }}
        <br />
        type:{{ type }}
        <br />
        partytime:{{ partytime }}
        <br />
        old:{{ oldday }}
        <br />
        <br />

        city:{{ city }}
      </p> -->
    </a-card>
  </div>

  <div class="center">
    <text>party活动的名字</text>
    <br />
    <a-input v-model:value="title" placeholder="标题" />
    <br />

    <text>party活动的类型</text>
    <br />
    <a-input v-model:value="type" autofocus placeholder="类型" />
    <br />

    <text> party活动的介绍</text>
    <br />
    <a-textarea v-model:value="content" placeholder="介绍" :rows="4" />
    <br />
    <a-range-picker v-model:value="partytime" />
    <!-- <a-textarea v-model:value="partytime" :rows="4" placeholder="maxlength is 6" :maxlength="6" />
    <br /> -->
    <br />
    <text>活动省份</text>
    <a-form-item>
      <a-cascader
        :options="cityOptions"
        placeholder="请选择地区"
        :fieldNames="fieldNames"
        v-model:value="city"
      />
    </a-form-item>
    <br />
    <br />
    <button class="pushable" @click="partycreate(title, type, content, city, partytime)">
      <span class="shadow"></span>
      <span class="edge"></span>
      <span class="front"> 创建party活动 </span>
    </button>
  </div>
</template>

<style scoped>
.center {
  display: grid;
  justify-items: center;
  margin-left: 5%;
  margin-right: 5%;
}
.travelscard {
  position: relative;
  margin-left: 5%;
}
.pushable {
  position: relative;
  background: transparent;
  padding: 0px;
  border: none;
  cursor: pointer;
  outline-offset: 4px;
  outline-color: deeppink;
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
