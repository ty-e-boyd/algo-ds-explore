package tools

import "github.com/charmbracelet/log"

type searchAlgo func(set []string) bool

func RunSearchTests(i []string, o []string, algo searchAlgo) bool {
	log.Info("Going to Run Tests")
	return true
}
