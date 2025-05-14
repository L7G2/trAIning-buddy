-- 1. Użytkownicy
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       role VARCHAR(20) NOT NULL CHECK (role IN ('trener', 'uczen'))
);

-- 2. Profil użytkownika
CREATE TABLE profiles (
                          user_id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
                          age INTEGER CHECK (age >= 12 AND age <= 100),
                          height INTEGER CHECK (height >= 50 AND height <= 300),
                          weight NUMERIC(5,2) CHECK (weight > 0),
                          gender VARCHAR(10) CHECK (gender IN ('male', 'female', 'other')),
                          goal VARCHAR(20) CHECK (goal IN ('reduction', 'recomposition', 'mass'))
);

-- 3. Plany treningowe
CREATE TABLE training_plans (
                                id SERIAL PRIMARY KEY,
                                client_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                trainer_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                name VARCHAR(255) NOT NULL,
                                description TEXT,
                                start_date DATE NOT NULL,
                                end_date DATE,
                                CHECK (end_date IS NULL OR start_date <= end_date)
);

-- 4. Treningi
CREATE TABLE workouts (
                          id SERIAL PRIMARY KEY,
                          plan_id INTEGER NOT NULL REFERENCES training_plans(id) ON DELETE CASCADE,
                          workout_date DATE NOT NULL,
                          notes TEXT
);

-- 5. Ćwiczenia
CREATE TABLE exercises (
                           id SERIAL PRIMARY KEY,
                           workout_id INTEGER NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
                           name VARCHAR(100) NOT NULL,
                           sets INTEGER CHECK (sets >= 1),
                           reps INTEGER CHECK (reps >= 1),
                           weight INTEGER CHECK (weight >= 0),
                           instructions TEXT
);

-- 6. Pliki multimedialne (filmy instruktażowe)
CREATE TABLE media_files (
                             id SERIAL PRIMARY KEY,
                             title VARCHAR(255) NOT NULL,
                             url TEXT NOT NULL,
                             uploaded_by INTEGER REFERENCES users(id) ON DELETE SET NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 7. Przypisanie filmów do ucznia (przez trenera)
CREATE TABLE user_media (
                            id SERIAL PRIMARY KEY,
                            user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                            media_id INTEGER NOT NULL REFERENCES media_files(id) ON DELETE CASCADE,
                            assigned_by INTEGER REFERENCES users(id), -- trener
                            assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 8. Plany dietetyczne
CREATE TABLE diet_plans (
                            id SERIAL PRIMARY KEY,
                            user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 9. Posiłki w diecie
CREATE TABLE meals (
                       id SERIAL PRIMARY KEY,
                       diet_plan_id INTEGER NOT NULL REFERENCES diet_plans(id) ON DELETE CASCADE,
                       meal_order INTEGER NOT NULL CHECK (meal_order BETWEEN 1 AND 5),
                       name VARCHAR(100) NOT NULL,
                       calories INTEGER CHECK (calories >= 0),
                       proteins INTEGER CHECK (proteins >= 0),
                       fats INTEGER CHECK (fats >= 0),
                       carbs INTEGER CHECK (carbs >= 0)
);

-- 10. Produkty spożywcze
CREATE TABLE food_products (
                               id SERIAL PRIMARY KEY,
                               name VARCHAR(255) NOT NULL,
                               calories NUMERIC(6,2) NOT NULL,
                               protein NUMERIC(6,2) NOT NULL,
                               fat NUMERIC(6,2) NOT NULL,
                               carbs NUMERIC(6,2) NOT NULL
);

-- 11. Produkty przypisane do posiłków
CREATE TABLE meal_products (
                               id SERIAL PRIMARY KEY,
                               meal_id INTEGER NOT NULL REFERENCES meals(id) ON DELETE CASCADE,
                               product_id INTEGER NOT NULL REFERENCES food_products(id) ON DELETE CASCADE,
                               amount_grams NUMERIC(6,2) NOT NULL CHECK (amount_grams > 0)
);

-- 12. Raporty postępów
CREATE TABLE progress_reports (
                                  id SERIAL PRIMARY KEY,
                                  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                                  date DATE NOT NULL DEFAULT CURRENT_DATE,
                                  weight NUMERIC(5,2) CHECK (weight > 0),
                                  notes TEXT
);

-- 13. Wiadomości (czat)
CREATE TABLE messages (
                          id SERIAL PRIMARY KEY,
                          sender_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          receiver_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          message TEXT NOT NULL,
                          sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
