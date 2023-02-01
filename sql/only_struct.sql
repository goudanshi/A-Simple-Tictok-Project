/*
 Navicat Premium Data Transfer

 Source Server         : docker-local
 Source Server Type    : MySQL
 Source Server Version : 50647
 Source Host           : 127.0.0.1:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50647
 File Encoding         : 65001

 Date: 01/02/2023 18:31:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `title` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
                            `user_id` bigint(20) NULL DEFAULT NULL,
                            `video_id` bigint(20) NULL DEFAULT NULL,
                            `create_at` datetime(3) NULL DEFAULT NULL,
                            `updated_at` datetime(3) NULL DEFAULT NULL,
                            `delete_at` datetime(3) NULL DEFAULT NULL,
                            PRIMARY KEY (`id`) USING BTREE,
                            INDEX `fk_video_comments`(`video_id`) USING BTREE,
                            INDEX `fk_user_comments`(`user_id`) USING BTREE,
                            CONSTRAINT `fk_user_comments` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                            CONSTRAINT `fk_video_comments` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for like
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like`  (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `user_id` bigint(20) NULL DEFAULT NULL,
                         `video_id` bigint(20) NULL DEFAULT NULL,
                         `create_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `delete_at` datetime(3) NULL DEFAULT NULL,
                         PRIMARY KEY (`id`) USING BTREE,
                         INDEX `fk_video_likes`(`video_id`) USING BTREE,
                         INDEX `fk_user_likes`(`user_id`) USING BTREE,
                         CONSTRAINT `fk_user_likes` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                         CONSTRAINT `fk_video_likes` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for relationship
-- ----------------------------
DROP TABLE IF EXISTS `relationship`;
CREATE TABLE `relationship`  (
                                 `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                 `user_id` bigint(20) NULL DEFAULT NULL,
                                 `subscribe_id` bigint(20) NULL DEFAULT NULL,
                                 `create_at` datetime(3) NULL DEFAULT NULL,
                                 `update_at` datetime(3) NULL DEFAULT NULL,
                                 `delete_at` datetime(3) NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`) USING BTREE,
                                 INDEX `fk_user_subscribes`(`subscribe_id`) USING BTREE,
                                 INDEX `fk_user_users`(`user_id`) USING BTREE,
                                 CONSTRAINT `fk_user_subscribes` FOREIGN KEY (`subscribe_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                                 CONSTRAINT `fk_user_users` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `username` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
                         `password` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
                         `user_count` bigint(20) NULL DEFAULT 0 COMMENT '女神数',
                         `subscribe_count` bigint(20) NULL DEFAULT 0 COMMENT '舔狗数',
                         `create_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `delete_at` datetime(3) NULL DEFAULT NULL,
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
                          `id` bigint(20) NOT NULL AUTO_INCREMENT,
                          `title` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
                          `play_url` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
                          `cover_url` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
                          `user_id` bigint(20) NULL DEFAULT NULL,
                          `like_count` bigint(20) NULL DEFAULT 0,
                          `comment_count` bigint(20) NULL DEFAULT 0,
                          `create_at` datetime(3) NULL DEFAULT NULL,
                          `updated_at` datetime(3) NULL DEFAULT NULL,
                          `delete_at` datetime(3) NULL DEFAULT NULL,
                          PRIMARY KEY (`id`) USING BTREE,
                          INDEX `fk_user_videos`(`user_id`) USING BTREE,
                          CONSTRAINT `fk_user_videos` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
