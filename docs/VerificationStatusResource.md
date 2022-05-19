# VerificationStatusResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User&#39;s identifier | 
**IsUserVerified** | **bool** | User&#39;s verification status | 

## Methods

### NewVerificationStatusResource

`func NewVerificationStatusResource(userId string, isUserVerified bool, ) *VerificationStatusResource`

NewVerificationStatusResource instantiates a new VerificationStatusResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationStatusResourceWithDefaults

`func NewVerificationStatusResourceWithDefaults() *VerificationStatusResource`

NewVerificationStatusResourceWithDefaults instantiates a new VerificationStatusResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *VerificationStatusResource) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VerificationStatusResource) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VerificationStatusResource) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsUserVerified

`func (o *VerificationStatusResource) GetIsUserVerified() bool`

GetIsUserVerified returns the IsUserVerified field if non-nil, zero value otherwise.

### GetIsUserVerifiedOk

`func (o *VerificationStatusResource) GetIsUserVerifiedOk() (*bool, bool)`

GetIsUserVerifiedOk returns a tuple with the IsUserVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserVerified

`func (o *VerificationStatusResource) SetIsUserVerified(v bool)`

SetIsUserVerified sets IsUserVerified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


