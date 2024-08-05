package main

import (
	"fmt"
	"taxation/filemanager"
	"taxation/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChannels[index], errorChannels[index])

		//if err != nil {
		//	fmt.Println("could not process job")
		//	fmt.Println(err)
		//}
	}

	for index := range taxRates {
		select {
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChannels[index]:
			fmt.Println("done")
		}
	}

}
