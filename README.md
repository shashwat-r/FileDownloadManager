# FileDownloadManager

Files
- models
    - requestPayload.go
        - for Download APIs
    - responsePayload.go
        - for Status API
    - [files.go]
- controller
    - handler.go
    - controls.go
- download
    - utils.go

Requirements

- FR1
    - Check Health API
        - returns status OK message
- FR2
    - Serial Download API
    - Concurrent Download API
- FR3
    - Status API
- FR4
    - Error API
- FR5
    - Browse Files API

Flow

- Each download request has a list of urls
- Each download request gets assigned a unique id
- Each download request keeps updating its status
- There should be a folder for all downloaded files
    - resources/tmp
- There should be a folder for all download statuses
    - resources/info
