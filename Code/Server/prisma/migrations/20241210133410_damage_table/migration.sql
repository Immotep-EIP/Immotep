-- CreateEnum
CREATE TYPE "priority" AS ENUM ('low', 'medium', 'high', 'urgent');

-- CreateTable
CREATE TABLE "damage" (
    "id" TEXT NOT NULL,
    "comment" TEXT NOT NULL,
    "room" TEXT NOT NULL,
    "priority" "priority" NOT NULL,
    "read" BOOLEAN NOT NULL DEFAULT false,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "fixed_at" TIMESTAMP(3),
    "property_id" TEXT NOT NULL,

    CONSTRAINT "damage_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "damage" ADD CONSTRAINT "damage_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "property"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
