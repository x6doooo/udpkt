package udpkt

import (
    goCache "github.com/patrickmn/go-cache"
    "github.com/satori/go.uuid"
    "sync"
    "time"
)

type pktCache struct {
    Total int
    Count int
    MsgPartials [][]byte
}

type Decoder struct {
    mu *sync.Mutex
    cache *goCache.Cache
}

func NewDecoder() *Decoder {
    return &Decoder{
        mu: &sync.Mutex{},
        cache: goCache.New(MAX_CACHE_TIME, LOOP_DURATION),
    }
}

func (me *Decoder) SetConf(timeout, interval time.Duration) {
    me.cache = goCache.New(timeout, interval)
}

func (me *Decoder) Do(pkt []byte) (fullMsg []byte, done bool) {
    pktIndex := int(pkt[0])
    pktTotal := int(pkt[1])
    uuidSize := int(pkt[2])
    uuidBytes := pkt[3: 3 + uuidSize]
    msgPartial := pkt[3 + uuidSize:]
    uuid, _ := uuid.FromBytes(uuidBytes)

    uuidStr := uuid.String()

    me.mu.Lock()
    x, found := me.cache.Get(uuidStr)

    var pktC *pktCache
    if found {
        pktC = x.(*pktCache)
    } else {
        pktC = &pktCache{}
        pktC.Total = pktTotal
        pktC.Count = 0
        pktC.MsgPartials = make([][]byte, pktTotal)
        me.cache.Set(uuidStr, pktC, goCache.DefaultExpiration)
    }

    pktC.MsgPartials[pktIndex] = msgPartial
    pktC.Count += 1
    me.mu.Unlock()

    if (pktC.Count == pktC.Total) {
        for _, partial := range pktC.MsgPartials {
            fullMsg = append(fullMsg, partial...)
        }
        me.cache.Delete(uuidStr)
        done = true
    }

    return
}




