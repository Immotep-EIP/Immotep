/*
  Warnings:

  - Added the required column `startDate` to the `contract` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "contract" ADD COLUMN     "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN     "startDate" TIMESTAMP(3) NOT NULL;

-- CreateTable
CREATE TABLE "waitingContract" (
    "id" TEXT NOT NULL,
    "tenantEmail" VARCHAR(255) NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "unlimited" BOOLEAN NOT NULL,
    "startDate" TIMESTAMP(3) NOT NULL,
    "endDate" TIMESTAMP(3),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "propertyId" TEXT NOT NULL,

    CONSTRAINT "waitingContract_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "waitingContract_tenantEmail_key" ON "waitingContract"("tenantEmail");

-- AddForeignKey
ALTER TABLE "waitingContract" ADD CONSTRAINT "waitingContract_propertyId_fkey" FOREIGN KEY ("propertyId") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
