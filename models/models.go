package models

//import "time"
import (
    //"fmt"
    // "time"
    "github.com/go-sql-driver/mysql"
    // "reflect"
    "fmt"
    "time"
    "encoding/json"

)
// User schema of the user table
type User struct {
    ID       int64  `json:"id"`
    Name     string `json:"name"`
    Email     string `json:"email"`
    DOB     	time.Time `json:"dob"`
    // DOB     	time.Time `default:"time.Now()" json:"dob"`
    // DOB     	time.Time `default:"0001-01-01" json:"dob"`
    // DOB     	mysql.NullTime `default:"0001-01-01" json:"dob"`
    // DOB     	NullTime `json:"dob"`
    PhoneNo      string `json:"phone_no"`
}

// NullTime is an alias for mysql.NullTime data type
// type NullTime mysql.NullTime

// // Scan implements the Scanner interface for NullTime
// func (nt *NullTime) Scan(value interface{}) error {
// 	var t mysql.NullTime
// 	if err := t.Scan(value); err != nil {
// 		return err
// 	}

// 	// if nil then make Valid false
// 	if reflect.TypeOf(value) == nil {
// 		*nt = NullTime{t.Time, false}
// 	} else {
// 		*nt = NullTime{t.Time, true}
// 	}

// 	return nil
// }

type NullTime struct {
	mysql.NullTime
}

// MarshalJSON for NullTime
// func (nt *NullTime) UnmarshalJSON() ([]byte, error) {

// 	fmt.Println("inside unmasrsha custom")
// 	if !nt.Valid {
// 		return []byte("null"), nil
// 	}
// 	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
// 	return []byte(val), nil
// }

// func (st *NullTime) UnmarshalJSON(b []byte) error {
//     //convert the bytes into an interface
//     //this will help us check the type of our value
//     //if it is a string that can be converted into a int we convert it
//     ///otherwise we return an error
//     var item interface{}
//     if err := json.Unmarshal(b, &item); err != nil {
//         return err
//     }
//     switch v := item.(type) {
//     // case int:
//     //     *st = StringInt(v)
//     // case float64:
//     //     *st = StringInt(int(v))
//     case string:
//         ///here convert the string into
//         ///an integer
//         // i, err := strconv.Atoi(v)
//         // if err != nil {
//         //     ///the string might not be of integer type
//         //     ///so return an error
//         //     return err

//         // }
//         t, err := time.Parse(time.RFC3339, str)

// 		if err != nil {
// 		    fmt.Println(err)
// 		}
//         // *st = NullTime(i)
//         // *st = t
//         *st = NullTime(t)

//     }
//     return nil
// }


// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
// nt = nt.Time.Format(time.RFC3339)
 err := json.Unmarshal(b, &nt.Time)
 nt.Valid = (err == nil)
 return err
}
