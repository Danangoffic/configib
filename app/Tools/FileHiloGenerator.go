package Tools

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// FileHiloGenerator
type fhiloGenerator struct {
}

type FHiloGenerator interface {
	GenerateUniqueCode(targetType string) string
}

func FileHiloGenerator() *fhiloGenerator {
	return &fhiloGenerator{}
}

func (h *fhiloGenerator) GenerateUniqueCode(targetType string) string {
	sdf2 := "yyyyMMdd"
	result := ""
	dateNow := time.Now().Format(sdf2)

	result = targetType + "-" + dateNow + time.Millisecond.String() + Generator().NextId()
	return result
}

// Id Generator
type IdGenerator interface {
	NextId() string
	TransRefNumberV3(transRef, prefix string) string
	TransRefNumSunline(transRef string) string
}

type generator struct {
}

func Generator() *generator {
	return &generator{}
}

func (g *generator) NextId() string {
	var (
		IdLength           = 20
		maxLowCtr          = 9999
		txIdFile           = "txid.counter"
		txIdFilePermission = fs.ModeType
	)

	var (
		ltxidCounter int
		htxidCounter int
		cdoy         int
	)

	cdate := time.Now()

	if ltxidCounter > maxLowCtr || cdate.YearDay() != cdoy {
		ltxidCounter = 0
		htxidCounter = -1
		cdoy = cdate.YearDay()
	}

	if htxidCounter == -1 {
		f, err := os.OpenFile(txIdFile, os.O_RDWR|os.O_CREATE, txIdFilePermission)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if fileInfo, err := f.Stat(); err == nil {
			fdate := fileInfo.ModTime()
			if fdate.YearDay() == cdoy {
				if bytes, err := ioutil.ReadFile(txIdFile); err == nil {
					htstr := string(bytes)
					if htstr != "" {
						htstr = "00000000000000000000" + htstr
						htstr = htstr[len(htstr)-12:]
						if parsedHtxidCounter, err := strconv.ParseInt(htstr, 10, 64); err == nil {
							varInInt := strconv.FormatInt(parsedHtxidCounter, 10)
							htxidCounter, err = strconv.Atoi(varInInt)
						}
					}
				}
			}
		}

		if htxidCounter == -1 {
			htxidCounter = 0
			ltxidCounter = 1
		} else {
			ltxidCounter = ltxidCounter + 1
		}

		htstr := fmt.Sprintf("%012d", htxidCounter+1)
		if _, err := f.WriteAt([]byte(htstr+"\n"), 0); err != nil {
			panic(err)
		}
	}

	str := "0000" + strconv.Itoa(ltxidCounter)
	str = fmt.Sprintf("%020d", htxidCounter) + str[len(str)-len(strconv.Itoa(maxLowCtr)):]
	ltxidCounter++

	return str[len(str)-IdLength:]
}

func (g *generator) TransRefNumberV3(transRef, prefix string) string {
	return transRefNumberV3(transRef, 12, prefix)
}

// Function ini berfungsi untuk generate Transaction Referrence Number untuk account sunline
//
//	args : transRef string
//	return : string
func (g *generator) TransRefNumSunline(transRef string) string {
	length := 12
	start := len(transRef) - length
	var transNumber string
	transNumber = fmt.Sprint("SL-", transRef[start])
	return transNumber
}

func transRefNumberV3(transRef string, length int, prefix string) string {
	start := len(transRef) - length
	var transNumber string
	if prefix == "" || strings.EqualFold(prefix, "") || len(prefix) == 0 {
		prefix = "MB"
	}
	transNumber = fmt.Sprint(prefix, "-", transRef[start])
	return transNumber
}
