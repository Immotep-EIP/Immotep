### **BETA TEST PLAN**

_This template serves as a structured example of what is expected for your Beta Test Plan._

## **1. Core Functionalities for Beta Version**

[List and describe the core functionalities that must be available for beta testing. Explain any changes made since the original Tech3 Action Plan.]

| **Feature Name** | **Description**     | **Priority (High/Medium/Low)** | **Changes Since Tech3**      |
| ---------------- | ------------------- | ------------------------------ | ---------------------------- |
| Feature 1        | [Brief description] | High                           | [Modifications or additions] |
| Feature 2        | [Brief description] | Medium                         | [Modifications or additions] |
| Feature 3        | [Brief description] | High                           | [Modifications or additions] |

---

## **2. Beta Testing Scenarios**

### **2.1 User Roles**

[Define the different user roles that will be involved in testing, e.g., Admin, Regular User, Guest, External Partner.]

| **Role Name**  | **Description**                                                                 |
|----------------|---------------------------------------------------------------------------------|
| Property Owner | User who owns and manages properties (create, edit, delete, invite tenants).    |
| Tenant         | User invited to rent a property, with limited access to property details and contract-related features. |

### **2.2 Test Scenarios**

For each core functionality, provide detailed test scenarios.

## Web scenarios

#### **Scenario 1: Property Creation**

- **Role Involved:** Property Owner
- **Objective:** Test the property creation functionality
- **Preconditions:** User is logged in with appropriate permissions
- **Test Steps:**

  1. Navigate to property page
  2. Click on add property button
  3. Fill in property details (address, description, price, ...)
  4. Upload property images
  5. Submit property listing

- **Expected Outcome:**

  - Property is successfully created
  - Property appears in user's property list
  - All property details are correctly saved

#### **Scenario 2: Inventory Report Creation**

- **Role Involved:** Property Owner
- **Objective:** Test the inventory report creation
- **Preconditions:**
  - User is logged in with appropriate permissions
  - Property exists in the system
- **Test Steps:**

  1. Navigate to property page
  2. Choose a property to which a inventory report will be created
  3. Navigate to the inventory tab
  4. Click on "Add Room" button and fill in the room name
  5. Click on the "+" icon next to the room to add items
  6. Fill in the name and number of items

- **Expected Outcome:**

  - Room is successfully created
  - Item is successfully created

#### **Scenario 3: Invite Tenant**

- **Role Involved:** Property Owner
- **Objective:** Test the invite tenant functionality
- **Preconditions:**
  - User is logged in with appropriate permissions
  - Property exists in the system
- **Test Steps:**

  1. Navigate to property page
  2. Choose a property to which a tenant will be invited
  3. Click on the drop-down menu at the top right of the property and add a tenant
  4. Fill in contract details (tenant email, start date of the contract, end date of the contract is optional)

- **Expected Outcome:**

  - Tenant is successfully invited
  - The property badge changes to **invitation sent**.
  - The tenant receives an e-mail affiliated with the property to create an Immotep account.

#### **Scenario 4: Property Modification**

- **Role Involved:** Property Owner
- **Objective:** Test the modification of an existing property
- **Preconditions:**
  - User is logged in with appropriate permissions
  - A property exists in the system
- **Test Steps:**

  1. Navigate to the property page
  2. Choose a property to updated
  3. Click on the drop-down menu at the top right of the property and update the property
  4. Update details (e.g., address, monthly rent, name, etc)
  5. Click "Save Changes"

- **Expected Outcome:**

  - Modifications are saved successfully
  - Updated details (address, rent) and image (if uploaded) appear immediately in the property list

#### **Scenario 5: Property Archiving**

- **Role Involved:** Property Owner
- **Objective:** Test the archiving of a property
- **Preconditions:**
  - User is logged in with appropriate permissions
  - A property exists in the system
- **Test Steps:**
  1. Navigate to the property page
  2. Choose a property to archived
  3. Click on the drop-down menu at the top right of the property and delete the property
  4. Confirm deletion in the alert dialog
- **Expected Outcome:**
  - Property is archived successfully
  - Property disappears from the active property list
  - No errors are displayed

## Mobile scenarios

#### **Scenario 1: Property Modification**

- **Role Involved:** Property Owner
- **Objective:** Test the modification of an existing property
- **Preconditions:**
  - User is logged in with appropriate permissions
  - A property exists in the system
- **Test Steps:**
  1. Navigate to the property view
  2. Swipe right on a property and click "Edit"
  3. Update details (e.g., address, monthly rent, name, etc)
  4. Click "Save Changes"
- **Expected Outcome:**
  - Modifications are saved successfully
  - Updated details (address, rent) and image (if uploaded) appear immediately in the property list

#### **Scenario 2: Property Archiving**

- **Role Involved:** Property Owner
- **Objective:** Test the archiving of a property
- **Preconditions:**
  - User is logged in with appropriate permissions
  - A property exists in the system
- **Test Steps:**
  1. Navigate to the property page
  2. Swipe right on a property and click "Delete"
  3. Confirm deletion in the alert dialog
- **Expected Outcome:**
  - Property is archived successfully
  - Property disappears from the active property list
  - No errors are displayed

#### **Scenario 3: Consulting an Inventory Report**

- **Role Involved:** Property Owner
- **Objective:** Test the ability to view an existing inventory report for a property
- **Preconditions:**
  - User is logged in with appropriate permissions
  - A property exists in the system with at least one inventory report made (entry or exit)
- **Test Steps:**
  1. Navigate to the property page
  2. Select a property with an existing inventory report
  3. Scroll down to the documents part
  4. Click on the document we want to open
- **Expected Outcome:**
  - The inventory report is displayed correctly in the pdf viewer
  - The interface is responsive and loads without errors

---

## **3. Success Criteria**

[Define the metrics and conditions that determine if the beta version is successful.]

---

## **4. Known Issues & Limitations**

[List any known bugs, incomplete features, or limitations that testers should be aware of.]

| **Issue** | **Description**     | **Impact** | **Planned Fix? (Yes/No)** |
| --------- | ------------------- | ---------- | ------------------------- |
| iOS Alert Display Bug             | On iOS, opening an alert too quickly after closing one grays out the background but the alert doesn’t appear | Medium     | Yes                       || Issue 2   | [Brief description] | Medium     | No                        |
Can't add document | Users unable to add new document in his property | Medium | Yes
Handling the damage section | Users unable to handle damage for the moment | Medium | Yes


### **Limitations**
- **Tenant Invitation Link Behavior:** If a tenant opens the invitation received by email on their phone, they will be directed to a web page rather than the Immotep application.

---

## **5. Conclusion**

This Beta Test Plan represents a crucial phase in the development of Immotep, focusing on validating core functionalities essential for property management and tenant interactions. Through structured testing scenarios across both web and mobile platforms, we aim to:

1. **Validate Core Features:** Ensure robust functionality of critical features including property creation, inventory management, and tenant invitation systems.

2. **Cross-Platform Consistency:** Verify seamless user experience across web and mobile interfaces, with special attention to platform-specific interactions.

3. **User Role Verification:** Confirm that both Property Owner and Tenant roles function as intended with appropriate access controls and permissions.

4. **Quality Assurance:** Identify and address potential issues before full release, with documented known limitations to guide future development priorities.

The successful execution of this test plan will provide valuable insights for final refinements and ensure Immotep meets the high standards required for a professional property management solution. Feedback gathered during this beta phase will be instrumental in delivering a polished, user-friendly platform that effectively serves the needs of property owners and tenants alike.
