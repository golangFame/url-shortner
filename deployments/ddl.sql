-- public.urls definition

-- Drop table

-- DROP TABLE public.urls;

CREATE TABLE public.urls
(
    id         bigserial NOT NULL,
    urls       text      NOT NULL,
    short_urls text      NOT NULL,
    created_at timetz    NOT NULL,
    updated_at timetz    NOT NULL,
    deleted_at timetz    NULL
)
    PARTITION BY LIST (deleted_at);

CREATE TABLE public.url__active PARTITION OF public.urls FOR VALUES IN (NULL);
CREATE TABLE public.url_inactive PARTITION OF public.urls DEFAULT;

CREATE INDEX urls_id_idx ON ONLY public.urls USING btree (id, deleted_at);
CREATE INDEX urls_short_urls_idx ON ONLY public.urls USING btree (short_urls, deleted_at);
CREATE INDEX urls_urls_idx ON ONLY public.urls USING btree (urls, deleted_at);

