CREATE TABLE Attendance (  
    AttendanceID SERIAL PRIMARY KEY,  
    StudentID INT,  
    LessonID INT,  
    Status VARCHAR(10),  
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID),  
    FOREIGN KEY (LessonID) REFERENCES Lessons(LessonID)  
); 