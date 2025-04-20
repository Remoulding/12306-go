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

 Date: 10/03/2025 23:07:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE = InnoDB AUTO_INCREMENT = 1677692017354547200 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '列车站点价格表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_train_station_price
-- ----------------------------
INSERT INTO `t_train_station_price` VALUES (1664877136352866304, 1, '北京南', '济南西', 0, 78200, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136424169472, 1, '北京南', '济南西', 1, 35700, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136453529600, 1, '北京南', '济南西', 2, 22300, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136474501120, 1, '北京南', '南京南', 0, 186400, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136487084032, 1, '北京南', '南京南', 1, 85200, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136499666944, 1, '北京南', '南京南', 2, 53300, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136516444160, 1, '北京南', '杭州东', 0, 231300, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136529027072, 1, '北京南', '杭州东', 1, 107700, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136541609984, 1, '北京南', '杭州东', 2, 67400, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136554192896, 1, '北京南', '宁波', 0, 253750, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136566775808, 1, '北京南', '宁波', 1, 119700, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136575164416, 1, '北京南', '宁波', 2, 74500, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136583553024, 1, '济南西', '南京南', 0, 116500, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136596135936, 1, '济南西', '南京南', 1, 53300, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136608718848, 1, '济南西', '南京南', 2, 33300, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136621301760, 1, '济南西', '杭州东', 0, 161400, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136633884672, 1, '济南西', '杭州东', 1, 75800, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136650661888, 1, '济南西', '杭州东', 2, 47400, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136663244800, 1, '济南西', '宁波', 0, 183850, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136675827712, 1, '济南西', '宁波', 1, 87800, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136684216320, 1, '济南西', '宁波', 2, 54500, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136696799232, 1, '南京南', '杭州东', 0, 44900, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136709382144, 1, '南京南', '杭州东', 1, 22500, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136721965056, 1, '南京南', '杭州东', 2, 14100, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136738742272, 1, '南京南', '宁波', 0, 67350, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136747130880, 1, '南京南', '宁波', 1, 34500, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136759713792, 1, '南京南', '宁波', 2, 21200, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136772296704, 1, '杭州东', '宁波', 0, 22450, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136784879616, 1, '杭州东', '宁波', 1, 12000, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1664877136797462528, 1, '杭州东', '宁波', 2, 7100, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730273529856, 2, '北京南', '南京南', 0, 186400, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730323861504, 2, '北京南', '南京南', 1, 85200, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730353221632, 2, '北京南', '南京南', 2, 53300, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730386776064, 2, '北京南', '杭州东', 0, 231300, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730403553280, 2, '北京南', '杭州东', 1, 106400, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730420330496, 2, '北京南', '杭州东', 2, 66300, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730432913408, 2, '南京南', '杭州东', 0, 44900, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730449690624, 2, '南京南', '杭州东', 1, 21200, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677231730466467840, 2, '南京南', '杭州东', 2, 11700, '2023-06-03 14:10:16', '2023-06-03 14:10:16', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290215301120, 3, '北京', '德州', 3, 7500, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290257244160, 3, '北京', '德州', 4, 15300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290286604288, 3, '北京', '德州', 5, 13000, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290307575808, 3, '北京', '德州', 13, 7500, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290328547328, 3, '北京', '南京', 3, 23300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290349518848, 3, '北京', '南京', 4, 44300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290362101760, 3, '北京', '南京', 5, 37700, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290374684672, 3, '北京', '南京', 13, 23300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290391461888, 3, '北京', '嘉兴', 3, 30900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290404044800, 3, '北京', '嘉兴', 4, 58700, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290416627712, 3, '北京', '嘉兴', 5, 49900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290429210624, 3, '北京', '嘉兴', 13, 30900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290441793536, 3, '北京', '海宁', 3, 31500, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290458570752, 3, '北京', '海宁', 4, 59700, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290483736576, 3, '北京', '海宁', 5, 50800, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290496319488, 3, '北京', '海宁', 13, 31500, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290508902400, 3, '北京', '杭州', 3, 32800, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290521485312, 3, '北京', '杭州', 4, 6220, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290529873920, 3, '北京', '杭州', 5, 52900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290542456832, 3, '北京', '杭州', 13, 32800, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290555039744, 3, '德州', '南京', 3, 15700, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290571816960, 3, '德州', '南京', 4, 29900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290584399872, 3, '德州', '南京', 5, 25400, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290596982784, 3, '德州', '南京', 13, 15700, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290609565696, 3, '德州', '嘉兴', 3, 23300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290622148608, 3, '德州', '嘉兴', 4, 44300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290634731520, 3, '德州', '嘉兴', 5, 37700, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290647314432, 3, '德州', '嘉兴', 13, 23300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290664091648, 3, '德州', '海宁', 3, 23900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290680868864, 3, '德州', '海宁', 4, 45400, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290689257472, 3, '德州', '海宁', 5, 38600, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290701840384, 3, '德州', '海宁', 13, 23900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290714423296, 3, '德州', '杭州', 3, 25200, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290727006208, 3, '德州', '杭州', 4, 47800, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290735394816, 3, '德州', '杭州', 5, 40700, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290747977728, 3, '德州', '杭州', 13, 25200, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290756366336, 3, '南京', '嘉兴', 3, 7600, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290768949248, 3, '南京', '嘉兴', 4, 15300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290785726464, 3, '南京', '嘉兴', 5, 13000, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290794115072, 3, '南京', '嘉兴', 13, 7600, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290806697984, 3, '南京', '海宁', 3, 8100, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290815086592, 3, '南京', '海宁', 4, 15500, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290827669504, 3, '南京', '海宁', 5, 13200, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290840252416, 3, '南京', '海宁', 13, 8100, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290848641024, 3, '南京', '杭州', 3, 9400, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290861223936, 3, '南京', '杭州', 4, 17900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290869612544, 3, '南京', '杭州', 5, 15300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290882195456, 3, '南京', '杭州', 13, 9400, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290903166976, 3, '嘉兴', '海宁', 3, 600, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290915749888, 3, '嘉兴', '海宁', 4, 15300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290928332800, 3, '嘉兴', '海宁', 5, 13000, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290940915712, 3, '嘉兴', '海宁', 13, 600, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290953498624, 3, '嘉兴', '杭州', 3, 1900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290970275840, 3, '嘉兴', '杭州', 4, 15300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290987053056, 3, '嘉兴', '杭州', 5, 13000, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591290999635968, 3, '嘉兴', '杭州', 13, 1900, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591291008024576, 3, '海宁', '杭州', 3, 1300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591291020607488, 3, '海宁', '杭州', 4, 15300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591291037384704, 3, '海宁', '杭州', 5, 13000, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677591291049967616, 3, '海宁', '杭州', 13, 1300, '2023-07-08 16:11:46', '2023-07-08 16:11:46', 0);
INSERT INTO `t_train_station_price` VALUES (1677692017258078208, 4, '北京南', '杭州东', 6, 32800, '2023-07-08 22:52:02', '2023-07-08 22:52:02', 0);
INSERT INTO `t_train_station_price` VALUES (1677692017308409856, 4, '北京南', '杭州东', 7, 6220, '2023-07-08 22:52:02', '2023-07-08 22:52:02', 0);
INSERT INTO `t_train_station_price` VALUES (1677692017337769984, 4, '北京南', '杭州东', 8, 52900, '2023-07-08 22:52:02', '2023-07-08 22:52:02', 0);
INSERT INTO `t_train_station_price` VALUES (1677692017354547200, 4, '北京南', '杭州东', 13, 32800, '2023-07-08 22:52:02', '2023-07-08 22:52:02', 0);

SET FOREIGN_KEY_CHECKS = 1;
