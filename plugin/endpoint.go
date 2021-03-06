package plugin

/*

ShieldEndpoints are used for store + targets. This code genericizes them and makes it easy for you to pull out arbitrary values from them. The plugin framework will feed your action methods with the appropriate endpoint, and you can pull whatever data out that you need.

*/

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ShieldEndpoint ...
type ShieldEndpoint map[string]interface{}

func getEndpoint(j string) (ShieldEndpoint, error) {
	if j == "" {
		return nil, fmt.Errorf("Missing required --endpoint flag")
	}
	endpoint := make(ShieldEndpoint)
	err := json.Unmarshal([]byte(j), &endpoint)
	if err != nil {
		return nil, JSONError{Err: fmt.Sprintf("Error trying parse --endpoint value as JSON: %s", err.Error())}
	}

	return endpoint, nil
}

// StringValue ...
func (endpoint ShieldEndpoint) StringValue(key string) (string, error) {
	_, ok := endpoint[key]
	if !ok {
		return "", EndpointMissingRequiredDataError{Key: key}
	}

	if reflect.TypeOf(endpoint[key]).Kind() != reflect.String {
		return "", EndpointDataTypeMismatchError{Key: key, DesiredType: "string"}
	}

	return endpoint[key].(string), nil
}

// StringValueDefault ...
func (endpoint ShieldEndpoint) StringValueDefault(key string, def string) (string, error) {
	s, err := endpoint.StringValue(key)
	if err == nil {
		return s, nil
	}
	if _, ok := err.(EndpointMissingRequiredDataError); ok {
		return def, nil
	}
	return "", err
}

// FloatValue ...
func (endpoint ShieldEndpoint) FloatValue(key string) (float64, error) {
	_, ok := endpoint[key]
	if !ok {
		return 0, EndpointMissingRequiredDataError{Key: key}
	}

	if reflect.TypeOf(endpoint[key]).Kind() != reflect.Float64 {
		return 0, EndpointDataTypeMismatchError{Key: key, DesiredType: "numeric"}
	}

	return endpoint[key].(float64), nil
}

// FloatValueDefault ...
func (endpoint ShieldEndpoint) FloatValueDefault(key string, def float64) (float64, error) {
	f, err := endpoint.FloatValue(key)
	if err == nil {
		return f, nil
	}
	if _, ok := err.(EndpointMissingRequiredDataError); ok {
		return def, nil
	}
	return 0.0, err
}

// BooleanValue ...
func (endpoint ShieldEndpoint) BooleanValue(key string) (bool, error) {
	_, ok := endpoint[key]
	if !ok {
		return false, EndpointMissingRequiredDataError{Key: key}
	}

	if reflect.TypeOf(endpoint[key]).Kind() != reflect.Bool {
		return false, EndpointDataTypeMismatchError{Key: key, DesiredType: "boolean"}
	}

	return endpoint[key].(bool), nil
}

// BooleanValueDefault ...
func (endpoint ShieldEndpoint) BooleanValueDefault(key string, def bool) (bool, error) {
	tf, err := endpoint.BooleanValue(key)
	if err == nil {
		return tf, nil
	}
	if _, ok := err.(EndpointMissingRequiredDataError); ok {
		return def, nil
	}
	return false, err
}

// ArrayValue ...
func (endpoint ShieldEndpoint) ArrayValue(key string) ([]interface{}, error) {
	_, ok := endpoint[key]
	if !ok {
		return nil, EndpointMissingRequiredDataError{Key: key}
	}

	if reflect.TypeOf(endpoint[key]).Kind() != reflect.Slice {
		return nil, EndpointDataTypeMismatchError{Key: key, DesiredType: "array"}
	}

	return endpoint[key].([]interface{}), nil
}

// ArrayValueDefault ...
func (endpoint ShieldEndpoint) ArrayValueDefault(key string, def []string) ([]string, error) {
	a, err := endpoint.ArrayValue(key)
	if err == nil {
		coerced := []string{}
		for _, elem := range a {
			coerced = append(coerced, elem.(string))
		}
		return coerced, nil
	}
	if _, ok := err.(EndpointMissingRequiredDataError); ok {
		return def, nil
	}
	return []string{}, err
}

// MapValue ...
func (endpoint ShieldEndpoint) MapValue(key string) (map[string]interface{}, error) {
	_, ok := endpoint[key]
	if !ok {
		return nil, EndpointMissingRequiredDataError{Key: key}
	}

	if reflect.TypeOf(endpoint[key]).Kind() != reflect.Map {
		return nil, EndpointDataTypeMismatchError{Key: key, DesiredType: "map"}
	}

	return endpoint[key].(map[string]interface{}), nil
}
