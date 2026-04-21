package models

type MasterType struct {
    ID         int64   
    Name       string  
    CategoryID int64   
    MinTemp    float64 
    MaxTemp    float64 
    Warmth     int64   
    Layer      *int64  
}
