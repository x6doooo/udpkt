/*
Package udpkt providers udp packets encode/decode
 */
package udpkt

/**

message protocol:

[]byte{
    []byte{ 第几个包, 总包数, uuid长度, uuid, 正文 },
    ......
}

*/

import (
    "time"
)


const (
    PKT_MAX_SIZE = 800
    MAX_CACHE_TIME = 2 * time.Minute
    LOOP_DURATION = 30 * time.Second
)




