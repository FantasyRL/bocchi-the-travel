<script setup>
import {CheckOutlined} from "@ant-design/icons-vue";
import {ref} from "vue";
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
  })
};
</script>

<script>
import axios from "axios";
import { HomeOutlined, CoffeeOutlined, EnvironmentOutlined, PlayCircleOutlined, QuestionCircleOutlined} from "@ant-design/icons-vue";
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
      data:[
        {
          id: 2,
          title: "1",
          founder_id: 6,
          action_type: 2,
          rectangle: "1",
          route_json: "1",
          remark: "1",
          sequence: 0,
          schedule_start_time: "2006-01-02 15:40",
          schedule_end_time: "2006-01-02 15:40",
          party_id: 5,
          is_merged: 1
        } // 假设这是 itinerary 的数据对象
      ]
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
          this.data = res.data.itinerary; // 假设这是返回的数据对象
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
            return '路线'; // route
          case 2:
            return '活动'; // activity
          case 3:
            return '住宿'; // accommodation
          case 4:
            return '餐饮'; // eat
          case 5:
            return '其他'; // other
          default:
            return '未知类型'; // unknown type
        }
      }
    },
  },
};
</script>
<template>

  <div class="travels">
    <a-page-header
        style="border: 1px solid rgb(235, 237, 240)"
        title="行程详情"
        @back="() => $router.go(-1)"
    />
  </div>

  <br>
  <div style="justify-content: center">
    <a-timeline stroke-width="100%" mode="alternate" v-for="item in data" :key="item">
      <a-timeline-item color="red">
        <template #dot><component :is="getIcon(item.action_type)" style="font-size: 16px" /></template>
        {{item.title}}
        <br>
        类型：{{getType(item.action_type)}}
        <br>
        时间：{{item.schedule_start_time}}
        <br>
        地点：{{item.rectangle}}
        <br>
        备注：{{item.remark}}
        <br>
        <div style="
            border-radius: 12px;
            border: 3px solid #f5f5f5;
            width: fit-content;
            margin-top: 5px;
          ">
          <a-button type="link" @click="showModal">查看线路</a-button>
        </div>
        <a-modal v-model:open="open" title="活动线路" :confirm-loading="confirmLoading" @ok="handleOk">
          // 这里放地图
        </a-modal>
      </a-timeline-item>
      <a-timeline-item color="green">
        <template #dot><CheckOutlined style="font-size: 16px" /></template>
        活动结束<br>{{item.schedule_end_time}}
      </a-timeline-item>
    </a-timeline>
  </div>

  itinerary debug part:
  <div>itinerary id:{{ id }}</div>
  <div>原始数据:{{ data }}</div>
  <br />
  <br />
  <br />
  <div>tittle:{{ data.title }}</div>
  <div>founder_id:{{ data.founder_id }}</div>
  <div>action_type:{{ data.action_type }}</div>
  <div>rectangle:{{ data.rectangle }}</div>
  <div>route_json:{{ data.route_json }}</div>
  <div>remark:{{ data.remark }}</div>
  <div>sequence:{{ data.sequence }}</div>
  <div>schedule_start_time:{{ data.schedule_start_time }}</div>
  <div>schedule_end_time:{{ data.schedule_end_time }}</div>
  <div>party_id:{{ data.party_id }}</div>
  <div>is_merged:{{ data.is_merged }}</div>
</template>
