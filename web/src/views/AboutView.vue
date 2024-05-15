<script setup>
import { UserOutlined } from '@ant-design/icons-vue';

</script>

<template>
  <a-page-header style="border: 1px solid rgb(235, 237, 240)" title="个人界面" sub-title="副标题"
    @back="() => $router.go(-1)" />

  <div class="info">

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

    <!-- 个人信息 start -->

    <a-empty />







    <!-- 个人信息 end -->
  </div>
</template>

<style scoped></style>
