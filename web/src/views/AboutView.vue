<script setup>
import axios from "axios";
import Cookies from "js-cookie";

/* import { useCounterStore } from '@/stores/counter' */
/* const counter = useCounterStore() */
</script>

<script>
export default {
  data() {
    return {
      /* 数据 */
      setname: "设置", // 用于设置用户名的输入框的值。
      setcount: 3,
      setshow: false, // 控制设置面板的显示与隐藏。
      id: 1,
      name: "用户名",
      avatar: "//xiey.work/640.jpg", // 假设返回的数据中有一个名为avatar的字段，表示用户头像。
      mail: "", // 假设返回的数据中有一个名为mail的字段，表示用户邮箱。
      signature: "这个人很懒，什么都没写。", // 假设返回的数据中有一个名为signature的字段，表示用户签名。
      access_token: "" // 假设返回的数据中有一个名为access_token的字段，表示用户访问令牌。
    };
  },
  methods: {
    settings(){
      this.setshow =! this.setshow
      this.setcount++ // 假设返回的数据中有一个名为name的字段，表示用户名。
    },
    PutAvatar() {
      this.access_token = Cookies.get("access_token");
      const file = this.$refs.fileInput.files[0];
      const formData = new FormData();
      formData.append("avatar_file", file);

      const headers = {
        access_token: this.access_token // 假设access_token是必需的。
      };
      axios
        .put("http://127.0.0.1:10001/bocchi/user/avatar/upload", formData, { headers })
        .then((response) => {
          console.log(response.data);
        })
        .catch((error) => {
          console.error(error);
        });
    },

    
    logout() {
      // 退出登录的函数，如果有的话。
      Cookies.remove("id"); // 删除id的cookie。
      Cookies.remove("access_token");
      window.location.href = "/";
    },
    init() {
      axios
        .get("http://127.0.0.1:10001/bocchi/user/info?user_id=" + this.id)
        .then((res) => {
          console.log(res);
          this.name = res.data.user.name; // 假设返回的数据中有一个名为name的字段，表示用户名。
          this.avatar = res.data.user.avatar; // 假设返回的数据中有一个名为avatar的字段，表示用户头像。
          this.mail = res.data.user.email; // 假设返回的数据中有一个名为mail的字段，表示用户邮箱。
          this.signature = res.data.user.signature; // 假设返回的数据中有一个名为signature的字段，表示用户签名。
          console.log(this.avatar);
          if (this.avatar === "") {
            // 如果用户没有头像，则使用默认头像。
            this.avatar = "//xiey.work/640.jpg"; // 假设默认头像的地址。
          }
          if (this.signature === "") {
            this.signature = "这个人很懒，什么都没写。";
          }
        })
        .catch((err) => {
          console.error(err);
        });
    } // 初始化函数，如果有的话
  },
  mounted() {
    /* 初始化 */
    this.id = Cookies.get("id");
    if (this.id) {
      // 如果id不存在，说明用户未登录。
      this.init(); // 跳转到登录页面。
    }
    // 初始化函数，如果有的话
  },
  components: {},
  computed: {},
  watch: {
    setcount(newsetcount, oldsetcount) {
            if (1) {
                console.log("变化前的值：", oldsetcount, "变化后的值：", newsetcount);
            }
            if (newsetcount%2==1) {
              this.setname ="设置"
            }
            if (newsetcount%2!=1) {
              this.setname ="关闭"
            }
        },
  },
  props: {},
  emits: {},
  setup() {}
};
</script>
<template>
  <link
    href="//netdna.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"
    rel="stylesheet"
  />
  <a-page-header
    style="border: 1px solid rgb(235, 237, 240)"
    title="个人界面"
    sub-title="副标题"
    @back="() => $router.go(-1)"
  />
  <div class="settings">
    <button class="button" @click="settings">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        viewBox="0 0 20 20"
        height="20"
        fill="none"
        class="svg-icon"
      >
        <g stroke-width="1.5" stroke-linecap="round" stroke="#5d41de">
          <circle r="2.5" cy="10" cx="10"></circle>
          <path
            fill-rule="evenodd"
            d="m8.39079 2.80235c.53842-1.51424 2.67991-1.51424 3.21831-.00001.3392.95358 1.4284 1.40477 2.3425.97027 1.4514-.68995 2.9657.82427 2.2758 2.27575-.4345.91407.0166 2.00334.9702 2.34248 1.5143.53842 1.5143 2.67996 0 3.21836-.9536.3391-1.4047 1.4284-.9702 2.3425.6899 1.4514-.8244 2.9656-2.2758 2.2757-.9141-.4345-2.0033.0167-2.3425.9703-.5384 1.5142-2.67989 1.5142-3.21831 0-.33914-.9536-1.4284-1.4048-2.34247-.9703-1.45148.6899-2.96571-.8243-2.27575-2.2757.43449-.9141-.01669-2.0034-.97028-2.3425-1.51422-.5384-1.51422-2.67994.00001-3.21836.95358-.33914 1.40476-1.42841.97027-2.34248-.68996-1.45148.82427-2.9657 2.27575-2.27575.91407.4345 2.00333-.01669 2.34247-.97026z"
            clip-rule="evenodd"
          ></path>
        </g>
      </svg>
      <span class="lable">{{ setname}}</span>
    </button>
  </div>

  <!-- info -->
  <div class="info">
    <!-- 头像 start -->
    <div class="touxiang">
      <div class="testdiv6">
        <!-- <img src="//xiey.work/640.jpg" class="touxiangimg rounded-image" alt="头像" /> -->
        <img :src="avatar" class="touxiangimg rounded-image" />
        <div>
          <text class="name">{{ name }}</text>
        </div>
      </div>
    </div>
    <!-- 头像 end -->

    <!-- 个人信息 -->
    <div class="shejiao">
      <text>uid:{{ id }}</text>
      <text>邮箱:{{ mail }}</text>
      <text>签名:{{ signature }}</text>
    </div>

    <!-- end -->
  </div>
  <!-- info end -->
  <!--   <div>
    <form @submit.prevent="uploadAvatar">
      <input type="file" ref="fileInput" />
      <button type="submit" @click="PutAvatar">上传头像</button>
    </form>
  </div>
  <button @click="PutAvatar">修改头像</button> -->
  <transition name="fade">
<div class="set" v-show="setshow">
  设置界面
  <mdui-menu>
  <mdui-menu-item>Item 1</mdui-menu-item>
  <mdui-menu-item>Item 2</mdui-menu-item>
</mdui-menu>
  <button @click="logout">退出登陆</button>
</div>
</transition>
</template>

<style scoped>

.set{
 
  position: fixed;
  top: 20%;
  margin-left: 10%;
  height: 60%;
  background-color: #c8c0f0;
  width: 80%;
  border-radius: 15px;
  align-items: center;
  display: flex;
  flex-direction: column;
  gap: 10px;
  padding: 30px 20px 10px 20px; /* 上下左右内边距 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); /* 添加阴影效果 */
  z-index: 1000; /* 确保设置界面在其它内容之上 */
}
.fade-enter {
  
  transition: opacity .5s;
}

.fade-enter-active {
  opacity: 1;
}

.fade-leave-active {
  transition: opacity .5s;
}

.fade-leave-to {
  opacity: 0;
}
/* .fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to
  opacity: 0;
} */
.settings{
  position:fixed;
  right: 10px;
  top: 70px;
}
.shejiao {
  /* 社交账号 */
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
.button {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 6px 12px;
  gap: 8px;
  height: 36px;
  width: 120px;
  border: none;
  background: #5e41de33;
  border-radius: 20px;
  cursor: pointer;
}

.lable {
  line-height: 20px;
  font-size: 17px;
  color: #5d41de;
  font-family: sans-serif;
  letter-spacing: 1px;
}

.button:hover {
  background: #5e41de4d;
}

.button:hover .svg-icon {
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}
</style>
