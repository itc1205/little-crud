
-- +migrate Up

CREATE INDEX goods_index
ON "GOODS" (id, project_id, name);

-- +migrate Down

DROP INDEX goods_index;