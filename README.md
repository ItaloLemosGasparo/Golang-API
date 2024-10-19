# Project Description: API in Golang

This project is a RESTful API developed in Go (Golang), designed to manage users, products, suppliers, and shopping carts. The API uses various popular libraries and frameworks to facilitate development and maintenance.

## Technologies Used
- **Go:** Main programming language.
- **Godotenv:** For environment variable management.
- **Gin:** Framework for building web APIs.
- **GORM:** ORM (Object Relational Mapping) for interacting with the database.
- **JWT-Go:** For implementing JWT token-based authentication.
- **bcrypt:** For password encryption.
- **PostgreSQL:** Database used for data storage.

## Main Features
- **User Management:**
  - Registration, updating, and deletion of users.
  - Updating passwords and additional information (such as phone and address).
  - User authentication via login.
  
- **Product Management:**
  - Registration, updating, and querying of products.
  
- **Supplier Management:**
  - Registration, updating, and querying of suppliers.
  
- **Cart and Favorites Management:**
  - Adding products to favorites and shopping cart.
  - Querying products in the cart and favorites.

## Project Structure Overview
### General Structure
The project is built in Go using the Gin framework for REST API development. The code is organized into packages, each with a specific responsibility.

### Main Packages
1. **controllers**
   - Contains the controllers that manage routes and business logic. Here are some of the main functions:
     - **User Management:** Manages the creation and updating of users, including password validation.
     - **Supplier Management:** Allows for the registration, updating, searching, and deletion of suppliers.
     - **Product Management:** Controls the operations of registration, updating, searching, and deletion of products.
     - **User Authentication:** Implements login and JWT token generation for authentication and encryption.

2. **models**
   - Defines the data structures representing the system's entities, such as:
     - **User:** Stores user information, including contact data and privileges.
     - **Passwords:** Stores the encrypted password associated with a user.
     - **Supplier:** Represents a supplier with their contact and identification details.
     - **Product:** Represents a product with related information, such as price and description.
     - **Cart:** Manages the items that a user adds for purchase.
     - **Favorites:** Allows users to mark products as favorites.

3. **initializers**
   - Contains functions for configuring and initializing the application, including:
     - **Database Connection:** Manages the connection to the database using GORM.
     - **Environment Variable Loading:** Loads the necessary configuration variables for the application.

4. **migration.go**
   - Responsible for database migration. Utilizes GORM's AutoMigrate method to create and update tables according to the defined models.

## Main Functionalities
- **User Registration and Authentication:** Allows users to register and log in using encrypted passwords.
- **Supplier and Product Management:** Offers complete CRUD (create, read, update, delete) functionality for suppliers and products.
- **Favorites and Shopping Cart:** Allows users to add products to favorites and manage their shopping cart.

## Conclusion
The project follows a modular architecture, where each package has a specific responsibility. This facilitates the maintenance and expansion of the system, allowing for the organized addition of new features.
