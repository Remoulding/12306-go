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

 Date: 10/03/2025 23:06:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '地区表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_region
-- ----------------------------
INSERT INTO `t_region` VALUES (1, '北京', '北京市', 'BJP', 'B', 'beijing', 1, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_region` VALUES (2, '济南', '济南市', 'JNK', 'J', 'jinan', 1, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_region` VALUES (3, '南京', '南京市', 'NJH', 'N', 'nanjing', 1, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_region` VALUES (4, '杭州', '杭州市', 'HZH', 'H', 'hangzhou', 1, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_region` VALUES (5, '宁波', '宁波市', 'NGH', 'N', 'ningbo', 0, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_region` VALUES (6, '天津', '天津市', 'TJP', 'T', 'tianjing', 1, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_region` VALUES (7, '嘉兴', '嘉兴市', 'JXH', 'J', 'jiaxing', 0, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);
INSERT INTO `t_region` VALUES (8, '海宁', '海宁市', 'HNH', 'H', 'haining', 0, '2023-06-01 20:50:00', '2023-06-01 20:50:00', 0);

SET FOREIGN_KEY_CHECKS = 1;
