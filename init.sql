-- =====================================================
-- Point of Sales Application Database Schema
-- PostgreSQL Database Initialization Script
-- =====================================================

-- =====================================================
-- DROP EXISTING TABLES (for clean migration)
-- =====================================================
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS admin_permissions CASCADE;
DROP TABLE IF EXISTS notifications CASCADE;
DROP TABLE IF EXISTS reservations CASCADE;
DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS customers CASCADE;
DROP TABLE IF EXISTS payment_methods CASCADE;
DROP TABLE IF EXISTS tables CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS otp_codes CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS roles CASCADE;

-- =====================================================
-- 1. ROLES TABLE
-- Tabel terpisah untuk role user
-- =====================================================
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 2. USERS TABLE
-- Menyimpan data staff dan admin
-- =====================================================
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE RESTRICT,
    profile_picture VARCHAR(500),
    salary DECIMAL(15, 2),
    date_of_birth DATE,
    shift_start TIME,
    shift_end TIME,
    address TEXT,
    additional_details TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- =====================================================
-- 3. OTP_CODES TABLE
-- Menyimpan OTP untuk login dan reset password
-- =====================================================
CREATE TABLE otp_codes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    code VARCHAR(6) NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('login', 'reset_password')),
    is_used BOOLEAN DEFAULT false,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 4. SESSIONS TABLE
-- Session management untuk tracking user login
-- =====================================================
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    ip_address VARCHAR(45),
    user_agent TEXT,
    is_active BOOLEAN DEFAULT true,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    last_activity_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 5. CATEGORIES TABLE
-- Kategori menu (Pizza, Burger, Chicken, dll)
-- =====================================================
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    icon VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 6. PRODUCTS TABLE
-- Produk makanan/minuman
-- =====================================================
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image VARCHAR(500),
    price DECIMAL(15, 2) NOT NULL CHECK (price >= 0),
    availability VARCHAR(20) DEFAULT 'in_stock' CHECK (availability IN ('in_stock', 'out_of_stock')),
    menu_type VARCHAR(50) DEFAULT 'normal' CHECK (menu_type IN ('normal', 'special_deals', 'new_year_special', 'desserts_and_drinks')),
    stock INTEGER NOT NULL DEFAULT 0 CHECK (stock >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);


-- =====================================================
-- 8. TABLES TABLE
-- Meja restoran (per lantai)
-- =====================================================
CREATE TABLE tables (
    id SERIAL PRIMARY KEY,
    table_number VARCHAR(10) UNIQUE NOT NULL,
    floor INTEGER NOT NULL DEFAULT 1 CHECK (floor >= 1),
    capacity INTEGER NOT NULL DEFAULT 4 CHECK (capacity >= 1),
    status VARCHAR(20) DEFAULT 'available' CHECK (status IN ('available', 'occupied', 'reserved')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 9. PAYMENT_METHODS TABLE
-- Metode pembayaran
-- =====================================================
CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 10. CUSTOMERS TABLE
-- Data pelanggan untuk reservasi
-- =====================================================
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    customer_id VARCHAR(50) UNIQUE NOT NULL,
    title VARCHAR(10) CHECK (title IN ('Mr', 'Mrs', 'Ms', 'Miss', 'Dr')),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100),
    phone VARCHAR(20),
    email VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 11. ORDERS TABLE
-- Data pesanan
-- =====================================================
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_number VARCHAR(20) UNIQUE NOT NULL,
    table_id INTEGER REFERENCES tables(id) ON DELETE SET NULL,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    payment_method_id INTEGER REFERENCES payment_methods(id) ON DELETE SET NULL,
    customer_name VARCHAR(255) NOT NULL,
    status VARCHAR(30) DEFAULT 'ready' CHECK (status IN ('ready', 'in_process', 'completed', 'cancelled', 'cooking_now', 'in_the_kitchen', 'ready_to_serve')),
    subtotal DECIMAL(15, 2) NOT NULL DEFAULT 0 CHECK (subtotal >= 0),
    tax_rate DECIMAL(5, 2) DEFAULT 5.00,
    tax_amount DECIMAL(15, 2) DEFAULT 0 CHECK (tax_amount >= 0),
    total DECIMAL(15, 2) NOT NULL DEFAULT 0 CHECK (total >= 0),
    notes TEXT,
    order_date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 12. ORDER_ITEMS TABLE
-- Detail item dalam pesanan
-- =====================================================
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    quantity INTEGER NOT NULL DEFAULT 1 CHECK (quantity >= 1),
    unit_price DECIMAL(15, 2) NOT NULL CHECK (unit_price >= 0),
    total_price DECIMAL(15, 2) NOT NULL CHECK (total_price >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 13. RESERVATIONS TABLE
-- Data reservasi meja
-- =====================================================
CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    table_id INTEGER NOT NULL REFERENCES tables(id) ON DELETE RESTRICT,
    customer_id INTEGER NOT NULL REFERENCES customers(id) ON DELETE RESTRICT,
    reservation_date DATE NOT NULL,
    reservation_time TIME NOT NULL,
    pax_number INTEGER NOT NULL DEFAULT 1 CHECK (pax_number >= 1),
    deposit_fee DECIMAL(15, 2) DEFAULT 0 CHECK (deposit_fee >= 0),
    status VARCHAR(20) DEFAULT 'confirmed' CHECK (status IN ('confirmed', 'awaited', 'cancelled')),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- 14. NOTIFICATIONS TABLE
-- Notifikasi untuk user
-- =====================================================
CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    message TEXT,
    type VARCHAR(50) DEFAULT 'info' CHECK (type IN ('info', 'warning', 'alert', 'success')),
    status VARCHAR(20) DEFAULT 'new' CHECK (status IN ('new', 'read')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    read_at TIMESTAMP WITH TIME ZONE
);

-- =====================================================
-- 15. ADMIN_PERMISSIONS TABLE
-- Hak akses admin (Manage Access feature)
-- =====================================================
CREATE TABLE admin_permissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    dashboard BOOLEAN DEFAULT false,
    reports BOOLEAN DEFAULT false,
    inventory BOOLEAN DEFAULT false,
    orders BOOLEAN DEFAULT false,
    customers BOOLEAN DEFAULT false,
    settings BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- =====================================================
-- INDEXES
-- Performance optimization
-- =====================================================

-- Roles indexes
CREATE INDEX idx_roles_name ON roles(name);

-- Users indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_role_id ON users(role_id);
CREATE INDEX idx_users_is_active ON users(is_active);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);

-- OTP codes indexes
CREATE INDEX idx_otp_codes_user_id ON otp_codes(user_id);
CREATE INDEX idx_otp_codes_type ON otp_codes(type);
CREATE INDEX idx_otp_codes_expires_at ON otp_codes(expires_at);
CREATE INDEX idx_otp_codes_is_used ON otp_codes(is_used);

-- Sessions indexes
CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE INDEX idx_sessions_token ON sessions(token);
CREATE INDEX idx_sessions_is_active ON sessions(is_active);
CREATE INDEX idx_sessions_expires_at ON sessions(expires_at);

-- Categories indexes
CREATE INDEX idx_categories_name ON categories(name);

-- Products indexes
CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_availability ON products(availability);
CREATE INDEX idx_products_menu_type ON products(menu_type);
CREATE INDEX idx_products_deleted_at ON products(deleted_at);
CREATE INDEX idx_products_created_at ON products(created_at);


-- Tables indexes
CREATE INDEX idx_tables_status ON tables(status);
CREATE INDEX idx_tables_floor ON tables(floor);

-- Orders indexes
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_table_id ON orders(table_id);
CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_order_date ON orders(order_date);
CREATE INDEX idx_orders_created_at ON orders(created_at);

-- Order items indexes
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_product_id ON order_items(product_id);

-- Reservations indexes
CREATE INDEX idx_reservations_table_id ON reservations(table_id);
CREATE INDEX idx_reservations_customer_id ON reservations(customer_id);
CREATE INDEX idx_reservations_date ON reservations(reservation_date);
CREATE INDEX idx_reservations_status ON reservations(status);

-- Notifications indexes
CREATE INDEX idx_notifications_user_id ON notifications(user_id);
CREATE INDEX idx_notifications_status ON notifications(status);
CREATE INDEX idx_notifications_created_at ON notifications(created_at);

-- =====================================================
-- FUNCTIONS
-- =====================================================

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- =====================================================
-- TRIGGERS
-- Auto-update updated_at on specific tables
-- =====================================================

CREATE TRIGGER update_roles_updated_at
    BEFORE UPDATE ON roles
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_categories_updated_at
    BEFORE UPDATE ON categories
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_products_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();


CREATE TRIGGER update_tables_updated_at
    BEFORE UPDATE ON tables
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_customers_updated_at
    BEFORE UPDATE ON customers
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_orders_updated_at
    BEFORE UPDATE ON orders
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_reservations_updated_at
    BEFORE UPDATE ON reservations
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_admin_permissions_updated_at
    BEFORE UPDATE ON admin_permissions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- =====================================================
-- SEEDER DATA
-- =====================================================

-- -- Seeder: Roles
-- INSERT INTO roles (name, description) VALUES
--     ('superadmin', 'Super Administrator dengan akses penuh ke semua fitur'),
--     ('admin', 'Administrator dengan akses terbatas sesuai permission'),
--     ('staff', 'Staff biasa dengan akses operasional dasar');

-- -- Seeder: Categories
-- INSERT INTO categories (name, description, icon) VALUES
--     ('All', 'Semua kategori menu', NULL),
--     ('Pizza', 'Berbagai macam pizza dengan topping pilihan', '/icons/pizza.png'),
--     ('Burger', 'Burger dengan daging sapi, ayam, atau ikan', '/icons/burger.png'),
--     ('Chicken', 'Ayam goreng, panggang, dan olahan ayam lainnya', '/icons/chicken.png'),
--     ('Bakery', 'Roti, kue, dan pastry segar', '/icons/bakery.png'),
--     ('Beverage', 'Minuman dingin dan panas', '/icons/beverage.png'),
--     ('Seafood', 'Hidangan laut segar', '/icons/seafood.png');

-- -- Seeder: Payment Methods
-- INSERT INTO payment_methods (name, is_active) VALUES
--     ('Cash', true),
--     ('Visa Card', true),
--     ('Master Card', true),
--     ('Debit Card', true);

-- -- Seeder: Tables (7 tables per floor, 3 floors)
-- INSERT INTO tables (table_number, floor, capacity, status) VALUES
--     -- 1st Floor
--     ('01', 1, 4, 'available'),
--     ('02', 1, 4, 'available'),
--     ('03', 1, 4, 'available'),
--     ('04', 1, 4, 'available'),
--     ('05', 1, 6, 'available'),
--     ('06', 1, 6, 'available'),
--     ('07', 1, 8, 'available'),
--     -- 2nd Floor
--     ('08', 2, 4, 'available'),
--     ('09', 2, 4, 'available'),
--     ('10', 2, 4, 'available'),
--     ('11', 2, 4, 'available'),
--     ('12', 2, 6, 'available'),
--     ('13', 2, 6, 'available'),
--     ('14', 2, 8, 'available'),
--     -- 3rd Floor
--     ('15', 3, 4, 'available'),
--     ('16', 3, 4, 'available'),
--     ('17', 3, 4, 'available'),
--     ('18', 3, 4, 'available'),
--     ('19', 3, 6, 'available'),
--     ('20', 3, 6, 'available'),
--     ('21', 3, 10, 'available');

-- -- Seeder: Superadmin User
-- -- Password: Admin@123 (hashed with bcrypt, you should generate this in your application)
-- INSERT INTO users (
--     email,
--     username,
--     password_hash,
--     full_name,
--     phone,
--     role_id,
--     salary,
--     shift_start,
--     shift_end,
--     is_active
-- ) VALUES (
--     'superadmin@posapp.com',
--     'superadmin',
--     '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', -- password: password (change this!)
--     'Super Administrator',
--     '+1 (123) 456-7890',
--     1, -- role_id = 1 (superadmin)
--     5000.00,
--     '09:00:00',
--     '18:00:00',
--     true
-- );

-- Create admin permissions for superadmin (all access)
-- INSERT INTO admin_permissions (
--     user_id,
--     dashboard,
--     reports,
--     inventory,
--     orders,
--     customers,
--     settings
-- ) VALUES (
--     1, -- user_id = 1 (superadmin)
--     true,
--     true,
--     true,
--     true,
--     true,
--     true
-- );



-- =====================================================
-- COMMENTS
-- Documentation for tables
-- =====================================================

-- COMMENT ON TABLE roles IS 'User roles (superadmin, admin, staff)';
-- COMMENT ON TABLE users IS 'Stores staff and admin user data';
-- COMMENT ON TABLE otp_codes IS 'Stores OTP codes for login and password reset';
-- COMMENT ON TABLE sessions IS 'Manages user login sessions';
-- COMMENT ON TABLE categories IS 'Product categories (Pizza, Burger, etc)';
-- COMMENT ON TABLE products IS 'Food and beverage products';
-- COMMENT ON TABLE inventory IS 'Product stock and inventory tracking';
-- COMMENT ON TABLE tables IS 'Restaurant tables information';
-- COMMENT ON TABLE payment_methods IS 'Available payment methods';
-- COMMENT ON TABLE customers IS 'Customer data for reservations';
-- COMMENT ON TABLE orders IS 'Order transactions';
-- COMMENT ON TABLE order_items IS 'Order line items/details';
-- COMMENT ON TABLE reservations IS 'Table reservation data';
-- COMMENT ON TABLE notifications IS 'User notifications';
-- COMMENT ON TABLE admin_permissions IS 'Admin access permissions for different modules';

-- =====================================================
-- END OF SCHEMA
-- =====================================================
