<template>
  <mdui-navigation-bar value="" label-visibility="labeled">
    <!-- <mdui-navigation-bar-item
      icon="place--outlined"
      value="item-1"
      @click="$router.push('/explore')"
      >探索</mdui-navigation-bar-item
    > -->
    <mdui-navigation-bar-item icon="home--outlined" value="item-2" @click="$router.push('/')"
      >主页</mdui-navigation-bar-item
    >
    <mdui-navigation-bar-item icon="add--outlined" value="item-3" @click="$router.push('/create')"
      >创建活动</mdui-navigation-bar-item
    >
    <mdui-navigation-bar-item
      icon="format_list_bulleted--outlined"
      value="item-4"
      @click="$router.push('/alltravels/')"
      >所有行程</mdui-navigation-bar-item
    >
    <mdui-navigation-bar-item
      icon="account_circle--outlined"
      value="item-5"
      @click="$router.push('/about')"
      >个人中心</mdui-navigation-bar-item
    >
  </mdui-navigation-bar>
</template>
<style>
mdui-navigation-bar {
  z-index: 554;
}
</style>
