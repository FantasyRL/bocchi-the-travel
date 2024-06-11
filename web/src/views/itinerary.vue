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
      datago:null,
      datadata:null,
      id: 1, // 假设这是 itinerary 的 id
      party: {},
      info: {},
      partynull: false // 假设这是 party 是否为空的标志
    };
  },
  methods: {
    gogogo(){
      window.location.replace("https://uri.amap.com/navigation?from="+this.datadata+",startpoint&to="+this.datago+",endpoint&via=&mode=car&policy=1&src=mypage&coordinate=gaode&callnative=0") 
    },
    deleteItinerary() {
      const url = "https://api.xiey.work/bocchi/party/itinerary/delete?itinerary_id=" + this.id; // 假设这是删除 itinerary 的 API 接口地址

      axios
        .get(url, {
          headers: {
            "access-token": this.access_token
          }
        })
        .then((res) => {
          console.log(res); // 假设这是删除成功的回调函数，可以在这里进行页面跳转等操作
        })
        .catch((err) => {
          console.error(err); // 假设这是删除失败的回调函数，可以在这里进行错误处理
        });
    },
    init() {
      const url = "https://api.xiey.work/bocchi/party/itinerary/info?itinerary_id=" + this.id;
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
    this.access_token = Cookies.get("access_token");
  },
  computed: {
    getlocal() {
      return (i) => {
        if (i !== "") {
          const go = `https://restapi.amap.com/v3/staticmap?zoom=17&size=250*250&key=eae4d0491385d75b43d247afaef4247d&location=`+i
          return go;
        }
      };
    },
    getroad() {
      return (i) => {
        if (i !== undefined) {
          
          const start = i.split(",")[0];
          const end = i.split(",")[1]; 
          
          axios.get("https://restapi.amap.com/v3/geocode/geo?key=eae4d0491385d75b43d247afaef4247d&address="+start)
  .then(res => {
    console.log(res.data.geocodes[0].location)
    this.datadata = res.data.geocodes[0].location
    
  });
          axios.get("https://restapi.amap.com/v3/geocode/geo?key=eae4d0491385d75b43d247afaef4247d&address="+end)
  .then(res => {
    console.log(res.data.geocodes[0].location)
    this.datago = res.data.geocodes[0].location
    
  })




          return "从 "+start+" 到 "+end;
          
        }
      };
    },
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
<!--   {{ datadata }}
  {{ datago }} -->
  <div v-if="partynull" class="itinerary">
    <el-timeline style="max-width: 600px; margin-left: 10%">
      <el-timeline-item>计划名:{{ info.title }} </el-timeline-item>
      <el-timeline-item>序列:{{ info.sequence }} </el-timeline-item>
      <el-timeline-item> 创建者:{{ info.founder_id }} </el-timeline-item>
      <el-timeline-item> 类型：{{ getType(info.action_type) }} </el-timeline-item>
      <el-timeline-item> 备注：{{ info.remark }} </el-timeline-item>
      <el-timeline-item> 地点：{{ info.rectangle }} </el-timeline-item>
      <el-timeline-item><div><img :src="getlocal(info.rectangle)"></img></div></el-timeline-item>
      
      <el-timeline-item>{{ getroad(info.route_json) }}</el-timeline-item>
      <el-timeline-item> 路线： <a-button type="primary" @click="gogogo()">点击跳转到大地图</a-button></el-timeline-item>
      <el-timeline-item> 开始时间：{{ info.schedule_start_time }} </el-timeline-item>
      <el-timeline-item>
        结束时间:
        {{ info.schedule_end_time }}
      </el-timeline-item>
    </el-timeline>
    <button @click="deleteItinerary">
      <span class="shadow"></span>
      <span class="edge"></span>
      <span class="front text"> 删除计划 </span>
    </button>
  </div>
  <div v-if="!partynull">
    <a-empty />
  </div>
</template>

<style>
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
    hsl(340deg 100% 16%) 0%,
    hsl(340deg 100% 32%) 8%,
    hsl(340deg 100% 32%) 92%,
    hsl(340deg 100% 16%) 100%
  );
}

.front {
  display: block;
  position: relative;
  padding: 12px 27px;
  border-radius: 12px;
  font-size: 1.1rem;
  color: white;
  background: hsl(345deg 100% 47%);
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
