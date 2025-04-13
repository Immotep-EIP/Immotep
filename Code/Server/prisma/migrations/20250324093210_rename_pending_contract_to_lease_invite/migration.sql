/*
  Warnings:

  - You are about to drop the `pendingContract` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "pendingContract" DROP CONSTRAINT "pendingContract_property_id_fkey";

-- DropTable
DROP TABLE "pendingContract";

-- CreateTable
CREATE TABLE "leaseInvite" (
    "id" TEXT NOT NULL,
    "tenant_email" VARCHAR(255) NOT NULL,
    "start_date" TIMESTAMP(3) NOT NULL,
    "end_date" TIMESTAMP(3),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "property_id" TEXT NOT NULL,

    CONSTRAINT "leaseInvite_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "leaseInvite_tenant_email_key" ON "leaseInvite"("tenant_email");

-- CreateIndex
CREATE UNIQUE INDEX "leaseInvite_property_id_key" ON "leaseInvite"("property_id");

-- AddForeignKey
ALTER TABLE "leaseInvite" ADD CONSTRAINT "leaseInvite_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
