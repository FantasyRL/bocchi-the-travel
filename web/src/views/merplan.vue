<script setup>
import Cookies from "js-cookie";
import { ref } from "vue";
const open = ref(false);
const confirmLoading = ref(false);
const showModal = () => {
  open.value = true;
};
const handleOk = () => {
  confirmLoading.value = true;
  setTimeout(() => {
    open.value = false;
    confirmLoading.value = false;
  });
};
</script>

<script>
import axios from "axios";
import {
  HomeOutlined,
  CoffeeOutlined,
  EnvironmentOutlined,
  PlayCircleOutlined,
  QuestionCircleOutlined
} from "@ant-design/icons-vue";
import "./itinerary.vue";
import Cookies from "js-cookie";
export default {
  props: {},
  data() {
    return {
      id: 1, // 假设这是 itinerary 的 id
      party: {
        start_time: "2006-01-02",
        end_time: "2006-01-03"
      },
      info: {},
      partynull: false // 假设这是 party 是否为空的标志
    };
  },
  methods: {
    goItinerary(iid) {
      const url =
        "https://api.xiey.work/bocchi/party/itinerary/merge?itinerary_id=" +
        iid +
        "&party_id=" +
        this.id;
      axios
        .get(url, {
          headers: {
            "access-token": this.access_token
          }
        })
        .then((res) => {
          console.log(res); // 假设这是删除成功的回调函数，可以在这里进行页面跳转等操作
          location.reload();
        })
        .catch((err) => {
          console.error(err); // 假设这是删除失败的回调函数，可以在这里进行错误处理
        });
    },
    init() {
      axios
        .get("https://api.xiey.work/bocchi/party/itinerary/my?party_id=" + this.id, {
          headers: {
            "access-token": this.access_token
          }
        })
        .then((res) => {
          console.log(res);
          this.partynull = true;
          this.info = res.data.itinerary_list;
          if (res.data.base.code == 10000) {
            this.partynull = 1;
            console.log(this.partynull);
          }
        })
        .catch((err) => {
          console.error(err);
        });
    }
  },

  mounted() {
    this.id = Number(this.$route.params.id);
    this.access_token = Cookies.get("access_token");
    this.init();
  },
  computed: {
    getIcon() {
      return (actionType) => {
        switch (actionType) {
          case 1:
            return EnvironmentOutlined; // route
          case 2:
            return PlayCircleOutlined; // activity
          case 3:
            return HomeOutlined; // accommodation
          case 4:
            return CoffeeOutlined; // eat
          case 5:
            return QuestionCircleOutlined; // other
        }
      };
    },
    getType() {
      return (actionType) => {
        switch (actionType) {
          case 1:
            return "路线"; // route
          case 2:
            return "活动"; // activity
          case 3:
            return "住宿"; // accommodation
          case 4:
            return "餐饮"; // eat
          case 5:
            return "其他"; // other
          default:
            return "未知类型"; // unknown type
        }
      };
    }
  }
};
</script>
<template>
  <div class="travels">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="审核计划 "
      @back="() => $router.go(-1)"
    />
  </div>

  <br />

  <div class="itinerary">
    <div v-for="item in info" :key="item.id" class="item">
      <el-card style="width: 90vw" v-if="item.is_merged">
        <div class="item-info">
          <div>ID:{{ item.id }}</div>
          <div>标题:{{ item.title }}</div>
          <div>是否通过审核:{{ item.is_merged }}</div>
          <div>备注:{{ item.remark }}</div>
          <div>路线:{{ item.route_json }}</div>
          <div>地址:{{ item.rectangle }}</div>
          <div>开始时间:{{ item.schedule_start_time }}</div>
          <div>结束时间:{{ item.schedule_end_time }}</div>

          <button @click="goItinerary(item.id)">
            <span class="shadow"></span>
            <span class="edge"></span>
            <span class="front text"> 通过计划 </span>
          </button>
        </div>
      </el-card>
      <br />
    </div>
  </div>
  <div v-if="partynull" class="itinerary"></div>
  <div v-if="!partynull">
    <a-empty />
  </div>
  <br />
  <br />
  <br />
  <br />
  <br />
  <br />
  <br />
</template>

<style>
.item-info {
  display: grid;
  justify-items: center;
  row-gap: 10px;

  grid-template-columns: repeat(1, 1fr);
}
.itinerary {
  display: grid;
  justify-content: center;
}
button {
  position: relative;
  border: none;
  background: transparent;
  padding: 0;
  cursor: pointer;
  outline-offset: 4px;
  transition: filter 250ms;
  user-select: none;
  touch-action: manipulation;
}

.shadow {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 12px;
  background: hsl(0deg 0% 0% / 0.25);
  will-change: transform;
  transform: translateY(2px);
  transition: transform 600ms cubic-bezier(0.3, 0.7, 0.4, 1);
}

.edge {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 12px;
  background: linear-gradient(
    to left,
    hsl(145, 100%, 70%) 0%,
    hsl(153, 42%, 56%) 8%,
    hsl(113, 100%, 72%) 92%,
    hsl(162, 78%, 64%) 100%
  );
}

.front {
  display: block;
  position: relative;
  padding: 12px 27px;
  border-radius: 12px;
  font-size: 1.1rem;
  color: white;
  background: hsl(135, 54%, 62%);
  will-change: transform;
  transform: translateY(-4px);
  transition: transform 600ms cubic-bezier(0.3, 0.7, 0.4, 1);
}

button:hover {
  filter: brightness(110%);
}

button:hover .front {
  transform: translateY(-6px);
  transition: transform 250ms cubic-bezier(0.3, 0.7, 0.4, 1.5);
}

button:active .front {
  transform: translateY(-2px);
  transition: transform 34ms;
}

button:hover .shadow {
  transform: translateY(4px);
  transition: transform 250ms cubic-bezier(0.3, 0.7, 0.4, 1.5);
}

button:active .shadow {
  transform: translateY(1px);
  transition: transform 34ms;
}

button:focus:not(:focus-visible) {
  outline: none;
}
</style>
