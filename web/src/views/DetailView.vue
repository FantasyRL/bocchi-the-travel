<script setup>
import {
  LikeFilled,
  LikeOutlined,
  DislikeFilled,
  ShoppingCartOutlined,
  DislikeOutlined,
  PlusOutlined
} from "@ant-design/icons-vue";
import { ref } from "vue";
import { RouterLink, useRoute } from "vue-router";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);

/* onMounted(async () => {
  await new Promise(resolve => setTimeout(resolve, 2000));
  useStore().commit('setLoading', false);
}); */

const baseUrl = "https://severj.top/img/";
const id = useRoute().params.id;
const likes = getLikes(Number(id));
const dislikes = getDislikes(Number(id));
const action = ref("");

const like = () => {
  likes.value++;
  dislikes.value--;
  action.value = "liked";
};

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
const getTags = (i, j) => {
  return i + j;
};
const getTagsLink = (i) => {
  return `${baseUrl}background${i}.webp`;
};
const getTagsSum = (i) => {
  return i;
};
const getRelativeID = (i) => {
  return i;
};
const dislike = () => {
  likes.value--;
  dislikes.value++;
  action.value = "disliked";
};

const visible = ref(false);
const handleClose = () => {
  visible.value = false;
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
        <h2 style="margin-bottom: 30px; font-weight: bold">{{ getName(Number(id)) }}</h2>
        <p>简介: {{ getDescription(Number(id)) }}</p>
        <p>评分：<a-rate :value="getStar(4.5)" allow-half disabled /></p>
        <p>
          标签：
          <a v-for="i in getTagsSum(3)" :key="i">
            <a-tag>
              <a :href="getTagsLink(Number(i))">{{ getTags(Number(id), i) }}</a>
            </a-tag>
          </a>
        </p>
      </div>
    </div>

    <a-divider orientation="center" class="separate">相关评论</a-divider>

    <div class="comments">
      <a-comment>
        <template #actions>
          <span key="comment-basic-like">
            <a-tooltip title="Like">
              <template v-if="action === 'liked'">
                <LikeFilled @click="like" />
              </template>
              <template v-else>
                <LikeOutlined @click="like" />
              </template>
            </a-tooltip>
            <span style="padding-left: 8px; cursor: auto">
              {{ likes }}
            </span>
          </span>
          <span key="comment-basic-dislike">
            <a-tooltip title="Dislike">
              <template v-if="action === 'disliked'">
                <DislikeFilled @click="dislike" />
              </template>
              <template v-else>
                <DislikeOutlined @click="dislike" />
              </template>
            </a-tooltip>
            <span style="padding-left: 8px; cursor: auto">
              {{ dislikes }}
            </span>
          </span>
        </template>
        <template #author><a>Severj</a></template>
        <template #avatar>
          <a-avatar src="https://severj.top/img/icon/logo.png" alt="Severj" />
        </template>
        <template #content>
          <p>
            We supply a series of design principles, practical patterns and high quality design
            resources, to help people create their product prototypes beautifully and efficiently.
          </p>
        </template>
        <template #datetime>
          <a-tooltip :title="dayjs().format('YYYY-MM-DD HH:mm:ss')">
            <span>{{ dayjs().fromNow() }}</span>
          </a-tooltip>
        </template>
      </a-comment>
    </div>

    <a-divider orientation="center" class="separate">相关行程</a-divider>

    <div style="padding: 20px" class="relative">
      <div v-for="i in 5" :key="i">
        <a-row>
          <a-card class="relative-card" :title="getName(getRelativeID(i))" :bordered="false">
            <a-row>
              <a-col class="inner-img"
                ><img :src="getImgUrl(Number(getRelativeID(i)))" alt=""
              /></a-col>
              <a-col style="margin-left: 20px"
                ><p>{{ getDescription(Number(getRelativeID(i))) }}</p></a-col
              >
            </a-row>
          </a-card>
        </a-row>
      </div>
    </div>

    <a-float-button-group style="margin-bottom: 60px">
      <router-link to="/alltravels">
        <a-float-button tooltip="查看行程" herf="/alltravels" style="margin-bottom: 15px">
          <template #icon>
            <ShoppingCartOutlined />
          </template>
        </a-float-button>
      </router-link>

      <a-float-button tooltip="加入行程" @click="visible = true">
        <template #icon>
          <PlusOutlined />
        </template>
      </a-float-button>
    </a-float-button-group>

    <a-alert
      class="alert"
      v-if="visible"
      message="成功！"
      description="已加入行程"
      type="success"
      show-icon
      closable
      banner
      :after-close="handleClose"
    />

    <a-back-top style="margin-bottom: 60px" />
  </div>
</template>

<script>
const getLikes = (i) => {
  return i;
};
const getDislikes = (i) => {
  return i;
};
</script>

<style scoped>
.ant-card-cover img {
  width: auto;
  height: 200px; /* 保持图片的宽高比 */
  object-fit: cover; /* 保持图片的宽高比，并裁剪多余的部分 */
}
.card1 {
  margin: auto;
  position: relative;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.card1 a {
  text-decoration: none;
}

.ant-card {
  margin: 15px;
  position: relative;
}
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
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.3);
}

.details {
  margin-left: 20px;
  margin-top: 25px;
}

.comments {
  margin-left: 20px;
  margin-right: 20px;
}

.separate {
  margin-top: 30px;
  font-size: 18px;
}

.alert {
  position: fixed;
  top: 10px;
  left: 10px;
  right: 10px;
  z-index: 1000; /* 确保它在其他元素之上 */
  padding: 10px;
  border-radius: 8px;
}

.inner-img img {
  width: 100px;
  height: 100px;
  object-fit: cover;
  object-position: center;
  border-radius: 8px;
}
.relative-card {
  width: 100%;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}
</style>
