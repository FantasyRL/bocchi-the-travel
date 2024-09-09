<script setup>
import { RouterLink } from "vue-router";
import { ref } from "vue";
import { useCounterStore } from "@/stores/api";
const current = ref(1);
</script>

<script>
import { defineComponent } from "vue";
import axios from "axios";

const baseUrl = "//upload.xiey.work/img/";
/* const baseUrl = "https://loremflickr.com/800/600"; */

export default defineComponent({
  data() {
    return {
      ipcity: "", // 用于存储IP地址的城市信息
      total: 10, // 假设的总数，实际应从后端获取
      current: 2, // 当前页码，默认为1
      items: {},
      searchText: "",
      images: [
        { id: 1, url: "https://severj.top/img/background1.webp", description: "Image 1" },
        { id: 2, url: "https://severj.top/img/background2.webp", description: "Image 2" }
        // Add more images as needed...
      ]
    };
  },

  props: {},
  methods: {
    getip() {
      axios
        .get("https://restapi.amap.com/v3/ip?key=eae4d0491385d75b43d247afaef4247d")
        .then((res) => {
          console.log(res);
          this.ipcity = res.data.city;
          this.searchprovince("", "", "", this.ipcity, "", "2", "1");
        })
        .catch((err) => {
          console.error(err);
        });
    },
    searchprovince(
      content,
      party_type,
      province,
      city,
      start_time_duration,
      search_type,
      page_num
    ) {
      if (content === "") {
        content = "nothing";
      }
      if (party_type === "") {
        party_type = "nothing";
      }
      if (province === "") {
        province = "nothing";
      }
      if (city === "") {
        city = "nothing";
      }
      if (start_time_duration === "") {
        start_time_duration = "-1";
      }
      const setapiurl = useCounterStore();
      const url =
        setapiurl.apiurl +
        "/party/search?content=" +
        content +
        "&party_type=" +
        party_type +
        "&province=" +
        province +
        "&city=" +
        city +
        "&start_time_duration=" +
        start_time_duration +
        "&search_type=" +
        search_type +
        "&page_num=" +
        page_num;
      const testurl = setapiurl.apiurl + "/party/search?content=" + content;
      axios
        .post(url)
        .then((res) => {
          console.log(res);
          this.data = res;
          this.items = res.data.party_list;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    onSearch(value) {
      console.log("Search:", value);
    },
    getImgUrl(i) {
      /* return `${baseUrl}background${i}.webp`; */
      return `${baseUrl}background${i}.webp`;
    },
    getDetailUrl(i) {
      return `/detail/${i}`;
    },
    getName(i) {
      return `name${i}`;
    },
    getDescription(i) {
      return `description${i}`;
    }
  },
  mounted() {
    this.getip();
    /* this.searchprovince("", "", "北京", "", "", "2", "1"); */
    /* if (this.ipcity != "牛逼") {
      this.searchprovince("", "", this.ipcity, "", "", "2", "1");
    } */
  },
  computed: {},
  watch: {}
});
</script>
<template>
  <div class="home">
    <div class="searchbar">
      <a-input-search
        class="custom-search"
        placeholder="请输入搜索内容"
        enterButton=""
        size="large"
        v-model="searchText"
        :bordered="false"
        @search="onSearch"
      />
    </div>

    <a-divider orientation="center" class="separate">每日推荐</a-divider>

    <div class="carousel">
      <a-carousel arrows dots-class="slick-dots slick-thumb">
        <template #customPaging="props">
          <a>
            <img :src="getImgUrl(props.i + 1)" alt="" />
          </a>
        </template>
        <div v-for="id in 5" :key="id">
          <router-link :to="getDetailUrl(id)">
            <img :src="getImgUrl(id)" alt="" style="border-radius: 15px; margin-bottom: 10px" />
          </router-link>
        </div>
      </a-carousel>
    </div>
    <!--  {{ ipcity }} -->
    <a-divider orientation="center" class="separate">{{ this.ipcity }}-附近活动</a-divider>

    <div class="searchcard">
      <div v-for="item in items" :key="item">
        <div style="justify-content: center; display: grid">
          <router-link :to="`/partys/${item.id}`">
            <el-card style="justify-content: center; width: 40vw">
              <p>活动名:{{ item.title }}</p>
              <p>{{ item.province }},{{ item.city }}</p>
              <p>已有{{ item.member_count }}人参加</p>
            </el-card>
          </router-link>
        </div>
      </div>
    </div>
    <a-divider orientation="center" class="separate">发现世界</a-divider>

    <div class="trcard">
      <div v-for="id in 8" :key="id">
        <router-link :to="getDetailUrl(id)">
          <a-card hoverable style="width: auto">
            <template #cover>
              <img :src="getImgUrl(id)" alt="example" />
            </template>
            <a-card-meta :title="getName(id)">
              <template #description>{{ getDescription(id) }}</template>
            </a-card-meta>
          </a-card>
        </router-link>
      </div>
    </div>

    <div class="page-switcher">
      <a-pagination :current="current" :total="total">
        <template #itemRender="{ type, originalElement }">
          <a v-if="type === 'prev'">前一页</a>
          <a v-else-if="type === 'next'">后一页</a>
          <component :is="originalElement" v-else></component>
        </template>
      </a-pagination>
    </div>
  </div>
</template>

<style scoped>
.ant-card-cover img {
  width: 100%;
  height: 10vh; /* 保持图片的宽高比 */
  object-fit: cover; /* 保持图片的宽高比，并裁剪多余的部分 */
}

.trcard {
  display: grid;
  grid-template-columns: 1fr 1fr;
  text-align: center;
  grid-gap: 10px 10px; /* 定义网格之间的水平和垂直间距 */
  margin: 10px 10px 10px 10px; /* 定义外边距 */
  justify-content: center;
}
.searchcard {
  display: grid;
  grid-template-columns: 1fr 1fr;
  text-align: center;
  grid-gap: 10px 10px;
  margin: 24px 24px 24px 24px;
}

.searchbar {
  display: flex;
  justify-content: center;
  margin: 24px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.2);
  border-radius: 5px;
}

.page-switcher {
  display: flex;
  justify-content: center;
  margin-top: 10px;
}

.image-grid img {
  max-width: 100%;
  max-height: 200px;
  object-fit: cover;
}

.separate {
  margin-top: 24px;
  font-size: 18px;
}

:deep(.slick-dots) {
  position: relative;
  height: auto;
}

:deep(.slick-slide img) {
  border: 5px solid #fff;
  display: block;
  margin: auto;
  max-width: 95%;
}

:deep(.slick-arrow) {
  display: none !important;
}

:deep(.slick-thumb) {
  bottom: 0;
}

:deep(.slick-thumb li) {
  width: 60px;
  height: 45px;
}

:deep(.slick-thumb li img) {
  width: 100%;
  height: 100%;
  filter: grayscale(100%);
  display: block;
}

:deep(.slick-thumb li.slick-active) img {
  filter: grayscale(0%);
}
</style>
