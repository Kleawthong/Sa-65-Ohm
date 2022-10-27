BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "scholar_admins" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "scholar_statuses" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"status_name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "scholar_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"type_name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "scholarships" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"scholar_name"	text,
	"scholar_admin_id"	integer,
	"scholar_status_id"	integer,
	"scholar_type_id"	integer,
	"scholar_detail"	text,
	CONSTRAINT "fk_scholar_admins_scholarships" FOREIGN KEY("scholar_admin_id") REFERENCES "scholar_admins"("id"),
	CONSTRAINT "fk_scholar_statuses_scholarships" FOREIGN KEY("scholar_status_id") REFERENCES "scholar_statuses"("id"),
	CONSTRAINT "fk_scholar_types_scholarships" FOREIGN KEY("scholar_type_id") REFERENCES "scholar_types"("id"),
	PRIMARY KEY("id")
);
CREATE INDEX IF NOT EXISTS "idx_scholar_admins_deleted_at" ON "scholar_admins" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_scholar_admins_email" ON "scholar_admins" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_scholar_statuses_deleted_at" ON "scholar_statuses" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_scholar_types_deleted_at" ON "scholar_types" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_scholarships_deleted_at" ON "scholarships" (
	"deleted_at"
);
COMMIT;
