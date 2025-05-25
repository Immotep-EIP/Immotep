/*
  Warnings:

  - You are about to drop the column `property_id` on the `inventoryReport` table. All the data in the column will be lost.
  - Added the required column `lease_id` to the `inventoryReport` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "inventoryReport" DROP CONSTRAINT "inventoryReport_property_id_fkey";

-- AlterTable
ALTER TABLE "inventoryReport" DROP COLUMN "property_id",
ADD COLUMN     "lease_id" TEXT NOT NULL;

-- AddForeignKey
ALTER TABLE "inventoryReport" ADD CONSTRAINT "inventoryReport_lease_id_fkey" FOREIGN KEY ("lease_id") REFERENCES "lease"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
