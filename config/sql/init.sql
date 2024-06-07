SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- BTREE的用途
--     匹配模糊查询;全值匹配的查询(where);

CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `user_name` varchar(255) NOT NULL COMMENT '用户名',
    `email` varchar(255) NOT NULL COMMENT '邮箱',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `avatar` varchar(255) NULL DEFAULT NULL COMMENT '头像',
    `signature` varchar(255) NULL DEFAULT NULL COMMENT '签名',
    `otp` varchar(255) NULL DEFAULT NULL COMMENT 'OTP',
    `type2fa` varchar(255) NOT NULL DEFAULT 0 COMMENT '关闭:0;开启:1',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `user_name_password_idx` (`user_name`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `party`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '活动id',
    `founder_id` bigint NOT NULL COMMENT '创建者id',
    `title` varchar(255) NOT NULL COMMENT '活动标题',
    `content` varchar(255) NOT NULL COMMENT '活动介绍',
    `type` varchar(255) NOT NULL COMMENT '活动类型',
    `province` varchar(255) NOT NULL COMMENT '活动省份',
    `city` varchar(255) NOT NULL COMMENT '活动城市',
    `rectangle` varchar(255) NOT NULL COMMENT 'rectangle',
    `start_time` timestamp NOT NULL COMMENT '开始时间',
    `end_time` timestamp NOT NULL COMMENT '结束时间',
    `status` bigint NOT NULL DEFAULT 0 COMMENT  '状态',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    constraint  `party_founder_user`
        foreign key (`founder_id`)
            references `user` (`id`)
            on delete cascade
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='活动表';

CREATE TABLE `member`(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `party_id` bigint NOT NULL COMMENT '活动id',
    `member_id` bigint NOT NULL COMMENT '成员id',
    `status` bigint NOT NULL DEFAULT 0 COMMENT '0:未加入,1:已加入,2:admin',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    constraint `member_party_id`
        foreign key (`party_id`)
            references `party`(id)
            on delete cascade
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='成员表';

CREATE TABLE `itinerary`(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL COMMENT '标题',
    `founder_id` bigint NOT NULL COMMENT '创建者id',
    `party_id` bigint NOT NULL COMMENT '活动id',
    `is_merged` bigint NOT NULL DEFAULT 0 COMMENT 'is_merged',
    `action_type` bigint NOT NULL COMMENT '类型',
    `rectangle` varchar(255) NULL DEFAULT NULL COMMENT 'rectangle',
    `route_json` varchar(2000) NULL DEFAULT NULL COMMENT 'route_json',
    `remark` varchar(255) NULL DEFAULT NULL COMMENT '备注',
    `schedule_start_time` varchar(255) NOT NULL COMMENT '开始时间',
    `schedule_end_time` varchar(255) NOT NULL COMMENT '结束时间',
    `sequence` bigint NOT NULL COMMENT 'sequence',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    constraint `itinerary_party_id`
        foreign key (`party_id`)
            references `party`(id)
            on delete cascade
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='行程表';

CREATE TABLE `comment`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论id' ,
    `poi_id` bigint NOT NULL COMMENT '被评论poi_id',
    `uid` bigint NOT NULL COMMENT '评论用户',
    `content` varchar(2048) NOT NULL COMMENT '评论内容',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '评论时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='评论表';