/*
  Warnings:

  - You are about to drop the column `contract_id` on the `document` table. All the data in the column will be lost.
  - You are about to drop the `contract` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `lease_id` to the `document` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "contract" DROP CONSTRAINT "contract_property_id_fkey";

-- DropForeignKey
ALTER TABLE "contract" DROP CONSTRAINT "contract_tenant_id_fkey";

-- DropForeignKey
ALTER TABLE "document" DROP CONSTRAINT "document_contract_id_fkey";

-- AlterTable
ALTER TABLE "document" DROP COLUMN "contract_id",
ADD COLUMN     "lease_id" TEXT NOT NULL;

-- DropTable
DROP TABLE "contract";

-- CreateTable
CREATE TABLE "lease" (
    "id" TEXT NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "start_date" TIMESTAMP(3) NOT NULL,
    "end_date" TIMESTAMP(3),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "tenant_id" TEXT NOT NULL,
    "property_id" TEXT NOT NULL,

    CONSTRAINT "lease_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "lease" ADD CONSTRAINT "lease_tenant_id_fkey" FOREIGN KEY ("tenant_id") REFERENCES "user"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "lease" ADD CONSTRAINT "lease_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "document" ADD CONSTRAINT "document_lease_id_fkey" FOREIGN KEY ("lease_id") REFERENCES "lease"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
