-- AlterTable
ALTER TABLE "furniture" ADD COLUMN     "archived" BOOLEAN NOT NULL DEFAULT false;

-- AlterTable
ALTER TABLE "property" ADD COLUMN     "archived" BOOLEAN NOT NULL DEFAULT false;

-- AlterTable
ALTER TABLE "room" ADD COLUMN     "archived" BOOLEAN NOT NULL DEFAULT false;
