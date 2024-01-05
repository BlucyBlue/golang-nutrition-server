CREATE TABLE Users (
                       UserID INT PRIMARY KEY AUTO_INCREMENT,
                       Username VARCHAR(50) NOT NULL,
                       Email VARCHAR(100) NOT NULL,
                       PasswordHash VARCHAR(255) NOT NULL
);
