<script setup>
import axios from 'axios';
import Cookies from 'js-cookie';


/* import { useCounterStore } from '@/stores/counter' */
/* const counter = useCounterStore() */





</script>

<script>
export default {
  data() {
    return {
      /* 数据 */
      id: 1,
      name: "用户名",
    }
  },
  methods: {
    /* 方法 */
    logout() { // 退出登录的函数，如果有的话。
      Cookies.remove('id'); // 删除id的cookie。
      Cookies.remove('access_token'); 
      window.location.href = '/';
     
    },
    init() { 
      axios.get('http://127.0.0.1:10001/bocchi/user/info?user_id='+this.id)
      .then(res => {
        console.log(res)
        this.name = res.data.user.name; // 假设返回的数据中有一个名为name的字段，表示用户名。
      })
      .catch(err => {
        console.error(err); 
      })
    }, // 初始化函数，如果有的话
  },
  mounted() { 
    /* 初始化 */
    this.id = Cookies.get('id');
    if (this.id) { // 如果id不存在，说明用户未登录。
      this.init(); // 跳转到登录页面。
    }
     // 初始化函数，如果有的话
  },
  components: {},
  computed: {},
  watch: {},
  props: {},
  emits: {},
  setup() { },

};
</script>
<template>
  <a-page-header style="border: 1px solid rgb(235, 237, 240)" title="个人界面" sub-title="副标题"
    @back="() => $router.go(-1)" />
id:{{ id }}
  <!-- info -->
  <div class="info">
    <!-- 头像 start -->
    <div class="touxiang">
      <div class="testdiv6">
        <img src="//xiey.work/640.jpg" class="touxiangimg rounded-image" alt="头像" />
        <div><text class="name">{{name}}</text></div>
      </div>
    </div>
    <!-- 头像 end -->
    <!-- 社交账号 -->
    <div class="shejiao">
      <text>社交账号</text>
      <text>社交账号</text>
      <text>社交账号</text>
      <text>社交账号</text>
    </div>
    <!-- 社交账号 end -->

    <!-- 个人信息 start -->
    <!--   <div class="ab2">
      <a-empty />
    </div> -->
    <!-- 个人信息 end -->
  </div>
  <!-- info end -->

  <button @click="logout">退出登陆</button>
</template>

<style scoped></style>
