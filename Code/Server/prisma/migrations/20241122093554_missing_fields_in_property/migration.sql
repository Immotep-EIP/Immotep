/*
  Warnings:

  - Added the required column `areaSqMeters` to the `property` table without a default value. This is not possible if the table is not empty.
  - Added the required column `depositPrice` to the `property` table without a default value. This is not possible if the table is not empty.
  - Added the required column `picture` to the `property` table without a default value. This is not possible if the table is not empty.
  - Added the required column `rentalPricePerMonth` to the `property` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "property" ADD COLUMN     "areaSqMeters" DOUBLE PRECISION NOT NULL,
ADD COLUMN     "depositPrice" INTEGER NOT NULL,
ADD COLUMN     "picture" TEXT NOT NULL,
ADD COLUMN     "rentalPricePerMonth" INTEGER NOT NULL;
