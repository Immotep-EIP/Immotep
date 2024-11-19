/*
  Warnings:

  - The primary key for the `contract` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `endDate` on the `contract` table. All the data in the column will be lost.
  - You are about to drop the column `propertyId` on the `contract` table. All the data in the column will be lost.
  - You are about to drop the column `startDate` on the `contract` table. All the data in the column will be lost.
  - You are about to drop the column `tenantId` on the `contract` table. All the data in the column will be lost.
  - You are about to drop the column `endDate` on the `pendingContract` table. All the data in the column will be lost.
  - You are about to drop the column `propertyId` on the `pendingContract` table. All the data in the column will be lost.
  - You are about to drop the column `startDate` on the `pendingContract` table. All the data in the column will be lost.
  - You are about to drop the column `tenantEmail` on the `pendingContract` table. All the data in the column will be lost.
  - You are about to drop the column `ownerId` on the `property` table. All the data in the column will be lost.
  - You are about to drop the column `postalCode` on the `property` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[tenant_email]` on the table `pendingContract` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `property_id` to the `contract` table without a default value. This is not possible if the table is not empty.
  - Added the required column `start_date` to the `contract` table without a default value. This is not possible if the table is not empty.
  - Added the required column `tenant_id` to the `contract` table without a default value. This is not possible if the table is not empty.
  - Added the required column `property_id` to the `pendingContract` table without a default value. This is not possible if the table is not empty.
  - Added the required column `start_date` to the `pendingContract` table without a default value. This is not possible if the table is not empty.
  - Added the required column `tenant_email` to the `pendingContract` table without a default value. This is not possible if the table is not empty.
  - Added the required column `owner_id` to the `property` table without a default value. This is not possible if the table is not empty.
  - Added the required column `postal_code` to the `property` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "contract" DROP CONSTRAINT "contract_propertyId_fkey";

-- DropForeignKey
ALTER TABLE "contract" DROP CONSTRAINT "contract_tenantId_fkey";

-- DropForeignKey
ALTER TABLE "pendingContract" DROP CONSTRAINT "pendingContract_propertyId_fkey";

-- DropForeignKey
ALTER TABLE "property" DROP CONSTRAINT "property_ownerId_fkey";

-- DropIndex
DROP INDEX "pendingContract_tenantEmail_key";

-- AlterTable
ALTER TABLE "contract" DROP CONSTRAINT "contract_pkey",
DROP COLUMN "endDate",
DROP COLUMN "propertyId",
DROP COLUMN "startDate",
DROP COLUMN "tenantId",
ADD COLUMN     "end_date" TIMESTAMP(3),
ADD COLUMN     "property_id" TEXT NOT NULL,
ADD COLUMN     "start_date" TIMESTAMP(3) NOT NULL,
ADD COLUMN     "tenant_id" TEXT NOT NULL,
ADD CONSTRAINT "contract_pkey" PRIMARY KEY ("tenant_id", "property_id");

-- AlterTable
ALTER TABLE "pendingContract" DROP COLUMN "endDate",
DROP COLUMN "propertyId",
DROP COLUMN "startDate",
DROP COLUMN "tenantEmail",
ADD COLUMN     "end_date" TIMESTAMP(3),
ADD COLUMN     "property_id" TEXT NOT NULL,
ADD COLUMN     "start_date" TIMESTAMP(3) NOT NULL,
ADD COLUMN     "tenant_email" VARCHAR(255) NOT NULL;

-- AlterTable
ALTER TABLE "property" DROP COLUMN "ownerId",
DROP COLUMN "postalCode",
ADD COLUMN     "owner_id" TEXT NOT NULL,
ADD COLUMN     "postal_code" TEXT NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX "pendingContract_tenant_email_key" ON "pendingContract"("tenant_email");

-- AddForeignKey
ALTER TABLE "contract" ADD CONSTRAINT "contract_tenant_id_fkey" FOREIGN KEY ("tenant_id") REFERENCES "user"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "contract" ADD CONSTRAINT "contract_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "property" ADD CONSTRAINT "property_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "user"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "pendingContract" ADD CONSTRAINT "pendingContract_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
