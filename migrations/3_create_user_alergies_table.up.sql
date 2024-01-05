CREATE TABLE UserAllergies (
                               UserID INT,
                               AllergyID INT,
                               PRIMARY KEY (UserID, AllergyID),
                               FOREIGN KEY (UserID) REFERENCES Users(UserID),
                               FOREIGN KEY (AllergyID) REFERENCES Allergies(AllergyID)
);
