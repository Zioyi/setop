package internal

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadSet(filePath string) []string {
	records := readCsv(filePath)
	set := []string{}
	for _, r := range records {
		set = append(set, r[0])
	}
	return set
}

func readCsv(filePath string) [][]string {
	fs, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
		os.Exit(0)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
		os.Exit(0)
	}

	return records
}

func WriteSet(set []string, filePath string) {
	records := [][]string{}
	for _, i := range set {
		records = append(records, []string{i})
	}
	writeCsv(records, filePath)
}

func writeCsv(records [][]string, filePath string) {
	if exists(filePath) {
		log.Fatalf("the file is existed, can not overwrite")
		os.Exit(0)
	}

	fs, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("can not create the file, err is %+v", err)
		os.Exit(0)
	}
	defer fs.Close()

	w := csv.NewWriter(fs)
	err = w.WriteAll(records)
	if err != nil {
		log.Fatalf("can not writeall, err is %+v", err)
		os.Exit(0)
	}
}

// 判断文件或文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
