use cend;

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


create table user_academic_course_preference
(
    id             int auto_increment primary key,
    created_at     timestamp default CURRENT_TIMESTAMP,
    updated_at     timestamp default CURRENT_TIMESTAMP,
    deleted_at     timestamp default '0000-00-00 00:00:00',
    user_academic_course_id int default 0,
    is_online      boolean,
    availability datetime
);

alter table user_academic_course_preference
    add constraint user_course_preference_user_academic_course_id_fk
        foreign key (user_academic_course_id) references user_academic_course_preference (id);



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