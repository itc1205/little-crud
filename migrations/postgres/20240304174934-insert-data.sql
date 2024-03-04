
-- +migrate Up

INSERT INTO "PROJECTS" (name) VALUES 
    ('Певрая запись');

-- +migrate Down

-- Firstly delete all goods of the project
DELETE FROM "GOODS"
WHERE project_id = (SELECT id FROM "PROJECTS" WHERE name='Первая запись');

-- Then delete project itself
DELETE FROM "PROJECTS"
WHERE name='Первая запись';