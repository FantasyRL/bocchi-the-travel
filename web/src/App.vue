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

  <!--<div class="toolbox">
  <Text>Debug box</Text>
  <RouterLink to="/">主页 </RouterLink>
  <RouterLink to="/about">个人界面</RouterLink>
  <button @click="counter.increment">counter +1</button>
  <div class="counter">Count: {{ counter.count }}</div>
</div>-->
</template>
