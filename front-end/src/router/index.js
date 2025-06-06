import { createRouter, createWebHistory } from 'vue-router'
import { message } from 'ant-design-vue'
import Login from '../views/login'
import TicketSearch from '../views/ticket-serach'
import Userinfo from '../views/user-info'
import Passenger from '../views/passenger'
import MyTicket from '../views/my-ticket'
import AddPassanger from '../views/add-passenger'
import BuyTicket from '../views/buy-ticket'
import Cookies from 'js-cookie'
const routes = [
  {
    path: '/',
    name: 'index',
    meta: { requiresAuth: false },
    redirect: (to) => {
      // 该函数接收目标路由作为参数
      // 相对位置不以`/`开头
      // 或 { path: 'profile'}
      return '/ticketSearch'
    }
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: { requiresAuth: false }
  },
  {
    label: '车票查询',
    path: '/ticketSearch',
    name: 'ticketSearch',
    component: TicketSearch,
    icon: 'icon-chaxun',
    meta: { requiresAuth: false }
  },
  {
    label: '用户信息',
    path: '/userInfo',
    name: 'userInfo',
    component: Userinfo,
    icon: 'icon-ic_account',
    meta: { requiresAuth: true }
  },
  {
    label: '人员信息',
    path: '/passenger',
    name: 'passenger',
    component: Passenger,
    icon: 'icon-renyuanfenbu',
    meta: { requiresAuth: true }
  },
  {
    label: '我的车票',
    path: '/myTicket',
    name: 'myTicket',
    component: MyTicket,
    icon: 'icon-dingdan',
    meta: { requiresAuth: true }
  },
  {
    label: '添加乘车人',
    path: '/addPassenger',
    name: 'addPassenger',
    component: AddPassanger,
    icon: 'icon-icon-adduser',
    meta: { requiresAuth: true }
  },
  {
    label: '预订车票',
    path: '/buyTicket',
    name: 'buyTicket',
    component: BuyTicket,
    icon: 'icon-goumai',
    meta: { requiresAuth: true }
  },
  // {
  //   label: '支付宝支付',
  //   path: '/aliPay',
  //   name: 'aliPay',
  //   component: AliPay,
  //   icon: 'icon-zhifubao0',
  //   meta: { requiresAuth: true }
  // },
  // {
  //   label: '订单列表',
  //   path: '/ticketList',
  //   name: 'ticketList',
  //   component: TikectList,
  //   icon: 'icon-zhifubao0',
  //   meta: { requiresAuth: true }
  // }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach(async (to, from) => {
  console.log(Cookies.get('username'))
  console.log(to, 'to')
  console.log(from, 'from')

  if (
    to.meta?.requiresAuth &&
    to.name !== 'login' &&
    (!Cookies.get('username') || !Cookies.get('token'))
  ) {
    message.error('用户未登录或已过期！')
    return {
      name: 'login'
    }
  }
})

export { routes }

export default router
