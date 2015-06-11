package main

import (
    "bytes"
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/json"
)

type data struct{
    file_name string
    md5checksum string
    url string
}
type Tracks struct {
    Toptracks data
}

func main() {
    apiUrl := "http://example.com"
    resource := "/path"
    data := url.Values{}
    data.Set("post_parameter1", "value 1")
    data.Set("post_parameter2", "value2")

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"

    client := &http.Client{}
    r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode())) // <-- URL-encoded payload
    r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    resp, _ := client.Do(r)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    var content map[string]interface{}
    json.Unmarshal(body, &content)
    fmt.Println(content)
}
