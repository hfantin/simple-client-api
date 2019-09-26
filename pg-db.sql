
-- Drop table

-- DROP TABLE public.licenses;

CREATE TABLE public.cli (
	id SERIAL PRIMARY KEY,
	"name" varchar(100) NOT NULL,
	birth_date date NOT NULL,
	email varchar(100) NULL
);
