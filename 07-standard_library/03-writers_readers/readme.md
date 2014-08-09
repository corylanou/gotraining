## Writers/Readers - Standard Library

The ability to stream and pass data around is incredibility important. Data is constantly coming at our programs whether over a socket, file, device, etc. Many times this data needs to just be moved from one stream. Sometimes it needs to be encrypted, hashed or stored for safe keeping. The Writer and Reader interfaces may be the most heavily used and supported interfaces in both the standard library and the community.

## Notes

* The standard library provide all the infrasture we need to streaming and working with data.
* If we implement the Reader and Writer interfaces in our own types, we get this functionality for free.
* Implementing interfaces to existing functionality saves us from both development time and bugs.

## Code Review

[io.Writer interface](example1/example1.go) ([Go Playground](http://play.golang.org/p/6pZ8RYzIN5))

[Simple curl with io.Reader and io.Writer](example2/example2.go)

[MultiWriters with curl example](example3/example3.go)

## Advanced Code Review

[TeeReader and io composition](advanced/example1/example1.go) ([Go Playground](http://play.golang.org/p/rLDpqYbnGR))

[Gzip and Md5 support with curl example](advanced/example2/example2.go) ([Go Playground](http://play.golang.org/p/rLDpqYbnGR))

## Exercises

### Exercise 1

Download any document from the web and display the content in the terminal and write it to a file at the same time.

___
[![GoingGo Training](../../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../../00-slides/images/ggb_logo.png)](http://www.goinggo.net)