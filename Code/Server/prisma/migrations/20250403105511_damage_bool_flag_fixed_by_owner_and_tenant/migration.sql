-- AlterTable
ALTER TABLE "damage" ADD COLUMN     "fixed_owner" BOOLEAN NOT NULL DEFAULT false,
ADD COLUMN     "fixed_tenant" BOOLEAN NOT NULL DEFAULT false;
