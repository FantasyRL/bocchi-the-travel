SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- BTREE的用途
--     匹配模糊查询;全值匹配的查询(where);

CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `user_name` varchar(255) NOT NULL COMMENT '用户名',
    `email` varchar(255) NOT NULL COMMENT '邮箱',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `avatar` varchar(255) NOT NULL COMMENT '头像',
    `otp` varchar(255) DEFAULT NULL COMMENT 'OTP',
    `type2fa` varchar(255) NOT NULL DEFAULT 0 COMMENT '关闭:0;开启:1',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `user_name_password_idx` (`user_name`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

