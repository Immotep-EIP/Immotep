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

| **Role Name** | **Description**     |
| ------------- | ------------------- |
| Role 1        | [Brief description] |
| Role 2        | [Brief description] |
| Role 3        | [Brief description] |

### **2.2 Test Scenarios**

For each core functionality, provide detailed test scenarios.

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
  2.
  3. Click on the drop-down menu at the top right of the property and add a tenant
  4. Fill in contract details (tenant email, start date of the contract, end date of the contract is optional)

- **Expected Outcome:**

  - Tenant is successfully invited
  - The property badge changes to **invitation sent**.
  - The tenant receives an e-mail affiliated with the property to create an Immotep account.

---

## **3. Success Criteria**

[Define the metrics and conditions that determine if the beta version is successful.]

---

## **4. Known Issues & Limitations**

[List any known bugs, incomplete features, or limitations that testers should be aware of.]

| **Issue** | **Description**     | **Impact** | **Planned Fix? (Yes/No)** |
| --------- | ------------------- | ---------- | ------------------------- |
| Issue 1   | [Brief description] | High       | Yes                       |
| Issue 2   | [Brief description] | Medium     | No                        |

---

## **5. Conclusion**

[Summarize the importance of this Beta Test Plan and what the team expects to achieve with it.]
