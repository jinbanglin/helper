package helper

import (
	"errors"
	"log"
	"os"
	"os/signal"

	"github.com/deckarep/golang-set"
)

var gsignal = mapset.NewSet()

var SigAlreadyRegisted = errors.New("sig already registed")

func RegisterSignal(sig os.Signal, process func()) error {
	sigChan, err := register(sig)
	if err != nil {
		return err
	}
	go func() {
		msg := <-sigChan
		log.Println("sig", msg, "received")
		process()
	}()
	return nil
}

func register(sig os.Signal) (chan os.Signal, error) {
	if gsignal.Contains(sig) {
		log.Fatal("signal", sig, "already registed")
		return nil, SigAlreadyRegisted
	}
	gsignal.Add(sig)
	sig_chan := make(chan (os.Signal))
	signal.Notify(sig_chan, sig)
	return sig_chan, nil
}

func RegisterContinueSignal(sig os.Signal, process func()) error {
	sig_chan, err := register(sig)
	if err != nil {
		return err
	}
	go func() {
		for {
			msg := <-sig_chan
			log.Println("sig", msg, "receveid")
			process()
		}
	}()
	return nil
}

func SignalProcessed(sig os.Signal) bool {
	return gsignal.Contains(sig)
}
