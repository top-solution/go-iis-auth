# go-iis-auth

This repository contains a collection of middleware functions that let you add informations about a Windows User (identified by Windows Authentication) to a Go context, or obtain the port used by IIS if that is how you serve your application.

There are two main packages:

* `goav1`: contains middleware functions that are specific for [Goa v1](https://goa.design/v1)
* `http`: contains standard HTTP middleware functions, that can be used with [Goa v3](https://goa.design/).

**IMPORTANT NOTE**: these middleware functions will work as expected only if compiled for Windows. If you try to use them on a Unix environment, you will obtain mock data.

## Prerequisites

1. Install [HttpPlatformHandler](https://www.iis.net/downloads/microsoft/httpplatformhandler) to use IIS as a reverse proxy

2. Configure your `web.config` so that Windows Auth tokens are forwarded to your API, for example:

    ```xml
    <?xml version="1.0" encoding="UTF-8"?>
    <configuration>
    <system.webServer>
        <handlers>
        <add name="httpPlatformHandler" path="*" verb="*"
            modules="httpPlatformHandler"
            resourceType="Unspecified"/>
        </handlers>
        <httpPlatform processPath="C:\path\to\your\binary.exe"
                    arguments="--arguments --for --your --binary"
                    startupRetryCount="3"
                    stdoutLogEnabled="true"
                    forwardWindowsAuthToken="true"/><!-- This is very important -->
    </system.webServer>
    </configuration>
    ```

3. Listen to the port reserved by IIS

    IIS will pass the expected port you should listen to via an environment variable. For convenience, you can use an helper function provided by this package:

    ```go
    import "github.com/top-solution/go-iis-auth/ad"
    ...
    portToListenTo := ad.GetIISPortWithFallback(9000) // this will return either 9000 or the actual port required by IIS
    ```

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