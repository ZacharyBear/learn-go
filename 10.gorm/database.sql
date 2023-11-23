create table `user`(
  `id` int PRIMARY KEY auto_increment,
  `name` varchar(50) not null,
  `birthday` date,
  `email` varchar(30)
);