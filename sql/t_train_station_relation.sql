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

 Date: 10/03/2025 23:07:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
  `departure_time` datetime NULL DEFAULT NULL COMMENT '出发时间',
  `arrival_time` datetime NULL DEFAULT NULL COMMENT '到达时间',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_train_id`(`train_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1677689610742865920 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '列车站点关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_train_station_relation
-- ----------------------------
INSERT INTO `t_train_station_relation` VALUES (1665025584123056128, 1, '北京南', '济南西', '北京', '济南', 1, 0, '2023-06-01 09:56:00', '2023-06-01 11:19:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584177582080, 1, '北京南', '南京南', '北京', '南京', 1, 0, '2023-06-01 09:56:00', '2023-06-01 13:20:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584206942208, 1, '北京南', '杭州东', '北京', '杭州', 1, 0, '2023-06-01 09:56:00', '2023-06-01 14:26:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584236302336, 1, '北京南', '宁波', '北京', '宁波', 1, 1, '2023-06-01 09:56:00', '2023-06-01 15:14:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584261468160, 1, '济南西', '南京南', '济南', '南京', 0, 0, '2023-06-01 11:19:00', '2023-06-01 15:14:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584282439680, 1, '济南西', '杭州东', '济南', '杭州', 0, 0, '2023-06-01 11:19:00', '2023-06-01 14:26:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584307605504, 1, '济南西', '宁波', '济南', '宁波', 0, 1, '2023-06-01 11:19:00', '2023-06-01 15:14:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584332771328, 1, '南京南', '杭州东', '南京', '杭州', 0, 0, '2023-06-01 13:20:00', '2023-06-01 14:26:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584357937152, 1, '南京南', '宁波', '南京', '宁波', 0, 1, '2023-06-01 13:20:00', '2023-06-01 15:14:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1665025584378908672, 1, '杭州东', '宁波', '杭州', '宁波', 0, 1, '2023-06-01 14:26:00', '2023-06-01 15:14:00', '2023-06-04 00:00:09', '2023-06-04 00:00:09', 0);
INSERT INTO `t_train_station_relation` VALUES (1677241394371317760, 2, '北京南', '南京南', '北京', '南京', 1, 0, '2023-06-01 19:04:00', '2023-06-01 22:17:00', '2023-07-07 17:01:25', '2023-07-07 17:01:25', 0);
INSERT INTO `t_train_station_relation` VALUES (1677241394434232320, 2, '北京南', '杭州东', '北京', '杭州', 1, 1, '2023-06-01 19:04:00', '2023-06-01 23:22:00', '2023-07-07 17:01:25', '2023-07-07 17:01:25', 0);
INSERT INTO `t_train_station_relation` VALUES (1677241394463592448, 2, '南京南', '杭州东', '南京', '杭州', 0, 1, '2023-06-01 22:20:00', '2023-06-01 23:22:00', '2023-07-07 17:01:25', '2023-07-07 17:01:25', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889368797184, 3, '北京', '德州', '北京', '德州', 1, 0, '2023-06-01 19:16:00', '2023-06-01 22:19:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889440100352, 3, '北京', '南京', '北京', '南京', 1, 0, '2023-06-01 19:16:00', '2023-06-02 04:46:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889461071872, 3, '北京', '嘉兴', '北京', '嘉兴', 1, 0, '2023-06-01 19:16:00', '2023-06-02 07:56:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889473654784, 3, '北京', '海宁', '北京', '海宁', 1, 0, '2023-06-01 19:16:00', '2023-06-02 08:14:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889503014912, 3, '北京', '杭州', '北京', '杭州', 1, 1, '2023-06-01 19:16:00', '2023-06-02 09:00:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889515597824, 3, '德州', '南京', '德州', '南京', 0, 0, '2023-06-01 22:21:00', '2023-06-02 04:46:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889528180736, 3, '德州', '嘉兴', '德州', '嘉兴', 0, 0, '2023-06-01 22:21:00', '2023-06-02 07:56:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889540763648, 3, '德州', '海宁', '德州', '海宁', 0, 0, '2023-06-01 22:21:00', '2023-06-02 08:14:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889557540864, 3, '德州', '杭州', '德州', '杭州', 0, 1, '2023-06-01 22:21:00', '2023-06-02 09:00:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889570123776, 3, '南京', '嘉兴', '南京', '嘉兴', 0, 0, '2023-06-02 04:52:00', '2023-06-02 07:56:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889582706688, 3, '南京', '海宁', '南京', '海宁', 0, 0, '2023-06-02 04:52:00', '2023-06-02 08:14:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889595289600, 3, '南京', '杭州', '南京', '杭州', 0, 1, '2023-06-02 04:52:00', '2023-06-02 09:00:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889607872512, 3, '嘉兴', '海宁', '嘉兴', '海宁', 0, 0, '2023-06-02 07:58:00', '2023-06-02 08:14:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889620455424, 3, '嘉兴', '杭州', '嘉兴', '杭州', 0, 1, '2023-06-02 07:58:00', '2023-06-02 09:00:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677579889633038336, 3, '海宁', '杭州', '海宁', '杭州', 0, 1, '2023-06-02 08:16:00', '2023-06-02 09:00:00', '2023-07-08 15:26:28', '2023-07-08 15:26:28', 0);
INSERT INTO `t_train_station_relation` VALUES (1677689610742865920, 4, '北京南', '杭州东', '北京', '杭州', 1, 1, '2023-06-01 18:56:00', '2023-06-01 18:56:00', '2023-07-08 22:42:28', '2023-07-08 22:42:28', 0);

SET FOREIGN_KEY_CHECKS = 1;
