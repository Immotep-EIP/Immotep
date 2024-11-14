const realProperties = [
  {
    id: 1,
    adress: "1234, rue de la paix",
    city: "Paris",
    zipCode: "75000",
    image: "https://picsum.photos/200",
    status: "pages.property.status.available",
    tenants: [
      {
        id: 1,
        name: "John Doe",
        email: "john.doe@immotep.fr"
      },
      {
        id: 2,
        name: "Jane Doe",
        email: "jane.doe@immotep.fr"
      }
    ],
    startDate: "2021-01-01",
    endDate: "2021-12-31",
    rent: 1000,
    deposit: 2000,
    area: 50,
    damages: [
      {
        id: 1,
        read: true,
        date: "2021-01-01",
        comment: "Commentaire",
        interventionDate: "2021-01-01",
        room: "Chambre",
        priority: "pages.property.damage.priority.high",
        photos: [
          {
            id: 1,
            url: "https://picsum.photos/200"
          }
        ]
      }
    ],
    documents: [
      {
        id: 1,
        name: "Bail",
        url: "https://picsum.photos/200"
      },
      {
        id: 2,
        name: "Etat des lieux",
        url: "https://picsum.photos/200"
      }
    ],
    inventory: [
      {
        room: "Chambre",
        furniture: [
          {
            id: 1,
            name: "Lit",
            quantity: 1
          },
          {
            id: 2,
            name: "Armoire",
            quantity: 1
          },
          {
            id: 3,
            name: "Table de chevet",
            quantity: 2
          }
        ]
      },
      {
        room: "Salon",
        furniture: [
          {
            id: 1,
            name: "Canapé",
            quantity: 1
          },
          {
            id: 2,
            name: "Table basse",
            quantity: 1
          }
        ]
      }
    ]
  },
  {
    id: 2,
    adress: "5678, avenue des Champs-Élysées",
    city: "Paris",
    zipCode: "75008",
    image: "https://picsum.photos/200",
    status: "pages.property.status.occupied",
    tenants: [
      {
        id: 1,
        name: "Alice Martin",
        email: "alice.martin@immotep.fr"
      }
    ],
    startDate: "2022-02-01",
    endDate: "2023-01-31",
    rent: 1500,
    deposit: 3000,
    area: 75,
    damages: [],
    documents: [
      {
        id: 1,
        name: "Contrat de location",
        url: "https://picsum.photos/200"
      }
    ],
    inventory: [
      {
        room: "Cuisine",
        furniture: [
          {
            id: 1,
            name: "Table",
            quantity: 1
          },
          {
            id: 2,
            name: "Chaises",
            quantity: 4
          }
        ]
      }
    ]
  },
  {
    id: 3,
    adress: "4321, rue du Commerce",
    city: "Lyon",
    zipCode: "69000",
    image: "https://picsum.photos/200",
    status: "pages.property.status.available",
    tenants: [],
    startDate: "2023-03-01",
    endDate: "2024-02-29",
    rent: 800,
    deposit: 1600,
    area: 45,
    damages: [
      {
        id: 1,
        read: false,
        date: "2023-05-10",
        comment: "Dégât des eaux",
        interventionDate: "2023-05-12",
        room: "Salle de bain",
        priority: "pages.property.damage.priority.medium",
        photos: [
          {
            id: 1,
            url: "https://picsum.photos/200"
          }
        ]
      }
    ],
    documents: [
      {
        id: 1,
        name: "Bail",
        url: "https://picsum.photos/200"
      }
    ],
    inventory: [
      {
        room: "Salon",
        furniture: [
          {
            id: 1,
            name: "Table basse",
            quantity: 1
          },
          {
            id: 2,
            name: "Télévision",
            quantity: 1
          }
        ]
      }
    ]
  },
  {
    id: 4,
    adress: "9876, rue de la Liberté",
    city: "Marseille",
    zipCode: "13001",
    image: "https://picsum.photos/200",
    status: "pages.property.status.available",
    tenants: [],
    startDate: "2023-06-01",
    endDate: "2024-05-31",
    rent: 950,
    deposit: 1900,
    area: 60,
    damages: [
      {
        id: 1,
        read: true,
        date: "2023-07-15",
        comment: "Fenêtre cassée",
        interventionDate: "2023-07-20",
        room: "Salon",
        priority: "pages.property.damage.priority.low",
        photos: [
          {
            id: 1,
            url: "https://picsum.photos/200"
          }
        ]
      }
    ],
    documents: [
      {
        id: 1,
        name: "Bail",
        url: "https://picsum.photos/200"
      }
    ],
    inventory: [
      {
        room: "Chambre",
        furniture: [
          {
            id: 1,
            name: "Lit",
            quantity: 1
          },
          {
            id: 2,
            name: "Commode",
            quantity: 1
          }
        ]
      }
    ]
  }
];

export default realProperties;
