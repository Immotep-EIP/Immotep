# Database Documentation

## Database Design

The database for this application is implemented using PostgreSQL. It stores all the persistent data required by the system, including user information, application data, and other related entities. The database design follows a relational model with well-defined tables and relationships.

### Entity-Relationship Diagram (ERD)

The following diagram illustrates the structure of the database, including the tables and their relationships:

![Entity-Relationship Diagram](prisma-erd.svg)

### Tables Overview

The database schema consists of several tables, each representing a distinct entity in the system. Below is a brief description of each table:

- **Users**: Stores user account information, including authentication details, roles, and profile data.
- **Contracts**: Represents rental agreements between tenants and properties, including start and end dates.
- **Damages**: Records information about damages reported in properties, including priority and status.
- **Properties**: Contains details about properties, such as address, rental price, and owner information.
- **Pending Contracts**: Holds information about contracts that are pending approval after an invitation from an owner to a tenant, including tenant email and property details.
- **Images**: Stores image data used for user profiles, property pictures, and inventory report state documentation.
- **Rooms**: Represents individual rooms within properties, including their names and associated property.
- **Furnitures**: Contains information about furniture items in rooms, including quantity and associated room.
- **Inventory Reports**: Records periodic reports on the state of rooms and furniture within properties.
- **Room States**: Documents the state and cleanliness of rooms as part of inventory reports.
- **Furniture States**: Documents the state and cleanliness of furniture items as part of inventory reports.

Each table is designed with foreign key constraints where applicable to maintain referential integrity across the database.

## Data Models

The Go application uses Prisma as the ORM (Object-Relational Mapping) to interact with the PostgreSQL database. The Prisma schema defines the data models and is used to generate type-safe database access methods.

The Prisma schema file, typically located in the `prisma` folder, defines the structure of the database models. For more information, [see the Prisma website](https://www.prisma.io/docs/orm).

## Migrations and Seeding

Prisma provides a built-in migration tool to manage changes to the database schema over time. Each migration captures a snapshot of the schema, allowing for safe and versioned updates to the database.

### Running Migrations

To apply new migrations to the database, use the following command:

```bash
cd Code/Server/
./migrate_db.sh
```

You then need to name your migration. This command creates a new migration file and applies it to the database.

### Seed

TODO

## Query Optimization and Indexes

TODO

## Backup and Recovery

TODO

## Migration Strategy

TODO
