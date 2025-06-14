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