# go-iis-auth

This repository contains a collection of middleware functions that let you add informations about a Windows User (identified by Windows Authentication) to a Go context, or obtain the port used by IIS if that is how you serve your application.

There are two main packages:

* `goav1`: contains middleware functions that are specific for [Goa v1](https://goa.design/v1)
* `http`: contains standard HTTP middleware functions, that can be used with [Goa v3](https://goa.design/).

**IMPORTANT NOTE**: these middleware functions will work as expected only if compiled for Windows. If you try to use them on a Unix environment, you will obtain mock data.

## How to use this package with Goa v1

1. Import the package by using:

    ```go
    import (
        admiddleware "github.com/top-solution/go-iis-auth/goav1"
    )
    ```

2. Mount the middleware as show in the following example:

    ```go
        service.Use(admiddleware.WithUser())
    ```

## How to use this package with Goa v3 or a standard HTTP handler

1. Import the package by using:

    ```go
        import (
            admiddleware "github.com/top-solution/go-iis-auth/http"
        )
    ```

2. Mount the middleware as shown in the following example:

    ```go
        var handler http.Handler = mux
        {
            handler = admiddleware.WithUser()(handler)
        }
    ```
