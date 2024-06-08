<script setup></script>
<script></script>

<template>
  <div class="grid">
    <div class="box">项目 1</div>
    <div class="box">项目 2</div>
    <div class="box">项目 3</div>
    <div class="box">项目 4</div>
  </div>
</template>

<style>
.grid {
  display: grid;
  grid-template-columns: 1fr 2fr 1fr 1fr;
}

.box {
  background-color: #f2f2f2;

  width: 100%;
  height: 200px;
}
</style>
