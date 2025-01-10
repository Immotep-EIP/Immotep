/*
  Warnings:

  - You are about to drop the column `picture` on the `property` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "property" DROP COLUMN "picture",
ADD COLUMN     "picture_id" TEXT;

-- AlterTable
ALTER TABLE "user" ADD COLUMN     "profile_picture_id" TEXT;

-- CreateTable
CREATE TABLE "image" (
    "id" TEXT NOT NULL,
    "data" BYTEA NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "image_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "user" ADD CONSTRAINT "user_profile_picture_id_fkey" FOREIGN KEY ("profile_picture_id") REFERENCES "image"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "property" ADD CONSTRAINT "property_picture_id_fkey" FOREIGN KEY ("picture_id") REFERENCES "image"("id") ON DELETE SET NULL ON UPDATE CASCADE;
