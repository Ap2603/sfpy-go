# CreateGuestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CreateJWTResponseData**](CreateJWTResponseData.md) |  | [optional] 
**Status** | Pointer to [**CreateJWTResponseStatus**](CreateJWTResponseStatus.md) |  | [optional] 

## Methods

### NewCreateGuestResponse

`func NewCreateGuestResponse() *CreateGuestResponse`

NewCreateGuestResponse instantiates a new CreateGuestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuestResponseWithDefaults

`func NewCreateGuestResponseWithDefaults() *CreateGuestResponse`

NewCreateGuestResponseWithDefaults instantiates a new CreateGuestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateGuestResponse) GetData() CreateJWTResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateGuestResponse) GetDataOk() (*CreateJWTResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateGuestResponse) SetData(v CreateJWTResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateGuestResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *CreateGuestResponse) GetStatus() CreateJWTResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateGuestResponse) GetStatusOk() (*CreateJWTResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateGuestResponse) SetStatus(v CreateJWTResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateGuestResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


