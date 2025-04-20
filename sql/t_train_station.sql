/*
 Navicat Premium Data Transfer

 Source Server         : mysqlDB
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : localhost:3306
 Source Schema         : 12306

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 10/03/2025 23:07:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
  `arrival_time` datetime NULL DEFAULT NULL COMMENT '到站时间',
  `departure_time` datetime NULL DEFAULT NULL COMMENT '出站时间',
  `stopover_time` int NULL DEFAULT NULL COMMENT '停留时间，单位分',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_train_id`(`train_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '列车站点表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_train_station
-- ----------------------------
INSERT INTO `t_train_station` VALUES (1, 1, 1, '01', '北京南', '济南西', '北京', '济南', '2023-06-01 09:56:00', '2023-06-01 09:56:00', NULL, '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train_station` VALUES (2, 1, 2, '02', '济南西', '南京南', '济南', '南京', '2023-06-01 11:19:00', '2023-06-01 11:21:00', 2, '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train_station` VALUES (3, 1, 3, '03', '南京南', '杭州东', '南京', '杭州', '2023-06-01 13:20:00', '2023-06-01 13:22:00', 2, '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train_station` VALUES (4, 1, 4, '04', '杭州东', '宁波', '杭州', '宁波', '2023-06-01 14:26:00', '2023-06-01 14:28:00', 2, '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train_station` VALUES (5, 1, 5, '05', '宁波', NULL, '宁波', NULL, '2023-06-01 15:14:00', '2023-06-01 15:14:00', NULL, '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train_station` VALUES (6, 2, 1, '01', '北京南', '南京南', '北京', '南京', '2023-06-01 19:04:00', '2023-06-01 19:04:00', NULL, '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train_station` VALUES (7, 2, 3, '02', '南京南', '杭州东', '南京', '杭州', '2023-06-01 22:17:00', '2023-06-01 22:20:00', 3, '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train_station` VALUES (8, 2, 4, '03', '杭州东', NULL, '杭州', NULL, '2023-06-01 23:22:00', '2023-06-01 23:22:00', NULL, '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train_station` VALUES (9, 3, 6, '01', '北京', '德州', '北京', '德州', '2023-06-01 19:16:00', '2023-06-01 19:16:00', NULL, '2023-07-08 14:05:05', '2023-07-08 14:05:05', 0);
INSERT INTO `t_train_station` VALUES (10, 3, 7, '02', '德州', '南京', '德州', '南京', '2023-06-01 22:19:00', '2023-06-01 22:21:00', 2, '2023-07-08 14:05:05', '2023-07-08 14:05:05', 0);
INSERT INTO `t_train_station` VALUES (11, 3, 8, '03', '南京', '嘉兴', '南京', '嘉兴', '2023-06-02 04:46:00', '2023-06-02 04:52:00', 6, '2023-07-08 14:05:05', '2023-07-08 14:05:05', 0);
INSERT INTO `t_train_station` VALUES (12, 3, 9, '04', '嘉兴', '海宁', '嘉兴', '海宁', '2023-06-02 07:56:00', '2023-06-02 07:58:00', 2, '2023-07-08 14:05:05', '2023-07-08 14:05:05', 0);
INSERT INTO `t_train_station` VALUES (13, 3, 10, '05', '海宁', '杭州', '海宁', '杭州', '2023-06-02 08:14:00', '2023-06-02 08:16:00', 2, '2023-07-08 14:05:05', '2023-07-08 14:05:05', 0);
INSERT INTO `t_train_station` VALUES (14, 3, 11, '06', '杭州', NULL, '杭州', NULL, '2023-06-02 09:00:00', '2023-06-02 09:00:00', NULL, '2023-07-08 14:05:05', '2023-07-08 14:05:05', 0);
INSERT INTO `t_train_station` VALUES (15, 4, 1, '01', '北京南', '杭州东', '北京', '杭州', '2023-06-01 18:56:00', '2023-06-01 18:56:00', NULL, '2023-07-08 14:05:05', '2023-07-08 14:05:05', 0);
INSERT INTO `t_train_station` VALUES (16, 4, 4, '02', '杭州东', '', '北京', '', '2023-06-01 18:56:00', '2023-06-02 10:57:00', NULL, '2023-07-08 14:05:05', '2023-07-08 14:05:05', 0);

SET FOREIGN_KEY_CHECKS = 1;
