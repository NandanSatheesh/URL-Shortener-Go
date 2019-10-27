# URL-Shortener-Go
URL shortener in Go using Redis and go-chi.    
[![forthebadge](https://forthebadge.com/images/badges/check-it-out.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/powered-by-responsibility.svg)](https://forthebadge.com)

## Redis 
Try Redis Enterprise free in the cloud [here](https://redislabs.com/). Get upto 30 MB of free storage on their enterprise cloud. :sunglasses:   

## Instructions 
After all the package setup done, run the `main.go` file.   
The app will run on `localhost:8080`   

### Endpoints 
``` POST - / ```   
Request - Make sure to include http:// or https://
``` 
{
    "url":"https://www.google.com"
}
```
Response -  
``` 
{
    "code": "Fek6nO0Zg",
    "url": "https://www.google.com",
    "created_at": 1572207289
}
```
``` GET - /{string} ``` - Will take you the webiste for which the URL is already shortened.   

## Other Notes 
The keys in this project are shown for reference only and are disposed with care :tada:   
There were a lot of projects running on my GoLand IDE. You may have to reconfigure few of the packages on your own. :confused:

## Contribute 

Want to contribute? Great!  
