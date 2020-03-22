package idgen

import (
	"fmt"
	"net/http"
)

var template = `
<?xml version="1.0"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:math="http://example.org/math">
   <soapenv:Header/>
   <soapenv:Body>
      <math:answer>
         <math:result>42</math:result>
      </math:answer>
   </soapenv:Body>
</soapenv:Envelope>
`

func Template() string {
	return template
}

func ListUsersHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, Template())
}
