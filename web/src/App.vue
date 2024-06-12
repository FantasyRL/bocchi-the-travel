<script setup>
import axios from "axios";
import Cookies from "js-cookie";
import { RouterLink, RouterView } from "vue-router";
import "mdui/components/icon.js";
</script>

<script>
export default {
  components: {},
  setup() {},
  data() {
    return {
      access_token: "",
      msg: "",
      username: "",
      password: "",
      email: "",
      emailRegex:
        /^(([^<>()[$$\\.,;:\s@"]+(\.[^<>()[$$\\.,;:\s@"]+)*)|(".+"))@(($$[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$$)|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
      startX: 0,
      login: 1,
      resgister: 1,
      overlay: 1,
      endX: 0,
      info: 0, // 登录状态
      id: null,
      avatar: ""
    };
  },
  mounted() {
    this.getinfo();

    this.access_token = Cookies.get("access_token");
    this.id = Cookies.get("id");
    this.avatar = Cookies.get("avatar");

    if (this.access_token && this.id) {
      // 如果access_token存在，则认为用户已经登录，更新info的值
      this.info = 1; // 假设登录成功后，将info设置为1表示已登录状态
      this.getinfo(); // 更新登录状态
    } else {
      this.info = 0; // 如果没有access_token，则认为用户未登录，将info设置为0表示未登录状态
      this.getinfo(); // 更新登录状态
    }
  },

  methods: {
    info1() {},
    loginto() {
      // 登录函数，需要根据后端接口进行调整
      axios
        .post(
          "https://api.xiey.work/bocchi/user/login/?username=" +
            this.username +
            "&password=" +
            this.password +
            "&" +
            "otp"
        )
        .then((response) => {
          console.log("结果:", response.data);
          this.msg = response.data.base.msg;

          if (response.data.base.code == 10000) {
            console.log("登录成功");
            this.cookiesSet = Cookies.set("access_token", response.data.access_token, {
              expires: 1
            });
            this.cookiesSet = Cookies.set("refresh_token", response.data.refresh_token, {
              expires: 7
            });
            this.cookiesSet = Cookies.set("id", response.data.user.id, { expires: 7 });

            this.getinfo(); // 更新登录状态
            window.location.reload();
          }
        });
    },

    registerto() {
      if (this.emailRegex.test(this.email)) {
        console.log("电子邮件格式正确");
        axios
          .post(
            "https://api.xiey.work/bocchi/user/register/?username=" +
              this.username +
              "&email=" +
              this.email +
              "&password=" +
              this.password
          )
          .then((response) => {
            console.log("结果:", response.data);
            this.msg = response.data.base.msg;
            if (response.data.base.code == "10008") {
              console.log(response.data.base.msg);
            }
          })
          .catch((error) => {
            console.error("Error", error);
          });
      } else {
        console.log("电子邮件格式错误");
        this.msg = "电子邮件格式错误";
      }
    },

    getinfo() {
      /* 获取登录状态 */
      if (this.info == 1) {
        console.log("已登录");
        this.login = 0;
        this.resgister = 0;
        this.overlay = 0;
      } else {
        console.log("未登录");
        this.login = 1;
        this.resgister = 1;
        this.overlay = 1;
      }
    },

    /* 显示注册页面 */
    userresgister() {
      this.login = 0;
      this.resgister = 1;
    },
    /* 显示登录页面 */
    userlogin() {
      this.login = 1;
      this.resgister = 0;
    },
    debuggerlogin() {
      this.info = 1; // 假设登录成功后，将info设置为1表示已登录状态
      this.getinfo(); // 更新登录状态
    }
  }
};
</script>

<template>
  <div style="width: 100%; height: 32px"></div>
  <router-view></router-view>
  <div class="overlay" v-show="overlay"></div>
  <div v-show="resgister" class="resgister-container">
    <div class="resgister">
      <div class="container">
        <div class="form">
          <div class="head">
            <span>注册</span>
            <p>输入你的名字和密码才能用哦</p>
            <text>{{ msg }}</text>
          </div>
          <div class="inputs">
            <input type="text" placeholder="用户名" v-model="username" />
            <input type="email" placeholder="邮箱" v-model="email" />
            <input type="password" placeholder="密码" v-model="password" />
          </div>
          <button @click="registerto">注册</button>
        </div>
        <div class="form-footer">
          <p>
            已经有了一个账号?
            <button
              @click="userlogin"
              style="border: none; cursor: pointer; background-color: transparent"
            >
              <a href style="pointer-events: none">登录</a>
            </button>
          </p>
        </div>
      </div>
    </div>
  </div>

  <div v-show="login" class="login-container">
    <div class="login">
      <div class="container">
        <div class="form">
          <div class="head">
            <span>登录</span>
            <p>输入你的名字和密码才能用哦</p>
            <text>{{ msg }}</text>
          </div>
          <div class="inputs">
            <input type="text" placeholder="用户名" v-model="username" />
            <input type="password" placeholder="密码" v-model="password" />
            <div class="bad"></div>
          </div>
          <button @click="loginto">登录</button>
          <button @click="debuggerlogin">开发者登录</button>
        </div>
        <div class="form-footer">
          <p>
            想注册一个账号?
            <button
              @click="userresgister"
              style="border: none; cursor: pointer; background-color: transparent"
            >
              <a href style="pointer-events: none">注册</a>
            </button>
          </p>
        </div>
      </div>
    </div>
  </div>

  <mdui-navigation-bar value="" label-visibility="labeled">
    <mdui-navigation-bar-item
      icon="place--outlined"
      value="item-1"
      @click="$router.push('/explore')"
      >探索</mdui-navigation-bar-item
    >
    <mdui-navigation-bar-item icon="home--outlined" value="item-2" @click="$router.push('/')"
      >主页</mdui-navigation-bar-item
    >
    <mdui-navigation-bar-item icon="add--outlined" value="item-3" @click="$router.push('/create')"
      >创建活动</mdui-navigation-bar-item
    >
    <mdui-navigation-bar-item
      icon="format_list_bulleted--outlined"
      value="item-4"
      @click="$router.push('/alltravels/')"
      >所有行程</mdui-navigation-bar-item
    >
    <mdui-navigation-bar-item
      icon="account_circle--outlined"
      value="item-5"
      @click="$router.push('/about')"
      >个人中心</mdui-navigation-bar-item
    >
  </mdui-navigation-bar>

  <div class="cecheimg">
    <img :src="avatar" alt="ceche" />
  </div>
</template>

<style>
.cecheimg {
  position: fixed;
  top: 550%;
  opacity: 0;
}
mdui-navigation-bar {
  z-index: 99999;
}

.menu-container {
  width: 100%;
  background-color: #fff;
}
.bad {
  background-color: #f1f7fe;
  height: 40px;
}

.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(83, 82, 82, 0.5);
  z-index: 555;
  backdrop-filter: blur(5px);
}

/* resgister */
/* resgister */
/* resgister */
.resgister {
  position: fixed;
  z-index: 1000;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 300px;
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.login {
  position: fixed;
  z-index: 1005;
  left: 50%;
  top: 50%;
  width: 300px;
  transform: translate(-50%, -50%);
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.container {
  max-width: 300px;
  background-color: #f1f7fe;
  border-radius: 15px;
  overflow: hidden;
}

.form {
  display: flex;
  align-items: center;
  flex-direction: column;
  padding: 32px 24px 24px;
  gap: 16px;
  text-align: center;
}

.form .head {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.form .head span {
  font-size: 1.6rem;
  font-weight: bolder;
  color: black;
}

.form .head p {
  font-size: 1.1rem;
  color: #7c6666;
}

.form .inputs {
  overflow: hidden;
  background-color: #fff;
  width: 100%;
  margin: 1rem 0.5rem;
  border-radius: 8px;
  border-bottom: none;
  outline: 0;
}

.form .inputs input {
  border: none;
  outline: none;
  padding: 8px 15px;
  /* علشان اقدر اخلي البلاس هولدر من اول الفورم بديله نفس الويدس */
  width: 100%;
  height: 40px;
  border-bottom: 1px solid rgba(128, 128, 128, 0.299);
  font-weight: 200;
}

.form button {
  background-color: #4287ef;
  color: white;
  width: 100%;
  height: 40px;
  padding-top: 8px;
  padding-bottom: 8px;
  border: 0;
  overflow: hidden;
  border-radius: 25px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 1s ease-in-out;
}

.form button:hover {
  background-color: #005ce6;
}

.form-footer {
  background-color: #e0ecfb;
  padding: 16px;
  font-size: 1rem;
  text-align: center;
}

.form-footer a {
  font-weight: bolder;
  color: #0066ff;
  transition: all 3s ease-in-out;
}

.form-footer a:hover {
  color: #005ce6;
}

/* resgister */
/* resgister */
/* resgister */
</style>
