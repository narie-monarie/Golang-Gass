
-- +migrate Up
CREATE TABLE `user` (
 `id`       int(11) PRIMARY KEY,
 `username` text,
 `password` text,
 `email`    text
);
-- +migrate Down
DROP TABLE user;

