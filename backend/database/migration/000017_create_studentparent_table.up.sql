CREATE TABLE StudentParents (  
    StudentID INT,  
    ParentID INT,  
    PRIMARY KEY (StudentID, ParentID),  
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID),  
    FOREIGN KEY (ParentID) REFERENCES Parents(ParentID)  
);  