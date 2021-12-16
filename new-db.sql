create table if not exists coach(
    idCoach        int not null
        constraint coach_pk
            primary key,
    idGame         int not null,
    licenseDate    text(30),
    idEmployeeData int not null
);

create table game
(
    idGame integer not null
        constraint game_pk
            primary key,
    name   text(30)
);


create table if not exists tournament
(
    idTournament integer  not null
        constraint tournament_pk
            primary key,
    idPlace      integer  not null,
    idGame       integer  not null,
    date         text(30) not null,
    duration     integer
);

create table if not exists staff
(
    idStaff        integer not null
        constraint staff_pk
            primary key,
    idEmployeeData integer
);
create table if not exists player
(
    idPlayer       integer not null
        constraint player_pk
            primary key,
    idGame         integer not null,
    ranking        integer,
    idEmployeeData integer not null
);
create table if not exists place
(
    idPlace integer not null
        constraint place_pk
            primary key autoincrement,
    name    text(30),
    address text(30),
    city    text(30)
);
create table if not exists employee_data
(
    idEmployee integer not null
        constraint employee_data_pk
            primary key autoincrement,
    lastname   text(30),
    firstname  text(30),
    gender     text(30),
    age        integer,
    wage       int
);