<script setup>
import axios from "axios";
import Cookies from "js-cookie";
import { ref, watch } from "vue";
import { onMounted, onUnmounted } from "vue";
import AMapLoader from "@amap/amap-jsapi-loader";
const title = ref("");
const action_type = ref("1");

const partytime = ref();

const roadstart = ref();
const roadend = ref();
const rectangletext = ref();
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
import AMapLoader from "@amap/amap-jsapi-loader";
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
      access_token: "",
      refresh_token: "",
      partyid: null,
      data: [],
      road: [],
      routejson: null,
      map: null,
      map2: null,
      resultjson: {}
    };
  },
  methods: {
    rero() {
      this.routejson = "";
      this.data = "";
      this.map = new AMap.Map("roadmap", {
        // 设置地图容器id
        viewMode: "3D", // 是否为3D地图模式
        zoom: 11, // 初始化地图级别
        center: [116.397428, 39.90923] // 初始化地图中心点位置
      });
    },

    remap() {
      this.routejson = "";
      this.data = "";
      this.map = new AMap.Map("rectangletmap", {
        // 设置地图容器id
        viewMode: "2D", // 是否为3D地图模式
        zoom: 6, // 初始化地图级别
        center: [116.397428, 39.90923], // 初始化地图中心点位置
        map: null
      });
    },
    savero(start, end) {
      const road = [
        { keyword: start }, //起始点坐标
        { keyword: end } //终点坐标
      ];
      this.resultjson = road;
      console.log(road);
      window._AMapSecurityConfig = {
        securityJsCode: "9a9a79b5c5fcc275c47bb5eafde2f7d3"
      };
      AMapLoader.load({
        key: "8c9bb5684ff803ee1c6efecc3ea36578",
        version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
        plugins: ["AMap.Scale"] //需要使用的的插件列表，如比例尺'AMap.Scale'，支持添加多个如：['...','...']
      })
        .then((AMap) => {
          let map = new AMap.Map("roadmap", {
            // 设置地图容器id
            viewMode: "2D", // 是否为3D地图模式
            zoom: 18, // 初始化地图级别
            center: [118.397428, 39.90923] // 初始化地图中心点位置
          });
          AMap.plugin(
            ["AMap.ToolBar", "AMap.Driving", "AMap.Polyline", "AMap.Marker"],
            function () {
              //异步同时加载多个插件
              var toolbar = new AMap.ToolBar();
              map.addControl(toolbar);

              var driving = new AMap.Driving({
                map: map,
                panel: "panel"
              }); //驾车路线规划
              driving.search(road, function (status, result) {
                // result 即是对应的驾车导航信息，相关数据结构文档请参考  https://lbs.amap.com/api/javascript-api/reference/route-search#m_DrivingResult
                if (status === "complete") {
                  console.log(result);

                  console.log(result.routes);
                } else {
                  console.log("获取驾车数据失败：" + result);
                }
              });
            }
          );
        })
        .catch((e) => {
          console.log(e);
        });
    },
    savemap(text) {
      axios
        .get(
          "https://restapi.amap.com/v3/geocode/geo?address=" +
            text +
            "&key=4a456acf68e96cfd42e35d8915c9cee0"
        )
        .then((res) => {
          console.log(res);
          this.data = res.data.geocodes[0].location.split(",");

          this.map = new AMap.Map("rectangletmap", {
            // 设置地图容器id
            viewMode: "2D", // 是否为3D地图模式
            zoom: 19, // 初始化地图级别
            center: this.data,
            map: null
          });
        })
        .catch((err) => {
          console.error(err);
        });
    },
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
    window._AMapSecurityConfig = {
      securityJsCode: "9a9a79b5c5fcc275c47bb5eafde2f7d3"
    };
    AMapLoader.load({
      key: "8c9bb5684ff803ee1c6efecc3ea36578",
      version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
      plugins: ["AMap.Scale"] //需要使用的的插件列表，如比例尺'AMap.Scale'，支持添加多个如：['...','...']
    })
      .then((AMap) => {
        this.map = new AMap.Map("rectangletmap", {
          // 设置地图容器id
          viewMode: "2D", // 是否为3D地图模式
          zoom: 18, // 初始化地图级别
          center: [116.397428, 39.90923] // 初始化地图中心点位置
        });
      })
      .catch((e) => {
        console.log(e);
      });
    AMapLoader.load({
      key: "8c9bb5684ff803ee1c6efecc3ea36578",
      version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
      plugins: ["AMap.Scale"] //需要使用的的插件列表，如比例尺'AMap.Scale'，支持添加多个如：['...','...']
    })
      .then((AMap) => {
        this.map2 = new AMap.Map("roadmap", {
          // 设置地图容器id
          viewMode: "2D", // 是否为3D地图模式
          zoom: 18, // 初始化地图级别
          center: [116.397428, 39.90923] // 初始化地图中心点位置
        });
      })
      .catch((e) => {
        console.log(e);
      });
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
    <a-divider orientation="left" class="separate" v-show="!(action_type - 1)">路线</a-divider>
    <div v-if="!(action_type - 1)">
      <div class="road">
        {{ resultjson }}
        <div id="panel"></div>
        <div id="roadmap"><button @click="rero()">刷新地图</button></div>
        <div class="input-box">
          <a-input v-model:value="roadstart" :bordered="false" size="large" placeholder="出发点" />
        </div>
        <div class="input-box">
          <a-input v-model:value="roadend" :bordered="false" size="large" placeholder="结束点" />
        </div>
        <button @click="rero()">刷新地图</button>
        <button @click="savero(roadstart, roadend)">保存路线</button>
      </div>
    </div>
    <a-divider orientation="left" class="separate" v-show="action_type - 1">地点</a-divider>
    <div v-if="action_type - 1">
      <!--  {{ data }} -->
      <div class="rectangle">
        <p>省份＋城市＋区县＋城镇＋乡村＋街道＋门牌号码</p>
        <div id="rectangletmap"><button @click="remap">刷新地图</button></div>
        <div class="input-box">
          <a-input
            v-model:value="rectangletext"
            :bordered="false"
            size="large"
            placeholder="输入地点保存后可获取参考地图(可选)"
          />
        </div>
        <button @click="remap">刷新地图</button>
        <button @click="savemap(rectangletext)">保存地点</button>
      </div>
    </div>
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
        @click="
          partycreate(title, action_type, partyid, this.data, this.routejson, remark, partytime)
        "
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
.road {
  display: grid;
  justify-content: center;
  margin-bottom: 20px;
}
#roadmap {
  width: 90%;
  height: 300px;
  margin: 5%;
}
#rectangletmap {
  width: 90%;
  height: 300px;
  margin: 5%;
}
.rectangle {
  display: grid;
  justify-content: center;
}
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
  margin: 10px;
  width: 90vw;
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
