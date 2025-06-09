/*
 Navicat Premium Dump SQL

 Source Server         : HL
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:3306
 Source Schema         : todolist_db

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 09/06/2025 16:07:42
*/

-- 1. 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS todolist_db
  DEFAULT CHARACTER SET utf8mb4
  COLLATE utf8mb4_general_ci;

-- 2. 切换到数据库
USE todolist_db;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for todos
-- ----------------------------
DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NULL DEFAULT NULL,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `status` tinyint NULL DEFAULT 0,
  `due_date` datetime NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `todos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of todos
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'hhnhhn', '123@gmail.com', '$2a$10$obsKIFtBcu884Ub3ZfeykOT0I.j8GvJHN2ZYdvNG2nly8B8LqtuEW', '2025-06-03 17:28:57');
INSERT INTO `users` VALUES (2, 'htyyy', '12345678@126.com', '$2a$10$.UX56QVqd/LcJRWVYktNNenWO9Y/8zFZTva9w5ohzHfwb2D92qb8i', '2025-06-04 16:31:19');
INSERT INTO `users` VALUES (3, 'haonan', 'deadman@126.com', '$2a$10$ALn6yaBpMZH5diWXdmS7b.dR5Kna3Y6W4b/sK..830wIxov1RVX.u', '2025-06-04 17:51:37');
INSERT INTO `users` VALUES (4, 'kkkkyy', 'oldman@126.com', '$2a$10$KeADGhsmz7ghS6vyUa4qKOlmvkTbcMt5NfLjeVkJQl9BGjLZ.uZ3W', '2025-06-04 18:02:06');

SET FOREIGN_KEY_CHECKS = 1;
