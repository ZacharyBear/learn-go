create table `user`(
  `id` int PRIMARY KEY auto_increment,
  `name` varchar(50) not null,
  `birthday` date,
  `email` varchar(30),
  `delete_at` date
);

drop table if exists `email`;
create table `email`(
  `id` int PRIMARY KEY auto_increment,
  `name` varchar(50) not null,
  `address` varchar(100) not null
);

insert `email`(`name`, `address`)
values ('Zenkie Bear', 'zq@zenkie.cn'),
('John Doe', 'john@doe.com');
