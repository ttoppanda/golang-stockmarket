package main 

import (
    "encoding/json"
    "net/http"
)

type Trade struct {
    Actor string `json:"actor"`
    Shares int `json:"shares"`
    Ticker string `json:"ticker"`
    Price float64 `json:"price"`
    Intent string `json:"intent"`
    Kind string `json:"kind"`
    State  string `json:"state"`
    Time float64 `json:"time"`
}

func main() {

    hash := make(map[string][]Trade)

    // only store minutes (we dont need 1 sec or 10 sec charts)

    // {
    //       timestamp_hour: ISODate("2013-10-10T23:00:00.000Z"),
    //       type: “memory_used”,
    //       values: {
    //         0: {vol: 1000, open: 10.05, close: 10.55, high: 11.00, low: 10.00 },
    //         1: {vol: 1000, open: 10.55, close: 10.60, high: 11.00, low: 10.50 },
    //         …,
    //         58: {vol: 1000, open: 10.65, close: 10.80, high: 11.00, low: 10.60 },
    //         59: {vol: 1000, open: 10.65, close: 11.55, high: 12.00, low: 10.00 } 
    //       }
    //     }
    // }
    // {$set: {“values.59”: {vol: 1000, open: 10.65, close: 11.55, high: 12.00, low: 10.00 } }

    // basic charting:
    // aggregate minute based data to correct format time ranges

    // what will an aggregation query look like? "I want the last <X>(5?) <range>(days?) <duration>(hour?) chart"
    // $match stock <X> * <range>
    // $unwind values
    // $groupby <duration> (want array of duration groups)
    // $sum vol, $min low, $max hi, open, close of first?
    // $map appropriate time stamps?
    
    // real time chart:
    // front end will merge published minute info into the chart's range

    // 1 second ticker data will be separate channel, not persisted in a collection

    // on price data
    //      - publish to ticker channel, rate limit to one / second
    //      - add to cache of last 60 seconds of trades 

    // every 60 seconds,
    //      for each stock
    //          - calculate high, open, low, close, vol for last 60 seconds (cache last state to solve missing values?)
    //          - update DB
    //          - publish latest minute info
    //          - clear cache

    // http://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        var payload [2]Trade
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&payload)
        if err != nil {
            panic(err)
        }

        ticker := payload[0].Ticker
        hash[ticker] = append(hash[ticker], payload[0])

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Status 200"))
    })   
}