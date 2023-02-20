ALTER TABLE "public"."co_fd_account"
DROP COLUMN "created_at",
  DROP COLUMN "updated_at",
  DROP COLUMN "deleted_at",
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "created_by" int8,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "updated_by" int8,
  ADD COLUMN "deleted_at" timestamp,
  ADD COLUMN "deleted_by" int8;



ALTER TABLE "public"."co_fd_account_bill"
DROP COLUMN "created_at",
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "created_by" int8;


ALTER TABLE "public"."co_fd_bank_card"
DROP COLUMN "created_at",
  DROP COLUMN "updated_at",
  DROP COLUMN "deleted_at",
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "created_by" int8,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "updated_by" int8,
  ADD COLUMN "deleted_at" timestamp,
  ADD COLUMN "deleted_by" int8;



ALTER TABLE "public"."co_fd_invoice"
DROP COLUMN "created_at",
  DROP COLUMN "updated_at",
  DROP COLUMN "deleted_at",
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "created_by" int8,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "updated_by" int8,
  ADD COLUMN "deleted_at" timestamp,
  ADD COLUMN "deleted_by" int8;


ALTER TABLE "public"."co_fd_invoice_detail"
DROP COLUMN "created_at",
  DROP COLUMN "updated_at",
  DROP COLUMN "deleted_at",
  ADD COLUMN "created_at" timestamp,
  ADD COLUMN "created_by" int8,
  ADD COLUMN "updated_at" timestamp,
  ADD COLUMN "updated_by" int8,
  ADD COLUMN "deleted_at" timestamp,
  ADD COLUMN "deleted_by" int8;
