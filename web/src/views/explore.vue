<script setup></script>
<script>
export default {
  data() {
    return {
        isShow: true
    }
  },
  methods: {
    submit() {
      this.isShow =! this.isShow; // 隐藏元素
    }
  }
}
</script>
<template>
<transition name="fade">
	<span v-show="isShow">动画效果</span>
</transition>
<button @click="submit">点击隐藏</button>
</template>

<style>
.fade-enter-active, .fade-leave-active {
	transition: opacity .15s
 }
 .fade-enter, .fade-leave-to /* .fade-leave-active in <2.1.8 */ {
	opacity: 0
} 


</style>