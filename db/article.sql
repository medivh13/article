CREATE TABLE article (
    id bigint NOT NULL,
    author text,
    title text,
    body text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);

CREATE SEQUENCE article_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE article_id_seq OWNED BY article.id;

ALTER TABLE ONLY article ALTER COLUMN id SET DEFAULT nextval('article_id_seq'::regclass);

ALTER TABLE ONLY article
    ADD CONSTRAINT article_pkey PRIMARY KEY (id);


CREATE INDEX articleindex ON article USING btree (deleted_at, id);


