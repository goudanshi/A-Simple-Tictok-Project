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

 Date: 01/02/2023 18:31:59
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
-- Records of like
-- ----------------------------
INSERT INTO `like` VALUES (1, 2, 1, '2023-01-31 22:25:31.377', '2023-01-31 22:25:31.377', NULL);
INSERT INTO `like` VALUES (2, 3, 1, '2023-01-31 22:26:02.011', '2023-01-31 22:26:02.011', NULL);
INSERT INTO `like` VALUES (3, 3, 2, '2023-01-31 22:26:09.321', '2023-01-31 22:26:09.321', NULL);
INSERT INTO `like` VALUES (4, 3, 3, '2023-01-31 22:26:12.866', '2023-01-31 22:26:12.866', NULL);
INSERT INTO `like` VALUES (13, 3, 4, '2023-02-01 15:08:55.054', '2023-02-01 15:08:55.054', NULL);
INSERT INTO `like` VALUES (14, 3, 5, '2023-02-01 15:22:26.873', '2023-02-01 15:22:26.873', '2023-02-01 15:37:29.770');

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
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'A-Simple-Tictok-Project', 'acc0aa297d5bbd540006807c57c0affb', 0, 0, '2023-01-31 11:59:59.440', '2023-01-31 11:59:59.440', NULL);
INSERT INTO `user` VALUES (2, 'tank', '9fe720581dcde5eab041673b9883da9f', 0, 0, '2023-01-31 12:00:31.626', '2023-01-31 12:00:31.626', NULL);
INSERT INTO `user` VALUES (3, 'alex', 'b75bd008d5fecb1f50cf026532e8ae67', 0, 0, '2023-01-31 12:00:51.289', '2023-01-31 12:00:51.289', NULL);
INSERT INTO `user` VALUES (4, 'egon', 'd2e3da2e9dd5f6d7ab0d90521ea910d2', 0, 0, '2023-01-31 12:01:02.862', '2023-01-31 12:01:02.862', NULL);
INSERT INTO `user` VALUES (5, 'eric', '4131f403beab0f4fa9e654b2ffa4f769', 0, 0, '2023-01-31 15:04:26.936', '2023-01-31 15:04:26.936', NULL);

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

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, '??A-Simple-Tictok-Project?????', 'http://127.0.0.1:8080/video/video.mp4', 'http://127.0.0.1:8080/video/cover.jpg', 1, 2, 0, '2023-01-31 21:09:48.044', '2023-01-31 21:09:48.044', NULL);
INSERT INTO `video` VALUES (2, '??A-Simple-Tictok-Project?????', 'http://127.0.0.1:8080/video/video.mp4', 'http://127.0.0.1:8080/video/cover.jpg', 1, 1, 0, '2023-01-31 21:10:01.906', '2023-01-31 21:10:01.906', NULL);
INSERT INTO `video` VALUES (3, '??A-Simple-Tictok-Project?????', 'http://127.0.0.1:8080/video/video.mp4', 'http://127.0.0.1:8080/video/cover.jpg', 1, 1, 0, '2023-01-31 21:10:09.454', '2023-01-31 21:10:09.454', NULL);
INSERT INTO `video` VALUES (4, '??A-Simple-Tictok-Project?????', 'http://127.0.0.1:8080/video/video.mp4', 'http://127.0.0.1:8080/video/cover.jpg', 1, 1, 0, '2023-01-31 21:12:50.675', '2023-01-31 21:12:50.675', NULL);
INSERT INTO `video` VALUES (5, '??A-Simple-Tictok-Project?????', 'http://127.0.0.1:8080/video/video.mp4', 'http://127.0.0.1:8080/video/cover.jpg', 1, 0, 0, '2023-01-31 21:12:58.220', '2023-01-31 21:12:58.220', NULL);
INSERT INTO `video` VALUES (6, '??A-Simple-Tictok-Project?????', 'http://127.0.0.1:8080/video/video.mp4', 'http://127.0.0.1:8080/video/cover.jpg', 1, 0, 0, '2023-01-31 21:13:06.677', '2023-01-31 21:13:06.677', NULL);

SET FOREIGN_KEY_CHECKS = 1;
