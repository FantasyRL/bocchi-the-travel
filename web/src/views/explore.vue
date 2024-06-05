<script setup>
import dayjs from "dayjs";
import { ref } from "vue";
const dateFormat = "YYYY/MM/DD";
const weekFormat = "MM/DD";
const monthFormat = "YYYY/MM";
const dateFormatList = ["DD/MM/YYYY", "DD/MM/YY"];
const customWeekStartEndFormat = (value) =>
  `${dayjs(value).startOf("week").format(weekFormat)} ~ ${dayjs(value)
    .endOf("week")
    .format(weekFormat)}`;
const value1 = ref(dayjs("2015/01/01", dateFormat));
const value2 = ref(dayjs("01/01/2015", dateFormatList[0]));
const value3 = ref(dayjs("2015/01", monthFormat));
const value4 = ref([dayjs("2015/01/01", dateFormat), dayjs("2015/01/01", dateFormat)]);
const value5 = ref(dayjs("2015/01/01", dateFormat));
const value6 = ref(dayjs());
const customFormat = (value) => `custom format: ${value.format(dateFormat)}`;
</script>
<script>
const dateTimeStr = "Thu, 31 Jul 2025 16:00:00 GMT";
const date = new Date(dateTimeStr);
const formattedDate = date.toISOString().slice(0, 10);

console.log(formattedDate); // 输出：2025-07-31
</script>
<template>
  <a-space direction="vertical" :size="12">
    <a-date-picker v-model:value="value1" :format="dateFormat" />
    <a-date-picker v-model:value="value2" :format="dateFormatList" />
    <a-date-picker v-model:value="value3" :format="monthFormat" picker="month" />
    <a-range-picker v-model:value="value4" :format="dateFormat" />
    <a-date-picker v-model:value="value5" :format="customFormat" />
    <a-date-picker v-model:value="value6" :format="customWeekStartEndFormat" picker="week" />
  </a-space>
  {{ value1 }}
</template>
