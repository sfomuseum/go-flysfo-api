# go-flysfo-api

Go package for accessing the developers.flysfo.com API.

## Documentation

Documentation is incomplete at this time.

## Design

The core of this package's approach to the `ExecuteMethod` method (which is attached to the `api.Client` struct) whose signature looks like this:

```
ExecuteMethod(context.Context, string, *url.Values) (io.ReadSeekCloser, error)
```

This package only defines [a handful of Go types or structs mapping to individual API responses](response). So far these are all specific to operations that have been SFO Museum in nature.

In time there may be others, along with helper methods for unmarshaling API responses in to typed responses but the baseline for all operations will remain: Query paramters (`url.Values`) sent over HTTP returning an `io.ReadSeekCloser` instance that is inspected and validated according to the needs and uses of the tools using the FlySFO API.

## Tools

### api

```
$> ./bin/api -h
Execute a method against the developers.flysfo.com API.
Usage:
	 ./bin/api [options]
Valid options are:
  -apikey string
    	A valid developers.flysfo.com API key encoded as a gocloud.dev/runtimevar URI. Supported schemes are: constant://, file://
  -method string
    	The relative URL of the method to be executed.
  -param value
    	Zero or more query parameters to append to the method being executed.
```

For example:

```
$>./bin/api -apikey 'constant://?val={APIKEY}' -method /offerings/facilities/terminals/ | json_pp | less

[
   {
      "airport" : {
         "airport_code" : "SFO",
         "airport_id" : 1876,
         "airport_location_id" : 1876,
         "airport_name" : "San Francisco International Airport",
         "customs_type" : 1,
         "iata_code" : "SFO",
         "icao_code" : "KSFO",
         "in_service" : false,
         "source" : 4,
         "utc_offset" : "-7.00"
      },
      "concourse" : [
         {
            "concourse_id" : 2,
            "concourse_location" : {
      ... and so on
   }
]
```

## See also

* https://developers.flysfo.com/