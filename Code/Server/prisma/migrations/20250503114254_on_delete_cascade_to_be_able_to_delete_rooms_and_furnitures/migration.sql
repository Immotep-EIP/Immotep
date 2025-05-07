-- DropForeignKey
ALTER TABLE "damage" DROP CONSTRAINT "damage_room_id_fkey";

-- DropForeignKey
ALTER TABLE "furniture" DROP CONSTRAINT "furniture_room_id_fkey";

-- DropForeignKey
ALTER TABLE "furnitureState" DROP CONSTRAINT "furnitureState_furniture_id_fkey";

-- DropForeignKey
ALTER TABLE "furnitureState" DROP CONSTRAINT "furnitureState_report_id_fkey";

-- DropForeignKey
ALTER TABLE "room" DROP CONSTRAINT "room_property_id_fkey";

-- DropForeignKey
ALTER TABLE "roomState" DROP CONSTRAINT "roomState_report_id_fkey";

-- DropForeignKey
ALTER TABLE "roomState" DROP CONSTRAINT "roomState_room_id_fkey";

-- AddForeignKey
ALTER TABLE "damage" ADD CONSTRAINT "damage_room_id_fkey" FOREIGN KEY ("room_id") REFERENCES "room"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "room" ADD CONSTRAINT "room_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "property"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "furniture" ADD CONSTRAINT "furniture_room_id_fkey" FOREIGN KEY ("room_id") REFERENCES "room"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "roomState" ADD CONSTRAINT "roomState_report_id_fkey" FOREIGN KEY ("report_id") REFERENCES "inventoryReport"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "roomState" ADD CONSTRAINT "roomState_room_id_fkey" FOREIGN KEY ("room_id") REFERENCES "room"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "furnitureState" ADD CONSTRAINT "furnitureState_report_id_fkey" FOREIGN KEY ("report_id") REFERENCES "inventoryReport"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "furnitureState" ADD CONSTRAINT "furnitureState_furniture_id_fkey" FOREIGN KEY ("furniture_id") REFERENCES "furniture"("id") ON DELETE CASCADE ON UPDATE CASCADE;
