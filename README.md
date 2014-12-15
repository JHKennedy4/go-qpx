go-qpx
======

A small library for querying Google's QPX flight data api.


This library only generates the query. It does not actually post it to the server.

Usage

`import "qpx"'`

`query := qpx.NewQuery(origin, destination, startDate, endDate)`

`http.Post("https://www.googleapis.com/qpxExpress/v1/trips/search?key=your_API_key_here", "application/json", query.ToJSON())`
