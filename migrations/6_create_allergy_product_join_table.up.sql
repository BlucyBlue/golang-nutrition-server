CREATE TABLE AllergyProducts (
                                 AllergyID INT NOT NULL,
                                 ProductID INT NOT NULL,
                                 PRIMARY KEY (AllergyID, ProductID),
                                 FOREIGN KEY (AllergyID) REFERENCES Allergies(AllergyID),
                                 FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);
