package udpkt

import (
    "testing"
    //"os"
)

//func Test_count(t *testing.T) {
//    var l int = MAX_LEN * 12 + 100
//    l = count(l)
//    if l != 13 {
//        t.Error(l)
//    }
//    l = MAX_LEN * 3
//    l = count(l)
//    if l != 3 {
//        t.Error(l)
//    }
//}
var longStr = []byte("10.232.5.164 - - [22/Aug/2016:17:36:18 +0800] fwf[-] abc.ksyun.com:3000 GET /api?Action=DescribeVpcs&Service=vpc&Version=2016-03-04&Region=cn-shanghai-3&_=1471858550988 HTTP/1.1 200  0.432 0.476 686 799 http://abc.ksyun.com:3000/ Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36 10.111.17.86 80 -10.232.5.164 - - [22/Aug/2016:17:36:18 +0800] fwf[-] abc.ksyun.com:3000 GET /api?Action=DescribeSubnets&Service=vpc&Version=2016-03-04&Region=cn-shanghai-3&Filter.1.Name=vpc-id&Filter.1.Value.1=2bdee368-f5b6-45ae-8fdd-5d918fdb1633&_=1471858550989 HTTP/1.1 200  0.240 0.271 764 859 http://abc.ksyun.com:3000/ Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36 10.111.17.86 80 -10.232.5.164 - - [22/Aug/2016:17:36:20 +0800] fwf[-] abc.ksyun.com:3000 GET /api?Action=DescribeVpcs&Service=vpc&Version=2016-03-04&Region=cn-shanghai-3&_=1471858553086 HTTP/1.1 200  0.459 0.491 686 799 http://abc.ksyun.com:3000/ Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36 10.111.17.86 80 -10.232.5.164 - - [22/Aug/2016:17:36:20 +0800] fwf[-] abc.ksyun.com:3000 GET /api?Action=DescribeSubnets&Service=vpc&Version=2016-03-04&Region=cn-shanghai-3&Filter.1.Name=vpc-id&Filter.1.Value.1=2bdee368-f5b6-45ae-8fdd-5d918fdb1633&_=1471858553087 HTTP/1.1 200  0.246 0.280 764 859 http://abc.ksyun.com:3000/ Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36 10.111.17.86 80 -10.232.5.189 - - [22/Aug/2016:17:36:31 +0800] fwf[-] network.console.ksyun.com GET /api?Action=DescribeAddressesOfBws&Service=eip&Version=2016-03-04&Region=cn-shanghai-3&BwsId=0eaa92ce-1af8-4821-b121-61662688858b&_=1471858571762 HTTP/1.1 200  0.452 0.452 680 982 http://network.console.ksyun.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -10.232.5.189 - - [22/Aug/2016:17:36:37 +0800] fwf[-] network.console.ksyun.com GET /api?Action=DescribePriceSystem&Region=abc&Version=2016-03-04&Service=eip&_=1471858571764 HTTP/1.1 200  0.037 0.037 624 679 http://network.console.ksyun.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -10.232.5.189 - - [22/Aug/2016:17:36:38 +0800] fwf[-] network.console.ksyun.com GET /api?Action=AddressesList&Service=eip&Version=2016-03-04&Region=cn-shanghai-3&time=1471858594647&MaxResults=100&_=1471858571763 HTTP/1.1 200  0.901 0.901 662 10873 http://network.console.ksyun.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -10.232.5.164 - - [22/Aug/2016:17:36:44 +0800] fwf[-] abc.ksyun.com:3000 GET /api?Action=DescribeVpcs&Service=vpc&Version=2016-03-04&Region=cn-shanghai-3&_=1471858577188 HTTP/1.1 200  0.465 0.501 686 799 http://abc.ksyun.com:3000/ Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36 10.111.17.86 80 -10.232.5.164 - - [22/Aug/2016:17:36:44 +0800] fwf[-] abc.ksyun.com:3000 GET /api?Action=DescribeSubnets&Service=vpc&Version=2016-03-04&Region=cn-shanghai-3&Filter.1.Name=vpc-id&Filter.1.Value.1=2bdee368-f5b6-45ae-8fdd-5d918fdb1633&_=1471858577189 HTTP/1.1 200  0.247 0.278 764 859 http://abc.ksyun.com:3000/ Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.110 Safari/537.36 10.111.17.86 80 -10.232.5.189 - - [22/Aug/2016:17:36:49 +0800] fwf[-] network.console.ksyun.com GET /api?Action=DescribePriceSystem&Region=abc&Version=2016-03-04&Service=bws&_=1471858571765 HTTP/1.1 200  0.034 0.034 624 378 http://network.console.ksyun.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -10.232.5.189 - - [22/Aug/2016:17:36:51 +0800] fwf[-] network.console.ksyun.com GET /api?Action=DescribeBws&Service=bws&Version=2016-03-04&Region=cn-shanghai-3&_=1471858571766 HTTP/1.1 200  1.543 1.543 626 5927 http://network.console.ksyun.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -10.232.5.189 - - [22/Aug/2016:17:36:53 +0800] fwf[-] network.console.ksyun.com GET /api?Action=DescribeAddressesOfBws&Service=eip&Version=2016-03-04&Region=cn-shanghai-3&BwsId=0eaa92ce-1af8-4821-b121-61662688858b&_=1471858571767 HTTP/1.1 200  0.455 0.455 680 982 http://network.console.ksyun.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -10.232.5.189 - - [22/Aug/2016:17:36:58 +0800] fwf[-] network.console.ksyun.com GET /api?Action=DescribeAddressesOfBws&Service=eip&Version=2016-03-04&Region=cn-shanghai-3&BwsId=eee6fc52-586e-44bb-ba21-ac06ef08787a&_=1471858571768 HTTP/1.1 200  0.768 0.768 680 1487 http://network.console.ksyun.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -")

func TestEncode(t *testing.T) {

    encoder := NewEncoder()
    decoder := NewDecoder()

    strEncoded := encoder.Do(longStr)

    for _, msg := range strEncoded {
        fullMsg, done := decoder.Do(msg)
        if done {
            if string(fullMsg) != string(longStr) {
                t.Error("encode or decode failed!!!")
            }
        }
    }
    //file, _ := os.Getwd()

}

var str = []byte("[22/Aug/2016:17:36:58 +0800] fwf[-] network.console.ksyun.com GET /api?Action=DescribeAddressesOfBws&Service=eip&Version=2016-03-04&Region=cn-shanghai-3&BwsId=eee6fc52-586e-44bb-ba21-ac06ef08787a&_=1471858571768 HTTP/1.1 200  0.768 0.768 680 1487 http://network.console.ksyun.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -")


func TestEncodeAndDecode(t *testing.T) {
    total := 100
    count := 0
    c := make(chan bool)
    encoder := NewEncoder()
    decoder := NewDecoder()
    for i := 0; i < total; i++ {
        tem := encoder.Do(longStr)
        for _, p := range tem {
            go func(ch chan bool, p []byte) {
                fullMsg, done := decoder.Do(p)
                if done {
                    if string(fullMsg) != string(longStr) {
                        t.Error("failed!")
                        panic("failed!!!!!!")
                    }
                    ch <- true
                }
            }(c, p)
        }
    }

    for i := 0; i < total; i++ {
        select {
        case <- c:
            if count += 1; count == total {
                return
            }
        }
    }
}


var str2 = []byte{123,34,70,105,108,101,110,97,109,101,34,58,34,47,117,115,114,47,108,111,99,97,108,47,110,103,105,110,120,47,108,111,103,115,47,97,99,99,101,115,115,46,108,111,103,34,44,34,84,101,120,116,34,58,34,49,50,55,46,48,46,48,46,49,32,45,32,45,32,91,50,52,47,65,117,103,47,50,48,49,54,58,49,54,58,50,50,58,49,57,32,43,48,56,48,48,93,32,92,34,71,69,84,32,47,63,48,49,50,51,52,53,54,55,56,57,48,45,49,49,50,51,52,53,54,55,56,57,48,45,50,49,50,51,52,53,54,55,56,57,48,45,51,49,50,51,52,53,54,55,56,57,48,45,52,49,50,51,52,53,54,55,56,57,48,45,53,49,50,51,52,53,54,55,56,57,48,45,54,49,50,51,52,53,54,55,56,57,48,45,55,49,50,51,52,53,54,55,56,57,48,45,56,49,50,51,52,53,54,55,56,57,48,45,57,49,50,51,52,53,54,55,56,57,48,45,49,48,49,50,51,52,53,54,55,56,57,48,45,49,49,49,50,51,52,53,54,55,56,57,48,45,49,50,49,50,51,52,53,54,55,56,57,48,45,49,51,49,50,51,52,53,54,55,56,57,48,45,49,52,49,50,51,52,53,54,55,56,57,48,45,49,53,49,50,51,52,53,54,55,56,57,48,45,49,54,49,50,51,52,53,54,55,56,57,48,45,49,55,49,50,51,52,53,54,55,56,57,48,45,49,56,49,50,51,52,53,54,55,56,57,48,45,49,57,49,50,51,52,53,54,55,56,57,48,45,50,48,49,50,51,52,53,54,55,56,57,48,45,50,49,49,50,51,52,53,54,55,56,57,48,45,50,50,49,50,51,52,53,54,55,56,57,48,45,50,51,49,50,51,52,53,54,55,56,57,48,45,50,52,49,50,51,52,53,54,55,56,57,48,45,50,53,49,50,51,52,53,54,55,56,57,48,45,50,54,49,50,51,52,53,54,55,56,57,48,45,50,55,49,50,51,52,53,54,55,56,57,48,45,50,56,49,50,51,52,53,54,55,56,57,48,45,50,57,49,50,51,52,53,54,55,56,57,48,45,51,48,49,50,51,52,53,54,55,56,57,48,45,51,49,49,50,51,52,53,54,55,56,57,48,45,51,50,49,50,51,52,53,54,55,56,57,48,45,51,51,49,50,51,52,53,54,55,56,57,48,45,51,52,49,50,51,52,53,54,55,56,57,48,45,51,53,49,50,51,52,53,54,55,56,57,48,45,51,54,49,50,51,52,53,54,55,56,57,48,45,51,55,49,50,51,52,53,54,55,56,57,48,45,51,56,49,50,51,52,53,54,55,56,57,48,45,51,57,49,50,51,52,53,54,55,56,57,48,45,52,48,49,50,51,52,53,54,55,56,57,48,45,52,49,49,50,51,52,53,54,55,56,57,48,45,52,50,49,50,51,52,53,54,55,56,57,48,45,52,51,49,50,51,52,53,54,55,56,57,48,45,52,52,49,50,51,52,53,54,55,56,57,48,45,52,53,49,50,51,52,53,54,55,56,57,48,45,52,54,49,50,51,52,53,54,55,56,57,48,45,52,55,49,50,51,52,53,54,55,56,57,48,45,52,56,49,50,51,52,53,54,55,56,57,48,45,52,57,49,50,51,52,53,54,55,56,57,48,45,53,48,49,50,51,52,53,54,55,56,57,48,45,53,49,49,50,51,52,53,54,55,56,57,48,45,53,50,49,50,51,52,53,54,55,56,57,48,45,53,51,49,50,51,52,53,54,55,56,57,48,45,53,52,49,50,51,52,53,54,55,56,57,48,45,53,53,49,50,51,52,53,54,55,56,57,48,45,53,54,49,50,51,52,53,54,55,56,57,48,45,53,55,49,50,51,52,53,54,55,56,57,48,45,53,56,49,50,51,52,53,54,55,56,57,48,45,53,57,49,50,51,52,53,54,55,56,57,48,45,54,48,49,50,51,52,53,54,55,56,57,48,45,54,49,49,50,51,52,53,54,55,56,57,48,45,54,50,49,50,51,52,53,54,55,56,57,48,45,54,51,49,50,51,52,53,54,55,56,57,48,45,54,52,49,50,51,52,53,54,55,56,57,48,45,32,72,84,84,80,47,49,46,49,92,34,32,50,48,48,32,49,54,52,55,32,92,34,45,92,34,32,92,34,99,117,114,108,47,55,46,52,51,46,48,92,34,32,92,34,45,92,34,34,44,34,84,105,109,101,34,58,34,49,52,55,50,48,50,54,57,51,57,34,125}
func TestEncodeAndDecode2(t *testing.T) {

    encoder := NewEncoder()
    decoder := NewDecoder()
    strEncoded := encoder.Do(str2)

    for _, msg := range strEncoded {
        fullMsg, done := decoder.Do(msg)
        if done {
            if string(fullMsg) != string(str2) {
                t.Error("encode or decode failed!!!")
            }
        }
    }
}

func BenchmarkEncode(b *testing.B) {
    encoder := NewEncoder()
    for i := 0; i < b.N; i++ {
        encoder.Do(str)
    }
}

func BenchmarkDecode(b *testing.B) {
    encoder := NewEncoder()
    decoder := NewDecoder()
    n := b.N
    for i := 0; i < n; i++ {
        msgs := encoder.Do(str)
        for _, m := range msgs {
            decoder.Do(m)
        }
    }
}