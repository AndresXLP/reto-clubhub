CREATE TABLE IF NOT EXISTS public.owners
(
    id         SERIAL NOT NULL UNIQUE PRIMARY KEY,
    first_name TEXT   NOT NULL,
    last_name  TEXT   NOT NULL,
    email      TEXT   NOT NULL,
    phone      TEXT   NOT NULL,
    address_id INT    NOT NULL,
    FOREIGN KEY (address_id) REFERENCES public.addresses (id)
);

CREATE TABLE IF NOT EXISTS public.companies
(
    id         SERIAL NOT NULL UNIQUE PRIMARY KEY,
    name       TEXT   NOT NULL UNIQUE,
    tax_number TEXT   NOT NULL,
    owner_id   INT    NOT NULL,
    address_id INT    NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES public.owners (id),
    FOREIGN KEY (address_id) REFERENCES public.addresses (id)
);

CREATE TABLE IF NOT EXISTS public.franchises
(
    id         SERIAL NOT NULL UNIQUE PRIMARY KEY,
    name       TEXT   NOT NULL UNIQUE,
    URL        TEXT   NOT NULL,
    address_id INT    NOT NULL,
    company_id INT    NOT NULL,
    FOREIGN KEY (address_id) REFERENCES public.addresses (id),
    FOREIGN KEY (company_id) REFERENCES public.companies (id)
);

CREATE TABLE IF NOT EXISTS public.additional_franchises_info
(
    franchise_id            INT NOT NULL,
    url_image               TEXT,
    protocol                TEXT,
    trace_routes            text[],
    domain_created_at       TIMESTAMP,
    domain_expired_at       TIMESTAMP,
    domain_registrant_name  TEXT,
    domain_registrant_email TEXT,
    FOREIGN KEY (franchise_id) references public.franchises (id)
);