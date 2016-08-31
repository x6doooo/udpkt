package udpkt

import (
    "github.com/satori/go.uuid"
    "errors"
)


type Encoder struct {
    pktMaxSize int
}

func NewEncoder() *Encoder {
    return &Encoder{
        pktMaxSize: PKT_MAX_SIZE,
    }
}

func (me *Encoder) SetPktMaxSize(size int) error {
    if size <= 0 {
        return errors.New("packet_max_size <= 0")
    }
    me.pktMaxSize = size;
    return nil
}

func (me *Encoder) count(len int) int {
    c := len / me.pktMaxSize
    if len % me.pktMaxSize != 0 {
        c += 1
    }
    return c
}

func (me *Encoder) Do(content []byte) [][]byte {
    length := len(content)
    pktNum := me.count(length)

    if pktNum > 255 {
        return [][]byte{}
    }

    id := uuid.NewV4()
    pkts := make([][]byte, pktNum)
    for i := 0; i < pktNum; i++ {
        // get range
        from := i * me.pktMaxSize
        to := (i + 1) * me.pktMaxSize
        if to > length {
            to = length
        }
        tem := content[from:to]

        // insert info
        idByte := id.Bytes()
        uuidSize := len(idByte)

        iByte := byte(i)
        pktNumByte := byte(pktNum)
        uuidSizeByte := byte(uuidSize)

        pkt := []byte{iByte, pktNumByte, uuidSizeByte}
        pkt = append(pkt, idByte...)
        pkt = append(pkt, tem...)
        pkts[i] = pkt
    }

    return pkts
}
