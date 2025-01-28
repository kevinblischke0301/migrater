USE `example`;

DROP TABLE IF EXISTS `foo_bar`;

CREATE TABLE `foo_bar` (
    `foo_id` INT NOT NULL,
    `bar_id` INT NOT NULL,
    PRIMARY KEY (`foo_id`, `bar_id`),
    FOREIGN KEY (`foo_id`) REFERENCES `foo`(`id`),
    FOREIGN KEY (`bar_id`) REFERENCES `bar`(`id`)
);
