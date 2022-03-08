CREATE TABLE tamiat_role_permission(
    id SERIAL PRIMARY KEY,
    role_id INTEGER REFERENCES tamiat_roles(id),
    permission_id INTEGER REFERENCES tamiat_permissions(id)
);