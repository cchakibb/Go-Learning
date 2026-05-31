package main

type PricingStrategy interface {
	Calculate(basePrice float64, nights int) float64
}

type StandardPricing struct {

}

type WeekendPricing struct {

}

func (p *StandardPricing) Calculate(basePrice float64, nights int) float64 {
	return basePrice * float64(nights)
}

func (p *WeekendPricing) Calculate(basePrice float64, nights int) float64 {
	return basePrice * float64(nights) * 1.5
}