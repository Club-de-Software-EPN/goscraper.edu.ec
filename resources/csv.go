package resources

import(
	"os"
	"encoding/csv"
	"fmt"
)

func SaveDataOnCSVFile(data [][]string, fileName string){
	// "fileName" param only refers to the name of the file but not to its extension

	// Create a file
	membersFile, err := os.Create(fmt.Sprintf("%s.csv", fileName))
	// Close file using defer
	defer membersFile.Close()

	if err != nil {
		fmt.Println("An error has ocurred")
	}

	// create a csv writer
	writerCsvFile := csv.NewWriter(membersFile)
	// buffer --> file "".csv
	defer writerCsvFile.Flush()

	// save data as csv format rows
	for _, rowMember := range data {
		_ = writerCsvFile.Write(rowMember)
	}
}