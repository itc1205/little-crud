-- +migrate Up

CREATE TABLE PROJECTS (
    id SERIAL NOT NULL CONSTRAINT projects_pkey PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE GOODS (
    id SERIAL NOT NULL CONSTRAINT goods_pkey PRIMARY KEY,
    project_id INT NOT NULL CONSTRAINT projects_id_fkey REFERENCES PROJECTS(id),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    priority INT
);

-- https://github.com/lib/pq/issues/974#issuecomment-642190369
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION goods_priority_insert_fnc() RETURNS trigger AS $$
    BEGIN
        IF NEW.priority IS NULL THEN
            UPDATE "GOODS"
            SET priority = (SELECT max(priority) FROM "GOODS") + 1
            WHERE id=NEW.id;
        END IF;
        RETURN NULL;
    END;
$$ 
LANGUAGE plpgsql;
-- +migrate StatementEnd


CREATE TRIGGER goods_priority_insert_trigger 
    AFTER INSERT 
    ON GOODS
    FOR EACH STATEMENT
    EXECUTE PROCEDURE goods_priority_insert_fnc();
    

-- +migrate Down
DROP FUNCTION goods_priority_insert_fnc();
DROP TRIGGER goods_priority_insert_trigger;
DROP TABLE PROJECTS;
DROP TABLE GOODS;