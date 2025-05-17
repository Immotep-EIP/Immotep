package main

import (
	"log"
	"time"

	"immotep/backend/prisma/db"
	"immotep/backend/services"
	"immotep/backend/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARNING: failed loading .env file")
	}

	c, err := services.ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = c.Client.Disconnect() }()

	log.Println("Connected to database, running seed...")

	password, err := utils.HashPassword("1234567890")
	if err != nil {
		log.Fatalln(err)
	}

	owner, err := c.Client.User.CreateOne(
		db.User.Email.Set("test.owner@example.com"),
		db.User.Password.Set(password),
		db.User.Firstname.Set("Test"),
		db.User.Lastname.Set("Owner"),
		db.User.Role.Set(db.RoleOwner),
		db.User.ID.Set("cm8prd41q00004f0zgo6y6fzw"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	tenant, err := c.Client.User.CreateOne(
		db.User.Email.Set("test.tenant@example.com"),
		db.User.Password.Set(password),
		db.User.Firstname.Set("Test"),
		db.User.Lastname.Set("Tenant"),
		db.User.Role.Set(db.RoleTenant),
		db.User.ID.Set("cm8prd41x00014f0z3q86vfxq"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, l, r := createMainProperty(c, *owner, *tenant)
	createReminder_1(c, *owner, password)
	createReminder_2(c, *owner, password)
	createReminder_3_4(c, *owner)
	createReminder_5(c, *owner)
	createReminder_6(c, l, r)
	createReminder_7(c, l, r)
	createReminder_8(c, l, r)
	createReminder_9(c, l, r)
	createReminder_10(c, l, r)
	createReminder_11(c, l, r)

	// _, err = c.Client.Property.CreateOne(
	// 	db.Property.Name.Set("Garage KM0"),
	// 	db.Property.Address.Set("30 rue François Spoerry"),
	// 	db.Property.City.Set("Mulhouse"),
	// 	db.Property.PostalCode.Set("68100"),
	// 	db.Property.Country.Set("France"),
	// 	db.Property.AreaSqm.Set(20),
	// 	db.Property.RentalPricePerMonth.Set(100),
	// 	db.Property.DepositPrice.Set(500),
	// 	db.Property.Owner.Link(db.User.ID.Equals(owner.ID)),
	// 	db.Property.ID.Set("cm8pr1ilc0003d8nhzlqxb2w4"),
	// ).Exec(c.Context)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	log.Println("Database seeded successfully.")
}

func createMainProperty(c *services.PrismaDB, owner db.UserModel, tenant db.UserModel) (db.PropertyModel, db.LeaseModel, db.RoomModel) {
	property, err := c.Client.Property.CreateOne(
		db.Property.Name.Set("Logement étudiant KM0"),
		db.Property.Address.Set("30 rue François Spoerry"),
		db.Property.City.Set("Mulhouse"),
		db.Property.PostalCode.Set("68100"),
		db.Property.Country.Set("France"),
		db.Property.AreaSqm.Set(1000),
		db.Property.RentalPricePerMonth.Set(1000),
		db.Property.DepositPrice.Set(5000),
		db.Property.Owner.Link(db.User.ID.Equals(owner.ID)),
		db.Property.ID.Set("cm8pr1il60002d8nhkyizcbhl"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	room1, err := c.Client.Room.CreateOne(
		db.Room.Name.Set("Chambre"),
		db.Room.Type.Set(db.RoomTypeBedroom),
		db.Room.Property.Link(db.Property.ID.Equals(property.ID)),
		db.Room.ID.Set("cm8pr1ile0004d8nhqddiph2t"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	room2, err := c.Client.Room.CreateOne(
		db.Room.Name.Set("Cuisine"),
		db.Room.Type.Set(db.RoomTypeKitchen),
		db.Room.Property.Link(db.Property.ID.Equals(property.ID)),
		db.Room.ID.Set("cm8pr1ilh0005d8nhqwtfnyq5"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Lit"),
		db.Furniture.Room.Link(db.Room.ID.Equals(room1.ID)),
		db.Furniture.Quantity.Set(1),
		db.Furniture.ID.Set("cm8pr1ilj0006d8nhn42fvjgx"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Table"),
		db.Furniture.Room.Link(db.Room.ID.Equals(room1.ID)),
		db.Furniture.Quantity.Set(1),
		db.Furniture.ID.Set("cm8pr1ilm0007d8nh3ibh8qlk"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Chaise"),
		db.Furniture.Room.Link(db.Room.ID.Equals(room1.ID)),
		db.Furniture.Quantity.Set(1),
		db.Furniture.ID.Set("cm8pr1iln0008d8nhxqqsy50s"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Frigo"),
		db.Furniture.Room.Link(db.Room.ID.Equals(room2.ID)),
		db.Furniture.Quantity.Set(1),
		db.Furniture.ID.Set("cm8pr1ilp0009d8nht09bpgr7"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Four"),
		db.Furniture.Room.Link(db.Room.ID.Equals(room2.ID)),
		db.Furniture.Quantity.Set(1),
		db.Furniture.ID.Set("cm8pr1ilq000ad8nhab8xaz3z"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Plaque de cuisson"),
		db.Furniture.Room.Link(db.Room.ID.Equals(room2.ID)),
		db.Furniture.Quantity.Set(1),
		db.Furniture.ID.Set("cm8pr1ilr000bd8nhwobl0ryl"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	lease, err := c.Client.Lease.CreateOne(
		db.Lease.StartDate.Set(time.Now()),
		db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
		db.Lease.Property.Link(db.Property.ID.Equals(property.ID)),
		db.Lease.EndDate.Set(time.Now().AddDate(1, 0, 0)),
		db.Lease.ID.Set("cm8pr1ils000cd8nhf1c7imt0"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.InventoryReport.CreateOne(
		db.InventoryReport.Type.Set(db.ReportTypeStart),
		db.InventoryReport.Lease.Link(db.Lease.ID.Equals(lease.ID)),
		db.InventoryReport.ID.Set("cm8pr1ilt000dd8nhq0j3x2g4"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	// TODO create room and furniture states

	return *property, *lease, *room1
}

func createReminder_1(c *services.PrismaDB, owner db.UserModel, password string) {
	tenant, err := c.Client.User.CreateOne(
		db.User.Email.Set("test.tenant.1@example.com"),
		db.User.Password.Set(password),
		db.User.Firstname.Set("Test"),
		db.User.Lastname.Set("Tenant"),
		db.User.Role.Set(db.RoleTenant),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	property, err := c.Client.Property.CreateOne(
		db.Property.Name.Set("Test reminder 1"),
		db.Property.Address.Set("30 rue François Spoerry"),
		db.Property.City.Set("Mulhouse"),
		db.Property.PostalCode.Set("68100"),
		db.Property.Country.Set("France"),
		db.Property.AreaSqm.Set(1000),
		db.Property.RentalPricePerMonth.Set(1000),
		db.Property.DepositPrice.Set(5000),
		db.Property.Owner.Link(db.User.ID.Equals(owner.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Room.CreateOne(
		db.Room.Name.Set("Test reminder 1"),
		db.Room.Type.Set(db.RoomTypeBedroom),
		db.Room.Property.Link(db.Property.ID.Equals(property.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	lease, err := c.Client.Lease.CreateOne(
		db.Lease.StartDate.Set(time.Now().AddDate(-1, 0, 0)),
		db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
		db.Lease.Property.Link(db.Property.ID.Equals(property.ID)),
		db.Lease.EndDate.Set(time.Now().AddDate(0, 0, 25)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.InventoryReport.CreateOne(
		db.InventoryReport.Type.Set(db.ReportTypeStart),
		db.InventoryReport.Lease.Link(db.Lease.ID.Equals(lease.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_2(c *services.PrismaDB, owner db.UserModel, password string) {
	tenant, err := c.Client.User.CreateOne(
		db.User.Email.Set("test.tenant.2@example.com"),
		db.User.Password.Set(password),
		db.User.Firstname.Set("Test"),
		db.User.Lastname.Set("Tenant"),
		db.User.Role.Set(db.RoleTenant),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	property, err := c.Client.Property.CreateOne(
		db.Property.Name.Set("Test reminder 2"),
		db.Property.Address.Set("30 rue François Spoerry"),
		db.Property.City.Set("Mulhouse"),
		db.Property.PostalCode.Set("68100"),
		db.Property.Country.Set("France"),
		db.Property.AreaSqm.Set(1000),
		db.Property.RentalPricePerMonth.Set(1000),
		db.Property.DepositPrice.Set(5000),
		db.Property.Owner.Link(db.User.ID.Equals(owner.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Room.CreateOne(
		db.Room.Name.Set("Test reminder 2"),
		db.Room.Type.Set(db.RoomTypeBedroom),
		db.Room.Property.Link(db.Property.ID.Equals(property.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Lease.CreateOne(
		db.Lease.StartDate.Set(time.Now().AddDate(0, 0, -7)),
		db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
		db.Lease.Property.Link(db.Property.ID.Equals(property.ID)),
		db.Lease.EndDate.Set(time.Now().AddDate(1, 0, 0)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_3_4(c *services.PrismaDB, owner db.UserModel) {
	_, err := c.Client.Property.CreateOne(
		db.Property.Name.Set("Test reminder 3/4"),
		db.Property.Address.Set("30 rue François Spoerry"),
		db.Property.City.Set("Mulhouse"),
		db.Property.PostalCode.Set("68100"),
		db.Property.Country.Set("France"),
		db.Property.AreaSqm.Set(1000),
		db.Property.RentalPricePerMonth.Set(1000),
		db.Property.DepositPrice.Set(5000),
		db.Property.Owner.Link(db.User.ID.Equals(owner.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_5(c *services.PrismaDB, owner db.UserModel) {
	property, err := c.Client.Property.CreateOne(
		db.Property.Name.Set("Test reminder 5"),
		db.Property.Address.Set("30 rue François Spoerry"),
		db.Property.City.Set("Mulhouse"),
		db.Property.PostalCode.Set("68100"),
		db.Property.Country.Set("France"),
		db.Property.AreaSqm.Set(1000),
		db.Property.RentalPricePerMonth.Set(1000),
		db.Property.DepositPrice.Set(5000),
		db.Property.Owner.Link(db.User.ID.Equals(owner.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Room.CreateOne(
		db.Room.Name.Set("Test reminder 1"),
		db.Room.Type.Set(db.RoomTypeBedroom),
		db.Room.Property.Link(db.Property.ID.Equals(property.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.LeaseInvite.CreateOne(
		db.LeaseInvite.TenantEmail.Set("tenant.invite@example.com"),
		db.LeaseInvite.StartDate.Set(time.Now().AddDate(0, 0, 7)),
		db.LeaseInvite.Property.Link(db.Property.ID.Equals(property.ID)),
		db.LeaseInvite.CreatedAt.Set(time.Now().AddDate(0, 0, -8)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_6(c *services.PrismaDB, lease db.LeaseModel, room db.RoomModel) {
	_, err := c.Client.Damage.CreateOne(
		db.Damage.Comment.Set("Test reminder 6"),
		db.Damage.Priority.Set(db.PriorityMedium),
		db.Damage.Lease.Link(db.Lease.ID.Equals(lease.ID)),
		db.Damage.Room.Link(db.Room.ID.Equals(room.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_7(c *services.PrismaDB, lease db.LeaseModel, room db.RoomModel) {
	_, err := c.Client.Damage.CreateOne(
		db.Damage.Comment.Set("Test reminder 7"),
		db.Damage.Priority.Set(db.PriorityMedium),
		db.Damage.Lease.Link(db.Lease.ID.Equals(lease.ID)),
		db.Damage.Room.Link(db.Room.ID.Equals(room.ID)),
		db.Damage.FixPlannedAt.Set(time.Now().AddDate(0, 0, 6)),
		db.Damage.Read.Set(true),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_8(c *services.PrismaDB, lease db.LeaseModel, room db.RoomModel) {
	_, err := c.Client.Damage.CreateOne(
		db.Damage.Comment.Set("Test reminder 8"),
		db.Damage.Priority.Set(db.PriorityUrgent),
		db.Damage.Lease.Link(db.Lease.ID.Equals(lease.ID)),
		db.Damage.Room.Link(db.Room.ID.Equals(room.ID)),
		db.Damage.Read.Set(true),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_9(c *services.PrismaDB, lease db.LeaseModel, room db.RoomModel) {
	_, err := c.Client.Damage.CreateOne(
		db.Damage.Comment.Set("Test reminder 9"),
		db.Damage.Priority.Set(db.PriorityMedium),
		db.Damage.Lease.Link(db.Lease.ID.Equals(lease.ID)),
		db.Damage.Room.Link(db.Room.ID.Equals(room.ID)),
		db.Damage.Read.Set(true),
		db.Damage.CreatedAt.Set(time.Now().AddDate(0, 0, -8)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_10(c *services.PrismaDB, lease db.LeaseModel, room db.RoomModel) {
	_, err := c.Client.Damage.CreateOne(
		db.Damage.Comment.Set("Test reminder 10"),
		db.Damage.Priority.Set(db.PriorityMedium),
		db.Damage.Lease.Link(db.Lease.ID.Equals(lease.ID)),
		db.Damage.Room.Link(db.Room.ID.Equals(room.ID)),
		db.Damage.Read.Set(true),
		db.Damage.FixedTenant.Set(true),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}

func createReminder_11(c *services.PrismaDB, lease db.LeaseModel, room db.RoomModel) {
	_, err := c.Client.Damage.CreateOne(
		db.Damage.Comment.Set("Test reminder 11"),
		db.Damage.Priority.Set(db.PriorityMedium),
		db.Damage.Lease.Link(db.Lease.ID.Equals(lease.ID)),
		db.Damage.Room.Link(db.Room.ID.Equals(room.ID)),
		db.Damage.Read.Set(true),
		db.Damage.FixPlannedAt.Set(time.Now().AddDate(0, 0, -1)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
}
