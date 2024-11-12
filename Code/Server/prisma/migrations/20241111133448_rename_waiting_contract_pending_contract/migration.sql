/*
  Warnings:

  - You are about to drop the `waitingContract` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "waitingContract" DROP CONSTRAINT "waitingContract_propertyId_fkey";

-- DropTable
DROP TABLE "waitingContract";

-- CreateTable
CREATE TABLE "pendingContract" (
    "id" TEXT NOT NULL,
    "tenantEmail" VARCHAR(255) NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "unlimited" BOOLEAN NOT NULL,
    "startDate" TIMESTAMP(3) NOT NULL,
    "endDate" TIMESTAMP(3),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "propertyId" TEXT NOT NULL,

    CONSTRAINT "pendingContract_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "pendingContract_tenantEmail_key" ON "pendingContract"("tenantEmail");

-- AddForeignKey
ALTER TABLE "pendingContract" ADD CONSTRAINT "pendingContract_propertyId_fkey" FOREIGN KEY ("propertyId") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
