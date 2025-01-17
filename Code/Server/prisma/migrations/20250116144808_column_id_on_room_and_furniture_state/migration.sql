/*
  Warnings:

  - The primary key for the `furnitureState` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - The primary key for the `roomState` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the `imageONfurniturestate` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `imageONroomstate` table. If the table is not empty, all the data it contains will be lost.
  - A unique constraint covering the columns `[report_id,furniture_id]` on the table `furnitureState` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[report_id,room_id]` on the table `roomState` will be added. If there are existing duplicate values, this will fail.
  - The required column `id` was added to the `furnitureState` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.
  - The required column `id` was added to the `roomState` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.

*/
-- DropForeignKey
ALTER TABLE "imageONfurniturestate" DROP CONSTRAINT "imageONfurniturestate_furniturestate_reportid_furnituresta_fkey";

-- DropForeignKey
ALTER TABLE "imageONfurniturestate" DROP CONSTRAINT "imageONfurniturestate_image_id_fkey";

-- DropForeignKey
ALTER TABLE "imageONroomstate" DROP CONSTRAINT "imageONroomstate_image_id_fkey";

-- DropForeignKey
ALTER TABLE "imageONroomstate" DROP CONSTRAINT "imageONroomstate_roomstate_reportid_roomstate_roomid_fkey";

-- AlterTable
ALTER TABLE "furnitureState" DROP CONSTRAINT "furnitureState_pkey",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "furnitureState_pkey" PRIMARY KEY ("id");

-- AlterTable
ALTER TABLE "roomState" DROP CONSTRAINT "roomState_pkey",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "roomState_pkey" PRIMARY KEY ("id");

-- DropTable
DROP TABLE "imageONfurniturestate";

-- DropTable
DROP TABLE "imageONroomstate";

-- CreateTable
CREATE TABLE "_imageToroomState" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "_furnitureStateToimage" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "_imageToroomState_AB_unique" ON "_imageToroomState"("A", "B");

-- CreateIndex
CREATE INDEX "_imageToroomState_B_index" ON "_imageToroomState"("B");

-- CreateIndex
CREATE UNIQUE INDEX "_furnitureStateToimage_AB_unique" ON "_furnitureStateToimage"("A", "B");

-- CreateIndex
CREATE INDEX "_furnitureStateToimage_B_index" ON "_furnitureStateToimage"("B");

-- CreateIndex
CREATE UNIQUE INDEX "furnitureState_report_id_furniture_id_key" ON "furnitureState"("report_id", "furniture_id");

-- CreateIndex
CREATE UNIQUE INDEX "roomState_report_id_room_id_key" ON "roomState"("report_id", "room_id");

-- AddForeignKey
ALTER TABLE "_imageToroomState" ADD CONSTRAINT "_imageToroomState_A_fkey" FOREIGN KEY ("A") REFERENCES "image"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_imageToroomState" ADD CONSTRAINT "_imageToroomState_B_fkey" FOREIGN KEY ("B") REFERENCES "roomState"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_furnitureStateToimage" ADD CONSTRAINT "_furnitureStateToimage_A_fkey" FOREIGN KEY ("A") REFERENCES "furnitureState"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_furnitureStateToimage" ADD CONSTRAINT "_furnitureStateToimage_B_fkey" FOREIGN KEY ("B") REFERENCES "image"("id") ON DELETE CASCADE ON UPDATE CASCADE;
