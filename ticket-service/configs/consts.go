package configs

import "time"

// 常量定义
const (
	CacheTTL = 60 * time.Minute
)

const (
	// UserIdKey 用户 ID Key
	UserIdKey = "UserId"

	// UserNameKey 用户名 Key
	UserNameKey = "Username"

	// RealNameKey 用户真实名称 Key
	RealNameKey = "RealName"

	// UserTokenKey 用户 Token Key
	UserTokenKey = "Token"
)

const (
	// JwtSecretKey JWT 密钥
	JwtSecretKey = "12306-jwt-secret-key"
)

const (
	TimeLayout = "2006-01-02"
)

const (
	// TrainInfo 列车基本信息，Key Prefix + 列车ID
	TrainInfo = "index12306TicketService:trainInfo:"

	// RegionTrainStationMapping 地区与站点映射查询
	RegionTrainStationMapping = "index12306TicketService:regionTrainStationMapping"

	// LockRegionTrainStationMapping 站点查询分布式锁 Key
	LockRegionTrainStationMapping = "index12306TicketService:lock:regionTrainStationMapping"

	// RegionTrainStation 站点查询，Key Prefix + 起始城市_终点城市_日期
	RegionTrainStation = "index12306TicketService:regionTrainStation:%s_%s"

	// LockRegionTrainStation 站点查询分布式锁 Key
	LockRegionTrainStation = "index12306TicketService:lock:regionTrainStation"

	// LockTrain 列车查询分布式锁 Key
	LockTrain = "index12306TicketService:lock:train"

	// TrainStationPrice 列车站点座位价格查询，Key Prefix + 列车ID_起始城市_终点城市
	TrainStationPrice = "index12306TicketService:trainStationPrice:%s_%s_%s"

	// LockTrainStationPrice 列车站点座位价格查询分布式锁 Key
	LockTrainStationPrice = "index12306TicketService:lock:trainStationPrice"

	// RegionStation 地区以及车站查询，Key Prefix + (车站名称或查询方式)
	RegionStation = "index12306TicketService:regionStation:"

	// TrainStationRemainingTicket 站点余票查询，Key Prefix + 列车ID_起始站点_终点
	TrainStationRemainingTicket = "index12306TicketService:trainStationRemainingTicket:%s_%s"

	// TrainCarriage 列车车厢查询，Key Prefix + 列车ID
	TrainCarriage = "index12306TicketService:trainCarriage:%s"

	// TrainStationCarriageRemainingTicket 车厢余票查询，Key Prefix + 列车ID_起始站点_终点
	TrainStationCarriageRemainingTicket = "index12306TicketService:trainStationCarriageRemainingTicket:%s_%s"

	// TrainStationDetail 站点详细信息查询，Key Prefix + 列车ID_起始站点_终点
	TrainStationDetail = "index12306TicketService:trainStationDetail:%s_%s"

	// TrainStationStopoverDetail 列车路线信息查询，Key Prefix + 列车ID
	TrainStationStopoverDetail = "index12306TicketService:trainStationStopoverDetail:%s"

	// StationAll 列车站点缓存
	StationAll = "index12306TicketService:allStation"

	// TrainCarriageSeatStatus 列车车厢状态，Key Prefix + 列车ID + 起始站点 + 目的站点 + 车厢编号
	TrainCarriageSeatStatus = "index12306TicketService:trainCarriageSeatStatus:%s_%s_%s_%s"

	// LockPurchaseTickets 用户购票分布式锁 Key
	LockPurchaseTickets = "${uniqueName:}index12306TicketService:lock:purchaseTickets_%s"

	// LockPurchaseTicketsV2 用户购票分布式锁 Key v2
	LockPurchaseTicketsV2 = "${uniqueName:}index12306TicketService:lock:purchaseTickets_%s_%d"

	// QueryAllRegionList 获取全部地点集合 Key
	QueryAllRegionList = "index12306TicketService:queryAllRegionList"

	// TicketAvailabilityTokenBucket 列车购买令牌桶，Key Prefix + 列车ID
	TicketAvailabilityTokenBucket = "index12306TicketService:ticketAvailabilityTokenBucket:%s"

	// LockQueryAllRegionList 获取全部地点集合分布式锁 Key
	LockQueryAllRegionList = "index12306TicketService:lock:queryAllRegionList"

	// LockQueryCarriageNumberList 获取列车车厢数量集合分布式锁 Key
	LockQueryCarriageNumberList = "index12306TicketService:lock:queryCarriageNumberList_%s"

	// LockQueryRegionStationList 获取地区以及站点集合分布式锁 Key
	LockQueryRegionStationList = "index12306TicketService:lock:queryRegionStationList_%s"

	LockQueryAllRegionStationList = "index12306TicketService:lock:queryAllRegionStationList"

	// LockSafeLoadSeatMarginGet 获取相邻座位余票分布式锁 Key
	LockSafeLoadSeatMarginGet = "index12306TicketService:lock:safeLoadSeatMargin_%s"

	// LockTicketAvailabilityTokenBucket 列车购买令牌桶加载数据 Key
	LockTicketAvailabilityTokenBucket = "index12306TicketService:lock:ticketAvailabilityTokenBucket:%s"
)
