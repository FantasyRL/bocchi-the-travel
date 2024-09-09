import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('apiurl', () => {
  const apiurl = ref("https://apiurl.com")

  function increment() {
    apiurl.value = "https://apiurl.com"
  }

  return { apiurl, increment }
})
