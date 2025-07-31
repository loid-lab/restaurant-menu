# Restaurant Menu API

This is a **Go-based backend API** for a local restaurant. It's designed to be a RESTful service to handle everything from menu management to customer orders and payments.

---

## 🔧 Tech Stack

- **Go**
- **PostgreSQL** (with Neon as a provider)
- **JWT Auth** for secure authentication
- **Cloudinary** for image management (e.g., menu item photos)

---

## 📦 Features

- ✅ **User Authentication**: Signup, login, and JWT-based authentication.
- 👥 **Role-Based Access Control**:
    - **User**: Can browse the menu, place orders, and view their history.
    - **Staff**: All user permissions, plus can manage menu items.
    - **Admin**: All staff permissions, plus can manage users and access sales reports.
-  меню **Menu Management**: Full CRUD (Create, Read, Update, Delete) operations for menu items, restricted to staff and admins.
- 🛒 **Shopping Cart**: Persistent cart for users to add items before checkout.
- 📦 **Order Management**: Track orders from placement to completion.
- 💳 **Payment Processing**: Integration with a payment provider to handle transactions.
- 🧾 **Invoicing**: Automatic generation and logging of invoices for each order.
- 📊 **Sales Reporting**: Endpoints for admins to view sales statistics.

---

## 🗂 Folder Structure

```
restaurant-menu/
├── controllers/    # HTTP handlers for routes
├── initializers/   # Database, Cloudinary, and environment variable configuration
├── middleware/     # JWT authentication and role-checking middleware
├── models/         # GORM models for the database schema
├── utils/          # Reusable utility functions (e.g., for invoices, mail)
├── main.go         # Application entry point
├── go.mod
├── go.sum
├── .env            # Environment variables (ignored by git)
├── docker-compose.yaml
└── README.md
```

---

## 📌 Setup Instructions

### 1. Clone the repo
```bash
git clone <https://github.com/loid-lab/restaurant-menu-git>
cd restaurant-menu
```

### 2. Setup Environment Variables
Create a `.env` file in the root directory. You'll need to add your specific credentials.

```env
# Neon Database URL
DATABASE_URL="postgres://user:password@host:5432/dbname"

# JWT Secret
SECRET="your-super-secret-key"

# SMTP Credentials (for sending emails like invoices)
SMTP_HOST="smtp.example.com"
SMTP_PORT=587
SMTP_USER="your-email@example.com"
SMTP_PASS="your-email-password"

# Cloudinary Credentials
CLOUDINARY_CLOUD_NAME="your_cloud_name"
CLOUDINARY_API_KEY="your_api_key"
CLOUDINARY_API_SECRET="your_api_secret"
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run the Application

#### Directly
```bash
go run main.go
```

#### With Docker Compose
This setup is configured to work with a remote database like Neon. Just ensure your `.env` file is populated with the `DATABASE_URL`.

```bash
docker-compose up --build
```
The application will be available at `http://localhost:8080`.

---

## 🔐 Authentication

- `POST /auth/signup` — Register a new user.
- `POST /auth/login` — Log in to receive a JWT.
- Authenticated routes require an `Authorization: Bearer <token>` header.
- Admin and Staff routes are protected by middleware that checks the user's role.

---

## 📝 API Endpoints

Here is a list of the available API endpoints:

### Authentication

- `POST /auth/signup`: Register a new user.
- `POST /auth/login`: Login and receive a JWT.

### User

- `GET /user/profile`: Get the current user's profile.

### Menu Items

- `GET /menu-items`: Get all menu items.
- `GET /menu-items/:id`: Get a single menu item by ID.
- `POST /menu-items`: Create a new menu item (Staff/Admin only).
- `PUT /menu-items/:id`: Update a menu item (Staff/Admin only).
- `DELETE /menu-items/:id`: Delete a menu item (Staff/Admin only).

### Cart

- `GET /cart`: Get the current user's cart.
- `POST /cart`: Add an item to the cart.
- `DELETE /cart/:id`: Remove an item from the cart.

### Orders

- `GET /orders`: Get all orders for the current user.
- `GET /orders/:id`: Get a single order by ID.
- `POST /orders`: Create a new order.

### Payments

- `POST /orders/:id/pay`: Create a Stripe checkout session for an order.

### Invoices

- `GET /invoices`: Get all invoices (Admin only).

### Sales

- `GET /sales/metrics`: Get sales metrics (Admin only).
- `GET /sales/stats`: Get order statistics (Admin only).

### Webhooks

- `POST /webhooks/stripe`: Stripe webhook for payment updates.

---

## 📘 License

This project is licensed under the MIT License.