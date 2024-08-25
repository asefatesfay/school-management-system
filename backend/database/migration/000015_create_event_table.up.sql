CREATE TABLE Events (  
    EventID SERIAL PRIMARY KEY,  
    EventName VARCHAR(100),  
    EventDate DATE,  
    StartTime TIME,  
    EndTime TIME,  
    Description TEXT  
);  