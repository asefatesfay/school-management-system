CREATE TABLE Students (  
    StudentID SERIAL PRIMARY KEY,  
    FirstName VARCHAR(50),  
    LastName VARCHAR(50),  
    BirthDate DATE,  
    Email VARCHAR(100),  
    PhoneNumber VARCHAR(15),  
    EnrollmentDate DATE,  
    ClassID INT  
);  