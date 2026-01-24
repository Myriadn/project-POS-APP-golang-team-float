```mermaid
erDiagram
ROLES {
  int id PK
  string name
  datetime created_at
  datetime updated_at
}

USERS {
  int id PK
  int role_id FK
  string name
  string email
  string phone
  string password
  string image
  string status
  datetime email_verified_at
  datetime created_at
  datetime updated_at
  datetime deleted_at
}

USER_SESSIONS {
  int id PK
  int user_id FK
  string token
  string ip_address
  string user_agent
  datetime last_activity
  datetime expires_at
  datetime created_at
}

OTP_CODES {
  int id PK
  int user_id FK
  string code
  string type
  datetime expires_at
  boolean is_used
  datetime created_at
}

PASSWORD_RESET_TOKENS {
  int id PK
  int user_id FK
  string token
  datetime expires_at
  datetime created_at
}

CATEGORIES {
  int id PK
  string name
  string description
  datetime created_at
  datetime updated_at
  datetime deleted_at
}

PRODUCTS {
  int id PK
  int category_id FK
  string name
  string description
  string image
  float price
  string type
  boolean is_active
  datetime created_at
  datetime updated_at
  datetime deleted_at
}

TABLES {
  int id PK
  string table_number
  int capacity
  boolean is_available
  datetime created_at
  datetime updated_at
}

PAYMENT_METHODS {
  int id PK
  string name
  boolean is_active
  datetime created_at
  datetime updated_at
}

ORDERS {
  int id PK
  string order_number
  int table_id FK
  int user_id FK
  int payment_method_id FK
  string customer_name
  float subtotal
  float tax
  float total_price
  string status
  datetime completed_at
  datetime created_at
  datetime updated_at
  datetime deleted_at
}

ORDER_ITEMS {
  int id PK
  int order_id FK
  int product_id FK
  int quantity
  float price
  float subtotal
  datetime created_at
  datetime updated_at
}

INVENTORIES {
  int id PK
  int product_id FK
  int stock
  int minimum_stock
  string unit
  datetime created_at
  datetime updated_at
}

NOTIFICATIONS {
  int id PK
  int user_id FK
  string type
  string message
  string status
  datetime read_at
  datetime created_at
  datetime updated_at
  datetime deleted_at
}

RESERVATIONS {
  int id PK
  int table_id FK
  string customer_name
  string customer_phone
  int number_of_guests
  datetime reservation_time
  string status
  text notes
  datetime created_at
  datetime updated_at
}

ROLES ||--o{ USERS : has
USERS ||--o{ USER_SESSIONS : has
USERS ||--o{ OTP_CODES : generates
USERS ||--o{ PASSWORD_RESET_TOKENS : requests
USERS ||--o{ ORDERS : creates
USERS ||--o{ NOTIFICATIONS : receives
CATEGORIES ||--o{ PRODUCTS : contains
PRODUCTS ||--o{ ORDER_ITEMS : included_in
PRODUCTS ||--o{ INVENTORIES : stored_in
ORDERS ||--o{ ORDER_ITEMS : contains
TABLES ||--o{ ORDERS : used_for
TABLES ||--o{ RESERVATIONS : booked_for
PAYMENT_METHODS ||--o{ ORDERS : paid_with
TABLES ||--o{ RESERVATIONS : booked
PAYMENT_METHODS ||--o{ ORDERS : paid_with
