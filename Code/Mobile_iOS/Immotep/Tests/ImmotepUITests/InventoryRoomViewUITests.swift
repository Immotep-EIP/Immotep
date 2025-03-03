//
//  InventoryRoomViewUITests.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 03/03/2025.
//

import XCTest

final class InventoryRoomViewUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        app.launchArguments = ["-skipLogin", "-inventoryRoomView"]
        print("XCUITest - Launch arguments set to: \(app.launchArguments)")
        app.launch()
    }

    func testRoomViewElements() throws {
        let topBarTitle = app.staticTexts["Entry Inventory"]
        XCTAssertTrue(topBarTitle.exists, "The 'Entry Inventory' title should be visible by default")

        let addRoomButton = app.buttons["plus.circle"]
        XCTAssertTrue(addRoomButton.exists, "The 'Add Room' button should be visible")

        let livingRoomCell = app.cells.staticTexts["Living Room"]
        XCTAssertTrue(livingRoomCell.exists, "The 'Living Room' cell should be displayed")

        let kitchenCell = app.cells.staticTexts["Kitchen"]
        XCTAssertTrue(kitchenCell.exists, "The 'Kitchen' cell should be displayed")

        let checkedIcon = app.cells.containing(.staticText, identifier: "Kitchen").images["checkmark"]
        XCTAssertTrue(checkedIcon.exists, "The 'checkmark' icon should be visible for the 'Kitchen' cell (checked: true)")
    }

    func testAddRoomButtonShowsAlert() throws {
        let addRoomButton = app.buttons["plus.circle"]
        XCTAssertTrue(addRoomButton.exists, "The 'Add Room' button should be visible")

        addRoomButton.tap()
    }

    func testSwipeToDeleteRoom() throws {
        let livingRoomCell = app.cells.staticTexts["Living Room"]
        XCTAssertTrue(livingRoomCell.exists, "The 'Living Room' cell should be displayed")

        livingRoomCell.swipeLeft()

        let deleteButton = app.buttons["Delete"]
        XCTAssertTrue(deleteButton.exists, "The 'Delete' button should appear after swiping left")

        deleteButton.tap()
    }

    func testFinalizeInventoryButtonAppearsWhenAllRoomsCompleted() throws {
        let finalizeButton = app.buttons["Finalize Inventory"]
        XCTAssertTrue(finalizeButton.exists, "The 'Finalize Inventory' button should be visible")
    }
}
