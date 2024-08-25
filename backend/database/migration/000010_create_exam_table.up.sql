CREATE TABLE Exams (  
    ExamID SERIAL PRIMARY KEY,  
    SubjectID INT,  
    ClassID INT,  
    ExamDate DATE,  
    StartTime TIME,  
    EndTime TIME,  
    FOREIGN KEY (SubjectID) REFERENCES Subjects(SubjectID),  
    FOREIGN KEY (ClassID) REFERENCES Classes(ClassID)  
);  