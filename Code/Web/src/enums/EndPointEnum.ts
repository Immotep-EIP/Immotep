const endpoints = {
  // Owner endpoints
  owner: {
    // Dashboard endpoints
    dashboard: {
      list: () => 'owner/dashboard/'
    },
    properties: {
      // Base property endpoints
      list: 'owner/properties/',
      picture: (propertyId: string) =>
        `owner/properties/${propertyId}/picture/`,
      details: (propertyId: string) => `owner/properties/${propertyId}/`,
      create: 'owner/properties/',
      archive: (propertyId: string) =>
        `owner/properties/${propertyId}/archive/`,
      update: (propertyId: string) => `owner/properties/${propertyId}/`,

      // Leases endpoints
      leases: {
        current: 'current',
        end: (propertyId: string) =>
          `owner/properties/${propertyId}/leases/current/end/`,
        byId: (propertyId: string, leaseId: string) =>
          `owner/properties/${propertyId}/leases/${leaseId}/`,
        list: (propertyId: string) => `owner/properties/${propertyId}/leases/`,
        documents: (propertyId: string, leaseId: string) =>
          `owner/properties/${propertyId}/leases/${leaseId}/docs/`,
        deleteDocument: (
          propertyId: string,
          leaseId: string,
          documentId: string
        ) =>
          `owner/properties/${propertyId}/leases/${leaseId}/docs/${documentId}/`
      },

      // Rooms endpoints
      rooms: {
        list: (propertyId: string) => `owner/properties/${propertyId}/rooms/`,
        byId: (propertyId: string, roomId: string) =>
          `owner/properties/${propertyId}/rooms/${roomId}/`,
        create: (propertyId: string) => `owner/properties/${propertyId}/rooms/`,
        archive: (propertyId: string, roomId: string) =>
          `owner/properties/${propertyId}/rooms/${roomId}/archive/`,
        furnitures: {
          list: (propertyId: string, roomId: string) =>
            `owner/properties/${propertyId}/rooms/${roomId}/furnitures/`,
          byId: (propertyId: string, roomId: string, furnitureId: string) =>
            `owner/properties/${propertyId}/rooms/${roomId}/furnitures/${furnitureId}/`,
          create: (propertyId: string, roomId: string) =>
            `owner/properties/${propertyId}/rooms/${roomId}/furnitures/`,
          archive: (propertyId: string, roomId: string, furnitureId: string) =>
            `owner/properties/${propertyId}/rooms/${roomId}/furnitures/${furnitureId}/archive/`
        }
      },

      // Inventory reports endpoints
      inventoryReports: {
        list: (propertyId: string) =>
          `owner/properties/${propertyId}/inventory-reports/`,
        byId: (propertyId: string, reportId: string) =>
          `owner/properties/${propertyId}/inventory-reports/${reportId}/`,
        create: (propertyId: string) =>
          `owner/properties/${propertyId}/inventory-reports/`
      },

      // Tenant management endpoints
      tenant: {
        invite: (propertyId: string) =>
          `owner/properties/${propertyId}/send-invite/`,
        cancelInvite: (propertyId: string) =>
          `owner/properties/${propertyId}/cancel-invite/`
      },

      // Damages endpoints
      damages: {
        list: (propertyId: string) =>
          `owner/properties/${propertyId}/leases/current/damages/`,
        byId: (propertyId: string, damageId: string) =>
          `owner/properties/${propertyId}/leases/current/damages/${damageId}/`,
        update: (propertyId: string, damageId: string) =>
          `owner/properties/${propertyId}/leases/current/damages/${damageId}/`
      }
    }
  },

  // Tenant endpoints
  tenant: {
    leases: {
      list: () => 'tenant/leases/',
      byId: (leaseId: string) => `tenant/leases/${leaseId}/`
    },
    invite: {
      accept: (leaseId: string) => `tenant/invite/${leaseId}/`
    }
  },

  // User endpoints
  user: {
    list: () => 'users/',
    profile: {
      get: () => 'profile/'
    },
    picture: {
      get: (id: string) => `user/${id}/picture/`,
      update: () => 'profile/picture/'
    },
    auth: {
      register: () => 'auth/register/',
      token: () => 'auth/token/',
      invite: (leaseId: string) => `auth/invite/${leaseId}/`
    }
  }
}

export default endpoints
