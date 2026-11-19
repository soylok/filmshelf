create schema filmshelf;

-- movies
create table filmshelf.movies (
	id bigint primary key generated always as identity,
	title varchar(255) not null,
	description text,
	release_year integer,
	rating numeric(3, 1),
	budget bigint,
	worldwide_box_office bigint,
	duration_minutes integer,
	poster_filename varchar(255),

	check (length(trim(title)) > 0),
	check (rating between 0 and 10),
	check (budget >= 0),
	check (worldwide_box_office >= 0),
	check (duration_minutes > 0),
	check (poster_filename is null or length(trim(poster_filename)) > 0)
);

-- countries
create table filmshelf.countries (
	id bigint primary key generated always as identity,
	code varchar(2) not null unique,

	check (code ~ '^[A-Z]{2}$')
);

create table filmshelf.movie_countries (
	movie_id bigint references filmshelf.movies(id) on delete cascade,
	country_id bigint references filmshelf.countries(id) on delete cascade,

	primary key (movie_id, country_id)
);

-- genres
create table filmshelf.genres (
	id bigint primary key generated always as identity,
	name varchar(100) not null unique,

	check (length(trim(name)) > 0)
);

create table filmshelf.movie_genres (
	movie_id bigint references filmshelf.movies(id) on delete cascade,
	genre_id bigint references filmshelf.genres(id) on delete cascade,

	primary key (movie_id, genre_id)
);

-- directors
create table filmshelf.directors (
	id bigint primary key generated always as identity,
	first_name varchar(100) not null,
	last_name varchar(100) not null,
	middle_name varchar(100),

	check(length(trim(first_name)) > 0),
	check(length(trim(last_name)) > 0),
	check(middle_name is null or length(trim(middle_name)) > 0)
);

create table filmshelf.movie_directors (
	movie_id bigint references filmshelf.movies(id) on delete cascade,
	director_id bigint references filmshelf.directors(id) on delete cascade,

	primary key (movie_id, director_id)
);

alter table filmshelf.movies
add column version integer default 1 not null;

alter table filmshelf.genres
add column version integer default 1 not null;

alter table filmshelf.directors
add column version integer default 1 not null;
