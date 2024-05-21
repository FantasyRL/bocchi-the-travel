<script setup>
import { useCounterStore } from '@/stores/counter'
const counter = useCounterStore()
</script>
<template>
  <div class="home">
    <h1>主页</h1>
    <button @click="counter.increment">+1</button>
  </div>
</template>

<style></style>
