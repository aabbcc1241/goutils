/*
 * reference : https://www.goinggo.net/2013/11/using-log-package-in-go.html (github.com/ardanlabs)
 */
package log

import (
  "io"
  "io/ioutil"
  "log"
  "os"
)

var (
  Info  *log.Logger
  Debug *log.Logger
  Error *log.Logger
)

func Init(infoVerbose bool, debugVerbose bool, errorVerbose bool) {
  commFlag := log.Ldate | log.Ltime | log.Lshortfile
  var infoLog io.Writer
  var debugLog io.Writer
  var errorLog io.Writer
  if infoVerbose {
    infoLog = os.Stdout
  } else {
    infoLog = ioutil.Discard
  }
  if debugVerbose {
    debugLog = os.Stdout
  } else {
    debugLog = ioutil.Discard
  }
  if errorVerbose {
    errorLog = os.Stderr
  } else {
    errorLog = ioutil.Discard
  }
  Info = log.New(infoLog, "Info: ", commFlag)
  Debug = log.New(debugLog, "Debug: ", commFlag)
  Error = log.New(errorLog, "Error: ", commFlag)
}

func init() {
  Init(true, true, true)
}
