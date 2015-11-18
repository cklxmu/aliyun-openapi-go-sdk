package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

func main() {
	meta_root_path := os.Args[1]
	output_root_path := os.Args[2]

	product_paths, err := ListDir(meta_root_path, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, product_path := range product_paths {
		version_paths, err := ListDir(product_path, "")
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		for _, version_path := range version_paths {
			infoFilePath := version_path + "\\" + "Version-Info.json"
			var infoJson VersionInfoJson
			err = UnmarshalVersionInfo(infoFilePath, &infoJson)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			outputPath := output_root_path + "\\" + strings.ToLower(strings.Replace(infoJson.Product, "-", "_", -1)) + "\\" + infoJson.Name
			err = os.MkdirAll(outputPath, 7777)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			err = GenerateVersionCode(&infoJson, outputPath+"\\"+"version.go")
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			apiFiles, err := ListDir(version_path+"\\Api", "")
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			for _, apiFile := range apiFiles {
				err = ApiJson2GoSource(apiFile, outputPath, infoJson.ApiStyle)
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
					os.Exit(1)
				}
			}
		}
	}

	// err := ApiJson2GoSource(os.Args[1], "")
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err.Error())
	// 	os.Exit(1)
	// }
	return
}

func UnmarshalVersionInfo(inputFile string, v *VersionInfoJson) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

func ApiJson2GoSource(inputFile, outputPath string, apiStyle string) error {
	var api ApiJson
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	err = json.Unmarshal(data, &api)
	if err != nil {
		return err
	}
	api.ApiStyle = apiStyle
	outputFile := outputPath + "\\" + api.Name + ".go"
	outputTestFile := outputPath + "\\" + api.Name + "_test.go"
	err = GenerateCode(&api, outputFile)
	if err != nil {
		return err
	}
	err = GenerateTestCode(&api, outputTestFile)
	if err != nil {
		return err
	}
	err = exec.Command("go", "fmt", outputFile, outputTestFile).Run()
	if err != nil {
		return err
	}
	return nil
}

func EndpointsXML2GoSource(inputFile, outputPath string) error {
	var endpoints EndpointsXML
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	err = xml.Unmarshal(data, &endpoints)
	if err != nil {
		return err
	}
	outputFile := outputPath + "endpoints.go"
	err = GenerateEndpointsCode(&endpoints, outputFile)
	if err != nil {
		return err
	}
	err = exec.Command("go", "fmt", outputFile).Run()
	if err != nil {
		return err
	}
	return nil
}
