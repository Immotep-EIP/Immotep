/*
  Warnings:

  - Added the required column `type` to the `document` table without a default value. This is not possible if the table is not empty.
  - Added the required column `type` to the `image` table without a default value. This is not possible if the table is not empty.

*/
-- CreateEnum
CREATE TYPE "docType" AS ENUM ('pdf', 'docx', 'xlsx');

-- CreateEnum
CREATE TYPE "imageType" AS ENUM ('png', 'jpeg');

-- AlterTable: Add column with default
ALTER TABLE "document" ADD COLUMN "type" "docType" NOT NULL DEFAULT 'pdf';

-- AlterTable: Add column with default
ALTER TABLE "image" ADD COLUMN "type" "imageType" NOT NULL DEFAULT 'jpeg';

-- Remove default constraint from "document"."type"
ALTER TABLE "document" ALTER COLUMN "type" DROP DEFAULT;

-- Remove default constraint from "image"."type"
ALTER TABLE "image" ALTER COLUMN "type" DROP DEFAULT;
