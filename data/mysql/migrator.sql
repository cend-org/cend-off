create database duval;

use duval;

create table user (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null ,
    updated_at DATETIME default CURRENT_TIMESTAMP not null ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    first_name varchar(100) default '',
    middle_name varchar(100) default '',
    last_name varchar(100) default '',
    nick_name varchar(100) default '',
    email varchar(100) default '',
    age int default 0
);

create table authorization (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null ,
    updated_at DATETIME default CURRENT_TIMESTAMP not null ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    user_id int default 0,
    level int default 0
);

create table person_link (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null ,
    updated_at DATETIME default CURRENT_TIMESTAMP not null ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    initiator_id int default 0,
    user_id int default 0,
    link_type int default 0
);

alter table person_link add constraint unique (initiator_id, user_id, link_type);

alter table authorization
    add constraint authorization_user_id_fk
        foreign key (user_id) references user (id)
            on update cascade on delete cascade;


alter table user
    add constraint user_pk
        unique (email);

alter table user add column birth_date datetime default '0000-00-00 00:00:00';

create table password (
                      id int primary key auto_increment not null,
                      created_at DATETIME default CURRENT_TIMESTAMP not null ,
                      updated_at DATETIME default CURRENT_TIMESTAMP not null ,
                      deleted_at DATETIME default '0000-00-00 00:00:00',
                      user_id int default 0 ,
                      psw varchar(1000) default ''
);

create index password_user_id_index
    on password (user_id);

alter table password
    add constraint password_user_id_fk
        foreign key (user_id) references user (id);

create table discussion (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null ,
    updated_at DATETIME default CURRENT_TIMESTAMP not null ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    name varchar(500) default ''
);

create table media (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null ,
    updated_at DATETIME default CURRENT_TIMESTAMP not null ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    file_name varchar(500) default '',
    extension varchar(10) default '',
    xid varchar(500) default '',
    user_id int default 0
);

create table discussion_actor (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null ,
    updated_at DATETIME default CURRENT_TIMESTAMP not null ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    discussion_id int default 0,
    user_id int default 0
);

create table message (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null ,
    updated_at DATETIME default CURRENT_TIMESTAMP not null ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    actor_id int default 0,
    content varchar(5000)  default '',
    status int default 0
);



alter table discussion_actor
    add constraint disc_act_fk
        foreign key (discussion_id) references discussion_actor(id)
on delete cascade on update cascade ;

alter table discussion_actor
    add constraint act_user_fk
        foreign key (user_id) references user(id)
on delete cascade on update cascade ;

create unique index disc_user_id
    on discussion_actor (discussion_id, user_id);

alter table message
    add constraint msg_actor
        foreign key (actor_id) references discussion_actor(id)
on delete cascade on update cascade ;

alter table discussion add column  last_message_sent_date DATETIME default '0000-00-00 00:00:00';

alter table user add column sex int default 0 after birth_date;
alter table user add column is_parent bool default false after sex;
alter table user add column is_student bool default false after is_parent;
alter table user add column is_tutor bool default false after is_student;
alter table user add column is_professor bool default false after is_tutor;
alter table user add column lang int default 0 after is_professor;

alter table media add column content_type int default 0 after user_id;

-- remove is_parent is_student is_tutor is_professor
alter table user drop column is_parent;
alter table user drop column is_student;
alter table user drop column is_tutor ;
alter table user drop column is_professor ;