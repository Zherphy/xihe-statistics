CREATE TABLE "public"."bigmodel_record" (
	 username VARCHAR(255) NOT NULL,
	 bigmodel VARCHAR(255) NOT NULL,
	 create_at int8 NOT NULL DEFAULT extract(epoch from now())
);