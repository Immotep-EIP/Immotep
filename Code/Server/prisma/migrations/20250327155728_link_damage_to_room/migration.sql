/*
  Warnings:

  - You are about to drop the column `room` on the `damage` table. All the data in the column will be lost.
  - Added the required column `room_id` to the `damage` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "damage" DROP COLUMN "room",
ADD COLUMN     "room_id" TEXT NOT NULL;

-- AddForeignKey
ALTER TABLE "damage" ADD CONSTRAINT "damage_room_id_fkey" FOREIGN KEY ("room_id") REFERENCES "room"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
