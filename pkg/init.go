package pkg

import (
	"io/ioutil"
	"net/http"
)

func Init(url string) []byte {

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic am9yZGFuLmthc3BlckBmb3JnZXJvY2suY29tL3Rva2VuOmtja0JTREx6YWs2V2NSWEZmQkt6eldCZjBNZ1pnWHJEQWFCbk1nRGc=")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	return body
}
