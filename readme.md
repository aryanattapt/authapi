# Auth API

This repository contains a simple authentication API built using Go. The API supports user signup and signin functionalities, utilizing JWT for session management and MongoDB for data storage.

## Structure of the Service

The service is structured as follows:
### Components:

- **controller:** Contains HTTP handler functions for managing requests and responses.
- **model:** Defines the data structures for users and request payloads.
- **repository:** Handles data operations with MongoDB.
- **service:** Contains business logic for signing up and signing in users.
- **utils:** Provides utility functions, including password hashing, JWT token generation and others.
- **route:** Defines the API routes mapping HTTP methods and endpoints to their corresponding controller functions, thereby facilitating structured and RESTful interactions within the application.
- **test:** Provides a suite of automated tests designed to validate the functionality and reliability of the application components

## Instructions to Run the Service

### Prerequisites

Before running the service, ensure you have the following installed:

- Go (1.16 or higher)
- MongoDB (local)

### Setup

1. **Install MongoDB:**
   - **For Windows:** Download the MongoDB installer from the [official MongoDB website](https://www.mongodb.com/try/download/community) and follow the installation instructions.
   - **For macOS:** You can use Homebrew:
     ```bash
     brew tap mongodb/brew
     brew install mongodb-community
     ```
   - **For Linux:** Follow the instructions on the [MongoDB installation documentation](https://docs.mongodb.com/manual/installation/).
   - Make sure your local mongodb run at port 27017 without any password

2. **Clone the Repository:**
   ```bash
   git clone https://github.com/yourusername/authapi.git
   cd authapi

3. **Install Dependency:**
   ```bash
   go mod tidy

4. **Run app:**
   ```bash
   go run main.go

5. **Run Tests:**
   - Make sure you have postman install in your machine. 
   - If its not installed then open your browser than paste this link https://www.postman.com/downloads/ and then click download.
   - If its already install open it. Then import the test file in folder test
   - After importing, you will see the collection in the left sidebar.
   - Click on the collection to view the individual requests.
   - You can run the entire collection by clicking on the **Run** button (▶) in the top right corner.
   - This will open the **Collection Runner** where you can execute the tests.