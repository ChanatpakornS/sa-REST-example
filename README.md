# REST Example APIs

### Model: Concert
- ID (uint, primary key)
- Title (string)
- Description (string)
- Location (string)
- Date (time.Time)
- Capacity (int)
- Price (float64)

### Endpoints
- `GET /health`: check health
- `GET /concerts`: Get all concerts
- `GET /concerts/:id`: Get a specific concert by ID
- `POST /concerts`: Create a new concert
- `PUT /concerts/:id`: Update an existing concert by ID
- `DELETE /concerts/:id`: Delete a concert by ID

### Stack
- Golang (Fiber V3)
- GORM (PostgreSQL)

### Run Locally

<code>make run</code> - Run the application locally
<code>make compose</code> - Run the Database