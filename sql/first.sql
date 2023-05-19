create database CarbonBasic;
use CarbonBasic;

create table user(
	uid int not null auto_increment comment '用户ID',
	username varchar(255) not null comment '用户名',
	passwd varchar(255) not null comment '密码',
	identity int not null comment '用户身份',
	state int not null comment '用户状态',
	user_desc varchar(255) not null comment '描述',
	primary key (uid)
) engine=InnoDB default charset=utf8 comment='用户表';

create table data(
	id int not null auto_increment comment '数据ID',
	plaintext varchar(255) not null comment '明文',
	ciphertext varchar(255) not null comment '密文',
	upload_date datetime not null comment '上传时间',
	state int not null comment '状态',
	uid int not null comment '上传者id',
	proof varchar(255) comment '证明',
	data_desc varchar(255) not null comment '描述',
	primary key (id)
) engine=InnoDB default charset=utf8 comment='数据表';

create table result(
	id int not null auto_increment comment '结果id',
	uid int not null comment '计算者id',
	res varchar(255) not null comment '结果',
	date datetime not null comment '查询时间',
	res_desc varchar(255) not null comment '描述',
	primary key (id)
) engine=InnoDB default charset=utf8 comment='结果表';