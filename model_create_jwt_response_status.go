/*
Safepay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateJWTResponseStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateJWTResponseStatus{}

// CreateJWTResponseStatus struct for CreateJWTResponseStatus
type CreateJWTResponseStatus struct {
	Errors []string `json:"errors,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewCreateJWTResponseStatus instantiates a new CreateJWTResponseStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateJWTResponseStatus() *CreateJWTResponseStatus {
	this := CreateJWTResponseStatus{}
	return &this
}

// NewCreateJWTResponseStatusWithDefaults instantiates a new CreateJWTResponseStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateJWTResponseStatusWithDefaults() *CreateJWTResponseStatus {
	this := CreateJWTResponseStatus{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateJWTResponseStatus) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJWTResponseStatus) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateJWTResponseStatus) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *CreateJWTResponseStatus) SetErrors(v []string) {
	o.Errors = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateJWTResponseStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJWTResponseStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateJWTResponseStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateJWTResponseStatus) SetMessage(v string) {
	o.Message = &v
}

func (o CreateJWTResponseStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateJWTResponseStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreateJWTResponseStatus struct {
	value *CreateJWTResponseStatus
	isSet bool
}

func (v NullableCreateJWTResponseStatus) Get() *CreateJWTResponseStatus {
	return v.value
}

func (v *NullableCreateJWTResponseStatus) Set(val *CreateJWTResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateJWTResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateJWTResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateJWTResponseStatus(val *CreateJWTResponseStatus) *NullableCreateJWTResponseStatus {
	return &NullableCreateJWTResponseStatus{value: val, isSet: true}
}

func (v NullableCreateJWTResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateJWTResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


