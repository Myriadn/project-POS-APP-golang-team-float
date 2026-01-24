```mermaid
erDiagram

ROLES {
  int id PK
  string name
}

USERS {
  int id PK
  int role_id FK
  string name
  string email
  string password
  string status
  datetime created_at
}

CATEGORIES {
  int id PK
  string name
  datetime created_at
}

PRODUCTS {
  int id PK
  int category_id FK
  string name
  float price
  string type
  boolean is_active
  datetime created_at
}

TABLES {
  int id PK
  string table_number
  int capacity
  boolean is_available
}

PAYMENT_METHODS {
  int id PK
  string name
}

ORDERS {
  int id PK
  int table_id FK
  int user_id FK
  int payment_method_id FK
  string customer_name
  float tax
  float total_price
  string status
  datetime created_at
}

ORDER_ITEMS {
  int id PK
  int order_id FK
  int product_id FK
  int quantity
  float price
}

INVENTORIES {
  int id PK
  int product_id FK
  int stock
  string unit
  datetime updated_at
}

NOTIFICATIONS {
  int id PK
  int user_id FK
  string message
  string status
  datetime created_at
}

RESERVATIONS {
  int id PK
  int table_id FK
  string customer_name
  datetime reservation_time
  boolean is_cancelled
}

ROLES ||--o{ USERS : has
CATEGORIES ||--o{ PRODUCTS : contains
PRODUCTS ||--o{ ORDER_ITEMS : included
PRODUCTS ||--o{ INVENTORIES : stored_in
USERS ||--o{ ORDERS : creates
USERS ||--o{ NOTIFICATIONS : receives
ORDERS ||--o{ ORDER_ITEMS : has
TABLES ||--o{ ORDERS : used
TABLES ||--o{ RESERVATIONS : booked
PAYMENT_METHODS ||--o{ ORDERS : paid_with
