-- drop database
DROP DATABASE IF EXISTS `goim`;

-- create database
create database if not exists goim;

-- define user table based on go structure User in current directory
DROP TABLE IF EXISTS goim.user;

create table if not exists goim.user (
	`id` bigint not null auto_increment,
	`uid` varchar(64) not null, -- 22 bytes of uuid
	`name` varchar(32) not null,
	`password` varchar(128) not null,
	`email` varchar(32) not null,
	`phone` varchar(32) not null,
	`avatar` varchar(128) not null,
	`status` tinyint not null DEFAULT 0,
	`created_at` int not null DEFAULT 0,
	`updated_at` int not null DEFAULT 0,
	primary key (`id`),
	unique key (`uid`),
    key (`email`),
    key (`phone`)
) auto_increment = 10000 engine = innodb charset = utf8mb4;

-- mock data
insert into goim.user (`id`, `uid`, `name`, `password`, `email`, `phone`, `avatar`, `status`, `created_at`, `updated_at`)
values
    (10000, '4F8DSQByUsEUMoETzTCabh', 'user1', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'user0@example.com', ' ', ' ', 1, 1528894200, 1528894200),
    (10001, 'C6CtUjpC6h5e5SW9tBFNVX', 'user2', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'user1@example.com', ' ', ' ', 0, 1528894200, 1528894200),
    (10002, '7mRZLYedtK1EwxzC5X1Lxf', 'user3', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'user2@example.com', ' ', ' ', 0, 1528894200, 1528894200),
    (10003, 'WmbtshDDMUgb3KWFisWZ4E', 'user4', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', 'user3@example.com', ' ', ' ', 0, 1528894200, 1528894200);

-- define friend table based on go structure Friend in current directory
DROP TABLE IF EXISTS goim.friend;

CREATE TABLE IF NOT EXISTS goim.friend (
    `id` bigint not null auto_increment,
    `uid` varchar(64) not null, -- 22 bytes of uuid
    `friend_uid` varchar(64) not null, -- 22 bytes of uuid
    `status` tinyint not null default 0 COMMENT '0: friend; 1: stranger; 2: blacked',
    `created_at` int not null default 0,
    `updated_at` int not null default 0,
    primary key (`id`),
    unique key (`uid`, `friend_uid`),
    key (`uid`)
) auto_increment = 10000 engine = innodb charset = utf8mb4;

-- mock data
insert into goim.friend (`id`, `uid`, `friend_uid`, `status`, `created_at`, `updated_at`)
values
    (10000, '4F8DSQByUsEUMoETzTCabh', 'C6CtUjpC6h5e5SW9tBFNVX', 1, 1528894200, 1528894200),
    (10001, 'C6CtUjpC6h5e5SW9tBFNVX', '4F8DSQByUsEUMoETzTCabh', 1, 1528894200, 1528894200),
    (10002, 'C6CtUjpC6h5e5SW9tBFNVX', '7mRZLYedtK1EwxzC5X1Lxf', 1, 1528894200, 1528894200),
    (10003, 'C6CtUjpC6h5e5SW9tBFNVX', 'WmbtshDDMUgb3KWFisWZ4E', 1, 1528894200, 1528894200),
    (10004, '7mRZLYedtK1EwxzC5X1Lxf', '4F8DSQByUsEUMoETzTCabh', 1, 1528894200, 1528894200),
    (10005, '7mRZLYedtK1EwxzC5X1Lxf', 'C6CtUjpC6h5e5SW9tBFNVX', 1, 1528894200, 1528894200),
    (10006, '7mRZLYedtK1EwxzC5X1Lxf', 'WmbtshDDMUgb3KWFisWZ4E', 1, 1528894200, 1528894200),
    (10007, 'WmbtshDDMUgb3KWFisWZ4E', '4F8DSQByUsEUMoETzTCabh', 1, 1528894200, 1528894200),
    (10008, 'WmbtshDDMUgb3KWFisWZ4E', 'C6CtUjpC6h5e5SW9tBFNVX', 1, 1528894200, 1528894200),
    (10009, 'WmbtshDDMUgb3KWFisWZ4E', '7mRZLYedtK1EwxzC5X1Lxf', 1, 1528894200, 1528894200);

-- define friend_request table based on go structure FriendRequest in current directory
DROP TABLE IF EXISTS goim.friend_request;

CREATE TABLE IF NOT EXISTS goim.friend_request (
    `id` bigint not null auto_increment,
    `uid` varchar(64) not null, -- 22 bytes of uuid
    `friend_uid` varchar(64) not null, -- 22 bytes of uuid
    `status` tinyint not null default 0 COMMENT '0: pending; 1: accepted; 2: rejected',
    `created_at` int not null default 0,
    `updated_at` int not null default 0,
    primary key (`id`),
    unique key (`uid`, `friend_uid`),
    key (`uid`)
) auto_increment = 10000 engine = innodb charset = utf8mb4;