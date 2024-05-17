drop database if exists cend;
create database cend;
use cend;

create table academic_course (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	academic_level_id	int default 0,
	name	varchar(255) default '',
	constraint academic_course_pk
		unique (id)
);

alter table academic_course
	add constraint academic_course_academic_level_id_fk
		foreign key (academic_level_id) references academic_level (id);

create table academic_level (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	name	varchar(255) default '',
	constraint academic_level_pk
		unique (id)
);

create table authorization (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	level	int default 0,
	constraint authorization_pk
		unique (id)
);

alter table authorization
	add constraint authorization_user_id_fk
		foreign key (user_id) references user (id);

create table bearer_token (
	id int auto_increment primary key,
	t	varchar(255) default '',
	constraint bearer_token_pk
		unique (t)
);

create table media (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	file_name	varchar(255) default '',
	extension	varchar(255) default '',
	xid	varchar(255) default '',
	content_type	int default 0,
	constraint media_pk
		unique (id)
);

create table media_thumb (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	extension	varchar(255) default '',
	media_xid	varchar(255) default '',
	xid	varchar(255) default '',
	constraint media_thumb_pk
		unique (id)
);

alter table media_thumb
	add constraint media_thumb_media_xid_fk
		foreign key (media_xid) references media (xid);

create table password (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	hash	varchar(255) default '',
	constraint password_pk
		unique (id)
);

alter table password
	add constraint password_user_id_fk
		foreign key (user_id) references user (id);

create table user (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	name	varchar(255) default '',
	family_name	varchar(255) default '',
	nick_name	varchar(255) default '',
	email	varchar(255) default '',
	matricule	varchar(255) default '',
	age	int default 0,
	birth_date	datetime default '0000-00-00 00:00:00',
	sex	int default 0,
	lang	int default 0,
	status	int default 0,
	profile_image_xid	varchar(255) default '',
	description	varchar(255) default '',
	cover_text	varchar(255) default '',
	profile	varchar(255) default '',
	experience_detail	varchar(255) default '',
	additional_description	varchar(255) default '',
	add_on_title	varchar(255) default '',
	constraint user_pk
		unique (id)
);

alter table user
	add constraint user_profile_image_xid_fk
		foreign key (profile_image_xid) references media (xid);

create table user_academic_course (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	course_id	int default 0,
	constraint user_academic_course_pk
		unique (id)
);

alter table user_academic_course
	add constraint user_academic_course_user_id_fk
		foreign key (user_id) references user (id);

alter table user_academic_course
	add constraint user_academic_course_course_id_fk
		foreign key (course_id) references course (id);

create table user_academic_course_preference (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_academic_course_id	int default 0,
	is_online	bool default '',
	availability	datetime default '0000-00-00 00:00:00',
	constraint user_academic_course_preference_pk
		unique (id)
);

alter table user_academic_course_preference
	add constraint user_academic_course_preference_user_academic_course_id_fk
		foreign key (user_academic_course_id) references user_academic_course (id);

create table user_authorization_link (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	link_type	int default 0,
	constraint user_authorization_link_pk
		unique (id)
);

create table user_authorization_link_actor (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_authorization_link_id	int default 0,
	authorization_id	int default 0,
	constraint user_authorization_link_actor_pk
		unique (id)
);

alter table user_authorization_link_actor
	add constraint user_authorization_link_actor_user_authorization_link_id_fk
		foreign key (user_authorization_link_id) references user_authorization_link (id);

alter table user_authorization_link_actor
	add constraint user_authorization_link_actor_authorization_id_fk
		foreign key (authorization_id) references authorization (id);

create table user_media_detail (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	owner_id	int default 0,
	document_type	int default 0,
	document_xid	varchar(255) default '',
	constraint user_media_detail_pk
		unique (id)
);

alter table user_media_detail
	add constraint user_media_detail_owner_id_fk
		foreign key (owner_id) references user (id);

alter table user_media_detail
	add constraint user_media_detail_document_xid_fk
		foreign key (document_xid) references media (xid);

