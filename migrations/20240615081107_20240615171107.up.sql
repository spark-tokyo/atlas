-- create "pets" table
CREATE TABLE `pets` (
 `id` bigint NOT NULL AUTO_INCREMENT,
 PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "users" table
CREATE TABLE `users` (
 `id` varchar(255) NOT NULL,
 `age` bigint NOT NULL,
 `name` varchar(255) NOT NULL,
 `nickname` varchar(255) NOT NULL,
 `email` varchar(255) NOT NULL,
 PRIMARY KEY (`id`),
 UNIQUE INDEX `email` (`email`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
