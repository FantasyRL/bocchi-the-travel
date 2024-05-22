<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { useCounterStore } from '@/stores/counter'
const counter = useCounterStore()
</script>



<template>











  <RouterView />
  <div class="toolbox">
    <Text>Debug box</Text>
    <RouterLink to="/">主页 </RouterLink>
    <RouterLink to="/about">个人界面</RouterLink>
    <button @click="counter.increment">counter +1</button>
    <div class="counter">Count: {{ counter.count }}</div>
  </div>
</template>
