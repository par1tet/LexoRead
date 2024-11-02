CREATE TABLE "books" (
    "id"          bigserial,
    "author"      text NOT NULL,
    "author_id"   bigint DEFAULT 0,
    "title"       text NOT NULL,
    "description" varchar(256),
    "likes"       bigint DEFAULT 0,
    "dis_likes"   bigint DEFAULT 0,
    "genre"       varchar(256),
    PRIMARY KEY ("id")
);

CREATE TABLE "files" (
    "id"          bigserial,
    "uploader_id" bigint NOT NULL,
    "file_url"    text NOT NULL,
    "for_id"      bigint NOT NULL,
    "for_type"    text DEFAULT 'book',
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_books_files" FOREIGN KEY ("for_id") REFERENCES "books" ("id") ON DELETE CASCADE
);

CREATE TABLE "comments" (
    "id"               bigserial,
    "book_id"         bigint NOT NULL,
    "content"         text,
    "user_id"         bigint DEFAULT 0,
    "likes"           bigint DEFAULT 0,
    "dislikes"        bigint DEFAULT 0,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_books_comments" FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE
);