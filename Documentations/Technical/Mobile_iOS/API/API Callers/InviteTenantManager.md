API Documentation for Tenant Invitation Management
==================================================

Overview
--------

This document provides details on the API calls related to tenant invitation management in the application. It includes information on sending and canceling tenant invitations for properties.

Tenant Invitation Management
----------------------------

### Overview

The TenantViewModel and OwnerPropertyViewModel classes handle tenant invitation-related operations. These include sending invitations to tenants for a specific property and canceling pending invitations.

### Data Models

#### InviteRequest

Represents the request body for sending a tenant invitation.

| Field        | Type   | Description                                                                 |
|--------------|--------|-----------------------------------------------------------------------------|
| tenantEmail  | String | The email address of the tenant to invite.                                  |
| startDate    | String | The start date of the lease in ISO 8601 format (e.g., "2025-07-08T00:00:00Z"). |
| endDate      | String? | The end date of the lease in ISO 8601 format (optional).                   |

#### InviteIDResponse

The response returned after successfully sending an invitation.

| Field | Type   | Description                              |
|-------|--------|------------------------------------------|
| id    | String | The unique identifier of the created invitation. |


### Functions

#### inviteTenant(propertyId: String, email: String, startDate: Date, endDate: Date?)

Sends an invitation to a tenant for a specific property.

*   **Endpoint**: POST /owner/properties/{propertyId}/send-invite/
    
*   **Description**: Sends a lease invitation to a tenant for the specified property, including the lease start date and an optional end date.
    
*   **Parameters**:
    
    *   propertyId: The unique identifier of the property.
        
    *   email: The email address of the tenant.
        
    *   startDate: The lease start date (converted to ISO 8601 format).
        
    *   endDate: The lease end date (optional, converted to ISO 8601 format).
        
*   **Request Headers**:
    
    *   Authorization: Bearer {token}: The authentication token.
        
    *   Content-Type: application/json
        
    *   Accept: application/json
        
*   { "tenant\_email": "string", "start\_date": "string (ISO 8601)", "end\_date": "string (ISO 8601, optional)"}
    
*   **Responses**:
    
    *   { "id": "string"}
        
        *   The invitation was successfully created, and the response includes the invitation ID.
            
    *   **400 Bad Request**:
        
        *   Error: Missing or invalid fields in the request body.
            
        *   Example: { "error": "Missing fields" }
            
    *   **403 Forbidden**:
        
        *   Error: The property does not belong to the authenticated user.
            
        *   Example: { "error": "Property is not yours" }
            
    *   **404 Not Found**:
        
        *   Error: The specified property was not found.
            
        *   Example: { "error": "Property not found" }
            
    *   **409 Conflict**:
        
        *   Error: An invitation for this email already exists for the property.
            
        *   Example: { "error": "Invite already exists for this email" }
            
    *   **Default**:
        
        *   Error: Unexpected server error with status code.
            
        *   Example: { "error": "Unexpected error: {statusCode}" }
            

#### cancelInvite(propertyId: String, token: String)

Cancels a pending tenant invitation for a specific property.

*   **Endpoint**: DELETE /owner/properties/{propertyId}/cancel-invite/
    
*   **Description**: Cancels a pending lease invitation for the specified property.
    
*   **Parameters**:
    
    *   propertyId: The unique identifier of the property.
        
    *   token: The authentication token.
        
*   **Request Headers**:
    
    *   Authorization: Bearer {token}: The authentication token.
        
    *   Accept: application/json
        
*   **Request Body**: None
    
*   **Responses**:
    
    *   **204 No Content**:
        
        *   The invitation was successfully canceled.
            
    *   **403 Forbidden**:
        
        *   Error: The property does not belong to the authenticated user.
            
        *   Example: { "error": "Property is not yours" }
            
    *   **404 Not Found**:
        
        *   Error: No pending lease invitation was found for the property.
            
        *   Example: { "error": "No pending lease" }
            
    *   **500 Internal Server Error**:
        
        *   Error: An internal server error occurred.
            
        *   Example: { "error": "Internal server error: {details}" }
            
    *   **Default**:
        
        *   Error: Unexpected server error with status code.
            
        *   Example: { "error": "Failed with status code: {statusCode} - {details}" }