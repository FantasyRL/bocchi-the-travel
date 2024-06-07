<script setup>
import { ref } from "vue";
</script>
<script>
export default {
  setup() {
    const options = ref([
      { id: 1, value: "路线" },
      { id: 2, value: "活动" },
      { id: 3, value: "住处" },
      { id: 4, value: "吃喝" },
      { id: 5, value: "其他" }
    ]);

    const selectedOption = ref(null);

    function sendData() {
      console.log("发送给后端的数字:", selectedOption.value);
    }

    return {
      options,
      selectedOption,
      sendData
    };
  }
};
</script>

<template>
  <div>
    <label for="options">选择：</label>
    <select id="options" v-model="selectedOption" @change="sendData">
      <option v-for="option in options" :value="option.value" :key="option.value">
        {{ option.value }}
      </option>
    </select>
  </div>
</template>
