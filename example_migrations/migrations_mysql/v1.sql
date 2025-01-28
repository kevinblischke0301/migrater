USE `example`;

DROP TABLE IF EXISTS `foo`;

CREATE TABLE `foo` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);
