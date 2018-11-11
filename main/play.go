package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://www.amemv.com/aweme/v1/aweme/post/?user_id=62318336592&count=21&max_cursor=0&aid=1128&_signature=Qa59FBASGuH-2IzyHQAZQEGufQ&dytk=cf68e1c738dca97bfa1905acac6ddf20"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "80d11cb3-c44c-4913-be7b-e09a4d5eeacc")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}