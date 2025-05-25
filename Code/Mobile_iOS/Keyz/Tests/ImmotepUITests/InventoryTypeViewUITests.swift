//
//  InventoryTypeViewUITests.swift
//  Immotep
//
//  Created by Liebenguth Alessio on 02/03/2025.
//

import XCTest
@testable import Keyz

final class InventoryTypeViewUITests: XCTestCase {
    let app = XCUIApplication()

    override func setUpWithError() throws {
        continueAfterFailure = false
        app.launchArguments = ["-skipLogin"]
        app.launchArguments += ["-inventoryTypeView"]
        Thread.sleep(forTimeInterval: 1.0)
        app.launch()
    }

    func testInventoryTypeViewElements() throws {
        let topBarTitle = app.staticTexts["Inventory"]
        XCTAssertTrue(topBarTitle.exists, "The 'Inventory' title should be visible")

        let entryInventoryButton = app.buttons["Entry Inventory"]
        XCTAssertTrue(entryInventoryButton.exists, "The 'Entry Inventory' button should be visible")

        let exitInventoryButton = app.buttons["Exit Inventory"]
        XCTAssertTrue(exitInventoryButton.exists, "The 'Exit Inventory' button should be visible")

        let entryIcon = app.images["figure.walk.arrival"]
        XCTAssertTrue(entryIcon.exists, "The entry icon should be visible")

        let exitIcon = app.images["figure.walk.departure"]
        XCTAssertTrue(exitIcon.exists, "The exit icon should be visible")
    }

    func testNavigateToEntryInventory() throws {
        let entryInventoryButton = app.buttons["Entry Inventory"]
        XCTAssertTrue(entryInventoryButton.exists, "The 'Entry Inventory' button should be visible")

        entryInventoryButton.tap()

        let inventoryRoomTitle = app.staticTexts["Entry Inventory"]
        XCTAssertTrue(inventoryRoomTitle.waitForExistence(timeout: 2), "The 'Entry Inventory' view should be displayed after navigation")
    }

    func testNavigateToExitInventory() throws {
        let exitInventoryButton = app.buttons["Exit Inventory"]
        XCTAssertTrue(exitInventoryButton.exists, "The 'Exit Inventory' button should be visible")

        exitInventoryButton.tap()

        let inventoryRoomTitle = app.staticTexts["Exit Inventory"]
        XCTAssertTrue(inventoryRoomTitle.waitForExistence(timeout: 2), "The 'Exit Inventory' view should be displayed after navigation")
    }

    func testInitialState() throws {
        let inventoryRoomTitle = app.staticTexts["Inventory Room"]
        XCTAssertFalse(inventoryRoomTitle.exists, "The 'Inventory Room' view should not be displayed initially")
    }
}
