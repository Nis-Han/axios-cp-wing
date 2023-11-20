-- +goose Up
ALTER TABLE "tasks" ADD FOREIGN KEY ("platform") REFERENCES "platform" ("id");

-- +goose Down
ALTER TABLE "tasks" DROP CONSTRAINT tasks_platform_fkey;