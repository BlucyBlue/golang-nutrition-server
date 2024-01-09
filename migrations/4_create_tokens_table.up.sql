CREATE TABLE Tokens (
                        TokenID SERIAL PRIMARY KEY,
                        UserID INT UNIQUE NOT NULL,
                        Token VARCHAR(255) UNIQUE NOT NULL,
                        CreatedAt TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        ExpiresAt TIMESTAMP WITH TIME ZONE NOT NULL,
                        FOREIGN KEY (UserID) REFERENCES Users(UserID)
);
