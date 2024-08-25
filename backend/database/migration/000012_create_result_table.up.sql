CREATE TABLE Results (  
    ResultID SERIAL PRIMARY KEY,  
    StudentID INT,  
    ExamID INT,  
    Score DECIMAL(5,2),  
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID),  
    FOREIGN KEY (ExamID) REFERENCES Exams(ExamID)  
);