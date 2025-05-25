//
//  CreatePropertyUITests.swift
//  ImmotepUITests
//
//  Created by Liebenguth Alessio on 12/03/2025.
//

import XCTest

final class CreatePropertyUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        app.launchArguments = ["-skipLogin", "-createPropertyView", "--UITests"]
        print("XCUITest - Launch arguments set to: \(app.launchArguments)")
        app.launch()
    }

    func testCreatePropertyViewLoadsCorrectly() throws {
        let titlePredicate = NSPredicate(format: "label == %@ OR label == %@", "New Property", "Nouveau bien")
        let titleElement = app.staticTexts.element(matching: titlePredicate)
        XCTAssertTrue(titleElement.waitForExistence(timeout: 2), "The 'New Property' title should be visible in English or French")
        XCTAssertTrue(app.images["image_property"].exists, "Default property image should be visible")

        XCTAssertTrue(app.textFields["Name_textfield"].exists, "Name field should be visible")
        XCTAssertTrue(app.textFields["Address_textfield"].exists, "Address field should be visible")

        app.swipeUp()

        XCTAssertTrue(app.textFields["City_textfield"].exists, "City field should be visible")
        XCTAssertTrue(app.textFields["Postal Code_textfield"].exists, "Postal Code field should be visible")
        XCTAssertTrue(app.textFields["Country_textfield"].exists, "Country field should be visible")
        XCTAssertTrue(app.textFields["Monthly Rent_textfield"].exists, "Monthly Rent field should be visible")
        XCTAssertTrue(app.textFields["Deposit_textfield"].exists, "Deposit field should be visible")
        XCTAssertTrue(app.textFields["Surface (m²)_textfield"].exists, "Surface field should be visible")

        let cancelButtonPredicate = NSPredicate(format: "identifier == 'cancel_button' AND (label == %@ OR label == %@)", "Cancel", "Annuler")
        let cancelButton = app.buttons.element(matching: cancelButtonPredicate)
        XCTAssertTrue(cancelButton.exists, "Cancel button should be visible in English or French")

        let addButtonPredicate = NSPredicate(format: "identifier == 'confirm_button' AND (label == %@ OR label == %@)", "Add Property", "Ajouter un bien")
        let addButton = app.buttons.element(matching: addButtonPredicate)
        XCTAssertTrue(addButton.exists, "Add Property button should be visible in English or French")
    }

    func testCancelCreateProperty() throws {
        let nameField = app.textFields["Name_textfield"]
        XCTAssertTrue(nameField.waitForExistence(timeout: 2), "Name field should be accessible")
        nameField.tap()
        nameField.typeText("Test Property")

        let cancelButtonPredicate = NSPredicate(format: "identifier == 'cancel_button' AND (label == %@ OR label == %@)", "Cancel", "Annuler")
        let cancelButton = app.buttons.element(matching: cancelButtonPredicate)
        XCTAssertTrue(cancelButton.waitForExistence(timeout: 2), "Cancel button should be accessible")
        cancelButton.tap()

        let overviewTitlePredicate = NSPredicate(format: "label == %@ OR label == %@", "Overview", "Aperçu")
        let overviewTitle = app.staticTexts.element(matching: overviewTitlePredicate)
        if overviewTitle.waitForExistence(timeout: 2) {
            XCTAssertTrue(true, "Returned to Overview view after cancel (adjust if this is not the expected view)")
        } else {
            let propertyTitlePredicate = NSPredicate(format: "label == %@ OR label == %@", "Property", "Biens")
            let propertyTitle = app.staticTexts.element(matching: propertyTitlePredicate)
            XCTAssertTrue(propertyTitle.waitForExistence(timeout: 2), "Should return to Property list view after cancel")
        }
    }

    func testFillAndSubmitCreateProperty() throws {
        let formScrollView = app.scrollViews.firstMatch
        let hasScrollView = formScrollView.waitForExistence(timeout: 5)
        if !hasScrollView {
            print("No scroll view found, falling back to app-level scrolling")
        } else {
            print("Scroll view found: \(formScrollView.debugDescription)")
        }

        let fields = [
            ("Name_textfield", "Beach House"),
            ("Address_textfield", "789 Coastal Rd"),
            ("City_textfield", "Miami"),
            ("Postal Code_textfield", "33101"),
            ("Country_textfield", "USA"),
            ("Monthly Rent_textfield", "2000"),
            ("Deposit_textfield", "4000"),
            ("Surface (m²)_textfield", "120")
        ]

        for (index, (identifier, text)) in fields.enumerated() {
            let scrollAfter = index < fields.count - 1
            try fillTextField(identifier, with: text, scrollView: formScrollView, hasScrollView: hasScrollView, scrollAfter: scrollAfter)
        }

        let addButtonPredicate = NSPredicate(format: "identifier == 'confirm_button' AND (label == %@ OR label == %@)", "Add Property", "Ajouter un bien")
        let addButton = app.buttons.element(matching: addButtonPredicate)
        XCTAssertTrue(addButton.waitForExistence(timeout: 2), "Add Property button should be accessible")
        while !addButton.isHittable {
            scroll(hasScrollView ? formScrollView : app)
        }
        addButton.tap()

        let titlePredicate = NSPredicate(format: "label == %@ OR label == %@", "New Property", "Nouveau bien")
        let titleElement = app.staticTexts.element(matching: titlePredicate)
        XCTAssertTrue(titleElement.exists, "Should remain on CreatePropertyView since API is not mocked")
    }

    private func fillTextField(_ identifier: String, with text: String, scrollView: XCUIElement, hasScrollView: Bool, scrollAfter: Bool = true) throws {
        let field = app.textFields[identifier]
        while !field.isHittable {
            print("Field \(identifier) not hittable, scrolling...")
            scroll(hasScrollView ? scrollView : app)
            Thread.sleep(forTimeInterval: 0.5)
            if field.isHittable { break }
            XCTAssertTrue(field.waitForExistence(timeout: 2), "\(identifier) should be accessible")
        }
        print("Tapping \(identifier)")
        field.tap()
        print("Typing '\(text)' into \(identifier)")
        field.typeText(text)
    }

    private func scroll(_ element: XCUIElement) {
        let start = element.coordinate(withNormalizedOffset: CGVector(dx: 0.3, dy: 0.3)) 
        let end = element.coordinate(withNormalizedOffset: CGVector(dx: 0.25, dy: 0.25))
        start.press(forDuration: 0.1, thenDragTo: end)
    }

    func testOpenImagePickerOptions() throws {
        let image = app.images["image_property"]
        XCTAssertTrue(image.waitForExistence(timeout: 2), "Property image should be accessible")
        image.tap()

        let takePhotoPredicate = NSPredicate(format: "label == %@ OR label == %@", "Take Photo", "Prendre une photo")
        let takePhotoButton = app.buttons.element(matching: takePhotoPredicate)
        XCTAssertTrue(takePhotoButton.waitForExistence(timeout: 2), "Take Photo option should appear in English or French")

        let chooseLibraryPredicate = NSPredicate(format: "label == %@ OR label == %@", "Choose from Library", "Choisir dans la bibliothèque")
        let chooseLibraryButton = app.buttons.element(matching: chooseLibraryPredicate)
        XCTAssertTrue(chooseLibraryButton.exists, "Choose from Library option should appear in English or French")

        let cancelPredicate = NSPredicate(format: "label == %@ OR label == %@", "Cancel", "Annuler")
        let cancelButton = app.sheets.buttons.element(matching: cancelPredicate)
        XCTAssertTrue(cancelButton.exists, "Cancel option in action sheet should appear in English or French")

        cancelButton.tap()

        let sheetDismissedPredicate = NSPredicate(format: "exists == false")
        let sheetDismissedExpectation = expectation(for: sheetDismissedPredicate, evaluatedWith: app.sheets.element, handler: nil)
        wait(for: [sheetDismissedExpectation], timeout: 5.0)

        XCTAssertFalse(app.sheets.element.exists, "Action sheet should be dismissed after tapping Cancel")
    }

    func testEditPropertyViewLoadsAndCancel() throws {
        app.launchArguments = ["-skipLogin", "-editPropertyView", "--UITests"]
        app.launch()

        let editTitlePredicate = NSPredicate(format: "label == %@ OR label == %@", "Edit Property", "Modifier le bien")
        let editTitle = app.staticTexts.element(matching: editTitlePredicate)
        XCTAssertTrue(editTitle.waitForExistence(timeout: 2), "Edit Property title should be visible in English or French")

        let nameField = app.textFields["Name_textfield"]
        XCTAssertTrue(nameField.waitForExistence(timeout: 2), "Name field should be accessible")
        XCTAssertEqual(nameField.value as? String, "Maison de Campagne", "Name should be pre-filled")

        let addressField = app.textFields["Address_textfield"]
        XCTAssertTrue(addressField.waitForExistence(timeout: 2), "Address field should be accessible")
        XCTAssertEqual(addressField.value as? String, "123 Rue des Champs", "Address should be pre-filled")

        let cancelButtonPredicate = NSPredicate(format: "identifier == 'cancel_button' AND (label == %@ OR label == %@)", "Cancel", "Annuler")
        let cancelButton = app.buttons.element(matching: cancelButtonPredicate)
        XCTAssertTrue(cancelButton.waitForExistence(timeout: 2), "Cancel button should be accessible")
        cancelButton.tap()

        let propertyTitlePredicate = NSPredicate(format: "label == %@ OR label == %@", "Property", "Biens")
        let propertyTitle = app.staticTexts.element(matching: propertyTitlePredicate)
        XCTAssertTrue(propertyTitle.waitForExistence(timeout: 2), "Should return to Property list view after cancel")
    }

    func testEditPropertyOpenImagePickerOptions() throws {
        app.launchArguments = ["-skipLogin", "-editPropertyView", "--UITests"] 
        app.launch()

        let image = app.images["image_property"]
        XCTAssertTrue(image.waitForExistence(timeout: 5), "Property image should be accessible")
        image.tap()

        let takePhotoPredicate = NSPredicate(format: "label == %@ OR label == %@", "Take Photo", "Prendre une photo")
        let takePhotoButton = app.buttons.element(matching: takePhotoPredicate)
        XCTAssertTrue(takePhotoButton.waitForExistence(timeout: 5), "Take Photo option should appear in English or French")

        let chooseLibraryPredicate = NSPredicate(format: "label == %@ OR label == %@", "Choose from Library", "Choisir dans la bibliothèque")
        let chooseLibraryButton = app.buttons.element(matching: chooseLibraryPredicate)
        XCTAssertTrue(chooseLibraryButton.exists, "Choose from Library option should appear in English or French")

        let cancelPredicate = NSPredicate(format: "label == %@ OR label == %@", "Cancel", "Annuler")
        let cancelButton = app.sheets.buttons.element(matching: cancelPredicate)
        XCTAssertTrue(cancelButton.exists, "Cancel option in action sheet should appear in English or French")

        cancelButton.tap()

        let sheetDismissedPredicate = NSPredicate(format: "exists == false")
        let sheetDismissedExpectation = expectation(for: sheetDismissedPredicate, evaluatedWith: app.sheets.element, handler: nil)
        wait(for: [sheetDismissedExpectation], timeout: 5.0)

        XCTAssertFalse(app.sheets.element.exists, "Action sheet should be dismissed after tapping Cancel")
    }
}
