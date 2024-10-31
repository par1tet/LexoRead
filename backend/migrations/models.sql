CREATE TABLE
    "books" (
        "id" bigserial,
        "author" text NOT NULL,
        "author_id" bigint DEFAULT 0,
        "title" text NOT NULL,
        "description" varchar(256),
        "likes" bigint DEFAULT 0,
        "dis_likes" bigint DEFAULT 0,
        PRIMARY KEY ("id")
    )
CREATE TABLE
    "files" (
        "id" bigserial,
        "uploader_id" bigint NOT NULL,
        "file_url" text NOT NULL,
        "for_id" bigint NOT NULL,
        "for_type" text DEFAULT 'book',
        PRIMARY KEY ("id"),
        CONSTRAINT "fk_books_files" FOREIGN KEY ("for_id") REFERENCES "books" ("id") ON DELETE CASCADE
    )
CREATE TABLE
    "comments" (
        "id" bigserial,
        "book_id" bigint NOT NULL,
        "content" varchar(256),
        "user_id" bigint DEFAULT 0,
        "likes" bigint DEFAULT 0,
        "dis_likes" bigint DEFAULT 0,
        "reply_comment_id" bigint DEFAULT null,
        PRIMARY KEY ("id"),
        CONSTRAINT "fk_comments_reply_comments" FOREIGN KEY ("reply_comment_id") REFERENCES "comments" ("id"),
        CONSTRAINT "fk_books_comments" FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE
    )
CREATE TABLE
    "users" (
        "id" bigserial,
        "isBanned" BOOLEAN,
        "username" VARCHAR(256),
        "email" VARCHAR(256),
        "description" TEXT,
        "hashPassword" TEXT,
        "avatarFileUrl" TEXT,
        "favBooks" BIGINT[] DEFAULT '{}',
        PRIMARY KEY ("id"),
    )