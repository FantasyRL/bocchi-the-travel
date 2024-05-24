<script setup>
import {RouterLink} from "vue-router";
import { PlusOutlined } from "@ant-design/icons-vue";
import { ref } from "vue";
const current = ref(1);


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
            <img :src="getImgUrl(props.i+1)" alt="" />
          </a>
        </template>
        <div v-for="id in 5" :key="id">
          <img :src="getImgUrl(id)" alt=""  style="border-radius: 15px;margin-bottom: 10px"/>
        </div>
      </a-carousel>
    </div>

    <a-divider orientation="center" class="separate">发现世界</a-divider>

    <div class="card1">
      <div v-for="id in 8" :key="id">
        <router-link :to="getDetailUrl(id)" @click="startLoading">
          <a-card hoverable style="width: 240px">
            <template #cover>
              <img :src="getImgUrl(id)" alt="example" />
            </template>
            <a-card-meta :title="getName(id)">
              <template #description>{{getDescription(id)}}</template>
            </a-card-meta>
          </a-card>
        </router-link>
      </div>
    </div>

    <router-link to="/create">
      <a-float-button tooltip="创建行程" herf="/create">
        <template #icon>
          <PlusOutlined />
        </template>
      </a-float-button>
    </router-link>

    <a-back-top />

    <div class="page-switcher">
      <a-pagination v-model:current="current" :total="id in images">
        <template #itemRender="{ type, originalElement }">
          <a v-if="type === 'prev'">前一页</a>
          <a v-else-if="type === 'next'">后一页</a>
          <component :is="originalElement" v-else></component>
        </template>
      </a-pagination>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import {useStore} from "vuex";

const baseUrl = "https://severj.top/img/";
export default defineComponent({
  data() {
    return {
      searchText: "",
      images: [
        { id: 1, url: "https://severj.top/img/background1.webp", description: "Image 1" },
        { id: 2, url: "https://severj.top/img/background2.webp", description: "Image 2" }
        // Add more images as needed...
      ]
    };
  },
  methods: {
    onSearch(value) {
      console.log("Search:", value);
    },
    getImgUrl(i) {
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
    },
    startLoading() {
      useStore().commit('setLoading', true);
    }
  }
});
</script>

<style scoped>

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
