ALTER TABLE "tasks" DROP CONSTRAINT IF EXISTS "tasks_pomodoro_id_fkey";
ALTER TABLE "tasks" DROP CONSTRAINT IF EXISTS "tasks_todo_id_fkey";
ALTER TABLE "todos" DROP CONSTRAINT IF EXISTS "todos_user_id_fkey";

DROP TABLE IF EXISTS "tasks";
DROP TABLE IF EXISTS "todos";
DROP TABLE IF EXISTS "pomodoros";
DROP TABLE IF EXISTS "users";

DROP EXTENSION IF EXISTS "uuid-ossp";