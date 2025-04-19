import { TFunction } from 'i18next'
import {
  ROOM_COLORS,
  ROOM_TYPES,
  ROOM_GROUPS,
  getRoomTypeOptions,
  isValidRoomType,
  getRoomColor
} from '../roomTypes'

describe('roomTypes', () => {
  describe('ROOM_COLORS', () => {
    it('should have a color for each room type', () => {
      ROOM_TYPES.forEach(type => {
        expect(ROOM_COLORS[type]).toBeDefined()
        expect(ROOM_COLORS[type]).toMatch(/^#[0-9A-F]{6}$/i)
      })
    })
  })

  describe('ROOM_TYPES', () => {
    it('should contain all expected room types', () => {
      const expectedTypes = [
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
      ]
      expect(ROOM_TYPES).toEqual(expectedTypes)
    })
  })

  describe('getRoomTypeOptions', () => {
    const mockT = jest.fn((key: string) => key) as unknown as TFunction

    it('should return options with "all" as first option', () => {
      const options = getRoomTypeOptions(mockT)
      expect(options[0]).toEqual({
        value: 'all',
        label: 'components.select.room_type.all'
      })
    })

    it('should return options for all room types with colors', () => {
      const options = getRoomTypeOptions(mockT)

      expect(options.length).toBe(ROOM_TYPES.length + 1)

      ROOM_TYPES.forEach(type => {
        const option = options.find(opt => opt.value === type)
        expect(option).toEqual({
          value: type,
          label: `components.select.room_type.${type}`,
          color: ROOM_COLORS[type]
        })
      })
    })
  })

  describe('isValidRoomType', () => {
    it('should return true for valid room types', () => {
      ROOM_TYPES.forEach(type => {
        expect(isValidRoomType(type)).toBe(true)
      })
    })

    it('should return false for invalid room types', () => {
      const invalidTypes = ['invalid', 'room', '123', '']
      invalidTypes.forEach(type => {
        expect(isValidRoomType(type)).toBe(false)
      })
    })
  })

  describe('ROOM_GROUPS', () => {
    it('should have valid room types in each group', () => {
      Object.values(ROOM_GROUPS).forEach(group => {
        group.forEach(type => {
          expect(ROOM_TYPES).toContain(type)
        })
      })
    })

    it('should have expected group categories', () => {
      const expectedGroups = [
        'living',
        'sleeping',
        'utility',
        'storage',
        'other'
      ]
      expect(Object.keys(ROOM_GROUPS)).toEqual(expectedGroups)
    })
  })

  describe('getRoomColor', () => {
    it('should return correct color for each room type', () => {
      ROOM_TYPES.forEach(type => {
        expect(getRoomColor(type)).toBe(ROOM_COLORS[type])
      })
    })

    it('should return consistent colors', () => {
      expect(getRoomColor('bedroom')).toBe('#9999FF')
      expect(getRoomColor('kitchen')).toBe('#FF6666')
      expect(getRoomColor('other')).toBe('#E0E0E0')
    })
  })
})
