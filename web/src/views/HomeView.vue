<script setup>
import { ref } from "vue";
const current = ref(1);
import { PlusOutlined } from "@ant-design/icons-vue";
</script>

<template>
  <a-config-provider :getPopupContainer="getPopupContainer">
    <app />
  </a-config-provider>

  <div class="home">
    <div class="searchbar">
      <a-input-search
        class="custom-search"
        placeholder="请输入搜索内容"
        enterButton="搜索"
        size="large"
        v-model="searchText"
        @search="onSearch"
      />
    </div>
    <div class="pubu">
      <a-carousel arrows dots-class="slick-dots slick-thumb">
        <template #customPaging="props">
          <a>
            <img :src="getImgUrl(props.i)" alt="" />
          </a>
        </template>
        <div v-for="item in 5" :key="item">
          <img :src="getImgUrl(item - 1)" alt="" />
        </div>
      </a-carousel>
    </div>
    <a-divider orientation="left">发现世界</a-divider>

    <div class="image-grid">
      <a-card hoverable style="width: 240px">
        <template #cover>
          <img alt="example" src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png" />
        </template>
        <a-card-meta title="Europe Street beat">
          <template #description>www.instagram.com</template>
        </a-card-meta>
      </a-card>

      <div v-for="item in 7" :key="item">
        <img :src="getImgUrl(item - 1)" alt="" />
      </div>
    </div>

    <a-float-button tooltip="创建行程" herf="/create">
      <template #icon>
        <PlusOutlined />
      </template>
    </a-float-button>

    <a-back-top />

    <div class="page-switcher">
      <a-pagination v-model:current="current" :total="item in images">
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
      return `${baseUrl}background${i + 1}.webp`;
    },
    getPopupContainer(el, dialogContext) {
      if (dialogContext) {
        return dialogContext.getDialogWrap();
      } else {
        return document.body;
      }
    }
  }
});
</script>

<style scoped>
.pubu {
  min-width: 50%;
}
.searchbar {
  display: flex;
  justify-content: center;
  margin: 20px;
}

.page-switcher {
  display: flex;
  justify-content: center;
  margin-top: 10px;
}

.image-grid {
  margin-left: 10px;
  margin-right: 10px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 10px;
}

.image-grid img {
  max-width: 100%;
  max-height: 200px;
  object-fit: cover;
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

:deep .slick-thumb li.slick-active img {
  filter: grayscale(0%);
}
</style>
