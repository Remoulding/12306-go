<template>
  <div class="buy-ticket-wrapper">
    <Space direction="vertical" :style="{ width: '100%' }" size="large">
      <Card>
        <template #title>
          <div class="title-wrapper">
            <h1>列表信息</h1>
            <h2>（以下余票信息仅供参考）</h2>
          </div>
        </template>
        <div>
          <span class="important-text">{{ query.departureDate }}</span>
          <span class="important-text">（{{ getWeekNumber(dayjs(query.departureDate).day()) }}）</span>
          <span class="important-text">{{ query.trainNumber }}</span>
          <span class="small-text">次</span>
          <span class="important-text">{{ state.currTrain?.departure }}</span>
          <span class="small-text">站</span>
          <span class="important-text">（{{ state.currTrain?.departureTime }}开）——</span>
          <span class="important-text">{{ state.currTrain?.arrival }}</span>
          <span class="small-text">站</span>
          <span class="important-text">（{{ state.currTrain?.arrivalTime }}到）</span>
        </div>
        <Divider dashed></Divider>
        <div class="seat-wrapper">
          <div v-for="item in state.currentSeat">
            <span>{{
              SEAT_CLASS_TYPE_LIST.find((seat) => seat.code === item.type)
                ?.label
            }}</span>
            （ <span class="price">￥{{ item?.price }}</span>）
            <span>{{ item.quantity >= 1 ? '有票' : '1张票' }}</span>
          </div>
        </div>
      </Card>
      <Card>
        <template #title>
          <div class="title-wrapper">
            <h1>乘客信息</h1>
            <h2>（填写说明）</h2>
          </div>
        </template>
        <div class="user-tip">
          <UserOutlined />
          <span>乘车人</span>
        </div>
        <div class="check-wrapper">
          <CheckboxGroup v-model:value="currPassenger">
            <Checkbox v-for="item in state.currPassengerList" :value="item.id">{{ item.realName }}</Checkbox>
          </CheckboxGroup>
        </div>
        <Divider></Divider>
        <Table :columns="columns" :data-source="state.dataSource.filter((item) => currPassenger?.includes(item.id))
          " :locale="{ emptyText: '请先选择乘车人' }" :pagination="false">
          <template #ticketType="{ text, record }">
            <Select :style="styleWidth" :value="text" @select="(value) => handleChooseTicketType(record.id, value)">
              <SelectOption v-for="item in TICKET_TYPE_LIST" :value="item.value">{{ item.label }}</SelectOption>
            </Select>
          </template>
          <template #seatType="{ text, record }">
            <Select :style="styleWidth" @select="(value) => handleChooseSeat(record.id, value)">
              <SelectOption v-for="item in state.currentSeat" :value="item.type">{{
                `${SEAT_CLASS_TYPE_LIST.find((seat) => seat.code === item.type)
                  ?.label
                }(￥${item.price})`
              }}</SelectOption>
            </Select>
          </template>
          <template #realName="{ text }">
            <Input :value="text" disabled :style="styleWidth"></Input>
          </template>
          <template #idType="{ text }">
            <Select disabled :style="styleWidth" :value="text">
              <SelectOption v-for="item in ID_CARD_TYPE" :value="item.value">{{
                item.label
                }}</SelectOption>
            </Select>
          </template>
          <template #idCard="{ text }">
            <Input disabled :value="text" :style="styleWidth"></Input>
          </template>
          <template #edit="{ text, record }">
            <CloseCircleOutlined :style="{ color: '#fd6a09' }" @click="() => handleDelete(record.id)" />
          </template>
        </Table>
      </Card>
      <!-- <div>
        <span>提交订单表示已阅读并同意</span>
        <span><a href="">《零元购协议》</a></span>
      </div> -->
      <div :style="{ width: '100%', textAlign: 'center' }">
        <Space size="large">
          <Button @click="() => router.push('/ticketSearch')">上一步</Button>
          <Button class="submit-btn" @click="
            () => {
              handleSubmit()
            }
          ">预定车票</Button>
        </Space>
      </div>
    </Space>
    <Modal :visible="state.open" title="请核对以下信息" wrap-class-name="check-info-wrapper" width="40%"
      @cancel="state.open = false" :footer="null">
      <Space direction="vertical" :style="{ width: '100%' }">
        <div>
          <span>{{ query.departureDate }}</span>
          <span>{{ getWeekNumber(dayjs(query.departureDate).day()) }}</span>
          <span>{{ query.trainNumber }}</span>
          <span>次</span>
          <span>{{ state.currTrain?.departure }}</span>
          <span>站</span>
          <span>({{ state.currTrain?.departureTime }}开)——</span>
          <span>{{ state.currTrain?.arrival }}</span>
          <span>站</span>
          <span>{{ state.currTrain?.arrivalTime }}到</span>
        </div>
        <Table :columns="checkColumns" :data-source="state.dataSource" :pagination="false">
          <template #seatType="{ text, record }">
            <span>
              {{
                SEAT_CLASS_TYPE_LIST.find((seat) => seat.code === text)?.label
              }}
            </span>
          </template>
          <template #discountType="{ text }">
            {{TICKET_TYPE_LIST.find((item) => item.value === text)?.label}}
          </template>
          <template #idType="{ text }">
            {{ID_CARD_TYPE.find((item) => item.value === text)?.label}}
          </template>
        </Table>
        <div v-if="
          state.isChooseSeat &&
          !(state.dataSource?.length > 5) &&
          state.dataSource?.length <= state.seatPosition.length
        ">
          <a href="">*如果本次列车剩余席位无法满足您的选座需求，系统将自动为您分配席位</a>
          <div class="seat-choose-wrapper">
            <div>
              <div class="tip">
                <IconFont type="icon-laba001"></IconFont>
                选座
              </div>
              <div>
                已选座{{ state.currentSeatCode.length }}/{{
                  state.dataSource.length
                }}
              </div>
            </div>
            <div>
              <div v-for="(item, index) in state.dataSource?.slice(0, 2)">
                <div class="action-wrapper">
                  <div>窗</div>
                  <Divider type="vertical"></Divider>
                  <div>
                    <div class="seat-img" v-for="(item, index, length) in state.seatPosition
                      .slice(
                        0 + index * state.seatNumber,
                        (1 + index) * state.seatNumber
                      )
                      .slice(0, state.seatLeft)" @click="() => handleSelectSeat(item)" :class="{
                          cur: state.currentSeatCode?.indexOf(item) !== -1
                        }">
                      {{ item.slice(0, 1) }}
                    </div>
                    <Divider type="vertical"></Divider>
                  </div>
                  <div>过道</div>
                  <Divider type="vertical"></Divider>
                  <div>
                    <div class="seat-img" v-for="(item, index, length) in state.seatPosition
                      .slice(
                        0 + index * state.seatNumber,
                        (1 + index) * state.seatNumber
                      )
                      .slice(state.seatLeft, state.seatNumber)" @click="() => handleSelectSeat(item)" :class="{
                          cur: state.currentSeatCode?.indexOf(item) !== -1
                        }">
                      {{ item.slice(0, 1) }}
                    </div>
                    <Divider type="vertical"></Divider>
                  </div>
                  <div>窗</div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-else>
          <a href="">*系统将随机为您申请席位，暂不支持自选席位</a>
        </div>
        <Divider dashed></Divider>
        <div class="info-tip">
          本次列车，<span v-for="item in state.currentSeat">
            {{
              SEAT_CLASS_TYPE_LIST.find((seat) => seat.code === item.type)
                ?.label
            }}余票{{ item.quantity }},</span>
        </div>
        <Space size="large" :style="{
          width: '100%',
          textAlign: 'center',
          justifyContent: 'center'
        }">
          <Button @click="state.open = false">返回修改</Button>
          <Button :loading="state.loading" :style="{
            backgroundColor: '#ff8001',
            color: '#fff',
            border: 'none'
          }" @click="handleSubmitBuyTicket">确认</Button>
        </Space>
      </Space>
    </Modal>
  </div>
</template>

<script setup>
import {
  Card,
  Divider,
  Space,
  Checkbox,
  CheckboxGroup,
  Table,
  Select,
  SelectOption,
  Input,
  Button,
  Modal,
  message
} from 'ant-design-vue'
import { useRoute } from 'vue-router'
import {
  fetchTicketSearch,
  fetchPassengerList,
  fetchBuyTicket
} from '@/service'
import { onMounted, reactive, toRaw, watch, ref } from 'vue'
import { getWeekNumber } from '@/utils'
import {
  UserOutlined,
  CloseCircleOutlined,
  ConsoleSqlOutlined
} from '@ant-design/icons-vue'
import { TICKET_TYPE_LIST, ID_CARD_TYPE } from '@/constants'
import IconFont from '@/components/icon-font'
import dayjs from 'dayjs'
import { SEAT_CLASS_TYPE_LIST } from '@/constants'
import Cookie from 'js-cookie'
import router from '@/router'
const styleWidth = { width: '150px' }

const { query } = useRoute()
const username = Cookie.get('username')
const state = reactive({
  currTrain: null,
  currPassengerList: [],
  dataSource: [],
  rawDataSource: [],
  currentSeat: [],
  open: false,
  currentSeatCode: [],
  seatPosition: [],
  isChooseSeat: true,
  seatLeft: 3,
  seatNumber: 5,
  loading: false
})
const currPassenger = ref([])

onMounted(() => {
  fetchTicketSearch({
    fromStation: query.fromStation,
    toStation: query.toStation,
    departureDate: query.departureDate
  }).then((res) => {
    if (res.success) {
      const currTrain = res.data.trainList.find(
        (item) => item.trainNumber === query.trainNumber
      )
      state.currentSeat = currTrain.seatClassList
      state.currTrain = currTrain
    }
  })
  fetchPassengerList({ username }).then((res) => {
    if (res.success) {
      state.currPassengerList = res.data
      state.dataSource = res.data.map((item, index) => ({
        ...item,
        keyNumber: index + 1
      }))
      state.rawDataSource = res.data.map((item, index) => ({
        ...item,
        keyNumber: index + 1
      }))
    }
  })
})

watch(
  () => state.dataSource,
  (newValue) => {
    let seatType = newValue?.length && newValue[0]?.seatType
    let seatPosition = ['A', 'B', 'C', 'D', 'E']
    let seatLeft = 3
    if (
      seatType ===
      SEAT_CLASS_TYPE_LIST.find((item) => item.label === '商务座')?.code
    ) {
      seatPosition = ['A', 'C', 'F']
      seatLeft = 2
    }
    if (
      seatType ===
      SEAT_CLASS_TYPE_LIST.find((item) => item.label === '一等座')?.code
    ) {
      seatPosition = ['A', 'C', 'D', 'F']
      seatLeft = 2
    }
    if (
      seatType ===
      SEAT_CLASS_TYPE_LIST.find((item) => item.label === '二等座')?.code
    ) {
      seatPosition = ['A', 'B', 'C', 'D', 'F']
      seatLeft = 3
    }
    let seatList = new Array(newValue?.length * seatPosition.length).fill('')

    seatList = seatList?.map((item, index) => {
      if (index < seatPosition.length - 1) {
        return `${seatPosition[index]}0`
      } else {
        return `${seatPosition[index % seatPosition.length]}${Math.floor(
          index / seatPosition.length
        )}`
      }
    })
    state.seatPosition = seatList
    state.seatLeft = seatLeft
    state.seatNumber = seatPosition.length
  },
  {
    deep: true
  }
)

watch([() => currPassenger.value], (newValue) => {
  const list = newValue.map((item) => toRaw(item))?.flat()
  const newDataSource = state.rawDataSource.filter((item) => {
    return list.includes(item.id)
  })
  state.dataSource = newDataSource
})

watch(
  () => state.dataSource,
  (newValue) => {
    let isChooseSeat = true
    let seatType = newValue?.length && newValue[0].seatType
    newValue?.length &&
      newValue.map((item) => {
        if (item.seatType !== seatType) {
          isChooseSeat = false
        }
      })
    state.isChooseSeat = isChooseSeat
  },
  { deep: true }
)

const columns = [
  {
    title: '序号',
    dataIndex: 'keyNumber'
  },
  // {
  //   title: '票种',
  //   dataIndex: 'discountType',
  //   slots: { customRender: 'ticketType' }
  // },
  {
    title: '席别',
    dataIndex: '',
    slots: { customRender: 'seatType' }
  },
  {
    title: '姓名',
    dataIndex: 'realName',
    slots: { customRender: 'realName' }
  },
  {
    title: '证件类型',
    dataIndex: 'idType',
    slots: { customRender: 'idType' }
  },
  {
    title: '证件号码',
    dataIndex: 'idCard',
    slots: { customRender: 'idCard' }
  },
  {
    title: '操作',
    dataIndex: '',
    slots: { customRender: 'edit' }
  }
]

const handleDelete = (id) => {
  state.dataSource = state.dataSource.filter((item) => {
    return item.id !== id
  })
  currPassenger.value = currPassenger.value.filter((item) => {
    return item !== id
  })
}

const handleChooseSeat = (id, value) => {
  let cIndex
  state.dataSource.find((item, index) => {
    if (item.id === id) {
      cIndex = index
      return true
    }
  })
  state.dataSource[cIndex]['seatType'] = value
}

const handleChooseTicketType = (id, value) => {
  let cIndex
  state.dataSource.find((item, index) => {
    if (item.id === id) {
      cIndex = index
      return true
    }
  })
  state.dataSource[cIndex]['idType'] = value
}

const handleSubmit = () => {
  const canNotSubmit = state.dataSource.some(
    (item) => item.seatType === null || item.seatType === undefined
  )
  if (canNotSubmit || state.dataSource?.length === 0) {
    return message.error('请补全信息')
  }
  state.open = true
}

const checkColumns = [
  {
    title: '序号',
    dataIndex: 'keyNumber'
  },
  {
    title: '席别',
    dataIndex: 'seatType',
    slots: { customRender: 'seatType' }
  },
  // {
  //   title: '票种',
  //   dataIndex: 'discountType',
  //   slots: { customRender: 'discountType' }
  // },
  {
    title: '姓名',
    dataIndex: 'realName'
  },
  {
    title: '证件类型',
    dataIndex: 'idType',
    slots: { customRender: 'idType' }
  },
  {
    title: '证件号码',
    dataIndex: 'idCard'
  }
]

const handleSelectSeat = (code) => {
  if (state.currentSeatCode?.indexOf(code) !== -1) {
    state.currentSeatCode = state.currentSeatCode?.filter(
      (item) => item !== code
    )
  } else {
    if (state.currentSeatCode.length === state.dataSource.length) {
      state.currentSeatCode[state.currentSeatCode.length - 1] = code
    } else {
      state.currentSeatCode.push(code)
    }
  }
}

const handleSubmitBuyTicket = () => {
  let params = { trainId: query?.trainId }
  const passengers = state.dataSource.map((item) => ({
    passengerId: item.id,
    seatType: item.seatType
  }))
  params = {
    ...params,
    passengers,
    chooseSeats: toRaw(state.currentSeatCode),
    departure: state.currTrain?.departure,
    arrival: state.currTrain?.arrival,
    departureDate: dayjs(query.departureDate).format('YYYY-MM-DD'),
    username: username,
  }
  state.loading = true
  fetchBuyTicket(params)
    .then((res) => {
      if (res.success) {
        message.success('下单成功，正在跳转至车票详情页')
        setTimeout(() => {
          router.push(`/myTicket`)
        }, 500)
      } else {
        message.error(res.message)
      }
      state.loading = false
    })
    .catch((error) => {
      state.loading = false
      console.log(error)
    })
}
</script>

<style lang="scss" scoped>
.buy-ticket-wrapper {
  .title-wrapper {
    display: flex;
    height: 20px;
    line-height: 20px;

    h1 {
      font-size: 16px;
      margin: 0;
      color: #fff;
    }

    h2 {
      font-size: 14px;
      margin: 0;
      color: #fff;
    }
  }

  .important-text {
    font-size: 16px;
    font-weight: bolder;
    // margin-right: 5px;
  }

  .seat-wrapper {
    display: flex;
    margin-bottom: 10px;

    >div {
      margin-right: 10px;

      .price {
        color: #f57531;
      }
    }
  }

  .tip {
    color: #1e71bd;
  }

  .user-tip {
    .anticon-user {
      color: #1e71bd;
      font-size: 20px;
      margin-right: 5px;
    }
  }

  .check-wrapper {
    padding: 5px 10px;
  }

  .submit-btn {
    background-color: #fd6a09;
    color: #fff;
  }

  .tips-txt {
    background: #fffbe5;
    border: 1px solid #fbd800;
    padding: 5px;
  }
}

::v-deep {
  .ant-card-head {
    background-color: #1e71bd;
  }

  .ant-card-body {
    background-color: #ebedf6;
    padding: 10px 15px;
  }

  .ant-divider-dashed {
    border-color: #d1d3da;
  }

  .ant-divider-horizontal {
    margin: 10px 0;
  }

  .ant-table-thead {
    .ant-table-cell {
      background: #f3f3f3 !important;
      background-color: #f3f3f3 !important;
    }

    >tr>th {
      padding: 8px;
    }
  }

  .check-info-wrapper {
    .ant-modal-header {
      background-color: #2386c8;
    }
  }
}

.check-info-wrapper {
  .important-text {
    font-size: 16px;
    font-weight: bolder;
  }

  .info-tip {
    color: #909090;

    span {
      font-size: 18px;
      color: #db0000;
      font-weight: bolder;
    }
  }
}

.seat-choose-wrapper {
  display: flex;
  width: 100%;
  background-color: #f3f3f3;
  padding: 20px;

  .tip {
    color: #ff8001;
    margin-right: 40px;
  }

  .action-wrapper {
    display: flex;
    align-items: center;

    padding: 20px;

    .seat-img {
      display: inline-block;
      text-align: center;
      width: 30px;
      height: 28px;
      line-height: 25px;
      background: url(https://kyfw.12306.cn/otn/resources/images/bg017.png) no-repeat;
      color: #fff;
      margin: 0 5px;
      cursor: pointer;
      background-position: -80px 0;
    }

    .cur {
      background-position: -40px 0;
    }
  }
}
</style>
