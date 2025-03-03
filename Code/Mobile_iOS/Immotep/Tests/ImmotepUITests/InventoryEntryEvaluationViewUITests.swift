//
//  InventoryEntryEvaluationViewUITests.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 03/03/2025.
//

import XCTest

final class InventoryEntryEvaluationViewUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        app.launchArguments = ["-skipLogin", "-inventoryEntryEvaluationView"]
        print("XCUITest - Launch arguments set to: \(app.launchArguments)")
        app.launch()
    }

    func testEntryEvaluationViewElements() throws {
        let topBarTitle = app.staticTexts["Inventory"]
        XCTAssertTrue(topBarTitle.exists, "The 'Inventory' title should be visible")

        let commentsLabel = app.staticTexts["Comments"]
        XCTAssertTrue(commentsLabel.exists, "The 'Comments' label should be visible")

        let textEditor = app.textViews.element
        XCTAssertTrue(textEditor.exists, "The text editor for comments should be visible")

        let statusLabel = app.staticTexts["Status"]
        XCTAssertTrue(statusLabel.exists, "The 'Status' label should be visible")

//        let statusPicker = app.otherElements["StatusPicker"]
//        XCTAssertTrue(statusPicker.exists, "The status picker should be visible")

        let sendReportButton = app.buttons["Send Report"]
        XCTAssertTrue(sendReportButton.exists, "The 'Send Report' button should be visible")
    }

    func testAddPictureButtonShowsOptions() throws {
        let addPictureButton = app.buttons["plus"]
        XCTAssertTrue(addPictureButton.exists, "The 'Add Picture' button should be visible")

        addPictureButton.tap()

        let actionSheet = app.sheets.element
        XCTAssertTrue(actionSheet.waitForExistence(timeout: 5), "The image source action sheet should appear within 5 seconds")

        let takePhotoOption = actionSheet.buttons["Take Photo"]
        XCTAssertTrue(takePhotoOption.exists, "The 'Take Photo' option should be visible")

        let chooseLibraryOption = actionSheet.buttons["Choose from Library"]
        XCTAssertTrue(chooseLibraryOption.exists, "The 'Choose from Library' option should be visible")
    }

    func testSendReportChangesToValidate() throws {
        let sendReportButton = app.buttons["Send Report"]
        XCTAssertTrue(sendReportButton.exists, "The 'Send Report' button should be visible")

//        sendReportButton.tap()
//
//        let validateButton = app.buttons["Validate"]
//        XCTAssertTrue(validateButton.waitForExistence(timeout: 5), "The 'Validate' button should appear within 5 seconds after sending the report")
    }
}
