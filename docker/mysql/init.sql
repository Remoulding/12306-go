/*
 Navicat Premium Data Transfer

 Source Server         : mysqlDB
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : localhost:3306
 Source Schema         : train_db

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 29/05/2025 16:11:49
*/

USE train_db;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_carriage
-- ----------------------------
DROP TABLE IF EXISTS `t_carriage`;
CREATE TABLE `t_carriage`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `train_id` bigint NULL DEFAULT NULL COMMENT '列车ID',
  `carriage_number` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '车厢号',
  `carriage_type` int NULL DEFAULT NULL COMMENT '车厢类型',
  `seat_count` int NULL DEFAULT NULL COMMENT '座位数',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_train_id`(`train_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '车厢表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_carriage
-- ----------------------------
INSERT INTO `t_carriage` VALUES (1, 1, '01', 0, 6, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (2, 1, '02', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (3, 1, '03', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (4, 1, '04', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (5, 1, '05', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (6, 1, '06', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (7, 1, '07', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (8, 1, '08', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (9, 1, '09', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (10, 1, '10', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (11, 1, '11', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (12, 1, '12', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (13, 1, '13', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (14, 1, '14', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (15, 1, '15', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (16, 1, '16', 0, 6, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (17, 2, '01', 0, 6, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (18, 2, '02', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (19, 2, '03', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (20, 2, '04', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (21, 2, '05', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (22, 2, '06', 1, 28, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (23, 2, '07', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (24, 2, '08', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (25, 2, '09', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (26, 2, '10', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (27, 2, '11', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (28, 2, '12', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (29, 2, '13', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (30, 2, '14', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (31, 2, '15', 2, 90, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (32, 2, '16', 0, 6, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);

-- ----------------------------
-- Table structure for t_passenger
-- ----------------------------
DROP TABLE IF EXISTS `t_passenger`;
CREATE TABLE `t_passenger`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户名',
  `real_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `id_type` int NULL DEFAULT NULL COMMENT '证件类型',
  `id_card` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '证件号码',
  `discount_type` int NULL DEFAULT NULL COMMENT '优惠类型',
  `phone` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '手机号',
  `create_date` datetime NULL DEFAULT NULL COMMENT '添加日期',
  `verify_status` int NULL DEFAULT NULL COMMENT '审核状态',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_id_card`(`id_card` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '乘车人表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_passenger
-- ----------------------------

-- ----------------------------
-- Table structure for t_region
-- ----------------------------
DROP TABLE IF EXISTS `t_region`;
CREATE TABLE `t_region`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '地区名称',
  `full_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '地区全名',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '地区编码',
  `initial` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '地区首字母',
  `spell` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '拼音',
  `popular_flag` tinyint(1) NULL DEFAULT NULL COMMENT '热门标识',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '地区表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_region
-- ----------------------------
INSERT INTO `t_region` VALUES (1, '北京', '北京市', 'BJP', 'B', 'beijing', 1, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_region` VALUES (2, '济南', '济南市', 'JNK', 'J', 'jinan', 1, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_region` VALUES (3, '南京', '南京市', 'NJH', 'N', 'nanjing', 1, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_region` VALUES (4, '杭州', '杭州市', 'HZH', 'H', 'hangzhou', 1, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_region` VALUES (5, '宁波', '宁波市', 'NGH', 'N', 'ningbo', 0, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_region` VALUES (6, '天津', '天津市', 'TJP', 'T', 'tianjing', 1, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_region` VALUES (7, '嘉兴', '嘉兴市', 'JXH', 'J', 'jiaxing', 0, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);
INSERT INTO `t_region` VALUES (8, '海宁', '海宁市', 'HNH', 'H', 'haining', 0, '2025-03-22 20:50:00', '2025-03-22 20:50:00', 0);

-- ----------------------------
-- Table structure for t_seat
-- ----------------------------
DROP TABLE IF EXISTS `t_seat`;
CREATE TABLE `t_seat`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `train_id` bigint NULL DEFAULT NULL COMMENT '列车ID',
  `carriage_number` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '车厢号',
  `seat_number` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '座位号',
  `seat_type` int NULL DEFAULT NULL COMMENT '座位类型',
  `start_station` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '起始站',
  `end_station` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '终点站',
  `price` int NULL DEFAULT NULL COMMENT '车票价格',
  `seat_status` int NULL DEFAULT NULL COMMENT '座位状态',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  `departure_date` date NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniq_train_seat`(`train_id` ASC, `carriage_number` ASC, `seat_number` ASC, `start_station` ASC, `end_station` ASC, `departure_date` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1685114673840531282 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '座位表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_seat
-- ----------------------------

-- ----------------------------
-- Table structure for t_station
-- ----------------------------
DROP TABLE IF EXISTS `t_station`;
CREATE TABLE `t_station`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '车站编号',
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '车站名称',
  `spell` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '拼音',
  `region` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '车站地区',
  `region_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '车站地区名称',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '车站表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_station
-- ----------------------------
INSERT INTO `t_station` VALUES (1, 'VNP', '北京南', 'beijingnan', 'BJP', '北京', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (2, 'JGK', '济南西', 'jinanxi', 'JNK', '济南', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (3, 'NKH', '南京南', 'nanjingnan', 'NJH', '南京', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (4, 'HGH', '杭州东', 'hangzhoudong', 'HZH', '杭州', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (5, 'NGH', '宁波', 'ningbo', 'NGH', '宁波', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (6, 'BJP', '北京', 'beijing', 'BJP', '北京', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (7, 'DZP', '德州', 'dezhou', 'DZP', '德州', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (8, 'NJH', '南京', 'nanjing', 'NJH', '南京', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (9, 'JXH', '嘉兴', 'jiaxing', 'JXH', '嘉兴', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (10, 'HNH', '海宁', 'haining', 'HNH', '海宁', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);
INSERT INTO `t_station` VALUES (11, 'HZH', '杭州', 'hangzhou', 'HZH', '杭州', '2025-03-22 20:54:00', '2025-03-22 20:54:00', 0);

-- ----------------------------
-- Table structure for t_ticket
-- ----------------------------
DROP TABLE IF EXISTS `t_ticket`;
CREATE TABLE `t_ticket`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户名',
  `carriage_number` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '车厢号',
  `seat_number` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '座位号',
  `passenger_id` bigint NULL DEFAULT NULL COMMENT '乘车人ID',
  `ticket_status` int NULL DEFAULT NULL COMMENT '车票状态',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  `departure_time` datetime NOT NULL COMMENT '出发时间，格式YYYY-MM-DD HH-MM-SS',
  `arrival_time` datetime NOT NULL COMMENT '出发时间，格式YYYY-MM-DD HH-MM-SS',
  `departure` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '出发站点',
  `arrival` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '到达站点',
  `train_number` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '列车车次',
  `train_id` bigint NULL DEFAULT NULL COMMENT '列车ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_passenger_id`(`passenger_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4421 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '车票表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_ticket
-- ----------------------------

-- ----------------------------
-- Table structure for t_train
-- ----------------------------
DROP TABLE IF EXISTS `t_train`;
CREATE TABLE `t_train`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `train_number` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '列车车次',
  `start_station` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '起始站',
  `end_station` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '终点站',
  `start_region` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '起始城市',
  `end_region` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '终点城市',
  `departure_time` time NOT NULL COMMENT '出发时间，仅保留时分秒',
  `arrival_time` time NOT NULL COMMENT '到达时间，仅保留时分秒',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  `cross_day` int NOT NULL DEFAULT 0 COMMENT '跨天数，表示到达时间相对出发时间的天数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '列车表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_train
-- ----------------------------
INSERT INTO `t_train` VALUES (1, 'G35', '北京南', '宁波', '北京', '宁波', '09:56:00', '15:14:00', '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);
INSERT INTO `t_train` VALUES (2, 'G39', '北京南', '杭州东', '北京', '杭州', '19:04:00', '23:22:00', '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);

-- ----------------------------
-- Table structure for t_train_station
-- ----------------------------
DROP TABLE IF EXISTS `t_train_station`;
CREATE TABLE `t_train_station`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `train_id` bigint NULL DEFAULT NULL COMMENT '车次ID',
  `station_id` bigint NULL DEFAULT NULL COMMENT '车站ID',
  `sequence` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '站点顺序',
  `departure` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '出发站点',
  `arrival` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '到达站点',
  `start_region` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '起始城市',
  `end_region` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '终点城市',
  `arrival_time` time NOT NULL COMMENT '到达时间，仅保留时分秒',
  `departure_time` time NOT NULL COMMENT '出发时间，仅保留时分秒',
  `stopover_time` int NULL DEFAULT NULL COMMENT '停留时间，单位分',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  `cross_day` int NOT NULL DEFAULT 0 COMMENT '跨天数，表示到达时间相对出发时间的天数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_train_id`(`train_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '列车站点表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_train_station
-- ----------------------------
INSERT INTO `t_train_station` VALUES (1, 1, 1, '01', '北京南', '济南西', '北京', '济南', '09:56:00', '09:56:00', NULL, '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);
INSERT INTO `t_train_station` VALUES (2, 1, 2, '02', '济南西', '南京南', '济南', '南京', '11:19:00', '11:21:00', 2, '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);
INSERT INTO `t_train_station` VALUES (3, 1, 3, '03', '南京南', '杭州东', '南京', '杭州', '13:20:00', '13:22:00', 2, '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);
INSERT INTO `t_train_station` VALUES (4, 1, 4, '04', '杭州东', '宁波', '杭州', '宁波', '14:26:00', '14:28:00', 2, '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);
INSERT INTO `t_train_station` VALUES (5, 1, 5, '05', '宁波', NULL, '宁波', NULL, '15:14:00', '15:14:00', NULL, '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);
INSERT INTO `t_train_station` VALUES (6, 2, 1, '01', '北京南', '南京南', '北京', '南京', '19:04:00', '19:04:00', NULL, '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);
INSERT INTO `t_train_station` VALUES (7, 2, 3, '02', '南京南', '杭州东', '南京', '杭州', '22:17:00', '22:20:00', 3, '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);
INSERT INTO `t_train_station` VALUES (8, 2, 4, '03', '杭州东', NULL, '杭州', NULL, '23:22:00', '23:22:00', NULL, '2025-03-22 20:45:00', '2025-03-22 20:45:00', 0, 0);

-- ----------------------------
-- Table structure for t_train_station_price
-- ----------------------------
DROP TABLE IF EXISTS `t_train_station_price`;
CREATE TABLE `t_train_station_price`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `train_id` bigint NULL DEFAULT NULL COMMENT '车次ID',
  `departure` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '出发站点',
  `arrival` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '到达站点',
  `seat_type` int NULL DEFAULT NULL COMMENT '座位类型',
  `price` int NULL DEFAULT NULL COMMENT '车票价格',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_train_id`(`train_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1677692017354547201 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '列车站点价格表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_train_station_price
-- ----------------------------
INSERT INTO `t_train_station_price` VALUES (1664877136352866304, 1, '北京南', '济南西', 0, 78200, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136424169472, 1, '北京南', '济南西', 1, 35700, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136453529600, 1, '北京南', '济南西', 2, 22300, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136474501120, 1, '北京南', '南京南', 0, 186400, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136487084032, 1, '北京南', '南京南', 1, 85200, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136499666944, 1, '北京南', '南京南', 2, 53300, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136516444160, 1, '北京南', '杭州东', 0, 231300, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136529027072, 1, '北京南', '杭州东', 1, 107700, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136541609984, 1, '北京南', '杭州东', 2, 67400, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136554192896, 1, '北京南', '宁波', 0, 253750, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136566775808, 1, '北京南', '宁波', 1, 119700, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136575164416, 1, '北京南', '宁波', 2, 74500, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136583553024, 1, '济南西', '南京南', 0, 116500, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136596135936, 1, '济南西', '南京南', 1, 53300, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136608718848, 1, '济南西', '南京南', 2, 33300, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136621301760, 1, '济南西', '杭州东', 0, 161400, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136633884672, 1, '济南西', '杭州东', 1, 75800, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136650661888, 1, '济南西', '杭州东', 2, 47400, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136663244800, 1, '济南西', '宁波', 0, 183850, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136675827712, 1, '济南西', '宁波', 1, 87800, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136684216320, 1, '济南西', '宁波', 2, 54500, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136696799232, 1, '南京南', '杭州东', 0, 44900, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136709382144, 1, '南京南', '杭州东', 1, 22500, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136721965056, 1, '南京南', '杭州东', 2, 14100, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136738742272, 1, '南京南', '宁波', 0, 67350, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136747130880, 1, '南京南', '宁波', 1, 34500, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136759713792, 1, '南京南', '宁波', 2, 21200, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136772296704, 1, '杭州东', '宁波', 0, 22450, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136784879616, 1, '杭州东', '宁波', 1, 12000, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136797462528, 1, '杭州东', '宁波', 2, 7100, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730273529856, 2, '北京南', '南京南', 0, 186400, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730323861504, 2, '北京南', '南京南', 1, 85200, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730353221632, 2, '北京南', '南京南', 2, 53300, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730386776064, 2, '北京南', '杭州东', 0, 231300, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730403553280, 2, '北京南', '杭州东', 1, 106400, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730420330496, 2, '北京南', '杭州东', 2, 66300, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730432913408, 2, '南京南', '杭州东', 0, 44900, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730449690624, 2, '南京南', '杭州东', 1, 21200, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730466467840, 2, '南京南', '杭州东', 2, 11700, '2025-03-22 14:10:16', '2025-03-22 14:10:16', 0);

-- ----------------------------
-- Table structure for t_train_station_relation
-- ----------------------------
DROP TABLE IF EXISTS `t_train_station_relation`;
CREATE TABLE `t_train_station_relation`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `train_id` bigint NULL DEFAULT NULL COMMENT '车次ID',
  `departure` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '出发站点',
  `arrival` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '到达站点',
  `start_region` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '起始城市名称',
  `end_region` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '终点城市名称',
  `departure_flag` tinyint(1) NULL DEFAULT NULL COMMENT '始发标识',
  `arrival_flag` tinyint(1) NULL DEFAULT NULL COMMENT '终点标识',
  `departure_time` time NOT NULL COMMENT '出发时间，仅保留时分秒',
  `arrival_time` time NOT NULL COMMENT '到达时间，仅保留时分秒',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  `cross_day` int NOT NULL DEFAULT 0 COMMENT '跨天数，表示到达时间相对出发时间的天数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_train_id`(`train_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1677689610742865921 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '列车站点关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_train_station_relation
-- ----------------------------
INSERT INTO `t_train_station_relation` VALUES (1665025584123056128, 1, '北京南', '济南西', '北京', '济南', 1, 0, '09:56:00', '11:19:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584177582080, 1, '北京南', '南京南', '北京', '南京', 1, 0, '09:56:00', '13:20:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584206942208, 1, '北京南', '杭州东', '北京', '杭州', 1, 0, '09:56:00', '14:26:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584236302336, 1, '北京南', '宁波', '北京', '宁波', 1, 1, '09:56:00', '15:14:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584261468160, 1, '济南西', '南京南', '济南', '南京', 0, 0, '11:19:00', '15:14:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584282439680, 1, '济南西', '杭州东', '济南', '杭州', 0, 0, '11:19:00', '14:26:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584307605504, 1, '济南西', '宁波', '济南', '宁波', 0, 1, '11:19:00', '15:14:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584332771328, 1, '南京南', '杭州东', '南京', '杭州', 0, 0, '13:20:00', '14:26:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584357937152, 1, '南京南', '宁波', '南京', '宁波', 0, 1, '13:20:00', '15:14:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584378908672, 1, '杭州东', '宁波', '杭州', '宁波', 0, 1, '14:26:00', '15:14:00', '2025-03-22 00:00:09', '2025-03-22 00:00:09', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1677241394371317760, 2, '北京南', '南京南', '北京', '南京', 1, 0, '19:04:00', '22:17:00', '2025-03-22 17:01:25', '2025-03-22 17:01:25', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1677241394434232320, 2, '北京南', '杭州东', '北京', '杭州', 1, 1, '19:04:00', '23:22:00', '2025-03-22 17:01:25', '2025-03-22 17:01:25', 0, 0);
INSERT INTO `t_train_station_relation` VALUES (1677241394463592448, 2, '南京南', '杭州东', '南京', '杭州', 0, 1, '22:20:00', '23:22:00', '2025-03-22 17:01:25', '2025-03-22 17:01:25', 0, 0);

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '密码',
  `real_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `region` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '0' COMMENT '国家/地区',
  `id_type` int NULL DEFAULT NULL COMMENT '证件类型',
  `id_card` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '证件号',
  `phone` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '手机号',
  `telephone` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '固定电话',
  `mail` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '邮箱',
  `user_type` int NULL DEFAULT NULL COMMENT '旅客类型',
  `verify_status` int NULL DEFAULT NULL COMMENT '审核状态',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '邮编',
  `address` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '地址',
  `deletion_time` bigint NULL DEFAULT 0 COMMENT '注销时间戳',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username` ASC, `deletion_time` ASC) USING BTREE,
  UNIQUE INDEX `idx_phone`(`phone` ASC, `deletion_time` ASC) USING BTREE,
  UNIQUE INDEX `idx_mail`(`mail` ASC, `deletion_time` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;

-- 创建canal用户
CREATE USER 'canal'@'%' IDENTIFIED BY 'canal';

-- 授予复制权限
GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'canal'@'%';

-- 授予数据库读取权限
GRANT SELECT ON train_db.* TO 'canal'@'%';

FLUSH PRIVILEGES;