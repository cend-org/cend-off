set sql_mode = ``;
drop database if exists cend;
create database cend;
use cend;

create table user
(
    id                     int auto_increment
        primary key,
    created_at             datetime     default CURRENT_TIMESTAMP     not null,
    updated_at             datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at             datetime     default '0000-00-00 00:00:00' null,
    name                   varchar(500) default ''                    null,
    family_name            varchar(500) default ''                    null,
    nick_name              varchar(100) default ''                    null,
    email                  varchar(100) default ''                    null,
    matricule              varchar(32)  default ''                    null,
    age                    int          default 0                     null,
    birth_date             datetime     default '0000-00-00 00:00:00' null,
    sex                    int          default 0                     null,
    lang                   int          default 0                     null,
    status                 int          default 0                     null,
    profile_image_xid      varchar(250) default ''                    null,
    description            varchar(500) default ''                    null,
    cover_text             varchar(500) default ''                    null,
    profile                varchar(500) default ''                    null,
    experience_detail      varchar(500) default ''                    null,
    additional_description varchar(500) default ''                    null,
    add_on_title           varchar(250) default ''                    null,
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
    id         int auto_increment
        primary key,
    created_at datetime     default CURRENT_TIMESTAMP     not null,
    updated_at datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at datetime     default '0000-00-00 00:00:00' null,
    user_id    int          default 0                     null,
    hash       varchar(500) default ''                    null,
    constraint password_user_id_fk
        foreign key (user_id) references user (id)
);

create index password_user_id_index
    on password (user_id);

create table media
(
    id         int auto_increment
        primary key,
    created_at datetime     default CURRENT_TIMESTAMP     not null,
    updated_at datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at datetime     default '0000-00-00 00:00:00' null,
    file_name  varchar(500) default ''                    null,
    extension  varchar(10)  default ''                    null,
    xid        varchar(100) default ''                    null,
    constraint user_pk
        unique (xid)
);

create table media_thumb
(
    id         int auto_increment
        primary key,
    created_at datetime     default CURRENT_TIMESTAMP,
    updated_at datetime     default CURRENT_TIMESTAMP,
    deleted_at datetime     default '0000-00-00 00:00:00',
    extension  varchar(10)  default '',
    media_xid  varchar(100) default '',
    xid        varchar(500) default ''
);

create table user_media_detail
(
    id            int auto_increment primary key,
    created_at    datetime            default CURRENT_TIMESTAMP,
    updated_at    datetime            default CURRENT_TIMESTAMP,
    deleted_at    datetime            default '0000-00-00 00:00:00',
    owner_id      int                 default 0,
    document_type int                 default 0,
    document_xid  varchar(100) unique default ''
);

alter table user_media_detail
    add constraint user_media_detail_user_id_fk
        foreign key (owner_id) references user (id)
            on delete cascade on update cascade;


create table user_authorization_link
(
    id         int auto_increment
        primary key,
    created_at datetime default CURRENT_TIMESTAMP,
    updated_at datetime default CURRENT_TIMESTAMP,
    deleted_at datetime default '0000-00-00 00:00:00',
    link_type  int
);

create table user_authorization_link_actor
(
    id                         int auto_increment
        primary key,
    created_at                 datetime default CURRENT_TIMESTAMP,
    updated_at                 datetime default CURRENT_TIMESTAMP,
    deleted_at                 datetime default '0000-00-00 00:00:00',
    user_authorization_link_id int      default 0,
    authorization_id           int      default 0
);

alter table user_authorization_link_actor
    add constraint user_authorization_link_actor_user_authorization_link_id_fk
        foreign key (user_authorization_link_id) references user_authorization_link (id)
            on delete cascade on update cascade;


alter table user_authorization_link_actor
    add constraint user_authorization_link_actor_authorization_id_fk
        foreign key (authorization_id) references authorization (id)
            on delete cascade on update cascade;


create table academic_level
(
    id         int auto_increment
        primary key,
    created_at datetime     default CURRENT_TIMESTAMP     not null,
    updated_at datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at datetime     default '0000-00-00 00:00:00' null,
    name       varchar(500) default ''                    null
);

create table academic_course
(
    id                int auto_increment
        primary key,
    created_at        datetime     default CURRENT_TIMESTAMP     not null,
    updated_at        datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at        datetime     default '0000-00-00 00:00:00' null,
    academic_level_id int          default 0                     null,
    name              varchar(500) default ''                    null
);

create table user_academic_course
(
    id         int auto_increment
        primary key,
    created_at datetime default CURRENT_TIMESTAMP     not null,
    updated_at datetime default CURRENT_TIMESTAMP     not null,
    deleted_at datetime default '0000-00-00 00:00:00' null,
    user_id    int      default 0                     null,
    course_id  int      default 0                     null
);

alter table user_academic_course
    add constraint user_academic_course_user_id_fk
        foreign key (user_id) references user (id)
            on delete cascade on update cascade;

alter table user_academic_course
    add constraint user_academic_course_course_id_fk
        foreign key (course_id) references academic_course (id)
            on delete cascade on update cascade;



create table user_academic_course_preference
(
    id         int auto_increment primary key,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp default '0000-00-00 00:00:00',
    user_id    int       default 0,
    is_online  boolean
);

alter table user_academic_course_preference
    add constraint user_academic_course_preference_user_id_fk
        foreign key (user_id) references user (id)
            on delete cascade on update cascade;


insert into academic_level (name)
values ('primaire 1'),
       ('primaire 2'),
       ('primaire 3'),
       ('primaire 4'),
       ('primaire 5'),
       ('primaire 6'),
       ('secondaire 1'),
       ('secondaire 2'),
       ('secondaire 3'),
       ('secondaire 4'),
       ('secondaire 5'),
       ('Cégep'),
       ('Universités');

--
--
--  NEW UPDATES RIGHT HERE
--
--
-- Primaire 1
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'primaire 1'), 'Mathématiques'),
       ((select academic_level.id from academic_level where name = 'primaire 1'), 'Français'),
       ((select academic_level.id from academic_level where name = 'primaire 1'), 'Anglais'),
       ((select academic_level.id from academic_level where name = 'primaire 1'), 'Science et technologie'),
       ((select academic_level.id from academic_level where name = 'primaire 1'), 'Histoire / Géographie'),
       ((select academic_level.id from academic_level where name = 'primaire 1'), 'Éthique  et culture religieuse'),
       ((select academic_level.id from academic_level where name = 'primaire 1'),
        'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 2
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'primaire 2'), 'Mathématiques'),
       ((select academic_level.id from academic_level where name = 'primaire 2'), 'Français'),
       ((select academic_level.id from academic_level where name = 'primaire 2'), 'Anglais'),
       ((select academic_level.id from academic_level where name = 'primaire 2'), 'Science et technologie'),
       ((select academic_level.id from academic_level where name = 'primaire 2'), 'Histoire / Géographie'),
       ((select academic_level.id from academic_level where name = 'primaire 2'), 'Éthique  et culture religieuse'),
       ((select academic_level.id from academic_level where name = 'primaire 2'),
        'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 3
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'primaire 3'), 'Mathématiques'),
       ((select academic_level.id from academic_level where name = 'primaire 3'), 'Français'),
       ((select academic_level.id from academic_level where name = 'primaire 3'), 'Anglais'),
       ((select academic_level.id from academic_level where name = 'primaire 3'), 'Science et technologie'),
       ((select academic_level.id from academic_level where name = 'primaire 3'), 'Histoire / Géographie'),
       ((select academic_level.id from academic_level where name = 'primaire 3'), 'Éthique  et culture religieuse'),
       ((select academic_level.id from academic_level where name = 'primaire 3'),
        'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 4
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'primaire 4'), 'Mathématiques'),
       ((select academic_level.id from academic_level where name = 'primaire 4'), 'Français'),
       ((select academic_level.id from academic_level where name = 'primaire 4'), 'Anglais'),
       ((select academic_level.id from academic_level where name = 'primaire 4'), 'Science et technologie'),
       ((select academic_level.id from academic_level where name = 'primaire 4'), 'Histoire / Géographie'),
       ((select academic_level.id from academic_level where name = 'primaire 4'), 'Éthique  et culture religieuse'),
       ((select academic_level.id from academic_level where name = 'primaire 4'),
        'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 5
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'primaire 5'), 'Mathématiques'),
       ((select academic_level.id from academic_level where name = 'primaire 5'), 'Français'),
       ((select academic_level.id from academic_level where name = 'primaire 5'), 'Anglais'),
       ((select academic_level.id from academic_level where name = 'primaire 5'), 'Science et technologie'),
       ((select academic_level.id from academic_level where name = 'primaire 5'), 'Histoire / Géographie'),
       ((select academic_level.id from academic_level where name = 'primaire 5'), 'Éthique  et culture religieuse'),
       ((select academic_level.id from academic_level where name = 'primaire 5'),
        'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 6
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'primaire 6'), 'Mathématiques'),
       ((select academic_level.id from academic_level where name = 'primaire 6'), 'Français'),
       ((select academic_level.id from academic_level where name = 'primaire 6'), 'Anglais'),
       ((select academic_level.id from academic_level where name = 'primaire 6'), 'Science et technologie'),
       ((select academic_level.id from academic_level where name = 'primaire 6'), 'Histoire / Géographie'),
       ((select academic_level.id from academic_level where name = 'primaire 6'), 'Éthique  et culture religieuse'),
       ((select academic_level.id from academic_level where name = 'primaire 6'),
        'Culture et citoyenneté québécoise (CCQ)');

-- Secondaire 1
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'secondaire 1'), 'Français'),
       ((select academic_level.id from academic_level where name = 'secondaire 1'), 'Histoire / Géographie'),
       ((select academic_level.id from academic_level where name = 'secondaire 1'), 'Sciences et technologies'),
       ((select academic_level.id from academic_level where name = 'secondaire 1'), 'Monde contemporain'),
       ((select academic_level.id from academic_level where name = 'secondaire 1'), 'academic_level financière'),
       ((select academic_level.id from academic_level where name = 'secondaire 1'), 'Éthique et culture religieuse'),
       ((select academic_level.id from academic_level where name = 'secondaire 1'), 'Mathématiques');

-- Secondaire 2
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'secondaire 2'), 'Français'),
       ((select academic_level.id from academic_level where name = 'secondaire 2'), 'Histoire / Géographie'),
       ((select academic_level.id from academic_level where name = 'secondaire 2'), 'Sciences et technologies'),
       ((select academic_level.id from academic_level where name = 'secondaire 2'), 'Monde contemporain'),
       ((select academic_level.id from academic_level where name = 'secondaire 2'), 'academic_level financière'),
       ((select academic_level.id from academic_level where name = 'secondaire 2'), 'Éthique et culture religieuse'),
       ((select academic_level.id from academic_level where name = 'secondaire 2'), 'Mathématiques');

-- secondaire 3
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'secondaire 3'), 'Sciences ST'),
       ((select academic_level.id from academic_level where name = 'secondaire 3'), 'Sciences ATS');

-- secondaire 4
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'secondaire 4'), 'Mathématiques CST'),
       ((select academic_level.id from academic_level where name = 'secondaire 4'), 'Mathématiques SN'),
       ((select academic_level.id from academic_level where name = 'secondaire 4'), 'Mathématiques TS'),
       ((select academic_level.id from academic_level where name = 'secondaire 4'), 'Sciences STE'),
       ((select academic_level.id from academic_level where name = 'secondaire 4'), 'Sciences SE');

-- secondaire 5
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'secondaire 5'), 'Mathématiques CST'),
       ((select academic_level.id from academic_level where name = 'secondaire 5'), 'Mathématiques SN'),
       ((select academic_level.id from academic_level where name = 'secondaire 5'), 'Mathématiques TS'),
       ((select academic_level.id from academic_level where name = 'secondaire 5'), 'Chimie'),
       ((select academic_level.id from academic_level where name = 'secondaire 5'), 'Physique');

-- Cégep
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'cégep'), 'Mathématiques'),
       ((select academic_level.id from academic_level where name = 'cégep'), 'Chimie'),
       ((select academic_level.id from academic_level where name = 'cégep'), 'Biologie'),
       ((select academic_level.id from academic_level where name = 'cégep'), 'Physique');

-- Universite
insert into academic_course(academic_level_id, name)
values ((select academic_level.id from academic_level where name = 'Universités'), 'Chimie'),
       ((select academic_level.id from academic_level where name = 'Universités'), 'Biologie'),
       ((select academic_level.id from academic_level where name = 'Universités'), 'Physique');