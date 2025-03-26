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

	property1, err := c.Client.Property.CreateOne(
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

	_, err = c.Client.Property.CreateOne(
		db.Property.Name.Set("Garage KM0"),
		db.Property.Address.Set("30 rue François Spoerry"),
		db.Property.City.Set("Mulhouse"),
		db.Property.PostalCode.Set("68100"),
		db.Property.Country.Set("France"),
		db.Property.AreaSqm.Set(20),
		db.Property.RentalPricePerMonth.Set(100),
		db.Property.DepositPrice.Set(500),
		db.Property.Owner.Link(db.User.ID.Equals(owner.ID)),
		db.Property.ID.Set("cm8pr1ilc0003d8nhzlqxb2w4"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	room1, err := c.Client.Room.CreateOne(
		db.Room.Name.Set("Chambre"),
		db.Room.Property.Link(db.Property.ID.Equals(property1.ID)),
		db.Room.ID.Set("cm8pr1ile0004d8nhqddiph2t"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	room2, err := c.Client.Room.CreateOne(
		db.Room.Name.Set("Cuisine"),
		db.Room.Property.Link(db.Property.ID.Equals(property1.ID)),
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

	_, err = c.Client.Lease.CreateOne(
		db.Lease.StartDate.Set(time.Now()),
		db.Lease.Tenant.Link(db.User.ID.Equals(tenant.ID)),
		db.Lease.Property.Link(db.Property.ID.Equals(property1.ID)),
		db.Lease.EndDate.Set(time.Now().AddDate(1, 0, 0)),
		db.Lease.ID.Set("cm8pr1ils000cd8nhf1c7imt0"),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Database seeded successfully.")
}
