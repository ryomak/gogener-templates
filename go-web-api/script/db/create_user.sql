drop database if exists [[.AppName]];
create database [[.AppName]];

create user if not exists '[[.AppName]]' identified by '[[.AppName]]';

grant all on [[.AppName]].* to '[[.AppName]]' with grant option;

flush privileges;