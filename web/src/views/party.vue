<script setup>
import axios from "axios";
</script>
<script>
export default {
  props: {},
  data() {
    return {
      id: 1, // 假设这是 party 的 id
      infodata: {
        id: 5,
        founder_id: 6,
        title: "1",
        content: "1",
        type: "",
        province: "1",
        city: "1",
        start_time: "2006-01-02",
        end_time: "2006-01-02",
        member_count: 0,
        status: 0,
        rectangle: ""
      }, // 假设这是 party 的数据对象
      itinerary: [
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
        }
      ]
    };
  },
  methods: {
    getin() {
      axios
        .get("/bocchi/party/itinerary/show?party_id=" + this.id)
        .then((res) => {
          console.log(res);
          this.itinerary = res.data.itineraries;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    partyinit() {
      const url = "/bocchi/party/info?party_id=" + this.id;
      const params = {};
      axios
        .get(url, params)
        .then((res) => {
          console.log(res);
          this.infodata = res.data.party; // 假设这是返回的数据对象
        })
        .catch((err) => {
          console.error(err);
        });
    }
  },
  mounted() {
    this.id = Number(this.$route.params.id);
    this.partyinit();
    this.getin();
  }
};
</script>
<template>
  party debug part:
  <div>party id:{{ id }}</div>
  <div>原始数据:</div>
  <div>party info:</div>
  {{ infodata }}
  <div>归属于这个party的itinerary</div>
  {{ itinerary }}
</template>
