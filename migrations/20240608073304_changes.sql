-- Create "pets" table
CREATE TABLE `pets` (`id` bigint NOT NULL AUTO_INCREMENT, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `age` bigint NOT NULL, `name` varchar(255) NOT NULL, `nickname` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `nickname` (`nickname`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
