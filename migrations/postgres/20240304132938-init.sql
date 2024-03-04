-- +migrate Up

CREATE TABLE "PROJECTS" (
    id SERIAL NOT NULL CONSTRAINT projects_pkey PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE "GOODS" (
    id SERIAL NOT NULL CONSTRAINT goods_pkey PRIMARY KEY,
    project_id INT NOT NULL CONSTRAINT projects_id_fkey REFERENCES "PROJECTS"(id),
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    priority INT
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION goods_priority_insert_fnc() RETURNS trigger AS $$
    BEGIN
        UPDATE "GOODS"
        -- Firstly select maximal priority from goods, 
        -- if not found we set it to 0, 
        -- and then increment by one
        SET priority = COALESCE((SELECT max(priority) FROM "GOODS" AS INT), 0) + 1
        WHERE id=NEW.id;
        RETURN NULL;
    END;
$$ 
LANGUAGE plpgsql;
-- +migrate StatementEnd


CREATE TRIGGER goods_priority_insert_trigger 
    AFTER INSERT 
    ON "GOODS"
    FOR EACH ROW
    WHEN (NEW.priority IS NULL)
    EXECUTE PROCEDURE goods_priority_insert_fnc();

INSERT INTO "PROJECTS" (name) VALUES 
    ('My first project!');

-- +migrate Down
DROP TRIGGER goods_priority_insert_trigger ON "GOODS";
DROP FUNCTION goods_priority_insert_fnc();
DROP TABLE "GOODS";
DROP TABLE "PROJECTS";