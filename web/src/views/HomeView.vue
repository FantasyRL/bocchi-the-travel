
<template>
  <div class="home">
    <a-input-search
        class="custom-search"
        placeholder="input search text"
        enterButton="Search"
        size="large"
        v-model="searchText"
        @search="onSearch"
    />

    <h3>情绪推荐</h3>

    <a-carousel arrows dots-class="slick-dots slick-thumb">
      <template #customPaging="props">
        <a>
          <img :src="getImgUrl(props.i)"  alt=""/>
        </a>
      </template>
      <div v-for="item in 5" :key="item">
        <img :src="getImgUrl(item - 1)"  alt=""/>
      </div>
    </a-carousel>

    <h4>发现世界></h4>

    <div class="image-grid">
      <div v-for="item in images" :key="item.id">
        <img :src="item.url" :alt="item.description" />
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue';

const baseUrl =
    'https://severj.top/img/';
export default defineComponent({
  data() {
    return {
      searchText: '',
      images: [
        { id: 1, url: 'https://severj.top/img/background1.webp', description: 'Image 1' },
        { id: 2, url: 'https://severj.top/img/background2.webp', description: 'Image 2' },
        // Add more images as needed...
      ],
    }
  },
  methods: {
    onSearch(value) {
      console.log('Search:', value);
    },
    getImgUrl(i) {
      return `${baseUrl}background${i + 1}.webp`;
    }
  }
});
</script>

<style scoped>
.custom-search {
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

h3 {
  text-align: center;
  font-size: 30px;
  color: #3b3b48;
  opacity: v-bind(0.8);
  margin-top: 15px;
  margin-bottom: 10px;
}
h4 {
  text-align: left;
  font-size: 25px;
  color: #3b3b48;
  opacity: v-bind(0.8);
  margin-top: 15px;
  margin-left: 10px;
  margin-bottom: 10px;
}

:deep(.slick-dots) {
  position: relative;
  height: auto;
}
:deep(.slick-slide img) {
  border: 5px solid #fff;
  display: block;
  margin: auto;
  max-width: 80%;
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
