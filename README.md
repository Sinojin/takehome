#Take Home Test
This project created for interview.Basically, It makes http requests and shows link and md5 of response.Project can work with synchronous and asynchronous.

## Installation
````
git clone https://github.com/sinojin/takeHome
cd takeHome
make build
cd bin/
````
#####OR

````
go get github.com/sinojin/takeHome
cd $GOPATH/src/github.com/sinojin/takehome
make build
cd bin/
````

##Usage 
Basic usage is below

`./takeHome  facebook.com yahoo.com yandex.com twitter.com`

and your response will be like this:
````
http://yandex.com - 7e7b113147febb92ecb924960839ebf8
http://facebook.com - 864d21343ba2900b7ccc35d4297d34f4
http://twitter.com - b8ed215e5edd46ff2cd33125d146bcfa
http://yahoo.com - f5b071e16664b7bb22ab09ebb61a6395
````

To use with one worker you need to run this command  
`./takeHome -parallel 1 facebook.com yahoo.com yandex.com twitter.com`
Your response will be like this: 
````
http://facebook.com - 71a8050ea2f9ce33e16586751be0abcd
http://yahoo.com - 315622cdca3b3a11541fc8c3402527af
http://yandex.com - 9443e4ed3224159da05964403119d6c2
http://twitter.com - 8c2cbf462ed85bdda33d24536275044f
````
#####Note:
If you want to use -parallel flag, you have to put after binary.Otherwise it doesn't take as parameter. It is a bug I didn't fix yet.

##Test
`go test ./...`