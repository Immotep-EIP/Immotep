import endpoints from '../EndPointEnum'

describe('EndPointEnum', () => {
  describe('owner endpoints', () => {
    describe('dashboard', () => {
      it('should return the correct dashboard endpoints', () => {
        expect(endpoints.owner.dashboard.list()).toBe('owner/dashboard/')
      })
    })

    describe('properties', () => {
      it('should return the correct base property endpoints', () => {
        expect(endpoints.owner.properties.list).toBe('owner/properties/')
        expect(endpoints.owner.properties.create).toBe('owner/properties/')
        expect(endpoints.owner.properties.picture('123')).toBe(
          'owner/properties/123/picture/'
        )
        expect(endpoints.owner.properties.details('123')).toBe(
          'owner/properties/123/'
        )
        expect(endpoints.owner.properties.archive('123')).toBe(
          'owner/properties/123/archive/'
        )
        expect(endpoints.owner.properties.update('123')).toBe(
          'owner/properties/123/'
        )
      })

      describe('leases', () => {
        it('should return the correct lease endpoints', () => {
          expect(endpoints.owner.properties.leases.current).toBe('current')
          expect(endpoints.owner.properties.leases.end('123')).toBe(
            'owner/properties/123/leases/current/end/'
          )
          expect(endpoints.owner.properties.leases.byId('123', '456')).toBe(
            'owner/properties/123/leases/456/'
          )
          expect(endpoints.owner.properties.leases.list('123')).toBe(
            'owner/properties/123/leases/'
          )
          expect(
            endpoints.owner.properties.leases.documents('123', '456')
          ).toBe('owner/properties/123/leases/456/docs/')
        })
      })

      describe('rooms', () => {
        it('should return the correct room endpoints', () => {
          expect(endpoints.owner.properties.rooms.list('123')).toBe(
            'owner/properties/123/rooms/'
          )
          expect(endpoints.owner.properties.rooms.byId('123', '456')).toBe(
            'owner/properties/123/rooms/456/'
          )
          expect(endpoints.owner.properties.rooms.create('123')).toBe(
            'owner/properties/123/rooms/'
          )
          expect(endpoints.owner.properties.rooms.archive('123', '456')).toBe(
            'owner/properties/123/rooms/456/archive/'
          )
        })

        describe('furnitures', () => {
          it('should return the correct furniture endpoints', () => {
            expect(
              endpoints.owner.properties.rooms.furnitures.list('123', '456')
            ).toBe('owner/properties/123/rooms/456/furnitures/')
            expect(
              endpoints.owner.properties.rooms.furnitures.byId(
                '123',
                '456',
                '789'
              )
            ).toBe('owner/properties/123/rooms/456/furnitures/789/')
            expect(
              endpoints.owner.properties.rooms.furnitures.create('123', '456')
            ).toBe('owner/properties/123/rooms/456/furnitures/')
            expect(
              endpoints.owner.properties.rooms.furnitures.archive(
                '123',
                '456',
                '789'
              )
            ).toBe('owner/properties/123/rooms/456/furnitures/789/archive/')
          })
        })
      })

      describe('inventoryReports', () => {
        it('should return the correct inventory report endpoints', () => {
          expect(endpoints.owner.properties.inventoryReports.list('123')).toBe(
            'owner/properties/123/inventory-reports/'
          )
          expect(
            endpoints.owner.properties.inventoryReports.byId('123', '456')
          ).toBe('owner/properties/123/inventory-reports/456/')
          expect(
            endpoints.owner.properties.inventoryReports.create('123')
          ).toBe('owner/properties/123/inventory-reports/')
        })
      })

      describe('tenant', () => {
        it('should return the correct tenant management endpoints', () => {
          expect(endpoints.owner.properties.tenant.invite('123')).toBe(
            'owner/properties/123/send-invite/'
          )
          expect(endpoints.owner.properties.tenant.cancelInvite('123')).toBe(
            'owner/properties/123/cancel-invite/'
          )
        })
      })

      describe('damages', () => {
        it('should return the correct damages endpoints', () => {
          expect(endpoints.owner.properties.damages.list('123')).toBe(
            'owner/properties/123/leases/current/damages/'
          )
          expect(endpoints.owner.properties.damages.byId('123', '456')).toBe(
            'owner/properties/123/leases/current/damages/456/'
          )
          expect(endpoints.owner.properties.damages.update('123', '456')).toBe(
            'owner/properties/123/leases/current/damages/456/'
          )
        })
      })
    })
  })

  describe('tenant endpoints', () => {
    describe('leases', () => {
      it('should return the correct tenant lease endpoints', () => {
        expect(endpoints.tenant.leases.list()).toBe('tenant/leases/')
        expect(endpoints.tenant.leases.byId('123')).toBe('tenant/leases/123/')
      })
    })

    describe('invite', () => {
      it('should return the correct tenant invite endpoints', () => {
        expect(endpoints.tenant.invite.accept('123')).toBe('tenant/invite/123/')
      })
    })
  })

  describe('user endpoints', () => {
    it('should return the correct user list endpoint', () => {
      expect(endpoints.user.list()).toBe('users/')
    })

    describe('profile', () => {
      it('should return the correct user profile endpoints', () => {
        expect(endpoints.user.profile.get()).toBe('profile/')
      })
    })

    describe('picture', () => {
      it('should return the correct user picture endpoints', () => {
        expect(endpoints.user.picture.get('123')).toBe('user/123/picture/')
        expect(endpoints.user.picture.update()).toBe('profile/picture/')
      })
    })

    describe('auth', () => {
      it('should return the correct auth endpoints', () => {
        expect(endpoints.user.auth.register()).toBe('auth/register/')
        expect(endpoints.user.auth.token()).toBe('auth/token/')
        expect(endpoints.user.auth.invite('123')).toBe('auth/invite/123/')
      })
    })
  })
})
