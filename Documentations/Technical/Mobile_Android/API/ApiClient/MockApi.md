# MockedApiService (Mock Implementation of ApiService)

This class simulates the behavior of the real `ApiService` for testing and development without actual backend calls. It returns predefined fake or static responses for all API methods.

---

### General Behavior

- Throws exceptions for some error simulation cases (e.g., login with `"error@gmail.com"`).
- Returns predefined "fake" data objects for most requests.
- Supports different user roles by varying returned data (e.g., tenant vs owner).


---

### Key Methods Overview

#### Authentication

- `login(grantType, username, password)`:  
  - Returns `fakeLoginResponse` or tenant-specific tokens.  
  - Throws on invalid credentials.

- `refreshToken(grantType, refreshToken)`:  
  - Returns `fakeLoginResponse`.

- `register(registrationInput)`:  
  - Returns `fakeRegistrationResponse`.

---

#### Profile

- `getProfile(authHeader)`:  
  - Returns `fakeProfileResponse`.  
  - If token is `"tenantAccessToken"`, returns profile with role `"tenant"`.

- `updateProfile(authHeader, profileUpdateInput)`:  
  - Throws on specific error email.  
  - Returns a constructed `ProfileResponse` based on input.

---

#### Property Management

- `getProperties(authHeader)`:  
  - Returns an array of fake properties (e.g., Paris, Marseille, Lyon, empty).

- `getProperty(authHeader, propertyId)`:  
  - Returns the Paris fake property.

- `addProperty(authHeader, AddPropertyInput)`:  
  - Returns a success response with the Paris fake property ID.

- `updateProperty(authHeader, AddPropertyInput, propertyId)`:  
  - Returns success response with the Paris fake property ID.

- `archiveProperty(authHeader, propertyId, ArchivePropertyInput)`:  
  - Returns success response with the given property ID.

- `getPropertyDocuments(authHeader, propertyId, leaseId)`:  
  - Returns an array containing a fake document.

- `getPropertyDocumentsTenant(authHeader, leaseId)`:  
  - Returns an array containing a fake document.

---

#### Rooms

- `getAllRooms(authHeader, propertyId)`:  
  - Returns an array with a single fake room.

- `addRoom(authHeader, propertyId, AddRoomInput)`:  
  - Returns success with the fake room ID.

- `getAllRoomsTenant(authHeader, leaseId)`:  
  - Returns an array with the fake room.

- `archiveRoom(authHeader, propertyId, roomId, ArchiveInput)`:  
  - Returns success with the room ID.

---

#### Furniture

- `getAllFurnitures(authHeader, propertyId, roomId)`:  
  - Returns an array with a single fake furniture.

- `addFurniture(authHeader, propertyId, roomId, FurnitureInput)`:  
  - Returns success with a new furniture ID.

---

#### Inventory Reports

- `inventoryReport(authHeader, propertyId, leaseId, InventoryReportInput)`:  
  - Returns a created inventory report with static values and input type.

- `getAllInventoryReports(authHeader, propertyId)`:  
  - Returns an array containing a fake inventory report.

- `getInventoryReportByIdOrLatest(authHeader, propertyId, reportId)`:  
  - Returns a fake inventory report.

---

#### AI Calls

- `aiSummarize(authHeader, propertyId, leaseId, AiCallInput)`:  
  - Returns a fake AI call output.

- `aiCompare(authHeader, propertyId, leaseId, oldReportId, AiCallInput)`:  
  - Returns a fake AI call output.

---

#### Tenant Invitations

- `inviteTenant(authHeader, propertyId, InviteInput)`:  
  - Returns success response with a fake invite ID.

- `cancelTenantInvitation(authHeader, propertyId)`:  
  - Returns a successful Retrofit `Unit` response.

---

#### Document Uploads

- `uploadDocument(authHeader, propertyId, leaseId, DocumentInput)`:  
  - Returns success with a new document ID.

- `uploadDocumentTenant(authHeader, leaseId, DocumentInput)`:  
  - Returns success with a new document ID.

---

#### Property Pictures

- `getPropertyPicture(authHeader, propertyId)`:  
  - Returns a Retrofit successful response containing a fake picture.

- `updatePropertyPicture(authHeader, propertyId, UpdatePropertyPictureInput)`:  
  - Returns success with a new picture ID.

---

#### Property Tenants & Damages

- `getPropertyTenant(authHeader, leaseId)`:  
  - Returns the Paris fake property.

- `getPropertyDamages(authHeader, propertyId, leaseId)`:  
  - Returns a fake damages array.

- `getPropertyDamagesTenant(authHeader, leaseId)`:  
  - Returns a fake damages array.

---

#### Damage Management

- `addDamage(authHeader, leaseId, DamageInput)`:  
  - Returns success with a new damage ID.

---

#### Dashboard

- `getDashboard(authHeader, lang)`:  
  - Returns a fake dashboard output.

---

### Notes

- All "fake" or mocked values (e.g., `fakeLoginResponse`, `parisFakeProperty`, `fakeRoom`, `fakeFurniture`, `fakeDocument`, `fakeInventoryReport`, `fakeAiCallOutput`, `fakeInviteOutput`, `fakeDamagesArray`, `fakeGetDashBoardOutput`) are predefined in the mock implementation (not detailed here).
- The class provides a full mock of backend API for integration testing or local development without real API dependency.
