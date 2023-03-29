<template>
  <div class="ml-2 mt-3.5 hidden lg:block">
    <a-row class="grid-demo" style="margin-right: 16px;">
      <a-col flex="100px" style="margin-right: 16px;">
        <a-select :style="{width:'120px'}" allow-search v-model="options.symbol" @change="resetSymbol">
          <a-option v-for="symbol in symbols" :value="symbol">{{ symbol }}</a-option>
        </a-select>
      </a-col>
      <a-col flex="auto">
        <a-select :style="{width:'180px'}" v-model="options.interval" @change="resetInterval">
          <a-option v-for="interval in intervals" :value="interval">{{ interval }}</a-option>

        </a-select>
      </a-col>
    </a-row>


  </div>
</template>

<script setup>
import {useRoute, useRouter} from 'vue-router'
import {useAppStore} from '@/store'
import {onMounted, reactive, ref} from "vue";
import {emitter} from "@/utils/bus";

const app = ref(window.go.gtools.App)
const appStore = useAppStore()
const intervals = ref([])
const symbols = ref([])
const options = reactive({
  symbol: "ETHUSDT",
  interval: "1m",
})
const resetInterval = (interval) => {
  options.interval = interval
  emitter.emit('symbolChange', options)
  // for (let sk of sockets.value.keys()) {
  //   sockets.value.get(sk)()
  // }

}
const resetSymbol = () => {
  emitter.emit('symbolChange', options)

}

const getData = () => {
  app.value.AppSetting().then(res => {
    console.log(res)
    if (res.code == 200) {
      intervals.value = res.data.intervals
      symbols.value = res.data.symbols

    } else {
      //message.error(res.msg)
    }
  })
}

onMounted(() => {
  getData()
})
</script>
