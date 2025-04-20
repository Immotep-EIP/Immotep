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
  'dressing',
  'laundryroom',
  'bedroom',
  'playroom',
  'bathroom',
  'toilet',
  'livingroom',
  'diningroom',
  'kitchen',
  'hallway',
  'balcony',
  'cellar',
  'garage',
  'storage',
  'office',
  'other'
] as const

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
