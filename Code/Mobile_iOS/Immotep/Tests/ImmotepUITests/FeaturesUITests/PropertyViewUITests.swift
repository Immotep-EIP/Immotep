//
//  PropertyViewUITests.swift
//  ImmotepUITests
//
//  Created by Liebenguth Alessio on 09/12/2024.
//

import XCTest

final class PropertyViewUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        app.launchArguments = ["-skipLogin", "-propertyList"]
        app.launch()
    }

    func testAddPropertyButton() throws {
        let addButton = app.buttons["add_property"]
        XCTAssertTrue(addButton.waitForExistence(timeout: 2), "Add property button should exist")
    }

    func testPropertyCardDisplay() throws {
        let firstPropertyCard = app.buttons["property_card_cm7gijdee000ly7i82uq0qf35"]
        XCTAssertTrue(firstPropertyCard.waitForExistence(timeout: 5), "First property card should exist")
        XCTAssertTrue(firstPropertyCard.isHittable, "First property card should be hittable")

        let secondPropertyCard = app.buttons["property_card_cm7gijdee000ly7i82uq0qf36"]
        XCTAssertTrue(secondPropertyCard.waitForExistence(timeout: 5), "Second property card should exist")
        XCTAssertTrue(secondPropertyCard.isHittable, "Second property card should be hittable")

        let addressLabel = secondPropertyCard.staticTexts["text_address"]
        XCTAssertTrue(addressLabel.exists, "Address should be visible")

        let tenantLabel = secondPropertyCard.staticTexts["text_tenant"]
        XCTAssertTrue(tenantLabel.exists, "Tenant should be visible")

        let leaseStartDateLabel = secondPropertyCard.staticTexts["text_started_on"]
        XCTAssertTrue(leaseStartDateLabel.exists, "Lease start date should be visible")

        let busyLabel = secondPropertyCard.staticTexts["text_busy"]
        XCTAssertTrue(busyLabel.exists, "Busy status should be visible")
    }

    func testDetailsButton() throws {
        let firstPropertyCard = app.buttons["property_card_cm7gijdee000ly7i82uq0qf35"]
        XCTAssertTrue(firstPropertyCard.waitForExistence(timeout: 2), "First property card should exist")
        firstPropertyCard.tap()

        let detailsTitle = app.staticTexts["Property Details"].exists || app.staticTexts["Détails du bien"].exists
        XCTAssertTrue(detailsTitle, "Property Details title should be visible")
    }

    func testPropertyDetailsView() throws {
        let firstPropertyCard = app.buttons["property_card_cm7gijdee000ly7i82uq0qf35"]
        XCTAssertTrue(firstPropertyCard.waitForExistence(timeout: 2), "First property card should exist")
        firstPropertyCard.tap()

        let aboutSectionHeader = app.staticTexts["About the property"].exists || app.staticTexts["À propos du bien"].exists
        XCTAssertTrue(aboutSectionHeader, "About section should be visible")

        let startInventoryButton = app.buttons["inventory_btn_start"]
        XCTAssertTrue(startInventoryButton.waitForExistence(timeout: 2), "Start inventory button should exist")
    }

    func testAboutCardView() throws {
        let secondPropertyCard = app.buttons["property_card_cm7gijdee000ly7i82uq0qf36"]
        XCTAssertTrue(secondPropertyCard.waitForExistence(timeout: 2), "Second property card should exist")
        secondPropertyCard.tap()

        let tenantInfo = app.staticTexts["Jean Dupont"]
        XCTAssertTrue(tenantInfo.waitForExistence(timeout: 2), "Tenant 'Jean Dupont' should be visible")

        let surfaceInfo = app.staticTexts["Area: 65 m²"].exists || app.staticTexts["Surface : 65 m²"].exists
        XCTAssertTrue(surfaceInfo, "Surface '65 m²' should be visible")

        let leaseStartDate = app.staticTexts["lease_start_date"] // Changement ici
        XCTAssertTrue(leaseStartDate.waitForExistence(timeout: 2), "Lease start date should be visible")

        let rentInfo = app.staticTexts["Rent: 1500€"].exists || app.staticTexts["Loyer: 1500€"].exists
        XCTAssertTrue(rentInfo, "Rent '1500€' should be visible")

        let leaseEndDate = app.staticTexts["End: No end date assigned"].exists || app.staticTexts["Fin : Pas de date renseignée"].exists
        XCTAssertTrue(leaseEndDate, "No end date should be visible")

        let depositInfo = app.staticTexts["Deposit: 3000€"].exists || app.staticTexts["Caution : 3000€"].exists
        XCTAssertTrue(depositInfo, "Deposit '3000€' should be visible")

        if !tenantInfo.exists { print(app.debugDescription) }
    }

    func testDocumentsGrid() throws {
        let firstPropertyCard = app.buttons["property_card_cm7gijdee000ly7i82uq0qf35"]
        XCTAssertTrue(firstPropertyCard.waitForExistence(timeout: 2), "First property card should exist")
        firstPropertyCard.tap()

        let documentsSectionHeader = app.staticTexts["documents_header"]
        XCTAssertTrue(documentsSectionHeader.waitForExistence(timeout: 2), "Documents section should exist")
    }

    func testStartInventoryButton() throws {
        let firstPropertyCard = app.buttons["property_card_cm7gijdee000ly7i82uq0qf35"]
        XCTAssertTrue(firstPropertyCard.waitForExistence(timeout: 2), "First property card should exist")
        firstPropertyCard.tap()

        let startInventoryButton = app.buttons["inventory_btn_start"]
        XCTAssertTrue(startInventoryButton.waitForExistence(timeout: 2), "Start inventory button should exist")
        startInventoryButton.tap()
    }

    func testCreatePropertyViewElements() throws {
        let addButton = app.buttons["add_property"]
        XCTAssertTrue(addButton.waitForExistence(timeout: 2), "Add property button should exist")
        addButton.tap()

        XCTAssertTrue(app.textFields["Name_textfield"].waitForExistence(timeout: 2), "Name field should exist")
        XCTAssertTrue(app.textFields["Address_textfield"].exists, "Address field should exist")
        XCTAssertTrue(app.textFields["City_textfield"].exists, "City field should exist")
        XCTAssertTrue(app.textFields["Postal Code_textfield"].exists, "Postal Code field should exist")
        XCTAssertTrue(app.textFields["Country_textfield"].exists, "Country field should exist")

        app.swipeUp()

        XCTAssertTrue(app.textFields["Monthly Rent_textfield"].waitForExistence(timeout: 2), "Rent field should exist")
        XCTAssertTrue(app.textFields["Deposit_textfield"].exists, "Deposit field should exist")
        XCTAssertTrue(app.textFields["Surface (m²)_textfield"].exists, "Surface field should exist")
        XCTAssertTrue(app.buttons["cancel_button"].exists, "Cancel button should exist")
        XCTAssertTrue(app.buttons["confirm_button"].exists, "Confirm button should exist")
    }

    func testErrorAddProperty() throws {
        let addButton = app.buttons["add_property"]
        XCTAssertTrue(addButton.waitForExistence(timeout: 2), "Add property button should exist")
        addButton.tap()

        let addPropertyButton = app.buttons["confirm_button"]
        XCTAssertTrue(addPropertyButton.waitForExistence(timeout: 2), "Confirm button should exist")
        addPropertyButton.tap()
    }

    func testCancelPropertyCreation() throws {
        let addButton = app.buttons["add_property"]
        XCTAssertTrue(addButton.waitForExistence(timeout: 2), "Add property button should exist")
        addButton.tap()

        let cancelButton = app.buttons["cancel_button"]
        XCTAssertTrue(cancelButton.waitForExistence(timeout: 2), "Cancel button should exist")
        cancelButton.tap()

        let firstPropertyCard = app.buttons["property_card_cm7gijdee000ly7i82uq0qf35"]
        XCTAssertTrue(firstPropertyCard.waitForExistence(timeout: 2), "First property card should exist after cancel")
    }
}
