# AXIOS-CP-WING
Platform for listing and tracking solved CP problems for students and the Competitive Programming Club of IIIT Lucknow.

## Features
- [x] User login/signup with `@iiitl.ac.in` email for tracking personal solved problems.
- [X] Listing of Coding Tasks under different lists (Getting Started/Internship Prep/ Pro CP) with problem Tags.
- [X] Unit Testing for API handlers using Go-Mock
- [X] Github Workflow for Testing Compilability and Running Unit Tests on Push
- [X] Admin User with permissions to dynamically create/edit problem lists.
- [ ] Email Verification and Password Reset feature through OTP on Gmail.
- [ ] Problems sortable with tags.
- [ ] Logging to keep track of Admin actions. 

- [ ] Deployment using proper CI/CD  


## Databse Schema (`Not fully implemented  as of now`)
![DB Schema](./resources/images/DB%20Schema.jpg)



# API Documentation (`To Be Updated`)



## HealthCheck APIs


### `1. GET  /ping`

#### Description

- Simple ping API to check if server is online.
- To be modified later to pull out system health logs etc.


#### Request

- Method: GET
- Path: `/ping`

#### Response

- Status Code: 200 (OK)
- Content Type: application/json


#### Example

**Request Body**
```json
[]
```

**cURL snippet**
```bash
curl  -X GET \
  'localhost:8080/ping'
```

**Response Body**
```json
{
    "message": "Pong"
}
```

