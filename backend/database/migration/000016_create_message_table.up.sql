CREATE TABLE Messages (  
    MessageID SERIAL PRIMARY KEY,  
    SenderID INT,  
    SenderType VARCHAR(10),  
    ReceiverID INT,  
    ReceiverType VARCHAR(10),  
    Subject VARCHAR(100),  
    Body TEXT,  
    SentDateTime TIMESTAMP,  
    IsRead BOOLEAN DEFAULT FALSE  
); 