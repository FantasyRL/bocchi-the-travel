<script setup>
import axios from "axios";
import Cookies from "js-cookie";
import { watch, ref } from "vue";
import cityOptions from "@/city";

const value5 = ref("");
const value1 = ref("");
const value2 = ref("");
const value3 = ref("");
const value4 = ref("");

watch(value3, () => {
  const dateTimeStr = value3.value[0];
  const date = new Date(dateTimeStr);
  const oldday = date.toISOString().slice(0, 10);
  console.log(oldday);
  const dateTimeStr2 = value3.value[1];
  const date2 = new Date(dateTimeStr2);
  const newday = date2.toISOString().slice(0, 10);
  console.log(newday);
  Cookies.set("oldday", oldday);
  Cookies.set("newday", newday);
});
</script>
<script>
export default {
  setup() {},
  props: {
    value: {}
  },
  data() {
    return {
      trnumber: 10,
      access_token: "", // 假设您将令牌存储在localStorage中
      refresh_token: "",
      fieldNames: { label: "label", value: "label", children: "children" }, //重置默认字段
      cityOptions: cityOptions, //数据
      newday: "",
      oldday: "",
      couter: 0,
      value1: "", // 标题
      value2: "", // 类型
      value3: "", // 日期选择器
      value4: "", // 地点选择器
      value5: "", // 标题
      value6: "" // 标题
    };
  },
  methods: {
    partycreate() {
      var myHeaders = new Headers();
      myHeaders.append(
        "access-token",
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc2MDM1NzUsIm9yaWdfaWF0IjoxNzE3NTE3MTc1LCJ1c2VyX2lkIjo2fQ.6AokVBT1vIFkA5Qrh16K91fNombLkkd69HqSqf9jdKc"
      );
      myHeaders.append("User-Agent", "Apifox/1.0.0 (https://apifox.com)");

      var requestOptions = {
        method: "POST",
        headers: myHeaders,
        redirect: "follow"
      };

      fetch(
        "/bocchi/party/create?title=1&content=1&type=1&province=1&city=1&start_time=2006-01-02&end_time=2006-01-02",
        requestOptions
      )
        .then((response) => response.text())
        .then((result) => console.log(result))
        .catch((error) => console.log("error", error));
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
    value3(newValue, oldValue) {
      console.log(`couter changed from ${oldValue} to ${newValue}`);
    }
  }
};
</script>
<template>
  <button @click="couter++">增加</button>
  <div class="create">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="创建行程"
      @back="() => $router.go(-1)"
    />
  </div>
  <div class="travelscard">
    <a-card title="已创建的行程" style="width: 90%">
      <template #extra><router-link to="/alltravels/">查看详细</router-link></template>
      <p>这里塞一个显示已经创建的行程信息</p>
      <p>
        {{ access_token }}
        {{ couter }}
        debug:
        {{ value5 }}
        {{ value1 }}
        {{ value2 }}
        {{ value3 }}
        <br />

        {{ value4 }}
      </p>
    </a-card>
  </div>

  <div class="center">
    <text>行程的名字</text>
    <br />
    <a-input v-model:value="value5" placeholder="标题" />
    <br />

    <text>行程的类型</text>
    <br />
    <a-input v-model:value.lazy="value2" autofocus placeholder="类型" />
    <br />

    <text> 行程的介绍</text>
    <br />
    <a-textarea v-model:value="value1" placeholder="介绍" :rows="4" />
    <br />
    <a-range-picker v-model:value="value3" />
    <!-- <a-textarea v-model:value="value3" :rows="4" placeholder="maxlength is 6" :maxlength="6" />
    <br /> -->
    <br />
    <text>活动省份</text>
    <a-form-item>
      <a-cascader
        :options="cityOptions"
        placeholder="请选择地区"
        :fieldNames="fieldNames"
        v-model:value.lazy="value4"
      />
    </a-form-item>
    <br />
    <br />
    <button class="pushable" @click="partycreate()">
      <span class="shadow"></span>
      <span class="edge"></span>
      <span class="front"> 创建行程 </span>
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
