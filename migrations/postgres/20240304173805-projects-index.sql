
-- +migrate Up

CREATE INDEX projects_index
ON "PROJECTS" (id);


-- +migrate Down

DROP INDEX projects_index;