<script setup></script>
<script>
export default {
  data() {
    return {
      trnumber: 10
    };
  }
};
</script>
<!-- 我的行程页面  -->
<template>
  <div class="travels">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="所有行程"
      @back="() => $router.go(-1)"
    />
  </div>
  <div v-for="trnumber in trnumber" :key="trnumber">
    <br />
    <el-card style="max-width: 480px">
      <template #header>
        <div class="card-header">
          <span>Card name {{ trnumber }}</span>
        </div>
      </template>
      <p v-for="o in 4" :key="o" class="text item">{{ "List item " + o }}</p>
      <template #footer>Footer content</template>
    </el-card>
    <br />
  </div>
  <br />
  <div class="null"></div>
  <br />
</template>

<style scoped>
.null {
  height: 100px;
}
</style>
