CREATE TABLE "public"."wukong_public_record" (
	 id bigserial PRIMARY KEY,
	 username VARCHAR(255) NOT NULL,
	 create_at int8 NOT NULL DEFAULT extract(epoch from now())
);