package utils

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

//type Config struct {
//	XMLName          xml.Name `xml:"config"`
//	Text             string   `xml:",chardata"`
//	T                string   `xml:"t,attr"`
//	ID               string   `xml:"id,attr"`
//	Configversion    string   `xml:"configversion,attr"`
//	Name             string   `xml:"name,attr"`
//	Version          string   `xml:"version,attr"`
//	Outputparameters struct {
//		Text      string `xml:",chardata"`
//		Parameter []struct {
//			Text        string `xml:",chardata"`
//			Name        string `xml:"name,attr"`
//			Description string `xml:"description,attr"`
//			Type        string `xml:"type,attr"`
//			Format      string `xml:"format,attr"`
//			Valid       string `xml:"valid,attr"`
//			Value       string `xml:"value,attr"`
//		} `xml:"parameter"`
//	} `xml:"outputparameters"`
//}

type SmbConfig struct {
	XMLName          xml.Name `xml:"config"`
	Xmlns            string   `xml:"xmlns,attr"`
	Outputparameters struct {
		Parameter []struct {
			Name  string `xml:"name,attr"`
			Value string `xml:"value"`
		} `xml:"parameter"`
	} `xml:"outputparameters"`
}
type IisConfig struct {
	XMLName          xml.Name `xml:"config"`
	Xmlns            string   `xml:"xmlns,attr"`
	Outputparameters struct {
		Parameter []struct {
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"parameter"`
	} `xml:"outputparameters"`
}

func ParseIisXml(FilePath string) map[string]string {
	// 打开 XML 文件
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// 解码 XML 数据
	outputParams := new(IisConfig)
	err = xml.NewDecoder(file).Decode(outputParams)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return nil
	}
	results := make(map[string]string)
	for _, v := range outputParams.Outputparameters.Parameter {
		results[v.Name] = v.Value
	}
	return results
}
func ParseSmbXml(FilePath string) map[string]string {
	// 打开 XML 文件
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// 解码 XML 数据
	outputParams := new(SmbConfig)
	err = xml.NewDecoder(file).Decode(outputParams)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return nil
	}
	results := make(map[string]string)
	for _, v := range outputParams.Outputparameters.Parameter {
		results[v.Name] = v.Value
	}
	return results
}
func CreateXmlFileName(Function string) string {
	now := time.Now()
	formattedTime := now.Format("2006-01-02.15.04.05.000000")
	return fmt.Sprintf("%s-%s.xml", formattedTime, Function)
}
