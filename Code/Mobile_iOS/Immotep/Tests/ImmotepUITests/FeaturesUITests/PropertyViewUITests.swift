//
//  PropertyViewUITests.swift
//  ImmotepUITests
//
//  Created by Liebenguth Alessio on 09/12/2024.
//

import XCTest
import Combine
import Foundation
import SwiftUI
@testable import Immotep

final class PropertyViewUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        app.launchArguments = ["-skipLogin"]
        app.launchArguments += ["-propertyList"]
        app.launch()
    }

    // Property View Tests

    func testAddPropertyButton() throws {
        let addButton = app.buttons["add_property"]
        XCTAssertTrue(addButton.exists)
    }

    func testPropertyCardDisplay() throws {
        let propertyCard = app.scrollViews.children(matching: .other).element(boundBy: 0)

        XCTAssertTrue(propertyCard.exists)

        let addressLabel = propertyCard.staticTexts["text_address"]
        XCTAssertTrue(addressLabel.exists)

        let tenantLabel = propertyCard.staticTexts["text_tenant"]
        XCTAssertTrue(tenantLabel.exists)

        let leaseStartDateLabel = propertyCard.staticTexts["text_started_on"]
        XCTAssertTrue(leaseStartDateLabel.exists)

        let availableLabel = propertyCard.staticTexts["text_available"]
        XCTAssertTrue(availableLabel.exists)
    }

    func testDetailsButton() throws {
        let firstPropertyCard = app.scrollViews.children(matching: .other).element(boundBy: 0)

        XCTAssertTrue(firstPropertyCard.exists)
        firstPropertyCard.tap()

        let detailsTitle = app.staticTexts["Property Details"].exists || app.staticTexts["Details"].exists
        XCTAssertTrue(detailsTitle)
    }

    // Property Details View Tests

    func testPropertyDetailsView() throws {
        let firstPropertyCard = app.scrollViews.children(matching: .other).element(boundBy: 0)
        XCTAssertTrue(firstPropertyCard.exists)
        firstPropertyCard.tap()

        let aboutSectionHeader = app.staticTexts["About the property"].exists || app.staticTexts["À propos"].exists
        XCTAssertTrue(aboutSectionHeader)

        //        let documentsSectionHeader = app.staticTexts["documents_header"]
        //        XCTAssertTrue(documentsSectionHeader.exists)

        let startInventoryButton = app.buttons["inventory_btn_start"]
        XCTAssertTrue(startInventoryButton.exists)
    }

    func testAboutCardView() throws {
        let firstPropertyCard = app.scrollViews.children(matching: .other).element(boundBy: 0)
        XCTAssertTrue(firstPropertyCard.exists)
        firstPropertyCard.tap()

        let tenantInfo = app.staticTexts["John Doe"]
        XCTAssertTrue(tenantInfo.exists)

        let surfaceInfo = app.staticTexts["Area: 60.50 m²"].exists || app.staticTexts["Surface : 60.50 m²"].exists
        XCTAssertTrue(surfaceInfo)

        let leaseStartDate = app.staticTexts["Start date: 10 déc. 2024"].exists || app.staticTexts["Date de début : 10 déc. 2024"].exists
        XCTAssertTrue(leaseStartDate)

        let rentInfo = app.staticTexts["Rent: 1500€"].exists || app.staticTexts["Loyer: 1500€"].exists
        XCTAssertTrue(rentInfo)

        let leaseEndDate = app.staticTexts["End: No end date assigned"].exists || app.staticTexts["Fin : Pas de date renseignée"].exists
        XCTAssertTrue(leaseEndDate)

        let depositInfo = app.staticTexts["Deposit: 3000€"].exists || app.staticTexts["Caution : 3000€"].exists
        XCTAssertTrue(depositInfo)
    }

    func testDocumentsGrid() throws {
        let firstPropertyCard = app.scrollViews.children(matching: .other).element(boundBy: 0)
        XCTAssertTrue(firstPropertyCard.exists)
        firstPropertyCard.tap()

        let document1 = app.staticTexts["Lease Agreement"]
        XCTAssertTrue(document1.exists)

        let document2 = app.staticTexts["Inspection Report"]
        XCTAssertTrue(document2.exists)
    }

    func testStartInventoryButton() throws {
        let firstPropertyCard = app.scrollViews.children(matching: .other).element(boundBy: 0)
        XCTAssertTrue(firstPropertyCard.exists)
        firstPropertyCard.tap()

        let startInventoryButton = app.buttons["inventory_btn_start"]
        XCTAssertTrue(startInventoryButton.exists)

        startInventoryButton.tap()
    }

    // Property Create View Tests

    func testCreatePropertyViewElements() throws {
        let addButton = app.buttons["add_property"]
        XCTAssertTrue(addButton.exists)
        addButton.tap()

        let nameField = app.textFields["Name_textfield"].exists
        XCTAssertTrue(nameField)

        let addressField = app.textFields["Address_textfield"].exists
        XCTAssertTrue(addressField)

        let cityField = app.textFields["City_textfield"].exists
        XCTAssertTrue(cityField)

        let postalCodeField = app.textFields["Postal Code_textfield"].exists
        XCTAssertTrue(postalCodeField)

        let countryField = app.textFields["Country_textfield"].exists
        XCTAssertTrue(countryField)

        let rentField = app.textFields["Monthly Rent_textfield"]
        if !rentField.exists {
            app.swipeUp()
        }

        XCTAssertTrue(rentField.exists)

        let depositField = app.textFields["Deposit_textfield"].exists
        XCTAssertTrue(depositField)

        let surfaceField = app.textFields["Surface (m²)_textfield"].exists
        XCTAssertTrue(surfaceField)

        let cancelButton = app.buttons["cancel_button"].exists
        XCTAssertTrue(cancelButton)

        let confirmButton = app.buttons["confirm_button"].exists
        XCTAssertTrue(confirmButton)
    }

    func testErrorAddProperty() throws {
        let addButton = app.buttons["add_property"]
        XCTAssertTrue(addButton.exists)
        addButton.tap()

        let addPropertyButton = app.buttons["confirm_button"]
        XCTAssert(addPropertyButton.exists)

        addPropertyButton.tap()
    }

    func testCancelPropertyCreation() throws {
        let addButton = app.buttons["add_property"]
        XCTAssertTrue(addButton.exists)
        addButton.tap()

        let cancelButton = app.buttons["cancel_button"]
        XCTAssertTrue(cancelButton.exists)
        cancelButton.tap()

        let firstPropertyCard = app.scrollViews.children(matching: .other).element(boundBy: 0)
        XCTAssertTrue(firstPropertyCard.exists)
    }

}
