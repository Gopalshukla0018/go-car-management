
CREATE TABLE IF NOT EXISTS engines (
    id UUID PRIMARY KEY,
    displacement BIGINT NOT NULL,
    no_of_cylinders BIGINT NOT NULL,
    car_range BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS cars (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    year TEXT NOT NULL,
    brand TEXT NOT NULL,
    fuel_type TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    engine_id UUID REFERENCES engines(id) ON DELETE SET NULL
);