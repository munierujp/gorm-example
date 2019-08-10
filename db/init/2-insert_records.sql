-- languages
INSERT INTO `languages` (`id`, `code`, `name`) VALUES
(1, 'en', 'English'),
(2, 'ja', 'Japanese');

-- users
INSERT INTO `users` (`id`, `name`, `language`) VALUES
(1, 'John Smith', 1),
(2, 'Tanaka Taro', 2);
