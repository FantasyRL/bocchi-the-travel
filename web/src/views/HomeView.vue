<script setup>
import { useCounterStore } from '@/stores/counter'
const counter = useCounterStore()
</script>
<template>
  <div class="home">
    <h1>666</h1>
    <button @click="counter.increment">+1</button>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
