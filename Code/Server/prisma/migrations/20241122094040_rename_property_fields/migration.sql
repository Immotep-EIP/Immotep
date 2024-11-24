/*
  Warnings:

  - You are about to drop the column `areaSqMeters` on the `property` table. All the data in the column will be lost.
  - You are about to drop the column `depositPrice` on the `property` table. All the data in the column will be lost.
  - You are about to drop the column `rentalPricePerMonth` on the `property` table. All the data in the column will be lost.
  - Added the required column `area_sqm` to the `property` table without a default value. This is not possible if the table is not empty.
  - Added the required column `deposit_price` to the `property` table without a default value. This is not possible if the table is not empty.
  - Added the required column `rental_price_per_month` to the `property` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "property" DROP COLUMN "areaSqMeters",
DROP COLUMN "depositPrice",
DROP COLUMN "rentalPricePerMonth",
ADD COLUMN     "area_sqm" DOUBLE PRECISION NOT NULL,
ADD COLUMN     "deposit_price" INTEGER NOT NULL,
ADD COLUMN     "rental_price_per_month" INTEGER NOT NULL;
