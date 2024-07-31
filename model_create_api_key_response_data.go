/*
Safepay API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the CreateAPIKeyResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAPIKeyResponseData{}

// CreateAPIKeyResponseData struct for CreateAPIKeyResponseData
type CreateAPIKeyResponseData struct {
	Token *string `json:"token,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	WebhookSecret *string `json:"webhook_secret,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewCreateAPIKeyResponseData instantiates a new CreateAPIKeyResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAPIKeyResponseData() *CreateAPIKeyResponseData {
	this := CreateAPIKeyResponseData{}
	return &this
}

// NewCreateAPIKeyResponseDataWithDefaults instantiates a new CreateAPIKeyResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAPIKeyResponseDataWithDefaults() *CreateAPIKeyResponseData {
	this := CreateAPIKeyResponseData{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateAPIKeyResponseData) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAPIKeyResponseData) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateAPIKeyResponseData) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateAPIKeyResponseData) SetToken(v string) {
	o.Token = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreateAPIKeyResponseData) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAPIKeyResponseData) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreateAPIKeyResponseData) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreateAPIKeyResponseData) SetClientId(v string) {
	o.ClientId = &v
}

// GetWebhookSecret returns the WebhookSecret field value if set, zero value otherwise.
func (o *CreateAPIKeyResponseData) GetWebhookSecret() string {
	if o == nil || IsNil(o.WebhookSecret) {
		var ret string
		return ret
	}
	return *o.WebhookSecret
}

// GetWebhookSecretOk returns a tuple with the WebhookSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAPIKeyResponseData) GetWebhookSecretOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookSecret) {
		return nil, false
	}
	return o.WebhookSecret, true
}

// HasWebhookSecret returns a boolean if a field has been set.
func (o *CreateAPIKeyResponseData) HasWebhookSecret() bool {
	if o != nil && !IsNil(o.WebhookSecret) {
		return true
	}

	return false
}

// SetWebhookSecret gets a reference to the given string and assigns it to the WebhookSecret field.
func (o *CreateAPIKeyResponseData) SetWebhookSecret(v string) {
	o.WebhookSecret = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateAPIKeyResponseData) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAPIKeyResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateAPIKeyResponseData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateAPIKeyResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateAPIKeyResponseData) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAPIKeyResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateAPIKeyResponseData) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CreateAPIKeyResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CreateAPIKeyResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAPIKeyResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.WebhookSecret) {
		toSerialize["webhook_secret"] = o.WebhookSecret
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableCreateAPIKeyResponseData struct {
	value *CreateAPIKeyResponseData
	isSet bool
}

func (v NullableCreateAPIKeyResponseData) Get() *CreateAPIKeyResponseData {
	return v.value
}

func (v *NullableCreateAPIKeyResponseData) Set(val *CreateAPIKeyResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAPIKeyResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAPIKeyResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAPIKeyResponseData(val *CreateAPIKeyResponseData) *NullableCreateAPIKeyResponseData {
	return &NullableCreateAPIKeyResponseData{value: val, isSet: true}
}

func (v NullableCreateAPIKeyResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAPIKeyResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


