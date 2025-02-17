/*
  Warnings:

  - The primary key for the `contract` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `contract_property_id` on the `document` table. All the data in the column will be lost.
  - You are about to drop the column `contract_tenant_id` on the `document` table. All the data in the column will be lost.
  - The required column `id` was added to the `contract` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.
  - Added the required column `contract_id` to the `document` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "document" DROP CONSTRAINT "document_contract_tenant_id_contract_property_id_fkey";

-- AlterTable
ALTER TABLE "contract" DROP CONSTRAINT "contract_pkey",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "contract_pkey" PRIMARY KEY ("id");

-- AlterTable
ALTER TABLE "document" DROP COLUMN "contract_property_id",
DROP COLUMN "contract_tenant_id",
ADD COLUMN     "contract_id" TEXT NOT NULL;

-- AddForeignKey
ALTER TABLE "document" ADD CONSTRAINT "document_contract_id_fkey" FOREIGN KEY ("contract_id") REFERENCES "contract"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
