<script setup>
import Cookies from "js-cookie";
import axios from "axios";
import { CheckOutlined } from "@ant-design/icons-vue";
import { useCounterStore } from "@/stores/api";
import {
  HomeOutlined,
  CoffeeOutlined,
  EnvironmentOutlined,
  PlayCircleOutlined,
  QuestionCircleOutlined
} from "@ant-design/icons-vue";
</script>
<script>
const setapiurl = useCounterStore();
export default {
  data() {
    return {
      partynull: 0,
      id: 1, // 假设这是 party 的 id
      infodata: {}, // 假设这是 party 的数据对象
      items: []
    };
  },
  methods: {
    ToEnd() {
      axios.get(setapiurl.apiurl + "/party/status?party_id=" + this.id + "&action_type=1", {
        headers: { "access-token": this.access_token }
      });
      this.$router.push(`/finish/${this.id}`);
    },
    Tomember() {
      this.$router.push(`/member/${this.id}`);
    },
    tomerplan() {
      this.$router.push(`/merplan/${this.id}`);
    },
    tocreate() {
      this.$router.push(`/Createplan/${this.id}`);
    },
    tomyitinerarys() {
      this.$router.push(`/myitinerarys/${this.id}`);
    },
    apply_party() {
      axios
        .get(setapiurl.apiurl + "/party/apply?party_id=" + this.id, {
          headers: {
            "access-token": this.access_token
          }
        })
        .then((res) => {
          console.log(res);
          alert("申请成功！");
        })
        .catch((err) => {
          console.error(err);
        });
    },
    getin() {
      axios
        .get(setapiurl.apiurl + "/party/itinerary/show?party_id=" + this.id, {
          headers: {
            "access-token": this.access_token
          }
        })
        .then((res) => {
          console.log(res);
          this.items = res.data.itineraries;
          console.log(res.data.itineraries);
          if (res.data.itineraries === undefined) {
            this.partynull = 1;
          }
        })
        .catch((err) => {
          console.error(err);
        });
    },
    partyinit() {
      const url = setapiurl.apiurl + "/party/get?party_id=" + this.id;
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
    console.log(this.$router);
    this.id = this.$route.params.id;
    this.access_token = Cookies.get("access_token");
    this.partyinit();
    this.getin();
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
      title="活动详情"
      @back="() => $router.go(-1)"
    />
  </div>

  <div class="info">
    <div class="card">
      <div class="tools">
        <div class="circle">
          <span class="red box"></span>
        </div>
        <div class="circle">
          <span class="yellow box"></span>
        </div>
        <div class="circle">
          <span class="green box"></span>
        </div>
      </div>
      <div class="card__content">
        <div class="party_info">
          <div class="title"><text class="info_name">活动名:</text><br />{{ infodata.title }}</div>
          <div class="title">
            <text class="info_name">创建者:</text><br />{{ infodata.founder_id }}
          </div>

          <div class="title">
            <text class="info_name">成员数:</text><br />{{ infodata.member_count }}
          </div>
        </div>
        <div class="content">
          <text class="content_text">介绍:</text><br />[本活动于"{{ infodata.province }},{{
            infodata.city
          }}"处进行] <br />{{ infodata.content }}
        </div>

        <div style="text-align: center; margin-top: 10px; margin-bottom: 10px">
          起止时间:{{ infodata.start_time }} - {{ infodata.end_time }}
        </div>
        <div style="text-align: center; margin-top: 10px; margin-bottom: 10px">
          <a-button type="dashed" @click="apply_party()">申请加入</a-button>
          <a-button style="margin-left: 10px" @click="Tomember()">查看成员</a-button>
          <a-button type="primary" @click="ToEnd()" style="margin-left: 10px">结束行程</a-button>
        </div>
      </div>
    </div>
    <div class="center">
      <text>所有计划</text>
    </div>
    <div class="sty">
      <a-timeline>
        <div v-for="item in items" :key="item.id">
          <a-timeline-item>
            <template #dot>
              <component :is="getIcon(item.action_type)" style="font-size: 16px" />
            </template>
            计划: {{ item.title }}
            <br />
            类型：{{ getType(item.action_type) }}
            <br />
            内容：{{ item.remark }}

            <br />
            时间：{{ item.schedule_start_time }}
          </a-timeline-item>
          <a-timeline-item>
            <template #dot>
              <CheckOutlined style="font-size: 16px" />
            </template>
            结束时间:

            {{ item.schedule_end_time }}
            <div
              style="
                border-radius: 12px;
                border: 3px solid #f5f5f5;
                width: fit-content;
                margin-top: 5px;
              "
            >
              <a-button type="link" @click="$router.push('/itinerarys/' + item.id)">
                查看详细
              </a-button>
            </div>
          </a-timeline-item>
        </div>
      </a-timeline>
    </div>
    <div v-show="partynull">
      <a-empty />
    </div>
  </div>
  <div class="foot">
    <div class="create">
      <button class="btn" @click="tocreate">创建计划</button><br />
      <button class="btn" @click="tomyitinerarys">查看撰写过的计划</button>
      <br />
      <button class="btn" @click="tomerplan">待通过计划</button>
    </div>
  </div>
  <br />
  <br />
  <br />
  <br />
  <br />
  <br />
</template>

<style scoped>
.foot {
  margin-top: 36px;
  margin-bottom: 36px;
  width: 100%;
  display: grid;
  justify-content: center;
}
.create {
  display: grid;
  justify-content: center;
  text-align: center;
  border-radius: 12px;
}
.title {
  font-size: 18px;
}
.content_text {
  font-size: 20px;
  font-weight: bold;
}
.content {
  text-align: center;
  margin-top: 10px;
  margin-bottom: 10px;

  width: 100%;
}
.info_name {
  font-size: 18px;
  font-weight: bold;
}
.party_info {
  display: grid;
  font-size: 20px;
  justify-content: center;
  text-align: center;
  grid-template-columns: repeat(3, 2fr);
  grid-column-gap: 5px;
  grid-row-gap: 10px;
}
.sty {
  margin-left: 7vw;
  margin-top: 20px;
}

.info {
  display: grid;
  justify-content: center;
  margin-top: 5vh;
}

.center {
  text-align: center;
  margin-bottom: 10px;
  font-size: 20px;
  font-weight: bold;
}

.card {
  width: 90vw;
  height: auto;
  margin: 0 auto;
  background-color: #f8fbfe;
  border-radius: 8px;
  z-index: 1;
  box-shadow: 1px 2px 2px 1px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
  outline-width: 3px;
  outline-style: solid;
  outline-color: rgba(124, 124, 124, 0.1);
}

.tools {
  display: flex;
  align-items: center;
  padding: 9px;
}

.circle {
  padding: 0 4px;
}

.box {
  display: inline-block;
  align-items: center;
  width: 10px;
  height: 10px;
  padding: 1px;
  border-radius: 50%;
}

.red {
  background-color: #ff605c;
}

.yellow {
  background-color: #ffbd44;
}

.green {
  background-color: #00ca4e;
}
.btn {
  background-color: #00bfa6;
  padding: 14px 40px;
  color: #fff;
  text-transform: uppercase;
  letter-spacing: 2px;
  cursor: pointer;
  border-radius: 10px;
  border: 2px dashed #00bfa6;
  box-shadow:
    rgba(50, 50, 93, 0.25) 0px 2px 5px -1px,
    rgba(0, 0, 0, 0.3) 0px 1px 3px -1px;
  transition: 0.4s;
}

.btn span:last-child {
  display: none;
}

.btn:hover {
  transition: 0.4s;
  border: 2px dashed #00bfa6;
  background-color: #fff;
  color: #00bfa6;
}

.btn:active {
  background-color: #87dbd0;
}
</style>
