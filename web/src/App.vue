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
          Explore
        </router-link>
      </a-menu-item>
      <a-menu-item key="home">
        <router-link to="/" active-class="active-link">
          Home
        </router-link>
      </a-menu-item>
      <a-menu-item key="about">
        <router-link to="/about" active-class="active-link">
          About
        </router-link>
      </a-menu-item>
    </a-menu>
  </header>

  <RouterView />
</template>
