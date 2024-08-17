CREATE TABLE profile_types (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
    CONSTRAINT unique_profile_type UNIQUE (name)
);

CREATE TABLE profiles (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    profile_type VARCHAR(36) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    CONSTRAINT fk_profile_type FOREIGN KEY (profile_type) REFERENCES profile_types(id),
    CONSTRAINT unique_profile UNIQUE (user_id, profile_type)
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
    CONSTRAINT fk_country FOREIGN KEY (country_id) REFERENCES countries(id),
    CONSTRAINT unique_province UNIQUE (country_id, code)
);

CREATE TABLE cities (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(2) NOT NULL,
    province_id VARCHAR(36) NOT NULL,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    description TEXT,
    CONSTRAINT fk_province FOREIGN KEY (province_id) REFERENCES provinces(id),
    CONSTRAINT unique_city UNIQUE (province_id, code)
);

CREATE TABLE sites (
    id VARCHAR(36) PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT
    image_url VARCHAR(255),
    CONSTRAINT unique_site UNIQUE (slug)
);

CREATE TABLE profile_sites (
    id VARCHAR(36) PRIMARY KEY,
    profile_id VARCHAR(36) NOT NULL,
    site_id VARCHAR(36) NOT NULL,
    CONSTRAINT fk_profile_site FOREIGN KEY (profile_id) REFERENCES profiles(id),
    CONSTRAINT fk_site_profile FOREIGN KEY (site_id) REFERENCES sites(id),
    CONSTRAINT unique_profile_site UNIQUE (profile_id, site_id)
);

CREATE TABLE business_types (
    id VARCHAR(36) PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    description TEXT
    image_url VARCHAR(255),
    CONSTRAINT unique_business_type UNIQUE (slug)
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
    CONSTRAINT fk_profile_business FOREIGN KEY (profile_id) REFERENCES profiles(id),
    CONSTRAINT fk_business_type FOREIGN KEY (business_type_id) REFERENCES business_types(id),
    CONSTRAINT fk_city_business FOREIGN KEY (city_id) REFERENCES cities(id),
    CONSTRAINT unique_business UNIQUE (profile_id, business_type_id, name)
);

CREATE TABLE business_images (
    id VARCHAR(36) PRIMARY KEY,
    business_id VARCHAR(36) NOT NULL,
    url VARCHAR(255) NOT NULL,
    CONSTRAINT fk_business_image FOREIGN KEY (business_id) REFERENCES businesses(id),
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
    CONSTRAINT fk_event_type FOREIGN KEY (event_type_id) REFERENCES event_types(id),
    CONSTRAINT fk_event_creator FOREIGN KEY (created_by) REFERENCES profiles(id),
    CONSTRAINT fk_event_city FOREIGN KEY (city_id) REFERENCES cities(id),
    CONSTRAINT unique_event UNIQUE (event_type_id, city_id, name)
);

CREATE TABLE event_schedules (
    id VARCHAR(36) PRIMARY KEY,
    event_id VARCHAR(36) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT true,
    start_at DATE NOT NULL,
    end_at DATE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_event_schedule FOREIGN KEY (event_id) REFERENCES events(id),
    CONSTRAINT unique_event_schedule UNIQUE (event_id, start_at, end_at)
);