<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
import { Menu } from 'ant-design-vue'
export default {
  components: {
    'a-menu': Menu,
    'a-menu-item': Menu.Item
  }
}
</script>

<template>
  <header>
    <a-menu mode="horizontal">
      <a-menu-item key="explore">
        <router-link to="/explore" active-class="active-link">
          探索
        </router-link>
      </a-menu-item>
      <a-menu-item key="home">
        <router-link to="/" active-class="active-link">
          主页
        </router-link>
      </a-menu-item>
      <a-menu-item key="about">
        <router-link to="/about" active-class="active-link">
          我的
        </router-link>
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
