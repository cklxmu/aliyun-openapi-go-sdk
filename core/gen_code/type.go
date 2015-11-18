package main

import (
	"fmt"
	"os"
	"strings"
)

// Version-Info.json
type VersionInfoJson struct {
	ApiStyle string `json:"apiStyle"`
	Apis     struct {
		Apis []struct {
			Name string `json:"name"`
		} `json:"apis"`
	} `json:"apis"`
	Name    string `json:"name"`
	Product string `json:"product"`
}

// Api/***.json
type ApiParameter struct {
	Required    string `json:"required"`
	TagName     string `json:"tagName"`
	TagPosition string `json:"tagPosition"`
	Type        string `json:"type"`
	ValueSwitch struct {
		Cases []struct {
			TagValue string `json:"tagValue"`
		} `json:"cases"`
	} `json:"valueSwitch"`
}
type ListElement struct {
	ItemName string `json:"itemName"`
	TagName  string `json:"tagName"`
}
type MemberElement struct {
	TagName     string `json:"tagName"`
	Type        string `json:"type"`
	ValueSwitch struct {
		Cases []struct {
			TagValue string `json:"tagValue"`
		} `json:"cases"`
	} `json:"valueSwitch"`
}
type StructElement struct {
	ResultMap
	TagName string `json:"tagName"`
}
type ArrayElement struct {
	ItemName string `json:"itemName"`
	StructElement
}
type ResultMap struct {
	Arrays  []ArrayElement  `json:"arrays"`
	Lists   []ListElement   `json:"lists"`
	Members []MemberElement `json:"members"`
	Structs []StructElement `json:"structs"`
}
type ApiJson struct {
	ApiStyle    string
	IsvProtocol struct {
		Method   string `json:"method"`
		Pattern  string `json:"pattern"`
		Protocol string `json:"protocol"`
	} `json:"isvProtocol"`

	Name string `json:"name"`

	Parameters struct {
		Parameters []ApiParameter `json:"parameters"`
	} `json:"parameters"`

	Product string `json:"product"`

	ResultMapping ResultMap `json:"resultMapping"`

	Version string `json:"version"`
}

var typeMap map[string]string = map[string]string{
	"string":  "string",
	"integer": "int",
	"long":    "int",
	"boolean": "bool",
	"float":   "float32",
	"double":  "float64",
}

func writef(file *os.File, format string, args ...interface{}) (int, error) {
	return file.WriteString(fmt.Sprintf(format, args...))
}

func parseElement(file *os.File, element *ResultMap, prefix string) error {
	prefix = "\t" + prefix
	for _, member := range element.Members {
		if _, exist := typeMap[strings.ToLower(member.Type)]; !exist {
			return fmt.Errorf("unknown response member type %s \n", member.Type)
		}
		writef(file, prefix+"%s %s `xml:\"%s\" json:\"%s\"`\n", member.TagName, typeMap[strings.ToLower(member.Type)], member.TagName, member.TagName)
	}
	for _, list := range element.Lists {
		writef(file, prefix+"%s struct {\n", list.TagName)
		writef(file, prefix+"\t%s []string `xml:\"%s\" json:\"%s\"`\n", list.ItemName, list.ItemName, list.ItemName)
		writef(file, prefix+"} `xml:\"%s\" json:\"%s\"`\n", list.TagName, list.TagName)
	}
	for _, array := range element.Arrays {
		writef(file, prefix+"%s struct {\n", array.TagName)
		writef(file, prefix+"\t%s []struct {\n", array.ItemName)
		err := parseElement(file, &array.ResultMap, prefix)
		if err != nil {
			return err
		}
		writef(file, prefix+"\t} `xml:\"%s\" json:\"%s\"`\n", array.ItemName, array.ItemName)
		writef(file, prefix+"} `xml:\"%s\" json:\"%s\"`\n", array.TagName, array.TagName)
	}
	for _, struc := range element.Structs {
		writef(file, prefix+"%s struct {\n", struc.TagName)
		err := parseElement(file, &struc.ResultMap, prefix)
		if err != nil {
			return err
		}
		writef(file, prefix+"} `xml:\"%s\" json:\"%s\"`\n", struc.TagName, struc.TagName)
	}

	return nil
}

func GenerateCode(api *ApiJson, outputFile string) error {
	fmt.Println("generate:", outputFile)
	ActionName := api.Name
	ProductName := api.Product
	requestType := ActionName + "Request"
	responseType := ActionName + "Response"

	codefile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer codefile.Close()

	//package & import
	importStrconv := false
	for _, param := range api.Parameters.Parameters {
		if param.Type != "String" && param.Type != "List" {
			importStrconv = true
			break
		}
	}
	_, err = writef(codefile, "package "+strings.ToLower(strings.Replace(ProductName, "-", "_", -1))+"\n")
	writef(codefile, "import (\n\t. \"aliyun-openapi-go-sdk/core\"\n")
	if importStrconv {
		writef(codefile, "\t\"strconv\"\n")
	}
	writef(codefile, ")\n")

	//request type definition
	paramTypes := make(map[string]string)
	writef(codefile, "\ntype "+requestType+" struct {\n")
	if api.ApiStyle == "ROA" {
		writef(codefile, "\tRoaRequest\n")
	} else {
		writef(codefile, "\tRpcRequest\n")
	}
	for _, param := range api.Parameters.Parameters {
		paramName := strings.Replace(param.TagName, ".", "_", -1)
		if paramName == "type" {
			paramName = "Type"
		}
		writef(codefile, "\t"+paramName+"\t")
		switch strings.ToLower(param.Type) {
		case "string":
			paramTypes[param.TagName] = "string"
			writef(codefile, "string\n")
		case "list": //TODO(K'):change list's type?
			paramTypes[param.TagName] = "string"
			writef(codefile, "string\n")
		case "integer":
			paramTypes[param.TagName] = "int"
			writef(codefile, "int\n")
		case "long":
			paramTypes[param.TagName] = "int"
			writef(codefile, "int\n")
		case "boolean":
			paramTypes[param.TagName] = "bool"
			writef(codefile, "bool\n")
		case "float":
			paramTypes[param.TagName] = "float32"
			writef(codefile, "float32\n")
		case "double":
			paramTypes[param.TagName] = "float64"
			writef(codefile, "float64\n")
		default:
			return fmt.Errorf("unknown request param type %s\n", param.Type)
		}
	}
	writef(codefile, "}\n\n")

	//request type setter & getter
	for _, param := range api.Parameters.Parameters {
		queryParams := "QueryParams"
		if param.TagPosition == "Path" {
			queryParams = "PathParams"
		}
		paramName := strings.Replace(param.TagName, ".", "_", -1)
		if paramName == "type" {
			paramName = "Type"
		}
		writef(codefile, "func (r *"+requestType+") Set"+paramName+"(value "+paramTypes[param.TagName]+") {\n")
		writef(codefile, "\tr."+paramName+" = value\n")
		switch paramTypes[param.TagName] {
		case "int":
			writef(codefile, "\tr.%s.Set(\"%s\", strconv.Itoa(value))\n", queryParams, param.TagName)
		case "bool":
			writef(codefile, "\tr.%s.Set(\"%s\", strconv.FormatBool(value))\n", queryParams, param.TagName)
		case "string":
			writef(codefile, "\tr.%s.Set(\"%s\", value)\n", queryParams, param.TagName)
		case "float32":
			writef(codefile, "\tr.%s.Set(\"%s\", strconv.FormatFloat(value, 'f', 5, 32))\n", queryParams, param.TagName)
		case "float64":
			writef(codefile, "\tr.%s.Set(\"%s\", strconv.FormatFloat(value, 'f', 5, 64))\n", queryParams, param.TagName)
		}
		writef(codefile, "}\n")
		writef(codefile, "func (r *%s) Get%s() %s {\n", requestType, paramName, paramTypes[param.TagName])
		writef(codefile, "\treturn r.%s\n", paramName)
		writef(codefile, "}\n")
	}

	//request init function definition
	writef(codefile, "\nfunc (r *%s) Init() {\n", requestType)
	if api.ApiStyle == "ROA" {
		writef(codefile, "\t r.RoaRequest.Init()\n")
		if api.IsvProtocol.Pattern != "" {
			writef(codefile, "\t r.PathPattern = \"%s\"\n", api.IsvProtocol.Pattern)
		}
	} else {
		writef(codefile, "\t r.RpcRequest.Init()\n")
		writef(codefile, "\t r.SetVersion(Version)\n")
		writef(codefile, "\t r.SetAction(\"%s\")\n", ActionName)
	}
	if !strings.Contains(api.IsvProtocol.Method, "|") {
		writef(codefile, "\t r.SetMethod(\"%s\")\n", api.IsvProtocol.Method)
	}
	if !strings.Contains(api.IsvProtocol.Protocol, "|") {
		writef(codefile, "\t r.SetProtocol(\"%s\")\n", api.IsvProtocol.Protocol)
	}
	writef(codefile, "\t r.SetProduct(Product)\n")
	writef(codefile, "}\n")

	//response type definition
	writef(codefile, "\ntype %s struct {\n", responseType)
	err = parseElement(codefile, &api.ResultMapping, "")
	if err != nil {
		return fmt.Errorf("parseElement: %s\n", err.Error())
	}
	writef(codefile, "}\n")

	//function definition
	writef(codefile, "func %s(req *%s, accessId, accessSecret string) (*%s, error) {\n", ActionName, requestType, responseType)
	writef(codefile, "\tvar pResponse %s\n", responseType)
	writef(codefile, "\tbody, err := ApiHttpRequest(accessId, accessSecret, req)\n")
	writef(codefile, "\tif err != nil {\n")
	writef(codefile, "\t\treturn nil,err\n")
	writef(codefile, "\t}\n")
	writef(codefile, "\tApiUnmarshalResponse(req.GetFormat(), body, &pResponse)\n")
	writef(codefile, "\treturn &pResponse, err\n")
	writef(codefile, "}\n")
	return nil
}

func GenerateTestCode(api *ApiJson, outputFile string) error {
	fmt.Println("generate:", outputFile)
	ActionName := api.Name
	ProductName := api.Product
	requestType := ActionName + "Request"

	codefile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer codefile.Close()

	writef(codefile, "package "+strings.ToLower(strings.Replace(ProductName, "-", "_", -1))+"\n")
	writef(codefile, `
import (
	"fmt"
	"testing"
)`)
	writef(codefile, "\nfunc Test%s(t *testing.T) {\n", ActionName)
	writef(codefile, "\tvar req %s\n", requestType)
	writef(codefile, "\treq.Init()\n")
	writef(codefile, "\treq.SetFormat(\"JSON\")\n")
	writef(codefile, "\treq.SetRegionId(\"cn-shenzhen\")\n")
	writef(codefile, "\tvar accessId = \"Ie65kUInu5GeAsma\"\n")
	writef(codefile, "\tvar accessSecret = \"8cCqoxdYU9zKUihwXFXiN1HEACBDwB\"\n")
	writef(codefile, "\tresp, err := %s(&req, accessId, accessSecret)\n", ActionName)
	writef(codefile, "\tif err != nil {\n")
	writef(codefile, "\t\tt.Errorf(\"Error: %%s\", err.Error())\n")
	writef(codefile, "\t}\n")
	writef(codefile, "\tfmt.Printf(\"Success: %%v\\n\", resp)\n")
	writef(codefile, "}\n")
	return nil
}

type EndpointsXML struct {
	Endpoint []struct {
		Name      string `xml:"name,attr"`
		RegionIds struct {
			RegionId []string `xml:"RegionId"`
		} `xml:"RegionIds"`
		Products struct {
			Product []struct {
				ProductName string `xml:"ProductName"`
				DomainName  string `xml:"DomainName"`
			} `xml:"Product"`
		} `xml:"Products"`
	} `xml:"Endpoint"`
}

func GenerateEndpointsCode(endpoints *EndpointsXML, outputFile string) error {
	codefile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer codefile.Close()

	writef(codefile, "package core\n")
	writef(codefile, "const endpointProductDomainMap = map[string]map[string]string{\n")
	for _, endpoint := range endpoints.Endpoint {
		writef(codefile, "\t\"%s\": {\n", endpoint.Name)
		for _, product := range endpoint.Products.Product {
			writef(codefile, "\t\t\"%s\": \"%s\",\n", product.ProductName, product.DomainName)
		}
		writef(codefile, "\t},\n")
	}
	writef(codefile, "}\n")

	return nil
}

func GenerateVersionCode(info *VersionInfoJson, outputFile string) error {
	fmt.Println("generate:", outputFile)
	Version := info.Name
	ProductName := info.Product
	ApiStyle := info.ApiStyle

	codefile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer codefile.Close()

	writef(codefile, "package "+strings.ToLower(strings.Replace(ProductName, "-", "_", -1))+"\n\n")
	writef(codefile, "var Version = \"%s\"\n", Version)
	writef(codefile, "var Product = \"%s\"\n", ProductName)
	writef(codefile, "var ApiStyle = \"%s\"\n", ApiStyle)
	return nil
}
