CREATE TABLE IF NOT EXISTS public.countries
(
    id   SERIAL NOT NULL UNIQUE PRIMARY KEY,
    name TEXT   NOT NULL
);

CREATE TABLE IF NOT EXISTS public.cities
(
    id         SERIAL NOT NULL UNIQUE PRIMARY KEY,
    name       TEXT   NOT NULL,
    country_id INT    NOT NULL,
    FOREIGN KEY (country_id) REFERENCES public.countries (id)
);

CREATE TABLE IF NOT EXISTS public.addresses
(
    id       SERIAL NOT NULL UNIQUE PRIMARY KEY,
    address  TEXT   NOT NULL,
    zip_code TEXT   NOT NULL,
    city_id  INT    NOT NULL,
    FOREIGN KEY (city_id) REFERENCES public.cities (id)
);