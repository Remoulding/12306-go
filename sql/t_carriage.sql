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

 Date: 10/03/2025 23:06:17
*/

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
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '车厢表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_carriage
-- ----------------------------
INSERT INTO `t_carriage` VALUES (1, 1, '01', 0, 5, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (2, 1, '02', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (3, 1, '03', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (4, 1, '04', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (5, 1, '05', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (6, 1, '06', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (7, 1, '07', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (8, 1, '08', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (9, 1, '09', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (10, 1, '10', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (11, 1, '11', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (12, 1, '12', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (13, 1, '13', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (14, 1, '14', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (15, 1, '15', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (16, 1, '16', 0, 5, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (17, 2, '01', 0, 5, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (18, 2, '02', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (19, 2, '03', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (20, 2, '04', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (21, 2, '05', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (22, 2, '06', 1, 28, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (23, 2, '07', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (24, 2, '08', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (25, 2, '09', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (26, 2, '10', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (27, 2, '11', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (28, 2, '12', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (29, 2, '13', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (30, 2, '14', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (31, 2, '15', 2, 90, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (32, 2, '16', 0, 5, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_carriage` VALUES (33, 3, '01', 3, 24, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (34, 3, '02', 3, 24, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (35, 3, '03', 3, 24, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (36, 3, '04', 3, 24, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (37, 3, '05', 4, 32, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (38, 3, '06', 4, 32, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (39, 3, '07', 4, 32, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (40, 3, '08', 4, 32, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (41, 3, '09', 4, 32, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (42, 3, '10', 4, 32, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (43, 3, '11', 5, 36, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (44, 3, '12', 5, 36, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (45, 3, '13', 5, 36, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (46, 3, '14', 5, 36, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (47, 3, '15', 5, 36, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (48, 3, '16', 5, 36, '2023-07-08 13:17:02', '2023-07-08 13:17:02', 0);
INSERT INTO `t_carriage` VALUES (49, 4, '01', 6, 24, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (50, 4, '02', 6, 24, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (51, 4, '03', 6, 24, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (52, 4, '04', 6, 24, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (53, 4, '05', 7, 32, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (54, 4, '06', 7, 32, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (55, 4, '07', 7, 32, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (56, 4, '08', 7, 32, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (57, 4, '09', 7, 32, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (58, 4, '10', 7, 32, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (59, 4, '11', 8, 36, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (60, 4, '12', 8, 36, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (61, 4, '13', 8, 36, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (62, 4, '14', 8, 36, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (63, 4, '15', 8, 36, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);
INSERT INTO `t_carriage` VALUES (64, 4, '16', 8, 36, '2023-07-08 22:41:14', '2023-07-08 22:41:14', 0);

SET FOREIGN_KEY_CHECKS = 1;
