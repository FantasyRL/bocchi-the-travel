<script setup>
import axios from "axios";
import Cookies from "js-cookie";
import { ref, watch } from "vue";

const title = ref("");
const action_type = ref("");

const partytime = ref();

const rectangle = ref();
const route_json = ref();
const remark = ref();
const options = ref([
  { value: "1", label: "路线" },
  { value: "2", label: "活动" },
  { value: "3", label: "住处" },
  { value: "4", label: "吃喝" },
  { value: "5", label: "其他" }
]);
</script>

<script>
import dayjs from "dayjs";
import "dayjs/locale/zh-cn";
import locale from "ant-design-vue/es/date-picker/locale/zh_CN";

dayjs.locale("zh-cn");
export default {
  setup() {
    return {
      value: dayjs("2015-01-01", "YYYY-MM-DD"),
      dayjs,
      locale
    };
  },
  data() {
    return {
      trnumber: 10,
      access_token: "", // 假设您将令牌存储在localStorage中
      refresh_token: "",
      partyid: null
    };
  },
  methods: {
    partycreate(title, action_type, party_id, rectangle, route_json, remark, time) {
      const dateTimeStr = time[0];
      const date = new Date(dateTimeStr);
      const oldday = date.toISOString().slice(0, 50);

      const dateTimeStr2 = time[1];
      const date2 = new Date(dateTimeStr2);
      const newday = date2.toISOString().slice(0, 50);
      console.log(time);
      axios
        .post(
          "/bocchi/party/itinerary/create?title=" +
            title +
            "&action_type=" +
            action_type +
            "&party_id=" +
            party_id +
            "&rectangle=" +
            rectangle +
            "&route_json=" +
            route_json +
            "&remark=" +
            remark +
            "&schedule_start_time=" +
            oldday +
            "&schedule_end_time=" +
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
            this.$router.go(-1); // 返回上一页
          }
        })
        .catch((error) => {
          console.log(error);
        });
    }
  },
  mounted() {
    this.partyid = Number(this.$route.params.id);
    this.access_token = Cookies.get("access_token");
    this.refresh_token = Cookies.get("refresh_token");
  },
  computed: {},
  watch: {}
};
</script>

<template>
  <div class="create">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="创建计划"
      @back="() => $router.go(-1)"
    />
  </div>

  <div class="create">
    <a-divider orientation="left" class="separate">计划名</a-divider>
    <div class="input-box">
      <a-input v-model:value="title" :bordered="false" size="large" placeholder="标题" />
    </div>
    <a-divider orientation="left" class="separate">活动类型</a-divider>
    <div class="input-box">
      <a-select
        v-model:value="action_type"
        size="large"
        style="width: 100%"
        :bordered="false"
        placeholder="请选择你的活动类型"
        :options="options"
      />
    </div>
    <a-divider orientation="left" class="separate" v-if="!(action_type - 1)">路线</a-divider>
    <div v-if="!(action_type - 1)">假装有一个路线图</div>
    <a-divider orientation="left" class="separate" v-if="action_type - 1">地点</a-divider>
    <div v-if="action_type - 1">假装有一个地图</div>
    <a-divider orientation="left" class="separate">备注</a-divider>
    <div class="input-box">
      <a-textarea
        v-model:value="remark"
        placeholder="备注"
        :rows="4"
        size="large"
        :bordered="false"
      />
    </div>

    <a-divider orientation="left" class="separate">起止时间</a-divider>
    <div class="time-box">
      <a-range-picker
        v-model:value="partytime"
        size="large"
        :bordered="false"
        :show-time="{ format: 'HH:mm' }"
      />
    </div>

    <br /><br />

    <div class="start-button">
      <button
        class="pushable"
        @click="partycreate(title, action_type, partyid, rectangle, route_json, remark, partytime)"
      >
        <span class="shadow"></span>
        <span class="edge"></span>
        <span class="front"> 创建 </span>
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
  max-width: 20px; /* 设置为固定宽度 */
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
