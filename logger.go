/*
システム全体で使うことのできるログを一つ作成しておきたい？
*/
package logging

import (
	"os"
	"log"
)

func GetLogger(format string) Logging {
	return New(
	os.Stdout,
	format,
	log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
}