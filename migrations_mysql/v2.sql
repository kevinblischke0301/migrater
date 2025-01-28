USE `example`;

DROP TABLE IF EXISTS `bar`;

CREATE TABLE `bar` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);
