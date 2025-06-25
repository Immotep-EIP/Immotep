import { TFunction } from 'i18next'

export const ROOM_COLORS = {
  livingroom: '#FF9966', // Orange chaleureux
  diningroom: '#FFB366', // Orange clair
  bedroom: '#9999FF', // Violet doux
  playroom: '#FF99CC', // Rose vif
  bathroom: '#66CCFF', // Bleu clair
  toilet: '#99CCFF', // Bleu très clair
  laundryroom: '#66CCCC', // Turquoise
  kitchen: '#FF6666', // Rouge chaleureux
  hallway: '#FFCC66', // Jaune doux
  dressing: '#CC99FF', // Violet clair
  storage: '#CCCCCC', // Gris clair
  cellar: '#999999', // Gris foncé
  garage: '#666666', // Gris très foncé
  balcony: '#99FF99', // Vert clair
  office: '#99CCFF', // Bleu clair
  other: '#E0E0E0' // Gris neutre
} as const

export const ROOM_TYPES = [
  'bedroom',
  'livingroom',
  'kitchen',
  'bathroom',
  'diningroom',
  'office',
  'toilet',
  'hallway',
  'dressing',
  'laundryroom',
  'playroom',
  'balcony',
  'storage',
  'garage',
  'cellar',
  'other'
] as const

export const ROOM_TEMPLATES: Record<
  string,
  { name: string; quantity: number }[]
> = {
  bedroom: [
    { name: 'Bed', quantity: 1 },
    { name: 'Wardrobe', quantity: 1 },
    { name: 'Nightstand', quantity: 2 }
  ],
  kitchen: [
    { name: 'Table', quantity: 1 },
    { name: 'Chair', quantity: 4 },
    { name: 'Fridge', quantity: 1 },
    { name: 'Oven', quantity: 1 },
    { name: 'Sink', quantity: 1 }
  ],
  livingroom: [
    { name: 'Sofa', quantity: 1 },
    { name: 'Coffee Table', quantity: 1 },
    { name: 'TV Stand', quantity: 1 },
    { name: 'Armchair', quantity: 2 }
  ],
  bathroom: [
    { name: 'Sink', quantity: 1 },
    { name: 'Shower', quantity: 1 },
    { name: 'Toilet', quantity: 1 },
    { name: 'Mirror', quantity: 1 }
  ],
  office: [
    { name: 'Desk', quantity: 1 },
    { name: 'Chair', quantity: 1 },
    { name: 'Bookshelf', quantity: 1 },
    { name: 'Lamp', quantity: 1 }
  ],
  diningroom: [
    { name: 'Dining Table', quantity: 1 },
    { name: 'Chair', quantity: 6 },
    { name: 'Buffet', quantity: 1 }
  ],
  playroom: [
    { name: 'Toy Box', quantity: 1 },
    { name: 'Play Mat', quantity: 1 },
    { name: 'Bookshelf', quantity: 1 }
  ],
  toilet: [
    { name: 'Toilet', quantity: 1 },
    { name: 'Sink', quantity: 1 }
  ],
  laundryroom: [
    { name: 'Washing Machine', quantity: 1 },
    { name: 'Dryer', quantity: 1 },
    { name: 'Laundry Basket', quantity: 1 }
  ],
  hallway: [
    { name: 'Coat Rack', quantity: 1 },
    { name: 'Shoe Cabinet', quantity: 1 }
  ],
  dressing: [
    { name: 'Wardrobe', quantity: 2 },
    { name: 'Mirror', quantity: 1 }
  ],
  storage: [
    { name: 'Shelf', quantity: 2 },
    { name: 'Storage Box', quantity: 4 }
  ],
  cellar: [
    { name: 'Wine Rack', quantity: 1 },
    { name: 'Shelf', quantity: 2 }
  ],
  garage: [
    { name: 'Workbench', quantity: 1 },
    { name: 'Tool Cabinet', quantity: 1 }
  ],
  balcony: [
    { name: 'Outdoor Chair', quantity: 2 },
    { name: 'Small Table', quantity: 1 }
  ]
}

export type RoomType = (typeof ROOM_TYPES)[number]

export const getRoomTypeOptions = (t: TFunction) => [
  { value: 'all', label: t('components.select.room_type.all') },
  ...ROOM_TYPES.map(type => ({
    value: type,
    label: t(`components.select.room_type.${type}`),
    color: ROOM_COLORS[type]
  }))
]

export const isValidRoomType = (type: string): type is RoomType =>
  ROOM_TYPES.includes(type as RoomType)

export const ROOM_GROUPS = {
  living: ['livingroom', 'diningroom'],
  sleeping: ['bedroom', 'playroom'],
  utility: ['bathroom', 'toilet', 'laundryroom', 'kitchen'],
  storage: ['dressing', 'storage', 'cellar', 'garage'],
  other: ['office', 'hallway', 'balcony', 'other']
} as const

export const getRoomColor = (type: RoomType) => ROOM_COLORS[type]
