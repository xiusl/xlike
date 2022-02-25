CREATE TABLE `user` (
    `id` bigint NOT NULL COMMENT '主键ID',
    `mobile` varchar(13) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '手机号码',
    `nick_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '用户昵称',
    `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '头像地址',
    `state` int NOT NULL DEFAULT '1' COMMENT '用户状态(-1-临时用户 1-注册用户 2-冻结 3-注销)',
    `memo` varchar(60) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '备注',
    `last_seen` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后登录时间',
    `version` int DEFAULT '0' COMMENT '乐观锁version',
    `del` tinyint(1) DEFAULT '0' COMMENT '逻辑删除标识 1-true-数据逻辑删除 0-false-未删除 默认-false',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间-只在创建时初始化一次',
    `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间-每次数据改动都会修改',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `mobile_index` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE `auth_code` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `mobile` varchar(13) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '手机号码',
    `code` varchar(12) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '验证码',
    `biz_code` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '业务代码',
    `status` int NOT NULL DEFAULT '0' COMMENT '状态(0-可用 1-使用过)',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间-只在创建时初始化一次',
    `expired_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;