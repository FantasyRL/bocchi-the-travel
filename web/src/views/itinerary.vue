<script setup>
import axios from "axios";
</script>
<script>
export default {
  props: {},
  data() {
    return {
      id: 1, // 假设这是 itinerary 的 id
      data: {
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
  }
};
</script>
<template>
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
