-- CreateTable
CREATE TABLE "imageONroomstate" (
    "roomstate_reportid" TEXT NOT NULL,
    "roomstate_roomid" TEXT NOT NULL,
    "image_id" TEXT NOT NULL,

    CONSTRAINT "imageONroomstate_pkey" PRIMARY KEY ("roomstate_reportid","roomstate_roomid","image_id")
);

-- CreateTable
CREATE TABLE "imageONfurniturestate" (
    "furniturestate_reportid" TEXT NOT NULL,
    "furniturestate_furnitureid" TEXT NOT NULL,
    "image_id" TEXT NOT NULL,

    CONSTRAINT "imageONfurniturestate_pkey" PRIMARY KEY ("furniturestate_reportid","furniturestate_furnitureid","image_id")
);

-- AddForeignKey
ALTER TABLE "imageONroomstate" ADD CONSTRAINT "imageONroomstate_roomstate_reportid_roomstate_roomid_fkey" FOREIGN KEY ("roomstate_reportid", "roomstate_roomid") REFERENCES "roomState"("report_id", "room_id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "imageONroomstate" ADD CONSTRAINT "imageONroomstate_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "image"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "imageONfurniturestate" ADD CONSTRAINT "imageONfurniturestate_furniturestate_reportid_furnituresta_fkey" FOREIGN KEY ("furniturestate_reportid", "furniturestate_furnitureid") REFERENCES "furnitureState"("report_id", "furniture_id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "imageONfurniturestate" ADD CONSTRAINT "imageONfurniturestate_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "image"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
