-- DropIndex
DROP INDEX "damage_fixed_at_created_at_idx";

-- CreateIndex
CREATE INDEX "damage_lease_id_idx" ON "damage"("lease_id");

-- CreateIndex
CREATE INDEX "damage_fixed_at_idx" ON "damage"("fixed_at");
