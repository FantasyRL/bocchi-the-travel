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
      avatar: "", // 假设返回的数据中有一个名为avatar的字段，表示用户头像。
      mail: "", // 假设返回的数据中有一个名为mail的字段，表示用户邮箱。
      signature: "", // 假设返回的数据中有一个名为signature的字段，表示用户签名。
      access_token: "", // 假设返回的数据中有一个名为access_token的字段，表示用户访问令牌。
      
    }
  },
  methods: {
    PutAvatar(){
      this.access_token = Cookies.get('access_token');
      const file = this.$refs.fileInput.files[0];
      const formData = new FormData();
      formData.append('avatar_file', file);
    
      const headers = {
        "access_token": this.access_token // 假设access_token是必需的。
      };
      axios.put('http://127.0.0.1:10001/bocchi/user/avatar/upload', formData,{ headers })
        .then(response => {
          console.log(response.data);
        })
        .catch(error => {
          console.error(error);
        });
    },
    
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
        this.avatar = res.data.user.avatar; // 假设返回的数据中有一个名为avatar的字段，表示用户头像。
        this.mail = res.data.user.email; // 假设返回的数据中有一个名为mail的字段，表示用户邮箱。
        this.signature = res.data.user.signature; // 假设返回的数据中有一个名为signature的字段，表示用户签名。
        console.log(this.avatar)
        if (this.avatar === "") { // 如果用户没有头像，则使用默认头像。
          this.avatar = "//xiey.work/640.jpg" // 假设默认头像的地址。
        } 
        if (this.signature === "") { 
          this.signature = "这个人很懒，什么都没写。" 
        }
          
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

  <!-- info -->
  <div class="info">
    <!-- 头像 start -->
    <div class="touxiang">
      <div class="testdiv6">
        <!-- <img src="//xiey.work/640.jpg" class="touxiangimg rounded-image" alt="头像" /> -->
        <img :src="avatar" class="touxiangimg rounded-image"/>
        <div><text class="name">{{name}}</text></div>
      </div>
    </div>
    <!-- 头像 end -->
    <!-- 社交账号 -->
    <div class="shejiao">
      <text>uid:{{ id }}</text>
      <text>邮箱:{{ mail }}</text>
      <text>签名:{{ signature }}</text>
    </div>
    <!-- 社交账号 end -->

    <!-- 个人信息 start -->
    <!--   <div class="ab2">
      <a-empty />
    </div> -->
    <!-- 个人信息 end -->
  </div>
  <!-- info end -->
  <div>
    <form @submit.prevent="uploadAvatar">
      <input type="file" ref="fileInput" />
      <button type="submit" @click="PutAvatar">上传头像</button>
    </form>
  </div>
  <button @click="PutAvatar">修改头像</button>
  <button @click="logout">退出登陆</button>
</template>

<style scoped>
.shejiao{ /* 社交账号 */
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #f0f0f0;
  position: relative;
  height: auto;
  width: 90%;
  margin: auto;
  text-align: center;
  
}
</style>
