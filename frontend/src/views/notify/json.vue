<template>
  <a-button @click="handleClick">添加</a-button>
  <a-modal v-model:visible="visible" title="Modal Form" @cancel="handleCancel" @before-ok="handleBeforeOk">
    <a-form :model="form">
      <a-form-item field="name" label="名称">
        <a-input v-model="form.name"/>
      </a-form-item>
      <a-form-item field="name" label="资产">
        <a-input v-model="form.symbol"/>
      </a-form-item>
      <a-form-item field="name" label="时间周期">
        <a-input v-model="form.interval"/>
      </a-form-item>
      <a-form-item field="name" label="策略文件">
        <a-input v-model="form.path"/>
      </a-form-item>

    </a-form>
  </a-modal>
  <a-table :data="data" style="margin-top: 30px">
    <template #columns>

      <a-table-column title="Salary" data-index="metadata.title"></a-table-column>
      <a-table-column title="Address" data-index="metadata.symbol"></a-table-column>
      <a-table-column title="Email" data-index="metadata.interval"></a-table-column>
      <a-table-column title="Email" data-index="metadata.path"></a-table-column>
      <a-table-column title="Email" data-index="metadata.status"></a-table-column>
      <a-table-column title="Optional">
        <template #cell="{ record }">
          <a-button @click="runBot(record)">运行</a-button>
          <a-button @click="closeBot(record)">停止</a-button>

          <a-button @click="deleteR(record)"> 删除</a-button>

        </template>

      </a-table-column>
    </template>
  </a-table>
</template>

<script setup>
import {onMounted, ref, reactive} from 'vue';

import {Message} from "@arco-design/web-vue";

const app = ref(window.go.gtools.App)

const show = ref(true)
const data = ref([])
const columns = [{
  title: '名称',
  dataIndex: 'title',
}, {
  title: '币',
  dataIndex: 'symbol',
}, {
  title: '时间周期',
  dataIndex: 'interval',
}];
const runBot = (record) => {
  app.value.RunAlert(record.id).then(res => {
    console.log(res)

  })
}
const closeBot = (record) => {
  app.value.CloseAlert(record.id).then(res => {
    console.log(res)

  })
}
const deleteR = (record) => {
  console.log(record)
  app.value.DelAlertItem(record.id).then(res => {
    getData ()

  })
}
const getData = () => {
  app.value.GetAlertList().then(res => {
    console.log(res)
    if (res.code == 200) {
      data.value = res.data.list

    } else {
      //message.error(res.msg)
    }
  })
}

const visible = ref(false);
const form = ref({
  title: '',
  symbol: '',
  interval: '',
  path: '',
});

const handleClick = () => {
  visible.value = true;

};
const handleBeforeOk = (done) => {
  console.log(form)

  done()
  const data={
    title: form.value.title,
    symbol: form.value.symbol,
    interval: form.value.interval,
    path: form.value.path,
  }
  console.log(data)
  let  res=app.value.AddAlertItem(data)
  console.log(res)
  getData ()

};
const handleCancel = () => {
  visible.value = false;
}

onMounted(() => {
  getData()
})
</script>
