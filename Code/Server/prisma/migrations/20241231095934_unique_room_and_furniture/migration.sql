/*
  Warnings:

  - A unique constraint covering the columns `[name,room_id]` on the table `furniture` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[name,property_id]` on the table `room` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateIndex
CREATE UNIQUE INDEX "furniture_name_room_id_key" ON "furniture"("name", "room_id");

-- CreateIndex
CREATE UNIQUE INDEX "room_name_property_id_key" ON "room"("name", "property_id");
