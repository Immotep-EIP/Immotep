/*
  Warnings:

  - You are about to drop the column `property_id` on the `damage` table. All the data in the column will be lost.
  - Added the required column `lease_id` to the `damage` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "damage" DROP CONSTRAINT "damage_property_id_fkey";

-- AlterTable
ALTER TABLE "damage" DROP COLUMN "property_id",
ADD COLUMN     "fix_planned_at" TIMESTAMP(3),
ADD COLUMN     "lease_id" TEXT NOT NULL;

-- CreateTable
CREATE TABLE "_damageToimage" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "_damageToimage_AB_unique" ON "_damageToimage"("A", "B");

-- CreateIndex
CREATE INDEX "_damageToimage_B_index" ON "_damageToimage"("B");

-- AddForeignKey
ALTER TABLE "damage" ADD CONSTRAINT "damage_lease_id_fkey" FOREIGN KEY ("lease_id") REFERENCES "lease"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_damageToimage" ADD CONSTRAINT "_damageToimage_A_fkey" FOREIGN KEY ("A") REFERENCES "damage"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_damageToimage" ADD CONSTRAINT "_damageToimage_B_fkey" FOREIGN KEY ("B") REFERENCES "image"("id") ON DELETE CASCADE ON UPDATE CASCADE;
