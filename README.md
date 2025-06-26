# Movie Collection

The objective of this project is to develop an API to manage movies and user's favourites list. The main features include:
- User registration and authentication
- Movie management (CRUD operations)

To develop this API, I have followed a Hexagonal Architecture approach, focusing on future scalability and maintainability.

The project was built on GOlang, using the Gin framework for the API and GORM for database interactions. The database used is PostgreSQL.

## Architecture

The Hexagonal Architecture, also known as Ports and Adapters, is a software design pattern that emphasizes the separation of concerns and the independence of the core business logic from external systems. This architecture allows for easier testing, maintainability, and scalability of the application.

Following the Hexagonal Architecture, the project is structured into several layers which correspond to the different subfolders inside the `app` directory:
- `domain` Layer/Folder: Contains the entities.
- `port` Layer/Folder: Defines the interfaces for the application services and infrastructure. This allows for a complete abstraction of the business logic from the infrastructure, and makes testing and mocking possible.
- `service` Layer/Folder: Implements the business logic. The packages in this layer implement the `service` interfaces defined in the `port` layer.
- `adapter` Layer/Folder: Contains the implementation of the infrastructure interfaces defined in the `port` layer. This includes the HTTP handlers, database repositories, and any other infrastructure-related code. This allows for easier separation of everything that is not business logic, allowing for a cleaner architecture and easier testing. There are two types of adapters:
    - *Driver adapters*: They start the communication with the core `services`. An example of this is the HTTP handlers that receive requests and call the appropriate service methods.
    - *Driven adapters*: They are called by the core `services` to perform some action, such as database operations. An example of this is the GORM repositories that implement the interfaces defined in the `port` layer.
- `util` layer: These are various utility functions that are used throughout the application, such as cryptography.

This makes the application extremely flexible. The database ORM could be changed, but as long as the interfaces in the `port` layer are respected, the application will continue to work. The same applies to the HTTP handlers, which could be replaced with a different framework or even a different protocol without affecting the core business logic.

The application also includes unit tests. These tests are located within the same package as the code they test, following the Go convention of placing tests in the same directory as the code.

## Installation

The project can (and should) be run using Docker. To use it, follow these steps:
1. Clone the repository:
   ```bash
   git clone https://github.com/Acova/movie-collection.git
   ```
2. Navigate to the project directory:
   ```bash
   cd movie-collection
   ```
3. Copy the `.env.example` file to `.env` and fill in the required environment variables:
   ```bash
   cp .env.example .env
   ```
   Make sure to set all the required environment variables in the `.env` file, such as database connection details and JWT secret key. For example:
   ```
   DATABASE_HOST=localhost
   DATABASE_PORT=5432
   DATABASE_USER=your_db_user
   DATABASE_PASSWORD=your_db_password
   DATABASE_NAME=movie_collection
   JWT_SECRET_KEY=your_jwt_secret
   ```

4. Build and run the Docker containers:
   ```bash
   docker-compose up --build
   ```

5. The app is running, but we need to execute the migrations to create the database tables. To do this, run the following command:
   ```bash
   docker exec -it movie-collection go run migration/migrate.go
   ```
   This will execute the migrations and create the necessary tables in the database.

6. (Optional) You can also populate the database with some initial data by running:
   ```bash
   docker exec -it movie-collection go run fixtures/populate.go
   ```

7. The API will be available at `http://localhost:8080`. You can use tools like Postman or cURL to interact with the API endpoints.

## Tests
The project includes unit tests for the core business logic. To run the tests, you can use the following command:
```bash
docker exec -it movie-collection go test -v ./...
```

## Usage

The API provides several endpoints for managing movies. All the endpoints, except for the `User Registration`, are protected by JWT authentication. To obtain your JWT token, you need to log in with your credentials on the `/login`. 

Because the app uses the "github.com/appleboy/gin-jwt/v2" middleware, the `/login` endpoint is not present in the Swagger documentation. This endpoint expects a POST request with the following JSON body:
```json
{
  "email": "user_email",
  "password": "user_password"
}
```
and will return a JWT token if the credentials are valid. You can then use this token to access the protected endpoints by including it in the `Authorization` header of your requests.

You can access the API documentation at `http://localhost:8080/swagger/index.html` to see the available endpoints and their usage. But here is a brief overview of the main endpoints:

### User Management
#### User Registration
- **POST** `/user`: Register a new user. The request body should contain the user's email, name, and password in JSON format:
```json
{
  "email": "user_email",
  "name": "user_name",
  "password": "user_password"
}
```
#### User List
- **GET** `/user`: Retrieve a list of all users.

### Movie Management
#### Movie List
- **GET** `/movie`: Retrieve a list of all movies. You can filter the results by title, director, genre, and cast using query parameters:

#### Movie Details
- **GET** `/movie/{id}`: Retrieve details of a specific movie by its ID.

#### Add Movie
- **POST** `/movie`: Add a new movie. The request body should contain the movie details in JSON format:
```json
{
  "title": "Movie Title",
  "director": "Movie Director",
  "synopsis": "Movie Synopsis",
  "release_year": 2023,
  "cast": "Actor 1, Actor 2",
  "genre": "Movie Genre",
  "rating": 8.5,
  "duration": 120,
  "poster_url": "https://example.com/movie-poster.jpg"
}
```

#### Update Movie
- **PUT** `/movie/{id}`: Update an existing movie by its ID. The request body should contain the updated movie details in JSON format:
```json
{
  "title": "Updated Movie Title",
  "director": "Updated Movie Director",
  "synopsis": "Updated Movie Synopsis",
  "release_year": 2023,
  "cast": "Updated Actor 1, Updated Actor 2",
  "genre": "Updated Movie Genre",
  "rating": 8.5,
  "duration": 120,
  "poster_url": "https://example.com/movie-poster.jpg"
}
```

#### Delete Movie
- **DELETE** `/movie/{id}`: Delete a movie by its ID.