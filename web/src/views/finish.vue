<script setup>
import {ref } from 'vue';
import {RouterLink, useRoute} from "vue-router";
import {HomeOutlined} from '@ant-design/icons-vue';
const id = useRoute().params.id;
const getName = (i) => {
  return `name${i}`;
};
const getSign = (i) => {
  return `sign${i}`;
};
const getFriendLink = (i) => {
  return `/about/${i}`;
};
const getPlacesLink = (i) => {
  return `/detail/${i}`;
};
const getPlacesImg = (i) => {
  return `https://severj.top/img/background${i}.webp`;
};
const value = ref(4.5);
</script>

<template>
  <div class="end">
    <a-page-header
        style="border: 1px solid rgb(235, 237, 240)"
        title="行程结束"
        :sub-title="`ID: ${id}`"
        @back="() => $router.go(-1)"
    />
  </div>

  <div class="view-img">
    <img src="@/assets/finish.svg" alt="" />
  </div>

  <div class="announce">
    <p>行程结束，动动手指打个分吧！</p>
  </div>

  <a-divider orientation="left" class="separate">伙伴评价</a-divider>

  <div class="List" v-for="i in 3" :key="i">
    <a-row class="Card">

      <a-col flex="70%">
        <router-link :to="getFriendLink(i)">
          <a-list item-layout="horizontal">
            <a-list-item>
              <a-list-item-meta :description="getSign(i)">
                <template #title>
                  <a>{{ getName(i) }}</a>
                </template>
                <template #avatar>
                  <a-avatar size="large" src="https://severj.top/img/icon/logo.png" />
                </template>
              </a-list-item-meta>
            </a-list-item>
          </a-list>
        </router-link>
      </a-col>

      <a-col flex="auto" style="margin-top: 10px">
        <a-rate v-model:value="value" allow-half />
      </a-col>

    </a-row>
  </div>

  <a-divider orientation="left" class="separate">景点评价</a-divider>

  <div class="List" v-for="i in 3" :key="i">
    <a-row class="Card">

      <a-col flex="70%">
        <router-link :to="getPlacesLink(i)">
          <a-list item-layout="horizontal">
            <a-list-item>
              <a-list-item-meta :description="getSign(i)">
                <template #title>
                  <a>{{ getName(i) }}</a>
                </template>
                <template #avatar>
                  <a-avatar shape="square" :size="64" :src="getPlacesImg(i)" />
                </template>
              </a-list-item-meta>
            </a-list-item>
          </a-list>
        </router-link>
      </a-col>

      <a-col flex="auto" style="margin-top: 10px">
        <a-rate v-model:value="value" allow-half />
      </a-col>

    </a-row>
  </div>

  <div style="display: flex; align-items: center; justify-content: center; margin: 20px">
    <router-link to="/">
      <a-button type="primary" shape="round" size="large">
        <template #icon>
          <HomeOutlined/>
        </template>
        返回首页
      </a-button>
    </router-link>
  </div>

</template>

<style scoped>
.view-img {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 40vh;
}
.announce {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
  font-weight: bold;
}
.separate {
  margin-top: 60px;
  font-size: 18px;
}
.List {
  margin: 10px;
  align-items: center;
}
</style>

<script>
export default {
  name: 'FinishView',
}
</script>
