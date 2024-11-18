/*
  Warnings:

  - A unique constraint covering the columns `[name,owner_id]` on the table `property` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "property_name_owner_id_key" ON "property"("name", "owner_id");
