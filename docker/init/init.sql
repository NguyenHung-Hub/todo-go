CREATE DATABASE IF NOT EXISTS todo;

CREATE TABLE IF NOT EXISTS `todo_items`(
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(255),
    `description` varchar(255),
    `status` ENUM('Doing','Done','Deleted'),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);