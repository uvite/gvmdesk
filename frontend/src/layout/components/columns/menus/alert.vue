<template>
  <div>
    <div class="menu-title flex items-center">
      <a-space>

        <a-button @click="createAlert">创建警报</a-button>
      </a-space>
    </div>
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

    <div>
      <a-split direction="vertical" :style="{height: '700px'}" v-model:size="size">
        <template #first>

          <a-list
              :style="{ width: `100%` }"
              :virtualListProps="{
                height: 560,
              }"
              :data="data"
          >
            <template #item="{ item, index }">

              <a-list-item :key="index" style="border-bottom: rgba(14,1,1,0.14) solid 1px">
                <a-row class="grid-demo" style="margin-bottom: 16px;">
                  <a-col flex="50px">
                    <div>{{ item.metadata.symbol }}</div>
                  </a-col>
                  <a-col flex="auto">
                    <a-button-group>
                      <a-button type="primary" status="success" size="small" @click="runBot(item)">运行</a-button>
                      <a-button type="primary" status="danger" size="small" @click="closeBot(item)">停止</a-button>
                      <a-button    size="small" @click="deleteR(item)"> 删除</a-button>
                    </a-button-group>
                  </a-col>
                </a-row>


              </a-list-item>
            </template>
          </a-list>


        </template>
        <template #second>
          afasdfasdf
        </template>
      </a-split>
    </div>

  </div>
</template>
<script setup>

import {onMounted, ref, reactive} from 'vue';

import {Message} from "@arco-design/web-vue";

const app = ref(window.go.gtools.App)

const show = ref(true)
const sizeList = ref('small');
const size = ref(0.5)
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

const createAlert = () => {
  visible.value = true;
};
const handleBeforeOk = (done) => {
  console.log(form)

  done()
  const data = {
    title: form.value.title,
    symbol: form.value.symbol,
    interval: form.value.interval,
    path: form.value.path,
  }
  console.log(data)
  let res = app.value.AddAlertItem(data)
  console.log(res)
  getData()

};
const handleCancel = () => {
  visible.value = false;
}

onMounted(() => {
  getData()
})

</script>
