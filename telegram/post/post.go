package post

import (
	"time"
)

func GetTime() string {
	//t := time.Now()
	var t = time.Now().Format("2006-01-02T15:04:05.05Z")
	//fmt.Println(t.Format("2006-01-02T15:04:05.05Z"))
	return t
}

//const postURL = "http://localhost:3001/order"
//const postCustomerURL = "http://localhost:3001/customer"
//
//func SendPost(data *telegram.Order) error {
//	params := fmt.Sprintf(
//		"{\"Name\": \"%s\", \"Telephone\": \"%s\", \"Product\": \"%s\", \"Amount\": \"%v\"}",
//		data.Name,
//		data.Telephone,
//		data.Product,
//		data.Amount)
//
//	buf := bytes.NewBufferString(params)
//	resp, err := http.Post(
//		postURL,
//		"application/json",
//		buf)
//	if err != nil {
//		return err
//	}
//	//otlojenniy vizov - prinuditelno zakrit pri lubom returne
//	defer resp.Body.Close()
//
//	if resp.StatusCode != 200 {
//		return errors.New("not 200 response")
//	}
//	return nil
//}
//
//func SendPostCustomerID(data *telegram.Order) error {
//	//resp, err := client.Do(req)
//	//sprintf("{ \"Name\": \"%s\"}")
//	params := fmt.Sprintf(
//		"{\"Phone\": \"%s\"}",
//		data.Telephone)
//
//	buf := bytes.NewBufferString(params)
//
//	resp, err := http.Post(
//		postCustomerURL,
//		"application/json",
//		buf)
//	if err != nil {
//		return err
//	}
//	//otlojenniy vizov - prinuditelno zakrit pri lubom returne
//	fmt.Printf("resp: %v\n", resp)
//
//	defer resp.Body.Close()
//
//	if resp.StatusCode != 200 {
//		return errors.New("not 200 response")
//	}
//	return nil
//}
