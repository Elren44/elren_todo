DROP TABLE IF EXISTS tasks CASCADE;
DROP TABLE IF EXISTS users CASCADE;


CREATE TABLE public.tasks (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100) NOT NULL,
    date TIMESTAMP NOT NULL,
    description TEXT
);

INSERT INTO public.tasks (title, date, description) VALUES ('first', '2022-02-07 02:31', 'test description');
INSERT INTO public.tasks (title, date, description) VALUES ('second', '2022-02-07 13:20', 'test description');
INSERT INTO public.tasks (title, date, description) VALUES ('x', '2022-02-08 03:30', 'test description');
INSERT INTO public.tasks (title, date, description) VALUES ('y', '2022-02-09 17:00', 'test description');

SELECT * FROM tasks;


CREATE TABLE public.users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(100) NOT NULL,
    password_hash VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc')
);