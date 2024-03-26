package config


DB *sql.DB


func Connect()error{
  db, err := sql.Open()

  if err!=nil{
    return err
  }


}
