CREATE TABLE Assignments (  
    AssignmentID SERIAL PRIMARY KEY,  
    SubjectID INT,  
    ClassID INT,  
    Title VARCHAR(100),  
    Description TEXT,  
    DueDate DATE,  
    FOREIGN KEY (SubjectID) REFERENCES Subjects(SubjectID),  
    FOREIGN KEY (ClassID) REFERENCES Classes(ClassID)  
); 