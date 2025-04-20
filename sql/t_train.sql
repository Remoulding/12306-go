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

 Date: 10/03/2025 23:07:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_train
-- ----------------------------
DROP TABLE IF EXISTS `t_train`;
CREATE TABLE `t_train`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `train_number` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '列车车次',
  `train_type` int NULL DEFAULT NULL COMMENT '列车类型 0：高铁 1：动车 2：普通车',
  `train_tag` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '列车标签 0：复兴号 1：智能动车组 2：静音车厢 3：支持选铺',
  `train_brand` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '列车品牌 0：GC-高铁/城际 1：D-动车 2：Z-直达 3：T-特快 4：K-快速 5：其他 6：复兴号 7：智能动车组',
  `start_station` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '起始站',
  `end_station` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '终点站',
  `start_region` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '起始城市',
  `end_region` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '终点城市',
  `sale_time` datetime NULL DEFAULT NULL COMMENT '销售时间',
  `sale_status` int NULL DEFAULT NULL COMMENT '销售状态 0：可售 1：不可售 2：未知',
  `departure_time` datetime NULL DEFAULT NULL COMMENT '出发时间',
  `arrival_time` datetime NULL DEFAULT NULL COMMENT '到达时间',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `del_flag` tinyint(1) NULL DEFAULT NULL COMMENT '删除标识',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '列车表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_train
-- ----------------------------
INSERT INTO `t_train` VALUES (1, 'G35', 0, '0,1,2', '0,6', '北京南', '宁波', '北京', '宁波', '2023-05-15 14:30:00', 0, '2023-06-01 09:56:00', '2023-06-01 15:14:00', '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train` VALUES (2, 'G39', 0, '0,3', '1,6', '北京南', '杭州东', '北京', '杭州', '2023-05-15 14:30:00', 0, '2023-06-01 19:04:00', '2023-06-01 23:22:00', '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train` VALUES (3, 'D717', 1, '0', '6', '北京', '杭州', '北京', '杭州', '2023-05-15 14:30:00', 0, '2023-06-01 19:16:00', '2023-06-02 09:00:00', '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);
INSERT INTO `t_train` VALUES (4, 'Z284', 2, '0', '6', '北京', '杭州', '北京', '杭州', '2023-05-15 14:30:00', 0, '2023-06-01 18:56:00', '2023-06-02 10:57:00', '2023-06-01 20:45:00', '2023-06-01 20:45:00', 0);

SET FOREIGN_KEY_CHECKS = 1;
