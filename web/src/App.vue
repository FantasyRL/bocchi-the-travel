<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { useCounterStore } from '@/stores/counter'
const counter = useCounterStore()
</script>



<template>



  <div>Count: {{ counter.count }}</div>

  <button @click="counter.increment">+1</button>


  <RouterLink to="/">主页 </RouterLink>
  <RouterLink to="/about">个人界面</RouterLink>


  <RouterView />
</template>
