import http from './axios'

const fetchLogin = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/v1/login',
    data: body
  })
  http.defaults.headers.common['Authorization'] = data.data?.accessToken
  http.defaults.headers.common['Username'] = data.data?.username
  return data
}

const fetchRegister = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/register',
    data: body
  })
  return data
}
;``
const fetchTicketSearch = async (params) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/ticket-service/ticket/query',
    params
  })
  return data
}

const fetchRegionStation = async (params) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/ticket-service/region-station/query',
    params
  })
  return data
}

const fetchPassengerList = async (params) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/passenger/query',
    params
  })
  return data
}
const fetchDeletePassenger = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/passenger/remove',
    data: body
  })
  return data
}

const fetchAddPassenger = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/passenger/save',
    data: body
  })
  return data
}

const fetchEditPassenger = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/passenger/update',
    data: body
  })
  return data
}
const fetchLogout = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/logout',
    data: body
  })
  http.defaults.headers.common['Authorization'] = null
  http.defaults.headers.common['Username'] = null
  return data
}

const fetchBuyTicket = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/ticket-service/ticket/purchase',
    data: body
  })

  return data
}

const fetchOrderBySn = async (params) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/order-service/order/ticket/query',
    params
  })
  return data
}

const fetchPay = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/pay-service/pay/create',
    data: body
  })
  return data
}

const fetchStationAll = async () => {
  const { data } = await http({
    method: 'POST',
    url: '/api/ticket-service/station/all'
  })
  return data
}

const fechUserInfo = async (params) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/query',
    params
  })
  return data
}

const fetchTrainStation = async (params) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/ticket-service/train-station/query',
    params
  })
  return data
}

const fetchMyTicketList = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/ticket-service/ticket/my-ticket',
    data: body
  })
  return data
}

const fetchTicketList = async (params) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/order-service/order/ticket/page',
    params
  })
  return data
}

const fetchOrderCancel = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/ticket-service/ticket/cancel',
    data: body
  })
  return data
}

const fetchUserUpdate = async (body) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/user-service/update',
    data: body
  })
  return data
}

const fetchOrderStatus = async (params) => {
  const { data } = await http({
    method: 'POST',
    url: '/api/pay-service/pay/query/order-sn',
    params
  })
  return data
}

export {
  fetchLogin,
  fetchRegister,
  fetchTicketSearch,
  fetchRegionStation,
  fetchPassengerList,
  fetchDeletePassenger,
  fetchAddPassenger,
  fetchEditPassenger,
  fetchLogout,
  fetchBuyTicket,
  fetchOrderBySn,
  fetchPay,
  fetchStationAll,
  fechUserInfo,
  fetchTrainStation,
  fetchTicketList,
  fetchMyTicketList,
  fetchOrderCancel,
  fetchOrderStatus,
  fetchUserUpdate
}
