<template>
  <Card>
    <Modal
      v-model:visible="cancelVisible"
      title="退票确认"
      @ok="handleCancelConfirm"
      @cancel="handleCancel"
    >
      <p>请再次确认是否要退票</p>
    </Modal>
    <Table :loading="loading" :style="{ marginTop: '5px' }" :columns="columns" :data-source="ticketList"
      :pagination="false">
      <template #ticketStatus="{ text }">
        <span :style="{ color: text === 0 ? 'green' : 'gray' }">
          {{ text === 0 ? '已购票' : '已取消' }}
        </span>
      </template>

      <template #actions="{ record }">
        <span v-if="record.ticketStatus === 0" @click="() => showCancelConfirm(record.id)"
          :style="{ color: 'red', cursor: 'pointer' }">
          退票
        </span>
      </template>
    </Table>
  </Card>
</template>

<script setup>
import { Card, Input, Button, Space, Table, message, Modal } from "ant-design-vue";
import { DeleteOutlined } from "@ant-design/icons-vue";
import { useRouter } from "vue-router";
import { fetchMyTicketList, fetchOrderCancel } from "@/service/index";
import { onMounted, reactive, ref } from "vue";
import Cookie from 'js-cookie'


const ticketList = ref([]);
const loading = ref(false);
const state = reactive({
  selectedRowKeys: [],
  searchUsername: Cookie.get('username')
});
// 新增响应式变量
const cancelVisible = ref(false)
const currentTicketId = ref(null)
onMounted(() => {
  getTicketList();
});

// 获取车票列表
const getTicketList = () => {
  loading.value = true;
  fetchMyTicketList({ username: state.searchUsername })
    .then((res) => {
      loading.value = false;
      if (!res.success) return message.error(res.message);
      ticketList.value = res.data.map((item) => ({ ...item, key: item.id })) ?? [];
    })
    .catch(() => {
      loading.value = false;
    });
};

// 取消单张车票
const cancelTicket = (id) => {
  fetchOrderCancel({
    ticketID: id
  })
    .then((res) => {
      if (!res.success) return message.error(res.message);
      getTicketList();
    });
};
const handleCancelConfirm = () => {
  if (!currentTicketId.value) return
  
  fetchOrderCancel({ ticketID: currentTicketId.value })
    .then((res) => {
      cancelVisible.value = false
      if (!res.success) return message.error(res.message)
      message.success('退票成功')
      getTicketList()
    })
}
// 修改actions模板中的点击事件
const showCancelConfirm = (id) => {
  currentTicketId.value = id
  cancelVisible.value = true
}

// 表格列定义
const columns = [
  { title: "乘车人", dataIndex: "passengerName" },
  { title: "列车", dataIndex: "trainNumber" },
  { title: "车厢号", dataIndex: "carriageNumber" },
  { title: "座位号", dataIndex: "seatNumber" },
  { title: "出发站", dataIndex: "departure" },
  { title: "到达站", dataIndex: "arrival" },
  { title: "出发时间", dataIndex: "departureTime" },
  { title: "到达时间", dataIndex: "arrivalTime" },
  { title: "状态", dataIndex: "ticketStatus", slots: { customRender: "ticketStatus" } },
  { title: "操作", dataIndex: "actions", slots: { customRender: "actions" } },
];

// 多选框选中变化
const onSelectChange = (keys) => {
  state.selectedRowKeys = keys;
};
</script>

<style lang="scss" scoped></style>
