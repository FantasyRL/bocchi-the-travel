<script setup>
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';

let id = ref(null);
const route = useRoute();
const baseUrl = "https://severj.top/img/";

onMounted(() => {
  id.value = route.params.id;
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
