package qpx

import "encoding/json"
import "fmt"

func ParseResponse(res []byte) Response {
	var r Response
	err := json.Unmarshal(res, &r)
	if err != nil {
		panic(err)
	}
	return r
}

func (r Response) PrettyPrint() string {
	output := "Flight Plans:\n"
	for i, e := range r.Trips.TripOption {
		output += fmt.Sprintf("%d. Price: %s \n", i+1, e.SaleTotal)

		for j := 0; j < len(e.Slice); j++ {
			output += "-------------------------------------------------\n"
			if j == 0 {
				output += "Outbound Flights: \n"
			} else if j == 1 {
				output += "Return Flights: \n"
			}

			for _, seg := range e.Slice[j].Segment {
				output += fmt.Sprintf("Duration: %d", seg.Duration)
				if seg.ConnectionDuration > 0 {
					output += fmt.Sprintf(", Layover %d", seg.ConnectionDuration)
				}
				output += "\n"

				for _, leg := range seg.Leg {
					output += fmt.Sprintf("|%-23s|%-23s|\n|%-23s|%-23s|\n",
						leg.Origin, leg.Destination,
						leg.DepartureTime, leg.ArrivalTime)
				}
			}

		}
		output += "-------------------------------------------------\n\n"
	}
	return output
}
