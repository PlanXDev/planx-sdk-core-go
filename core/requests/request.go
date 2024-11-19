package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	AppIDName = "appid"
	NonceName = "nonce"
	SignName  = "sign"

	GET  = "GET"
	POST = "POST"

	Header = "Header"
	Query  = "Query"
	Body   = "Body"
	Path   = "Path"
)

type ComRequest interface {
	GetMethod() string
	GetPathParams() map[string]string
	GetHeaders() map[string]string
	GetQueryParams() map[string]string
	GetFormParams() map[string]interface{}
	GetBodyReader() io.Reader
	GetReadTimeout() time.Duration
	GetConnectTimeout() time.Duration
	SetReadTimeout(readTimeout time.Duration)
	SetConnectTimeout(connectTimeout time.Duration)

	GetBuildUrl() string
	SetBuildUrl(url string)

	addHeaderParam(key, value string)
	addQueryParam(key, value string)
	addFormParam(key string, value interface{})
	addPathParam(key, value string)
}

type BaseRequest struct {
	Method         string
	BuildUrl       string
	ReadTimeout    time.Duration
	ConnectTimeout time.Duration

	PathParams  map[string]string
	QueryParams map[string]string
	Headers     map[string]string
	FormParams  map[string]interface{}
}

func (request *BaseRequest) GetPathParams() map[string]string {
	return request.PathParams
}

func (request *BaseRequest) GetQueryParams() map[string]string {
	return request.QueryParams
}

func (request *BaseRequest) GetFormParams() map[string]interface{} {
	return request.FormParams
}

func (request *BaseRequest) GetReadTimeout() time.Duration {
	return request.ReadTimeout
}

func (request *BaseRequest) GetConnectTimeout() time.Duration {
	return request.ConnectTimeout
}

func (request *BaseRequest) SetReadTimeout(readTimeout time.Duration) {
	request.ReadTimeout = readTimeout
}

func (request *BaseRequest) SetConnectTimeout(connectTimeout time.Duration) {
	request.ConnectTimeout = connectTimeout
}

func (request *BaseRequest) addHeaderParam(key, value string) {
	request.Headers[key] = value
}

func (request *BaseRequest) addQueryParam(key, value string) {
	request.QueryParams[key] = value
}

func (request *BaseRequest) addPathParam(key, value string) {
	request.PathParams[key] = value
}

func (request *BaseRequest) addFormParam(key string, value interface{}) {
	request.FormParams[key] = value
}

func (request *BaseRequest) GetHeaders() map[string]string {
	return request.Headers
}

func (request *BaseRequest) GetBuildUrl() string {
	url := request.buildPath()
	querystring := request.buildQueryString()
	if len(querystring) > 0 {
		url = fmt.Sprintf("%s?%s", url, querystring)
	}
	return url
}

func (request *BaseRequest) SetBuildUrl(url string) {
	request.BuildUrl = url
}

func (request *BaseRequest) buildQueryString() string {
	queryParams := request.QueryParams
	// sort QueryParams by key
	q := url.Values{}
	for key, value := range queryParams {
		q.Add(key, value)
	}
	return q.Encode()
}

func (request *BaseRequest) buildPath() string {
	path := request.BuildUrl
	for key, value := range request.PathParams {
		path = strings.Replace(path, "["+key+"]", value, 1)
	}
	return path
}

func (request *BaseRequest) GetBodyReader() io.Reader {
	if request.FormParams != nil && len(request.FormParams) > 0 {
		marshal, _ := json.Marshal(request.FormParams)
		return strings.NewReader(string(marshal))
	} else {
		return nil
	}
}

func (request *BaseRequest) GetMethod() string {
	return request.Method
}

func NewGetRequest(path string) (request *BaseRequest) {
	request = &BaseRequest{
		Method:      GET,
		BuildUrl:    path,
		QueryParams: make(map[string]string),
		Headers:     make(map[string]string),
		FormParams:  make(map[string]interface{}),
		PathParams:  make(map[string]string),
	}
	return
}

func NewPostRequest(path string) (request *BaseRequest) {
	request = &BaseRequest{
		Method:      POST,
		BuildUrl:    path,
		QueryParams: make(map[string]string),
		Headers:     make(map[string]string),
		FormParams:  make(map[string]interface{}),
		PathParams:  make(map[string]string),
	}
	return
}

func InitParams(request ComRequest) (err error) {
	requestValue := reflect.ValueOf(request).Elem()
	err = flatRepeatedList(requestValue, request, "", "")
	return
}

func flatRepeatedList(dataValue reflect.Value, request ComRequest, position, prefix string) (err error) {
	dataType := dataValue.Type()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		name, containsNameTag := field.Tag.Lookup("name")
		fieldPosition := position
		if fieldPosition == "" {
			fieldPosition, _ = field.Tag.Lookup("position")
		}
		typeTag, containsTypeTag := field.Tag.Lookup("type")
		bindingTag, containsBindingTag := field.Tag.Lookup("binding")
		if containsBindingTag {
			err = checkParam(bindingTag, dataValue, name, i)
			if err != nil {
				return
			}
		}
		if containsNameTag {
			if !containsTypeTag {
				// simple param
				key := prefix + name
				if dataValue.Field(i).Kind().String() == "map" || dataValue.Field(i).Kind().String() == "slice" {
					err = addParam(request, fieldPosition, key, dataValue.Field(i).Interface())
					if err != nil {
						return
					}
				} else {
					value := dataValue.Field(i).String()
					err = addParam(request, fieldPosition, key, value)
					if err != nil {
						return
					}
				}

			} else if typeTag == "Repeated" {
				// repeated param
				err = handleRepeatedParams(request, dataValue, prefix, name, fieldPosition, i)
				if err != nil {
					return
				}
			} else if typeTag == "Struct" {
				err = handleStruct(request, dataValue, prefix, name, fieldPosition, i)
				if err != nil {
					return
				}
			} else if typeTag == "Map" {
				err = handleMap(request, dataValue, prefix, name, fieldPosition, i)
				if err != nil {
					return err
				}
			} else if typeTag == "Json" {
				byt, err := json.Marshal(dataValue.Field(i).Interface())
				if err != nil {
					return err
				}
				key := prefix + name
				err = addParam(request, fieldPosition, key, string(byt))
				if err != nil {
					return err
				}
			}
		}
	}
	return
}

func checkParam(bindingTag string, dataValue reflect.Value, name string, i int) error {
	tags := strings.Split(bindingTag, ",")
	if len(tags) > 0 {
		for _, tag := range tags {
			if tag == "required" {
				if !dataValue.Field(i).IsValid() {
					return errors.New(fmt.Sprintf("param %s is required", name))
				} else if dataValue.Field(i).Kind().String() == "string" {
					if dataValue.Field(i).String() == "" {
						return errors.New(fmt.Sprintf("param %s is required", name))
					}
				}
			}
			if dataValue.Field(i).IsValid() {
				if strings.Contains(tag, "lt") || strings.Contains(tag, "gt") ||
					strings.Contains(tag, "lte") || strings.Contains(tag, "gte") {
					err := checkNum(tag, dataValue, name, i)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func checkNum(tag string, dataValue reflect.Value, name string, i int) error {
	tags := strings.Split(tag, "=")
	if len(tags) != 2 {
		return errors.New(fmt.Sprintf("param %s tag format error", name))
	}
	cmp := 0
	if dataValue.Field(i).Kind().String() == "string" {
		v, err := strconv.ParseFloat(dataValue.Field(i).String(), 64)
		if err != nil {
			return errors.New(fmt.Sprintf("param %s format number error", name))
		}
		fltn, err := strconv.ParseFloat(tags[1], 64)
		if err != nil {
			return errors.New(fmt.Sprintf("param %s format number error", name))
		}
		if v < fltn {
			cmp = -1
		}
		if v > fltn {
			cmp = 1
		}
	}
	if strings.Contains(dataValue.Field(i).Kind().String(), "int") {
		intn, err := strconv.ParseInt(tags[1], 10, 64)
		if err != nil {
			return errors.New(fmt.Sprintf("param %s format number error", name))
		}
		if dataValue.Field(i).Int() < intn {
			cmp = -1
		}
		if dataValue.Field(i).Int() > intn {
			cmp = 1
		}
	}
	if strings.Contains(dataValue.Field(i).Kind().String(), "float") {
		fltn, err := strconv.ParseFloat(tags[1], 64)
		if err != nil {
			return errors.New(fmt.Sprintf("param %s format number error", name))
		}
		if dataValue.Field(i).Float() < fltn {
			cmp = -1
		}
		if dataValue.Field(i).Float() > fltn {
			cmp = 1
		}
	}
	switch tags[0] {
	case "lt":
		if cmp >= 0 {
			return errors.New(fmt.Sprintf("param %s must be less than to %s", name, tags[1]))
		}
	case "gt":
		if cmp <= 0 {
			return errors.New(fmt.Sprintf("param %s must be greater than to %s", name, tags[1]))
		}
	case "lte":
		if cmp > 0 {
			return errors.New(fmt.Sprintf("param %s must be less than or equal to %s", name, tags[1]))
		}
	case "gte":
		if cmp < 0 {
			return errors.New(fmt.Sprintf("param %s must be greater than or equal to %s", name, tags[1]))
		}
	}
	return nil
}

func handleRepeatedParams(request ComRequest, dataValue reflect.Value, prefix, name, fieldPosition string, index int) (err error) {
	repeatedFieldValue := dataValue.Field(index)
	if repeatedFieldValue.Kind() != reflect.Slice {
		// possible value: {"[]string", "*[]struct"}, we must call Elem() in the last condition
		repeatedFieldValue = repeatedFieldValue.Elem()
	}
	if repeatedFieldValue.IsValid() && !repeatedFieldValue.IsNil() {
		for m := 0; m < repeatedFieldValue.Len(); m++ {
			elementValue := repeatedFieldValue.Index(m)
			key := prefix + name + "." + strconv.Itoa(m+1)
			if elementValue.Type().Kind().String() == "string" {
				value := elementValue.String()
				err = addParam(request, fieldPosition, key, value)
				if err != nil {
					return
				}
			} else {
				err = flatRepeatedList(elementValue, request, fieldPosition, key+".")
				if err != nil {
					return
				}
			}
		}
	}
	return nil
}

func handleParam(request ComRequest, dataValue reflect.Value, prefix, key, fieldPosition string) (err error) {
	if dataValue.Type().String() == "[]string" {
		if dataValue.IsNil() {
			return
		}
		for j := 0; j < dataValue.Len(); j++ {
			err = addParam(request, fieldPosition, key+"."+strconv.Itoa(j+1), dataValue.Index(j).String())
			if err != nil {
				return
			}
		}
	} else {
		if dataValue.Type().Kind().String() == "string" {
			value := dataValue.String()
			err = addParam(request, fieldPosition, key, value)
			if err != nil {
				return
			}
		} else if dataValue.Type().Kind().String() == "struct" {
			err = flatRepeatedList(dataValue, request, fieldPosition, key+".")
			if err != nil {
				return
			}
		} else if dataValue.Type().Kind().String() == "int" {
			value := dataValue.Int()
			err = addParam(request, fieldPosition, key, strconv.Itoa(int(value)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func handleMap(request ComRequest, dataValue reflect.Value, prefix, name, fieldPosition string, index int) (err error) {
	valueField := dataValue.Field(index)
	if valueField.IsValid() && !valueField.IsNil() {
		iter := valueField.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			key := prefix + name + ".#" + strconv.Itoa(k.Len()) + "#" + k.String()
			if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
				elementValue := v.Elem()
				err = handleParam(request, elementValue, prefix, key, fieldPosition)
				if err != nil {
					return err
				}
			} else if v.IsValid() && v.IsNil() {
				err = handleParam(request, v, prefix, key, fieldPosition)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func handleStruct(request ComRequest, dataValue reflect.Value, prefix, name, fieldPosition string, index int) (err error) {
	valueField := dataValue.Field(index)
	if valueField.IsValid() && valueField.String() != "" {
		valueFieldType := valueField.Type()
		for m := 0; m < valueFieldType.NumField(); m++ {
			fieldName := valueFieldType.Field(m).Name
			elementValue := valueField.FieldByName(fieldName)
			key := prefix + name + "." + fieldName
			if elementValue.Type().String() == "[]string" {
				if elementValue.IsNil() {
					continue
				}
				for j := 0; j < elementValue.Len(); j++ {
					err = addParam(request, fieldPosition, key+"."+strconv.Itoa(j+1), elementValue.Index(j).String())
					if err != nil {
						return
					}
				}
			} else {
				if elementValue.Type().Kind().String() == "string" {
					value := elementValue.String()
					err = addParam(request, fieldPosition, key, value)
					if err != nil {
						return
					}
				} else if elementValue.Type().Kind().String() == "struct" {
					err = flatRepeatedList(elementValue, request, fieldPosition, key+".")
					if err != nil {
						return
					}
				} else if !elementValue.IsNil() {
					repeatedFieldValue := elementValue.Elem()
					if repeatedFieldValue.IsValid() && !repeatedFieldValue.IsNil() {
						for m := 0; m < repeatedFieldValue.Len(); m++ {
							elementValue := repeatedFieldValue.Index(m)
							if elementValue.Type().Kind().String() == "string" {
								value := elementValue.String()
								err := addParam(request, fieldPosition, key+"."+strconv.Itoa(m+1), value)
								if err != nil {
									return err
								}
							} else {
								err = flatRepeatedList(elementValue, request, fieldPosition, key+"."+strconv.Itoa(m+1)+".")
								if err != nil {
									return
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func addParam(request ComRequest, position, name string, value interface{}) (err error) {
	if value != nil {
		switch position {
		case Header:
			request.addHeaderParam(name, value.(string))
		case Query:
			request.addQueryParam(name, value.(string))
		case Path:
			request.addPathParam(name, value.(string))
		case Body:
			request.addFormParam(name, value)
		default:
			err = errors.New("unknown position")
		}
	}
	return
}
