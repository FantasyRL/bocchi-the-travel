<script setup>
import { CheckOutlined } from "@ant-design/icons-vue";
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
    init() {
      const url = "/bocchi/party/itinerary/info?itinerary_id=" + this.id;
      const params = {};
      axios
        .get(url, params)
        .then((res) => {
          console.log(res);
          this.partynull = true;
          this.info = res.data.itinerary; // 假设这是返回的数据对象
          if (res.data.base.code == 10007) {
            this.partynull = 0;
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
      title="计划详情"
      @back="() => $router.go(-1)"
    />
  </div>

  <br />
  <div style="justify-content: center" v-if="partynull">
    <el-timeline style="max-width: 600px; margin-left: 10%">
      <el-timeline-item>计划名:{{ info.title }} </el-timeline-item>
      <el-timeline-item>序列:{{ info.sequence }} </el-timeline-item>
      <el-timeline-item> 创建者:{{ info.founder_id }} </el-timeline-item>
      <el-timeline-item> 类型：{{ getType(info.action_type) }} </el-timeline-item>
      <el-timeline-item> 备注：{{ info.remark }} </el-timeline-item>
      <el-timeline-item> 地点：{{ info.rectangle }} </el-timeline-item>
      <el-timeline-item> 开始时间：{{ info.schedule_start_time }} </el-timeline-item>
      <el-timeline-item>
        结束时间:
        {{ info.schedule_end_time }}
      </el-timeline-item>
    </el-timeline>
  </div>
  <div v-if="!partynull">
    <a-empty />
  </div>

  <div id="debug" v-if="0">
    <br />
    <br />
    <br />
    itinerary debug part:
    <div>itinerary id:{{ id }}</div>
    <div>原始数据:{{ info }}</div>
    <div>title:{{ info.title }}</div>
    <div>founder_id:{{ info.founder_id }}</div>
    <div>action_type:{{ info.action_type }}</div>
    <div>rectangle:{{ info.rectangle }}</div>
    <div>route_json:{{ info.route_json }}</div>
    <div>remark:{{ info.remark }}</div>
    <div>sequence:{{ info.sequence }}</div>
    <div>schedule_start_time:{{ info.schedule_start_time }}</div>
    <div>schedule_end_time:{{ info.schedule_end_time }}</div>
    <div>party_id:{{ info.party_id }}</div>
    <div>is_merged:{{ info.is_merged }}</div>
  </div>
</template>
