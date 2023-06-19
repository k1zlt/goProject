Create TABLE IF NOT EXISTS "user"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "surname" VARCHAR(255) NOT NULL,
    "age" INTEGER
);

Create TABLE IF NOT EXISTS "category"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL
);

Create TABLE IF NOT EXISTS "lesson"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "content" TEXT NOT NULL,
    "category_id" INTEGER REFERENCES "category" ("id")
);

CREATE TABLE IF NOT EXISTS "CompletedLesson"(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER REFERENCES "user" ("id"),
    "lesson_id" INTEGER REFERENCES "lesson" ("id")
);