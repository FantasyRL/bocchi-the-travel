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

CREATE TABLE `video` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频ID',
    `uid` bigint NOT NULL COMMENT '用户ID',
    `play_url` varchar(255) NOT NULL COMMENT '视频url',
    `cover_url` varchar(255) NOT NULL COMMENT '封面url',
    `title` varchar(255) DEFAULT NULL COMMENT '标题',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '发布时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `time` (`created_at`) USING BTREE,
    KEY `author` (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

CREATE TABLE `like` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `uid` bigint NOT NULL COMMENT '点赞用户id',
    `video_id` bigint NOT NULL COMMENT '被点赞的视频id',
    `status` bigint NOT NULL DEFAULT 1 COMMENT '点赞：1 取消：0',#因为取消点赞不是真的删除行所以就多个status
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '点赞时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    constraint `favorite_user`
        foreign key (`uid`)
            references `user` (`id`)
            on delete cascade,
    constraint `favorite_video`
        foreign key (`video_id`)
            references video (`id`)
            on delete cascade,
    KEY `userIdtoVideoIdIdx` (`uid`,`video_id`) USING BTREE,
    KEY `uid_idx` (`uid`) USING BTREE,
    KEY `video_idx` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='点赞表';

CREATE TABLE `comment_like` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `uid` bigint NOT NULL COMMENT '点赞用户id',
    `comment_id` bigint NOT NULL COMMENT '被点赞的评论id',
    `status` bigint NOT NULL DEFAULT 1 COMMENT '点赞：1 取消：0',#因为取消点赞不是真的删除行所以就多个status
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '点赞时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    constraint `favorite_comment_user`
        foreign key (`uid`)
            references `user` (`id`)
            on delete cascade,
    constraint `favorite_comment`
        foreign key (`comment_id`)
            references comment (`id`)
            on delete cascade
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='评论点赞表';

CREATE TABLE `comment`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论id' ,
    `video_id` bigint NOT NULL COMMENT '被评论视频id',
    `parent_id` bigint NULL DEFAULT NULL COMMENT '父评论id' ,
    `uid` bigint NOT NULL COMMENT '评论用户',
    `content` varchar(2048) NOT NULL COMMENT '评论内容',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '评论时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    constraint `video_comment`
        foreign key (`video_id`)
            references video (`id`)
            on delete cascade
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

CREATE TABLE `follow`(
    `id` bigint NOT NULL AUTO_INCREMENT ,
    `uid` bigint NOT NULL COMMENT 'user_id',
    `followed_id` bigint NOT NULL COMMENT '被关注者',
    `status` bigint NOT NULL DEFAULT 1 COMMENT '1:关注;0:取消关注',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp ,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `follower_user`
        FOREIGN KEY (`uid`)
            REFERENCES user (`id`)
            ON DELETE CASCADE,
    CONSTRAINT `followed_user`
        FOREIGN KEY (`followed_id`)
            REFERENCES user (`id`)
            ON DELETE CASCADE
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='关注表';

CREATE TABLE `message`(
    `id` bigint NOT NULL AUTO_INCREMENT ,
    `uid` bigint NOT NULL COMMENT 'user_id',
    `target_id` bigint NOT NULL COMMENT '发送对象id',
    `content` varchar(2048) NOT NULL COMMENT '信息',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp ,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `message_user`
        FOREIGN KEY (`uid`)
            REFERENCES user (`id`)
            ON DELETE CASCADE
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='关注表';