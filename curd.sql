 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50716
 Source Host           : 127.0.0.1:3306
 Source Schema         : curd

 Target Server Type    : MySQL
 Target Server Version : 50716
 File Encoding         : 65001

 Date: 14/12/2021 11:36:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_age` int(11) DEFAULT NULL,
  `user_pwd` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_createtime` datetime(0) DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 57763 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'gzp', 24, 'gzp', '2021-07-23 15:55:08');
INSERT INTO `users` VALUES (2, 'lxc', 2, 'lxc', '2021-07-15 16:26:59');
INSERT INTO `users` VALUES (3, 'gzp2', 30, 'c208130a40bfc075e83cabfdfec24276', '0000-00-00 00:00:00');
INSERT INTO `users` VALUES (20, 'gzp3', 301, '8374876ed8f1b95f1725d01d4f32c8d8', '0000-00-00 00:00:00');
INSERT INTO `users` VALUES (21, 'gzp888', 301, '8374876ed8f1b95f1725d01d4f32c8d8', '0000-00-00 00:00:00');
INSERT INTO `users` VALUES (22, 'gzp13', 301, '8374876ed8f1b95f1725d01d4f32c8d8', '0000-00-00 00:00:00');