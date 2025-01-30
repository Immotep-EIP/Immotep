-- CreateEnum
CREATE TYPE "reportType" AS ENUM ('start', 'middle', 'end');

-- CreateEnum
CREATE TYPE "state" AS ENUM ('broken', 'needsRepair', 'bad', 'medium', 'good', 'new');

-- CreateEnum
CREATE TYPE "cleanliness" AS ENUM ('dirty', 'medium', 'clean');

-- CreateTable
CREATE TABLE "room" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "property_id" TEXT NOT NULL,

    CONSTRAINT "room_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "furniture" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "quantity" INTEGER NOT NULL DEFAULT 1,
    "room_id" TEXT NOT NULL,

    CONSTRAINT "furniture_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "inventoryReport" (
    "id" TEXT NOT NULL,
    "date" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "type" "reportType" NOT NULL,
    "property_id" TEXT NOT NULL,

    CONSTRAINT "inventoryReport_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "roomState" (
    "cleanliness" "cleanliness" NOT NULL,
    "state" "state" NOT NULL,
    "note" TEXT NOT NULL,
    "report_id" TEXT NOT NULL,
    "room_id" TEXT NOT NULL,

    CONSTRAINT "roomState_pkey" PRIMARY KEY ("report_id","room_id")
);

-- CreateTable
CREATE TABLE "furnitureState" (
    "cleanliness" "cleanliness" NOT NULL,
    "state" "state" NOT NULL,
    "note" TEXT NOT NULL,
    "report_id" TEXT NOT NULL,
    "furniture_id" TEXT NOT NULL,

    CONSTRAINT "furnitureState_pkey" PRIMARY KEY ("report_id","furniture_id")
);

-- AddForeignKey
ALTER TABLE "room" ADD CONSTRAINT "room_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "furniture" ADD CONSTRAINT "furniture_room_id_fkey" FOREIGN KEY ("room_id") REFERENCES "room"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "inventoryReport" ADD CONSTRAINT "inventoryReport_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "roomState" ADD CONSTRAINT "roomState_report_id_fkey" FOREIGN KEY ("report_id") REFERENCES "inventoryReport"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "roomState" ADD CONSTRAINT "roomState_room_id_fkey" FOREIGN KEY ("room_id") REFERENCES "room"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "furnitureState" ADD CONSTRAINT "furnitureState_report_id_fkey" FOREIGN KEY ("report_id") REFERENCES "inventoryReport"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "furnitureState" ADD CONSTRAINT "furnitureState_furniture_id_fkey" FOREIGN KEY ("furniture_id") REFERENCES "furniture"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
