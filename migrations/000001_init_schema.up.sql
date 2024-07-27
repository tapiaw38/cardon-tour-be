CREATE TABLE profile_types (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE profiles (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    profile_type VARCHAR(36) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    FOREIGN KEY (profile_type) REFERENCES profile_types(id)
);

CREATE TABLE countries (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(2) NOT NULL,
    description TEXT
);

CREATE TABLE provinces (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(2) NOT NULL,
    country_id VARCHAR(36) NOT NULL,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    description TEXT,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);

CREATE TABLE cities (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(2) NOT NULL,
    province_id VARCHAR(36) NOT NULL,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    description TEXT,
    FOREIGN KEY (province_id) REFERENCES provinces(id)
);

CREATE TABLE profile_locations (
    profile_id VARCHAR(36) NOT NULL,
    country_id VARCHAR(36) NOT NULL,
    province_id VARCHAR(36) NOT NULL,
    city_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profiles(id),
    FOREIGN KEY (country_id) REFERENCES countries(id),
    FOREIGN KEY (province_id) REFERENCES provinces(id),
    FOREIGN KEY (city_id) REFERENCES cities(id)
);

CREATE TABLE business_types (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE businesses (
    id VARCHAR(36) PRIMARY KEY,
    profile_id VARCHAR(36) NOT NULL,
    business_type_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    description TEXT,
    active BOOLEAN NOT NULL DEFAULT true,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    city_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (profile_id) REFERENCES profiles(id),
    FOREIGN KEY (business_type_id) REFERENCES business_types(id),
    FOREIGN KEY (city_id) REFERENCES cities(id)
);

CREATE TABLE business_images (
    id VARCHAR(36) PRIMARY KEY,
    business_id VARCHAR(36) NOT NULL,
    url VARCHAR(255) NOT NULL,
    FOREIGN KEY (business_id) REFERENCES businesses(id),
    CONSTRAINT unique_business_image UNIQUE (business_id, url)
);

CREATE TABLE event_types (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE events (
    id VARCHAR(36) PRIMARY KEY,
    event_type_id VARCHAR(36) NOT NULL,
    city_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    created_by VARCHAR(36) NOT NULL,
    FOREIGN KEY (event_type_id) REFERENCES event_types(id),
    FOREIGN KEY (created_by) REFERENCES profiles(id),
    FOREIGN KEY (city_id) REFERENCES cities(id)
);

CREATE TABLE event_schedules (
    id VARCHAR(36) PRIMARY KEY,
    event_id VARCHAR(36) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    start_at DATE NOT NULL,
    end_at DATE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events(id),
    CONSTRAINT unique_event_schedule UNIQUE (event_id, start_at, end_at)
)