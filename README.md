python -m SimpleHTTPServer
go run getandcompare.go

Verify that the go compare endpoint works:

curl 'http://localhost:8000/random1.bin' | curl --data-binary @- 'http://localhost:9999/compare?f=random1.bin'
Matches!

curl 'http://localhost:8000/random2.bin' | curl --data-binary @- 'http://localhost:9999/compare?f=random1.bin'
Uploaded data doesn't match

Verify that the go binary GET works
curl 'http://localhost:9999/binary?f=random1.bin' | curl --data-binary @- 'http://localhost:9999/compare?f=random1.bin'
Matches!

k6 run binarybug.js
