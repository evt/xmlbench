package xmlbench

type FlightResults struct {
	FlightResponse FlightResponse
}

type FlightResponse struct {
	OffersGroup OffersGroup `xml:"OffersGroup"`
}

type OffersGroup struct {
	AirlineOffers []AirlineOffer `xml:"AirlineOffers>AirlineOffer"`
}

type AirlineOffer struct {
	OfferID    OfferID `xml:"OfferID"`
	TimeLimits TimeLimits
}

type OfferID struct {
	Owner string `xml:",attr"`
	Value string `xml:",chardata"`
}

type TimeLimits struct {
	OfferExpiration OfferExpiration
}

type OfferExpiration struct {
	DateTime string `xml:",attr"`
}
