/*
  Warnings:

  - Added the required column `type` to the `room` table without a default value. This is not possible if the table is not empty.

*/
-- CreateEnum
CREATE TYPE "roomType" AS ENUM ('dressing', 'laundryroom', 'bedroom', 'playroom', 'bathroom', 'toilet', 'livingroom', 'diningroom', 'kitchen', 'hallway', 'balcony', 'cellar', 'garage', 'storage', 'office', 'other');

-- AlterTable
ALTER TABLE "room" ADD COLUMN     "type" "roomType" NOT NULL;
