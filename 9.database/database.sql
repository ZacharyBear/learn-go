drop table if exists t_user;
create table t_user(
	`id` int primary key auto_increment comment 'User ID',
	`name` varchar(50) not null,
	`birthday` date comment 'Birth Day'
);

insert into t_user(name, birthday)
values ('Zenkie Bear', '2004-03-04'),
  ('Taylor Swift', '1989-10-13'),
  -- Chaning is a deliberate typo
  ('Caroline Chaning', '1987-05-28');
