<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex'

let id = ref(null);
const route = useRoute();
const baseUrl = "https://severj.top/img/";

onMounted(async () => {
  id.value = route.params.id;
  await new Promise(resolve => setTimeout(resolve, 2000));
  useStore().commit('setLoading', false);
});

const getImgUrl = (i) => {
  return `${baseUrl}background${i}.webp`;
};
const getName = (i) => {
  return `name${i}`;
};
const getDescription = (i) => {
  return `description${i}`;
};
const getStar = (i) => {
  return i;
};
const getTags = (i,j) => {
  return i+j;
};
const getTagsLink = (i) => {
  return `${baseUrl}background${i}.webp`;
};
const getTagsSum = (i) => {
  return i;
};
</script>

<template>
  <div>
    <a-page-header
        style="border: 1px solid rgb(235, 237, 240)"
        title="景点详情"
        :sub-title="`ID: ${id}`"
        @back="() => $router.go(-1)"
    />
    <div class="content">
      <div class="view-img">
        <img :src="getImgUrl(Number(id))" alt="" />
      </div>
      <div class="details">
        <h2>{{getName(Number(id))}}</h2>
        <p>简介: {{getDescription(Number(id))}}</p>
        <p>评分：<a-rate :value="getStar(4.5)" allow-half disabled /></p>
        <p>标签：
          <a v-for="i in getTagsSum(3)" :key="i">
            <a-tag>
              <a :href="getTagsLink(Number(i))">{{ getTags(Number(id),i) }}</a>
            </a-tag>
          </a>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.content {
  display: flex;
  justify-content: flex-start;
  align-items: start;
}

.view-img img {
  margin-left: 20px;
  margin-top: 25px;
  width: 200px;
  height: 200px;
  object-fit: cover;
  object-position: center;
  border-radius: 8px;
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.5); /* 添加阴影效果 */
}

.details {
  margin-left: 20px;
  margin-top: 25px;
}
</style>
