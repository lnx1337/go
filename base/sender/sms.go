package sender

import (
    "bytes"
    "fmt"
    "net/http"
    "net/url"
)

type Sms struct {
    Lada      string
    Telephone string 
    Message   string
}

func NewSms(_telephone string ,_message string) Sms {
    self := Sms{
        Telephone:_telephone,
        Message:_message,
    }
    return self
}

// Send sms via twilio API
func (self *Sms) Send() *http.Response {    
    apiUrl := "https://api.twilio.com"
    resource := "/2010-04-01/Accounts/AC18bea9586c89e878ae963c655fbca950/SMS/Messages.json"
    
    data := url.Values{}
    data.Add("From", "+12018775547")
    data.Add("To", self.Telephone)
    data.Add("Body",self.Message)

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := fmt.Sprintf("%v", u)

    client := &http.Client{}
    r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
    r.SetBasicAuth("AC18bea9586c89e878ae963c655fbca950","68f3a83dc6ba71a985b9b1a932acd4e8")
    r.PostForm = data
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    
    resp, _ := client.Do(r)  
    return resp
}


/*
curl -X POST https://api.twilio.com/2010-04-01/Accounts/AC18bea9586c89e878ae963c655fbca950/SMS/Messages.json \
    -u AC18bea9586c89e878ae963c655fbca950:68f3a83dc6ba71a985b9b1a932acd4e8 \
    --data-urlencode "From=+12018775547" \
    --data-urlencode "To=+525545130805" \
    --data-urlencode 'Body=lnx1337 say hello'
*/
    // r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
    // r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    // r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    // data.Set("From", "+12018775547")
