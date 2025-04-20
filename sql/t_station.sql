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

 Date: 10/03/2025 23:06:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '车站表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_station
-- ----------------------------
INSERT INTO `t_station` VALUES (1, 'VNP', '北京南', 'beijingnan', 'BJP', '北京', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (2, 'JGK', '济南西', 'jinanxi', 'JNK', '济南', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (3, 'NKH', '南京南', 'nanjingnan', 'NJH', '南京', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (4, 'HGH', '杭州东', 'hangzhoudong', 'HZH', '杭州', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (5, 'NGH', '宁波', 'ningbo', 'NGH', '宁波', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (6, 'BJP', '北京', 'beijing', 'BJP', '北京', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (7, 'DZP', '德州', 'dezhou', 'DZP', '德州', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (8, 'NJH', '南京', 'nanjing', 'NJH', '南京', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (9, 'JXH', '嘉兴', 'jiaxing', 'JXH', '嘉兴', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (10, 'HNH', '海宁', 'haining', 'HNH', '海宁', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);
INSERT INTO `t_station` VALUES (11, 'HZH', '杭州', 'hangzhou', 'HZH', '杭州', '2023-06-01 20:54:00', '2023-06-01 20:54:00', 0);

SET FOREIGN_KEY_CHECKS = 1;
