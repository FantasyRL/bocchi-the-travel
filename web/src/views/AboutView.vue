<script setup>
import { UserOutlined } from '@ant-design/icons-vue';
import { useCounterStore } from '@/stores/counter'
const counter = useCounterStore()
</script>

<template>
  <a-page-header style="border: 1px solid rgb(235, 237, 240)" title="个人界面" sub-title="副标题"
    @back="() => $router.go(-1)" />
  <button @click="counter.increment">+1</button>

  <div class="info">


    <div class="ab1">
      <!-- 头像 start -->
      <div class="touxiang">
        <a-space direction="vertical" :size="32">
          <a-space :size="16">
            <a-avatar :size="64" src="https://xiey.work/640.jpg">
              <template #icon>
                <UserOutlined />
              </template>
            </a-avatar>
          </a-space>
        </a-space>
      </div>
      <!-- 头像 end -->

      <div class="name">我的名字</div>


    </div>


    <!-- 个人信息 start -->
    <div class="ab2">

      <a-empty />
    </div>
    <!-- 个人信息 end -->








  </div>
  <!-- info end -->


</template>

<style scoped></style>
