-- 1. grant grs
-- create database grs;
-- create user grs with password 'grs';
-- grant all privileges on database grs to grs;

-- 2. migrate grs
create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);
