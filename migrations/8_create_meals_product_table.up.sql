CREATE TABLE MealProducts (
                              MealID INT NOT NULL,
                              ProductID INT NOT NULL,
                              PRIMARY KEY (MealID, ProductID),
                              FOREIGN KEY (MealID) REFERENCES Meals(MealID),
                              FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);
