CREATE TABLE tamiat_users(
                             id SERIAL PRIMARY KEY,
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             deleted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                             name VARCHAR(20) NOT NULL,
                             email TEXT NOT NULL UNIQUE ,
                             password TEXT NOT NULL,
                             role_id INTEGER REFERENCES tamiat_roles(id)
)
