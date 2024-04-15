drop database if exists duval;
create database duval;
use duval;

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

create table code
(
    id                int auto_increment
        primary key,
    created_at        datetime default CURRENT_TIMESTAMP     not null,
    updated_at        datetime default CURRENT_TIMESTAMP     not null,
    deleted_at        datetime default '0000-00-00 00:00:00' null,
    user_id           int      default 0                     null,
    verification_code int      default 0                     null
);

create index code_user_val
    on code (user_id, verification_code);

create table message
(
    id         int auto_increment
        primary key,
    created_at datetime      default CURRENT_TIMESTAMP     not null,
    updated_at datetime      default CURRENT_TIMESTAMP     not null,
    deleted_at datetime      default '0000-00-00 00:00:00' null,
    identifier varchar(100)  default ''                    null,
    number     int           default 0                     null,
    xid        varchar(500)  default ''                    null,
    label      varchar(5000) default ''                    null
);

alter table message
    add column language int default 0 after label;

create unique index id_lang on message (identifier, number, language);

create table menu
(
    id             int auto_increment
        primary key,
    created_at     datetime     default CURRENT_TIMESTAMP     not null,
    updated_at     datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at     datetime     default '0000-00-00 00:00:00' null,
    identifier     varchar(100) default ''                    null,
    number         int          default 0                     null,
    message_number int          default 0                     null
);

create table menu_item
(
    id                        int auto_increment
        primary key,
    created_at                datetime     default CURRENT_TIMESTAMP     not null,
    updated_at                datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at                datetime     default '0000-00-00 00:00:00' null,
    identifier                varchar(100) default ''                    null,
    number                    int          default 0                     null,
    menu_title_message_number int,
    message_number            int          default 0                     null
);

create table address
(
    id           int auto_increment primary key not null,
    created_at   datetime     default CURRENT_TIMESTAMP,
    updated_at   datetime     default CURRENT_TIMESTAMP,
    deleted_at   datetime     default '0000-00-00 00:00:00',
    country      varchar(100) default '',
    city         varchar(100) default '',
    latitude     float        default 0,
    longitude    float        default 0,
    street       varchar(100),
    full_address varchar(600),
    xid          varchar(500) default ''
);

create table user_address
(
    id           int auto_increment primary key not null,
    created_at   datetime     default CURRENT_TIMESTAMP,
    updated_at   datetime     default CURRENT_TIMESTAMP,
    deleted_at   datetime     default '0000-00-00 00:00:00',
    user_id      int unique,
    address_id   int unique,
    address_type varchar(100) default '',
    foreign key (user_id) references user (id),
    foreign key (address_id) references address (id)
);

create table thumb
(
    id           int auto_increment
        primary key,
    created_at   datetime     default CURRENT_TIMESTAMP     not null,
    updated_at   datetime     default CURRENT_TIMESTAMP     not null,
    deleted_at   datetime     default '0000-00-00 00:00:00' null,
    file_name    varchar(500) default ''                    null,
    extension    varchar(10)  default ''                    null,
    media_xid    varchar(500) default ''                    null,
    content_type int          default 0                     null
);

alter table user
    add profile_image_xid varchar(500) default '' after status;

alter table thumb
    rename media_thumb,
    add column xid varchar(500) default '',
    drop column content_type,
    drop column file_name;

drop table menu;
drop table menu_item;
drop table message;

create table message
(
    id              int auto_increment
        primary key,
    created_at      datetime      default CURRENT_TIMESTAMP     not null,
    updated_at      datetime      default CURRENT_TIMESTAMP     not null,
    deleted_at      datetime      default '0000-00-00 00:00:00' null,
    resource_type   int           default 0,
    resource_number int           default 0,
    resource_value  int           default 0,
    resource_label  varchar(5000) default ''
);

alter table message
    add column resource_language int default 0 after resource_label;
create unique index msg_type_nb_val_lang
    on message (resource_type, resource_number, resource_value, resource_language);

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
        foreign key (owner_id) references user (id);


alter table media
    drop column user_id;

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

-- update 26/03/2024
create table qr_code_registry
(
    id         int primary key auto_increment,
    created_at datetime            default CURRENT_TIMESTAMP,
    deleted_at datetime            default '0000-00-00 00:00:00',
    user_id    int unique          default 0,
    xid        varchar(100) unique default '',
    is_used    boolean,
    foreign key (user_id) references user (id)
);

create table calendar_planning
(
    id               int primary key auto_increment,
    created_at       datetime     default CURRENT_TIMESTAMP,
    updated_at       datetime     default CURRENT_TIMESTAMP,
    deleted_at       datetime     default '0000-00-00 00:00:00',
    authorization_id int          default 0,
    start_date_time  datetime     default CURRENT_TIMESTAMP,
    end_date_time    datetime     default CURRENT_TIMESTAMP,
    description      varchar(100) default '',
    foreign key (authorization_id) references authorization (id)

);

create table calendar_planning_actor
(
    id                   int primary key auto_increment,
    created_at           datetime default CURRENT_TIMESTAMP,
    updated_at           datetime default CURRENT_TIMESTAMP,
    deleted_at           datetime default '0000-00-00 00:00:00',
    authorization_id     int      default 0,
    calendar_planning_id int      default 0,
    foreign key (authorization_id) references authorization (id),
    foreign key (calendar_planning_id) references calendar_planning (id)
);


--
--
--  NEW UPDATES RIGHT HERE
--
--


create table education
(
    id         int primary key auto_increment,
    created_at datetime     default CURRENT_TIMESTAMP,
    updated_at datetime     default CURRENT_TIMESTAMP,
    deleted_at datetime     default '0000-00-00 00:00:00',
    name       varchar(500) default ''
);

create table subject
(
    id                 int primary key auto_increment,
    created_at         datetime     default CURRENT_TIMESTAMP,
    updated_at         datetime     default CURRENT_TIMESTAMP,
    deleted_at         datetime     default '0000-00-00 00:00:00',
    education_level_id int          default 0,
    name               varchar(500) default ''
);

alter table subject
    add constraint subject_education_id_fk
        foreign key (education_level_id) references education (id);

-- update 06/04/2024
insert into education (name)
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
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'primaire 1'), 'Mathématiques'),
       ((select education.id from education where name = 'primaire 1'), 'Français'),
       ((select education.id from education where name = 'primaire 1'), 'Anglais'),
       ((select education.id from education where name = 'primaire 1'), 'Science et technologie'),
       ((select education.id from education where name = 'primaire 1'), 'Histoire / Géographie'),
       ((select education.id from education where name = 'primaire 1'), 'Éthique  et culture religieuse'),
       ((select education.id from education where name = 'primaire 1'), 'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 2
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'primaire 2'), 'Mathématiques'),
       ((select education.id from education where name = 'primaire 2'), 'Français'),
       ((select education.id from education where name = 'primaire 2'), 'Anglais'),
       ((select education.id from education where name = 'primaire 2'), 'Science et technologie'),
       ((select education.id from education where name = 'primaire 2'), 'Histoire / Géographie'),
       ((select education.id from education where name = 'primaire 2'), 'Éthique  et culture religieuse'),
       ((select education.id from education where name = 'primaire 2'), 'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 3
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'primaire 3'), 'Mathématiques'),
       ((select education.id from education where name = 'primaire 3'), 'Français'),
       ((select education.id from education where name = 'primaire 3'), 'Anglais'),
       ((select education.id from education where name = 'primaire 3'), 'Science et technologie'),
       ((select education.id from education where name = 'primaire 3'), 'Histoire / Géographie'),
       ((select education.id from education where name = 'primaire 3'), 'Éthique  et culture religieuse'),
       ((select education.id from education where name = 'primaire 3'), 'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 4
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'primaire 4'), 'Mathématiques'),
       ((select education.id from education where name = 'primaire 4'), 'Français'),
       ((select education.id from education where name = 'primaire 4'), 'Anglais'),
       ((select education.id from education where name = 'primaire 4'), 'Science et technologie'),
       ((select education.id from education where name = 'primaire 4'), 'Histoire / Géographie'),
       ((select education.id from education where name = 'primaire 4'), 'Éthique  et culture religieuse'),
       ((select education.id from education where name = 'primaire 4'), 'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 5
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'primaire 5'), 'Mathématiques'),
       ((select education.id from education where name = 'primaire 5'), 'Français'),
       ((select education.id from education where name = 'primaire 5'), 'Anglais'),
       ((select education.id from education where name = 'primaire 5'), 'Science et technologie'),
       ((select education.id from education where name = 'primaire 5'), 'Histoire / Géographie'),
       ((select education.id from education where name = 'primaire 5'), 'Éthique  et culture religieuse'),
       ((select education.id from education where name = 'primaire 5'), 'Culture et citoyenneté québécoise (CCQ)');

-- Primaire 6
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'primaire 6'), 'Mathématiques'),
       ((select education.id from education where name = 'primaire 6'), 'Français'),
       ((select education.id from education where name = 'primaire 6'), 'Anglais'),
       ((select education.id from education where name = 'primaire 6'), 'Science et technologie'),
       ((select education.id from education where name = 'primaire 6'), 'Histoire / Géographie'),
       ((select education.id from education where name = 'primaire 6'), 'Éthique  et culture religieuse'),
       ((select education.id from education where name = 'primaire 6'), 'Culture et citoyenneté québécoise (CCQ)');

-- Secondaire 1
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'secondaire 1'), 'Français'),
       ((select education.id from education where name = 'secondaire 1'), 'Histoire / Géographie'),
       ((select education.id from education where name = 'secondaire 1'), 'Sciences et technologies'),
       ((select education.id from education where name = 'secondaire 1'), 'Monde contemporain'),
       ((select education.id from education where name = 'secondaire 1'), 'Education financière'),
       ((select education.id from education where name = 'secondaire 1'), 'Éthique et culture religieuse'),
       ((select education.id from education where name = 'secondaire 1'), 'Mathématiques');

-- Secondaire 2
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'secondaire 2'), 'Français'),
       ((select education.id from education where name = 'secondaire 2'), 'Histoire / Géographie'),
       ((select education.id from education where name = 'secondaire 2'), 'Sciences et technologies'),
       ((select education.id from education where name = 'secondaire 2'), 'Monde contemporain'),
       ((select education.id from education where name = 'secondaire 2'), 'Education financière'),
       ((select education.id from education where name = 'secondaire 2'), 'Éthique et culture religieuse'),
       ((select education.id from education where name = 'secondaire 2'), 'Mathématiques');

-- secondaire 3
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'secondaire 3'), 'Sciences ST'),
       ((select education.id from education where name = 'secondaire 3'), 'Sciences ATS');

-- secondaire 4
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'secondaire 4'), 'Mathématiques CST'),
       ((select education.id from education where name = 'secondaire 4'), 'Mathématiques SN'),
       ((select education.id from education where name = 'secondaire 4'), 'Mathématiques TS'),
       ((select education.id from education where name = 'secondaire 4'), 'Sciences STE'),
       ((select education.id from education where name = 'secondaire 4'), 'Sciences SE');

-- secondaire 5
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'secondaire 5'), 'Mathématiques CST'),
       ((select education.id from education where name = 'secondaire 5'), 'Mathématiques SN'),
       ((select education.id from education where name = 'secondaire 5'), 'Mathématiques TS'),
       ((select education.id from education where name = 'secondaire 5'), 'Chimie'),
       ((select education.id from education where name = 'secondaire 5'), 'Physique');

-- Cégep
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'cégep'), 'Mathématiques'),
       ((select education.id from education where name = 'cégep'), 'Chimie'),
       ((select education.id from education where name = 'cégep'), 'Biologie'),
       ((select education.id from education where name = 'cégep'), 'Physique');

-- Universite
insert into subject(education_level_id, name)
values ((select education.id from education where name = 'Universités'), 'Chimie'),
       ((select education.id from education where name = 'Universités'), 'Biologie'),
       ((select education.id from education where name = 'Universités'), 'Physique');

--
--
--  ADD EXTRA FIELD TO USER
--
--

alter table user
    add (
        description varchar(500) default '',
        cover_text varchar(255) default '',
        profile varchar(255) default '',
        experience_detail varchar(500) default '',
        additional_description varchar(255) default '',
        add_on_title varchar(100) default ''
        );

--
--
--  INDIRECTION TABLE BETWEEN SUBJECT AND USER
--
--
create table user_education_level_subject
(
    id         int primary key auto_increment,
    created_at datetime default CURRENT_TIMESTAMP,
    updated_at datetime default CURRENT_TIMESTAMP,
    deleted_at datetime default '0000-00-00 00:00:00',
    user_id    int      default 0 unique,
    subject_id int      default 0 unique
);


alter table user_education_level_subject
    add constraint user_education_level_subject_user_id_fk
        foreign key (user_id) references user (id);


alter table user_education_level_subject
    add constraint user_education_level_subject_subject_id_fk
        foreign key (subject_id) references subject (id);

--
--
--  ADD TABLE USER_MARK
--
--
create table user_mark
(
    id             int primary key auto_increment,
    created_at     datetime     default CURRENT_TIMESTAMP,
    updated_at     datetime     default CURRENT_TIMESTAMP,
    deleted_at     datetime     default '0000-00-00 00:00:00',
    user_id        int          default 0,
    author_id      int          default 0 unique,
    author_comment varchar(255) default '',
    author_mark    int          default 0
);

alter table user_mark
    add constraint user_mark_user_id_fk
        foreign key (user_id) references user (id);

alter table user_mark
    add constraint user_mark_author_id_fk
        foreign key (author_id) references user (id);
-- *


--  update 09/04/2024
--
--
--  CONTRACT
--
--

create table contract
(
    id             int primary key auto_increment,
    created_at     datetime default CURRENT_TIMESTAMP,
    updated_at     datetime default CURRENT_TIMESTAMP,
    deleted_at     datetime default '0000-00-00 00:00:00',
    tutor_id       int      default 0 unique,
    parent_id      int      default 0 unique,
    student_id     int      default 0 unique,
    start_date     datetime default '0000-00-00 00:00:00',
    end_date       datetime default '0000-00-00 00:00:00',
    payment_type   int      default 0,
    salary_value   float    default 0,
    payment_method int      default 0
);

alter table contract
    add constraint contract_tutor_id_fk
        foreign key (tutor_id) references user (id);


alter table contract
    add constraint contract_parent_id_fk
        foreign key (parent_id) references user (id);


alter table contract
    add constraint contract_student_id_fk
        foreign key (student_id) references user (id);


create table contract_timesheet_detail
(
    id          int primary key auto_increment,
    created_at  datetime   default CURRENT_TIMESTAMP,
    updated_at  datetime   default CURRENT_TIMESTAMP,
    deleted_at  datetime   default '0000-00-00 00:00:00',
    contract_id int  default 0,
    date        date       default '0000-00-00',
    hours       time        default 0
);

alter table contract_timesheet_detail
    add constraint contract_timesheet_detail_contract_id_fk
        foreign key (contract_id) references contract (id);
-- *