
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
    <h3>情绪推荐：</h3>
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
  </div>
</template>

<script>
import { defineComponent } from 'vue';

const baseUrl =
    'https://severj.top/img/';
export default defineComponent({
  data() {
    return {
      searchText: ''
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
