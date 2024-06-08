<script setup>
import { ref } from "vue";
import { RouterLink, useRoute } from "vue-router";
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

const CommentText = ref("");
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
  }, 1000);
};
</script>
<script>
export default {
  data() {
    return {};
  },
  methods: {},
  mounted() {
    window.scrollTo({
      top: 0,
      behavior: "smooth"
    });
  }
};
</script>

<template>
  <div class="end">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240); z-index: 99999"
      title="行程结束"
      :sub-title="`ID: ${id}`"
      @back="() => $router.go(-1)"
    />
  </div>

  <br />
  <br />
  <br />

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
                  <a style="font-size: 18px">{{ getName(i) }}</a>
                </template>
                <template #avatar>
                  <a-avatar
                    style="margin-top: 3px"
                    :size="54"
                    src="https://severj.top/img/icon/logo.png"
                  />
                </template>
              </a-list-item-meta>
            </a-list-item>
          </a-list>
        </router-link>
      </a-col>

      <a-col flex="auto" style="margin-top: 5px">
        <a-rate v-model:value="value" allow-half />
        <div style="display: flex; justify-content: center">
          <div
            style="
              display: flex;
              border-radius: 12px;
              border: 3px solid #f5f5f5;
              margin: 5% auto auto;
              text-decoration: none;
            "
          >
            <a-button type="link" @click="showModal">评价</a-button>
          </div>
          <a-modal
            v-model:open="open"
            title="请输入你的评价"
            :confirm-loading="confirmLoading"
            @ok="handleOk"
          >
            <a-textarea
              v-model:value="CommentText"
              placeholder="请在此输入"
              :rows="4"
              size="large"
              :bordered="false"
            />
          </a-modal>
        </div>
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
                  <a style="font-size: 18px">{{ getName(i) }}</a>
                </template>
                <template #avatar>
                  <a-avatar shape="square" :size="64" :src="getPlacesImg(i)" />
                </template>
              </a-list-item-meta>
            </a-list-item>
          </a-list>
        </router-link>
      </a-col>

      <a-col flex="auto" style="margin-top: 5px">
        <a-rate v-model:value="value" allow-half />
        <div style="display: flex; justify-content: center">
          <div
            style="
              display: flex;
              border-radius: 12px;
              border: 3px solid #f5f5f5;
              margin: 5% auto auto;
              text-decoration: none;
            "
          >
            <a-button type="link" @click="showModal">评价</a-button>
          </div>
          <a-modal
            v-model:open="open"
            title="请输入你的评价"
            :confirm-loading="confirmLoading"
            @ok="handleOk"
          >
            <a-textarea
              v-model:value="CommentText"
              placeholder="请在此输入"
              :rows="4"
              size="large"
              :bordered="false"
            />
          </a-modal>
        </div>
      </a-col>
    </a-row>
  </div>

  <br /><br /><br />
</template>

<style scoped>
.end {
  width: 100vw;
  color: rgb(255, 255, 255);
  background-color: rgb(255, 255, 255);
}
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
