package main

func StartDecipher(senderChan chan string, decipherer func(encrypted string) string) chan string {
	receiverChan := make(chan string, 5)
	go decode(senderChan, receiverChan, decipherer)
	return receiverChan
}

func decode(sender <-chan string, receiver chan<- string, decipherer func(encrypted string) string) {
	for {
		encrypted, isOpen := <-sender
		if !isOpen {
			return
		}

		receiver <- decipherer(encrypted)
	}
}
