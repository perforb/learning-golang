package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

// &{Status:200 OK StatusCode:200 Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Accept-Ranges:[bytes] Age:[504] Connection:[keep-alive] Content-Length:[6308] Content-Type:[text/html] Date:[Thu, 29 Sep 2022 11:07:39 GMT] Etag:["67e4bf16110c35a2ae75d1f776f61e55"] Last-Modified:[Fri, 23 Sep 2022 13:13:54 GMT] Server:[AmazonS3] Via:[1.1 50d80cbc4f2c3fd4b5c67fa188a4e928.cloudfront.net (CloudFront)] X-Amz-Cf-Id:[gguYrHv6NPfsDgGvDaN7j41bjiJ2lX3CM2wsEEeceSXWEBSq2d_xXg==] X-Amz-Cf-Pop:[NRT57-P4] X-Cache:[Hit from cloudfront]] Body:0x1400015c400 ContentLength:6308 TransferEncoding:[] Close:false Uncompressed:false Trailer:map[] Request:0x1400017a000 TLS:0x140000ee000}
// Response status: 200 OK
// <!DOCTYPE html>
// <html>
//
//	<head>
//	  <meta charset="utf-8">
//	  <title>Go by Example</title>
func main() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		e := Body.Close()
		if e != nil {
			panic(e)
		}
	}(resp.Body)

	fmt.Printf("%+v\n", resp)
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if e := scanner.Err(); e != nil {
		panic(e)
	}
}
