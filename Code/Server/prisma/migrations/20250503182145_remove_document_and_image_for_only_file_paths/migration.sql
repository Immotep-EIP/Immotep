/*
  Warnings:

  - You are about to drop the column `picture_id` on the `property` table. All the data in the column will be lost.
  - You are about to drop the column `profile_picture_id` on the `user` table. All the data in the column will be lost.
  - You are about to drop the `_damageToimage` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `_furnitureStateToimage` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `_imageToroomState` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `document` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `image` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "_damageToimage" DROP CONSTRAINT "_damageToimage_A_fkey";

-- DropForeignKey
ALTER TABLE "_damageToimage" DROP CONSTRAINT "_damageToimage_B_fkey";

-- DropForeignKey
ALTER TABLE "_furnitureStateToimage" DROP CONSTRAINT "_furnitureStateToimage_A_fkey";

-- DropForeignKey
ALTER TABLE "_furnitureStateToimage" DROP CONSTRAINT "_furnitureStateToimage_B_fkey";

-- DropForeignKey
ALTER TABLE "_imageToroomState" DROP CONSTRAINT "_imageToroomState_A_fkey";

-- DropForeignKey
ALTER TABLE "_imageToroomState" DROP CONSTRAINT "_imageToroomState_B_fkey";

-- DropForeignKey
ALTER TABLE "document" DROP CONSTRAINT "document_lease_id_fkey";

-- DropForeignKey
ALTER TABLE "property" DROP CONSTRAINT "property_picture_id_fkey";

-- DropForeignKey
ALTER TABLE "user" DROP CONSTRAINT "user_profile_picture_id_fkey";

-- AlterTable
ALTER TABLE "damage" ADD COLUMN     "pictures" TEXT[];

-- AlterTable
ALTER TABLE "furnitureState" ADD COLUMN     "pictures" TEXT[];

-- AlterTable
ALTER TABLE "lease" ADD COLUMN     "documents" TEXT[];

-- AlterTable
ALTER TABLE "property" DROP COLUMN "picture_id",
ADD COLUMN     "picture" TEXT;

-- AlterTable
ALTER TABLE "roomState" ADD COLUMN     "pictures" TEXT[];

-- AlterTable
ALTER TABLE "user" DROP COLUMN "profile_picture_id",
ADD COLUMN     "profile_picture" TEXT;

-- DropTable
DROP TABLE "_damageToimage";

-- DropTable
DROP TABLE "_furnitureStateToimage";

-- DropTable
DROP TABLE "_imageToroomState";

-- DropTable
DROP TABLE "document";

-- DropTable
DROP TABLE "image";
