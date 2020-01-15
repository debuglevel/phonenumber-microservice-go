Minimal clone of https://github.com/debuglevel/phonenumber-microservice in Golang.

```
curl -X POST -d '{"Phonenumber":"0921123456789"}' -H "Content-Type: application/json" -H "Accept: application/json" http://localhost:80/format/
```

Does not really handle any bad input, is not well tested et ceteraL; but only 10MB in size.