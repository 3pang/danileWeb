-- 导出  表 mysql223.users 结构
DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
                                       `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                       `name` varchar(25) NOT NULL,
                                       `age` tinyint(3) unsigned NOT NULL DEFAULT '0',
                                       PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- 正在导出表  mysql223.users 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `name`, `age`) VALUES
(1, '张三', 25),
(2, '李四', 22),
(3, '田七', 25);