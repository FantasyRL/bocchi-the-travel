<script setup>
import { RouterLink, RouterView } from "vue-router";
</script>

<script>
import { Menu } from "ant-design-vue";
export default {
  components: {
    "a-menu": Menu,
    "a-menu-item": Menu.Item
  },

  //不知道有没有用的左右页面切换
  data() {
    return {
      startX: 0,
      endX: 0
    };
  },
  mounted() {
    window.addEventListener('touchstart', this.touchStart);
    window.addEventListener('touchend', this.touchEnd);
    
  },
  beforeUnmount() {
    window.removeEventListener('touchstart', this.touchStart);
    window.removeEventListener('touchend', this.touchEnd);
  },
  methods: {
    touchStart(e) {
      this.startX = e.touches[0].clientX;
    },
    touchEnd(e) {
      this.endX = e.changedTouches[0].clientX;
      this.handleSwipe();
    },
    handleSwipe() {
      const diffX = this.startX - this.endX;
      if (Math.abs(diffX) > 100) { // 阈值可以根据需要调整
        if (diffX > 0) {
          // 向左滑动
          this.$router.push('/explore');
        } else {
          // 向右滑动
          this.$router.push('/about');
        }
      }
    }
  }
};
</script>

<template>
  <header>
    <a-menu mode="horizontal">
      <a-menu-item key="explore">
        <router-link to="/explore" active-class="active-link"> 探索 </router-link>
      </a-menu-item>
      <a-menu-item key="home">
        <router-link to="/" active-class="active-link"> 主页 </router-link>
      </a-menu-item>
      <a-menu-item key="about">
        <router-link to="/about" active-class="active-link"> 我的 </router-link>
      </a-menu-item>
    </a-menu>
  </header>

  <RouterView />

   <div class="overlay"></div>
    <div class="login">
    <div class="container">
      <form>
        <div class="head">
          <span>注册</span>
          <p>输入你的名字和密码才能用哦</p>
        </div>
        <div class="inputs">
          <input type="text" placeholder="用户名">
          <input type="email" placeholder="邮箱">
          <input type="password" placeholder="密码">
        </div>
        <button>注册</button>
      </form>
      <div class="form-footer">
        <p>已经有了一个账号? <a href="#">登录</a></p>
      </div>

    </div>
  </div>

  <!--<div class="toolbox">
  <Text>Debug box</Text>
  <RouterLink to="/">主页 </RouterLink>
  <RouterLink to="/about">个人界面</RouterLink>
  <button @click="counter.increment">counter +1</button>
  <div class="counter">Count: {{ counter.count }}</div>
</div>-->
</template>


<style>
.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(83, 82, 82, 0.5);
  z-index: 1000;
  backdrop-filter: blur(5px);
}

/* login */
/* login */
/* login */
.login {
  position: fixed;
  z-index: 1000;
  left: 50%;
  top: 45%;
  transform: translate(-50%, -50%);
  width: 300px;
  /* 设置宽度 */
  height: 300px;
  /* 设置高度 */
  background-color: #fff;
  /* 设置背景颜色 */
  border-radius: 10px;
  /* 设置圆角 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  /* 设置阴影 */
}

.container {
  max-width: 300px;
  background-color: #F1F7FE;
  border-radius: 15px;
  overflow: hidden;
}

form {
  display: flex;
  align-items: center;
  flex-direction: column;
  padding: 32px 24px 24px;
  gap: 16px;
  text-align: center;
}

form .head {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

form .head span {
  font-size: 1.6rem;
  font-weight: bolder;
  color: black;
}

form .head p {
  font-size: 1.1rem;
  color: #7C6666;
}

form .inputs {
  overflow: hidden;
  background-color: #fff;
  width: 100%;
  margin: 1rem 0.5rem;
  border-radius: 8px;
  border-bottom: none;
  outline: 0;
}

form .inputs input {
  border: none;
  outline: none;
  padding: 8px 15px;
  /* علشان اقدر اخلي البلاس هولدر من اول الفورم بديله نفس الويدس */
  width: 100%;
  height: 40px;
  border-bottom: 1px solid rgba(128, 128, 128, 0.299);
  font-weight: 200;
}

form button {
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

form button:hover {
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

/* login */
/* login */
/* login */
</style>
