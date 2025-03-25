-- CreateTable
CREATE TABLE "document" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "data" BYTEA NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "contract_tenant_id" TEXT NOT NULL,
    "contract_property_id" TEXT NOT NULL,

    CONSTRAINT "document_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "document" ADD CONSTRAINT "document_contract_tenant_id_contract_property_id_fkey" FOREIGN KEY ("contract_tenant_id", "contract_property_id") REFERENCES "contract"("tenant_id", "property_id") ON DELETE RESTRICT ON UPDATE CASCADE;
