/*
  Warnings:

  - A unique constraint covering the columns `[property_id]` on the table `pendingContract` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "pendingContract_property_id_key" ON "pendingContract"("property_id");
