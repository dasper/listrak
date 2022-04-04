# Listrak REST API Integration
 A Go library to make integration with Listrak's REST API easier to implement and maintain.

## Features

### No Third-Party Dependencies
Everything uses Go's standard library

### Go 1.11+
Each major Go release is supported until there are two newer major releases. It is strongly advised to stay within the security window of updates for Go if possible.
However, this is not always possible in all projects. For that reason, there is a low barrier of entry for this to be included in existing projects.


### Defining Expected Errors
The error json responses have been saved to variables (Sentinel errors).
As many of the errors as exclusive to particular status codes this can infer response status codes as well as handle errors.

## Implementation Status/Progress
### SMS REST API
All API endpoints are implemented

### Email REST API
Transactional emails only

### Cross Channel REST API
Not Implemented

### Data Import REST API
Not Implemented

## Use

`go get github.com/dasper/listrak`

Once included, how you store and pass your Listrak's id and secret is up to you.
Only one integration is allowed to be initialized at a time.
This limitation is in place to ensure the integrity of the Oauth2 bearer token as well as provide passive refreshing of the token.

### Example 
```go
err := listrak.Initialize("id", "secret")
if err != nil {
    return
}
smsRequest := sms.NewClient()
shortCodes, err := smsRequest.GetShortCodeCollection()
if err != nil {
    return
}
for _, item := range shortCodes.Data {
    phoneList, err := smsRequest.GetListCollection(item.ShortCodeID)
    if err != nil {
        return
    }
    for _, phone := range phoneList.Data {
        fmt.Println("Name", phone.PhoneListName)
        fmt.Println("Created At", phone.CreateDate)
    }
}
```

## Todo

### Finish all APIs
### Create tests
### Add Validation for properties in POST/PUT requests
### Upping minimum version to 1.13
Wrapping errors and utilizing `errors.Is()` could be beneficial as REST APIs have many potential points of failures and wrapped errors could allow attempting graceful recoveries in certain situations while still returning all relivant messages.
