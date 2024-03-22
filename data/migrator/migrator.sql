drop database if exists duval;
create database duval;
use duval;
create table discussion
(
    id                     int auto_increment
        primary key,
    created_at             datetime     default CURRENT_TIMESTAMP     not null,
    updated_at             datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at             datetime     default '0000-00-00 00:00:00' null,
    name                   varchar(500) default ''                    null,
    last_message_sent_date datetime     default '0000-00-00 00:00:00' null
);

create table media
(
    id           int auto_increment
        primary key,
    created_at   datetime     default CURRENT_TIMESTAMP     not null,
    updated_at   datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime     default '0000-00-00 00:00:00' null,
    file_name    varchar(500) default ''                    null,
    extension    varchar(10)  default ''                    null,
    xid          varchar(500) default ''                    null,
    user_id      int          default 0                     null,
    content_type int          default 0                     null
);

create table user
(
    id          int auto_increment
        primary key,
    created_at  datetime     default CURRENT_TIMESTAMP     not null,
    updated_at  datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at  datetime     default '0000-00-00 00:00:00' null,
    name        varchar(500) default ''                    null,
    family_name varchar(500) default ''                    null,
    nick_name   varchar(100) default ''                    null,
    email       varchar(100) default ''                    null,
    matricule   varchar(32)  default ''                    null,
    age         int          default 0                     null,
    birth_date  datetime     default '0000-00-00 00:00:00' null,
    sex         int          default 0                     null,
    lang        int          default 0                     null,
    status      int          default 0                     null,
    constraint user_pk
        unique (email)
);

create table authorization
(
    id         int auto_increment
        primary key,
    created_at datetime default CURRENT_TIMESTAMP     not null,
    updated_at datetime default CURRENT_TIMESTAMP     not null,
    deleted_at datetime default '0000-00-00 00:00:00' null,
    user_id    int      default 0                     null,
    level      int      default 0                     null,
    constraint authorization_pk
        unique (user_id, level),
    constraint authorization_user_id_fk
        foreign key (user_id) references user (id)
            on update cascade on delete cascade
);

create table password
(
    id           int auto_increment
        primary key,
    created_at   datetime      default CURRENT_TIMESTAMP     not null,
    updated_at   datetime      default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime      default '0000-00-00 00:00:00' null,
    user_id      int           default 0                     null,
    psw          varchar(1000) default ''                    null,
    content_hash varchar(500)  default ''                    null,
    constraint password_user_id_fk
        foreign key (user_id) references user (id)
);

create index password_user_id_index
    on password (user_id);

-- upgrade end  on 13/03/24
create table code
(
    id           int auto_increment
        primary key,
    created_at   datetime      default CURRENT_TIMESTAMP     not null,
    updated_at   datetime      default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime      default '0000-00-00 00:00:00' null,
    user_id      int           default 0                     null,
    verification_code          int default 0                    null
);
create index code_user_val on code(user_id,verification_code);


create table message
(
    id           int auto_increment
        primary key,
    created_at   datetime      default CURRENT_TIMESTAMP     not null,
    updated_at   datetime      default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime      default '0000-00-00 00:00:00' null,
    identifier varchar(100) default '' null,
    number int default 0 null,
    xid varchar(500) default '' null,
    label varchar(5000) default '' null
);

alter table message add column  language int default 0 after label;

create unique index id_lang on message(identifier,number, language);

create table menu
(
    id           int auto_increment
        primary key,
    created_at   datetime      default CURRENT_TIMESTAMP     not null,
    updated_at   datetime      default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime      default '0000-00-00 00:00:00' null,
    identifier varchar(100) default '' null,
    number int default 0 null,
    message_number int default 0 null
);

create table menu_item
(
    id           int auto_increment
        primary key,
    created_at   datetime      default CURRENT_TIMESTAMP     not null,
    updated_at   datetime      default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime      default '0000-00-00 00:00:00' null,
    identifier varchar(100) default '' null,
    number int default 0 null,
    menu_title_message_number int,
    message_number int default 0 null
);



-- update 13/03/24
create table address
(
    id           int auto_increment      primary key not null,
    created_at   datetime      default CURRENT_TIMESTAMP    ,
    updated_at   datetime      default CURRENT_TIMESTAMP    ,
    deleted_at   datetime      default '0000-00-00 00:00:00' ,
    country varchar(100) default '' ,
    city varchar(100) default '',
    latitude float default 0,
    longitude float default 0,
    street varchar(100) ,
    full_address varchar(600),
    xid          varchar(500) default ''
);

create table user_address
(
    id           int auto_increment      primary key not null,
    created_at   datetime      default CURRENT_TIMESTAMP    ,
    updated_at   datetime      default CURRENT_TIMESTAMP    ,
    deleted_at   datetime      default '0000-00-00 00:00:00' ,
    user_id int unique,
    address_id int unique,
    address_type varchar(100) default '',
    foreign key (user_id) references user(id),
    foreign key (address_id) references address(id)
);

-- update 16/03/24
create table thumb
(
    id           int auto_increment
        primary key,
    created_at   datetime     default CURRENT_TIMESTAMP     not null,
    updated_at   datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime     default '0000-00-00 00:00:00' null,
    file_name    varchar(500) default ''                    null,
    extension    varchar(10)  default ''                    null,
    media_xid          varchar(500) default ''                    null,
    content_type int          default 0                     null
);


-- update 18/03/24
alter table user add profile_image_xid varchar(500) default '' after status ;
-- *

-- update 19/03/24
alter table thumb
    rename media_thumb,
    add column xid varchar(500) default '',
    drop column content_type ,
    drop column file_name ;

-- ***** --
drop table menu;
drop table menu_item;
drop table message;

create table message
                   (
                       id           int auto_increment
                           primary key,
                       created_at   datetime     default CURRENT_TIMESTAMP     not null,
                       updated_at   datetime     default CURRENT_TIMESTAMP     not null,
                       deleted_at   datetime     default '0000-00-00 00:00:00' null,
                       resource_type int default 0,
                       resource_number int default 0,
                       resource_value int default 0,
                       resource_label varchar(5000) default ''
                   );

alter table message add column  resource_language int default 0 after resource_label;
create unique index msg_type_nb_val_lang
    on message (resource_type, resource_number, resource_value, resource_language);

-- update 21/03/2024
create table user_media_detail
(
   id     int auto_increment       primary key,
   created_at   datetime     default CURRENT_TIMESTAMP ,
   updated_at   datetime     default CURRENT_TIMESTAMP ,
   deleted_at   datetime     default '0000-00-00 00:00:00',
   owner_id int ,
   document_type int,
   document_xid varchar(100) unique,
   foreign key (owner_id) references user(id),
   foreign key (document_xid) references media(id)
);

alter table media drop column user_id;
-- *