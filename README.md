Indian-Mobile-Circle-Go
=======================

Small Golang library to display the mobile operator and circle for Indian numbers.

This can work for more cases except a few:
* The mobile number owner has changed to a different operator using Mobile Number Portability (MNP)
* The mobile number owner has shifted to another state without changing the number.

Based on [this](https://github.com/rahulchordiya/Indian-mobile-circle) library.

How to use this as a library?
-----------------------------

* Get the source code of the library.

```go 
go get github.com/R4CHI7/Indian-Mobile-Circle-Go
```

* Import the library.
```go 
import "github.com/R4CHI7/Indian-Mobile-Circle-Go"
```

* Call the function.
```go
imcg.GetMobileNumberData("9711XXXXXX")
```
