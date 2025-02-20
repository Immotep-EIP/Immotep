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

	o, err := c.Client.User.CreateOne(
		db.User.Email.Set("test.owner@example.com"),
		db.User.Password.Set(password),
		db.User.Firstname.Set("Test"),
		db.User.Lastname.Set("Owner"),
		db.User.Role.Set(db.RoleOwner),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	t, err := c.Client.User.CreateOne(
		db.User.Email.Set("test.tenant@example.com"),
		db.User.Password.Set(password),
		db.User.Firstname.Set("Test"),
		db.User.Lastname.Set("Tenant"),
		db.User.Role.Set(db.RoleTenant),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	p, err := c.Client.Property.CreateOne(
		db.Property.Name.Set("Logement étudiant KM0"),
		db.Property.Address.Set("30 rue François Spoerry"),
		db.Property.City.Set("Mulhouse"),
		db.Property.PostalCode.Set("68100"),
		db.Property.Country.Set("France"),
		db.Property.AreaSqm.Set(1000),
		db.Property.RentalPricePerMonth.Set(1000),
		db.Property.DepositPrice.Set(5000),
		db.Property.Owner.Link(db.User.ID.Equals(o.ID)),
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
		db.Property.Owner.Link(db.User.ID.Equals(o.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	r1, err := c.Client.Room.CreateOne(
		db.Room.Name.Set("Chambre"),
		db.Room.Property.Link(db.Property.ID.Equals(p.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	r2, err := c.Client.Room.CreateOne(
		db.Room.Name.Set("Cuisine"),
		db.Room.Property.Link(db.Property.ID.Equals(p.ID)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Lit"),
		db.Furniture.Room.Link(db.Room.ID.Equals(r1.ID)),
		db.Furniture.Quantity.Set(1),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Table"),
		db.Furniture.Room.Link(db.Room.ID.Equals(r1.ID)),
		db.Furniture.Quantity.Set(1),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Chaise"),
		db.Furniture.Room.Link(db.Room.ID.Equals(r1.ID)),
		db.Furniture.Quantity.Set(1),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Frigo"),
		db.Furniture.Room.Link(db.Room.ID.Equals(r2.ID)),
		db.Furniture.Quantity.Set(1),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Four"),
		db.Furniture.Room.Link(db.Room.ID.Equals(r2.ID)),
		db.Furniture.Quantity.Set(1),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = c.Client.Furniture.CreateOne(
		db.Furniture.Name.Set("Plaque de cuisson"),
		db.Furniture.Room.Link(db.Room.ID.Equals(r2.ID)),
		db.Furniture.Quantity.Set(1),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = c.Client.Contract.CreateOne(
		db.Contract.StartDate.Set(time.Now()),
		db.Contract.Tenant.Link(db.User.ID.Equals(t.ID)),
		db.Contract.Property.Link(db.Property.ID.Equals(p.ID)),
		db.Contract.EndDate.Set(time.Now().AddDate(1, 0, 0)),
	).Exec(c.Context)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Database seeded successfully.")
}
