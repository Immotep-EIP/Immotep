-- AlterTable
ALTER TABLE "property" ADD COLUMN     "apartment_number" TEXT,
ALTER COLUMN "deposit_price" SET DATA TYPE DOUBLE PRECISION,
ALTER COLUMN "rental_price_per_month" SET DATA TYPE DOUBLE PRECISION;
