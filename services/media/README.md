# MediaService

This service has two main functions

1. User can upload temporary images, temp images will be cleaned up after some duration.
2. Other services confirm the usage of the temporary images and make them temporary.

Upload images with port 8080 using multipart form file with key "image"  
Confirm usages with port 8081 using gprc.
