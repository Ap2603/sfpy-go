/*
Safepay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateJWTRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateJWTRequest{}

// CreateJWTRequest struct for CreateJWTRequest
type CreateJWTRequest struct {
	Client string `json:"client"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type _CreateJWTRequest CreateJWTRequest

// NewCreateJWTRequest instantiates a new CreateJWTRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateJWTRequest(client string, email string, password string) *CreateJWTRequest {
	this := CreateJWTRequest{}
	this.Client = client
	this.Email = email
	this.Password = password
	return &this
}

// NewCreateJWTRequestWithDefaults instantiates a new CreateJWTRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateJWTRequestWithDefaults() *CreateJWTRequest {
	this := CreateJWTRequest{}
	return &this
}

// GetClient returns the Client field value
func (o *CreateJWTRequest) GetClient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *CreateJWTRequest) GetClientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value
func (o *CreateJWTRequest) SetClient(v string) {
	o.Client = v
}

// GetEmail returns the Email field value
func (o *CreateJWTRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateJWTRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateJWTRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *CreateJWTRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateJWTRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateJWTRequest) SetPassword(v string) {
	o.Password = v
}

func (o CreateJWTRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateJWTRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client"] = o.Client
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *CreateJWTRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client",
		"email",
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateJWTRequest := _CreateJWTRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateJWTRequest)

	if err != nil {
		return err
	}

	*o = CreateJWTRequest(varCreateJWTRequest)

	return err
}

type NullableCreateJWTRequest struct {
	value *CreateJWTRequest
	isSet bool
}

func (v NullableCreateJWTRequest) Get() *CreateJWTRequest {
	return v.value
}

func (v *NullableCreateJWTRequest) Set(val *CreateJWTRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateJWTRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateJWTRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateJWTRequest(val *CreateJWTRequest) *NullableCreateJWTRequest {
	return &NullableCreateJWTRequest{value: val, isSet: true}
}

func (v NullableCreateJWTRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateJWTRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


