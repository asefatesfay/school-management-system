CREATE TABLE Lessons (  
    LessonID SERIAL PRIMARY KEY,  
    SubjectID INT,  
    TeacherID INT,  
    ClassID INT,  
    LessonDate DATE,  
    StartTime TIME,  
    EndTime TIME,  
    FOREIGN KEY (SubjectID) REFERENCES Subjects(SubjectID),  
    FOREIGN KEY (TeacherID) REFERENCES Teachers(TeacherID),  
    FOREIGN KEY (ClassID) REFERENCES Classes(ClassID)  
); 