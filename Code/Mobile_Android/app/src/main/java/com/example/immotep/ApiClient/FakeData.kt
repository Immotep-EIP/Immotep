package com.example.immotep.ApiClient

import com.example.immotep.realProperty.details.DetailedProperty
import java.util.Date

val newTestArray = arrayOf(
    DetailedProperty(
    "last",
    "https://plus.unsplash.com/premium_photo-1661915661139-5b6a4e4a6fcc?fm=jpg&q=60&w=3000&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aG91c2V8ZW58MHx8MHx8fDA%3D",
    "19 rue de la paix, Paris 75000",
    "John Doe",
    true,
    Date(),
    null,
    45,
    1000,
    2000,
    arrayOf("july quittance", "old inventory", "oven invoice", "august quittance")
),
    DetailedProperty(
        "abc",
        "https://plus.unsplash.com/premium_photo-1661915661139-5b6a4e4a6fcc?fm=jpg&q=60&w=3000&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aG91c2V8ZW58MHx8MHx8fDA%3D",
        "30 rue de la source, Lyon 69000",
        "Tom Nook",
        false,
        Date(),
        null,
        70,
        1300,
        3000,
        arrayOf("old inventory", "invoice1")
    ),
    DetailedProperty(
        "one",
        "https://plus.unsplash.com/premium_photo-1661915661139-5b6a4e4a6fcc?fm=jpg&q=60&w=3000&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aG91c2V8ZW58MHx8MHx8fDA%3D",
        "1 rue de la companie des indes, Marseille 13000",
        "Crash Bandicoot",
        false,
        Date(),
        Date(),
        150,
        2000,
        4000,
        arrayOf("may quittance", "july quittance")
    )
)
