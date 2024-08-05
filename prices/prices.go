package prices

import (
	"fmt"
	"taxation/conversion"
	"taxation/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChannel chan bool, errorChannel chan error) {
	err := job.LoadData()
	if err != nil {
		errorChannel <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result

	err = job.IOManager.WriteResult(job)

	if err != nil {
		errorChannel <- err
		return
	}

	doneChannel <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
