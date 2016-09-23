# go-jsonapi-exercise
Implementation of JSONAPI protocol in Golang

### Usage:

`./go-jsonapi-exercise -p :8081`

where `-p` is `port` option. You can omit it, default value is :8080.

You can test this app using CURL:

`$ curl -i -H "Accept: application/vnd.api+json" \
-H 'Content-Type:application/vnd.api+json' http://localhost:<port>/tags`

Application endpoints are:

`/posts`

`/comments`

`/tags`

They represent following BD structure:

![screenshot.jpg](https://raw.githubusercontent.com/thefivekey/go-jsonapi-exercise/master/screenshot.jpg)
