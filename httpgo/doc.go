/*
HTTPgo is Go clone of HTTPie, which is a CLI, cURL like tool for human written in Python.
Like as HTTPie, HTTPgo is a command line HTTP client. Its goal is same as HTTPie, that is,
to make CLI interation with web services as human friendly as possible, since this tool is
aiming at complete clone of HTTPie as of now, but further more, with writing in Go.
cross platform eiligibility and binary distribution will be acheived.

It provides a simple "httpgo" command that is one binary and allows for sending arbitraty
HTTP requests using a simple and natural syntax, and displays colorized responses, which is
almost same explanation as HTTPie's introduction.

HTTPgo can be used for testing, debugging, and generally interacting with HTTP servers.

Main Features

HTTPgo is trying to implement following feature implemented in HTTPie:

    * Expressive and intuitive syntax
    * Formatted and colorized terminal output
    * Built-in JSON support
    * Forms and file uploads
    * HTTPS, proxies, and authentication
    * Arbitrary request data
    * Custom headers
    * Persistent sessions
    * Wget-like downloads
    * Linux, Mac OS X and Windows support
    * ARM support
    * Documentation
    * Test coverage

Installation

Just copy "httpgo" binary into the directory set in $PATH.

**** From now on, all following document are taken from HTTPie original one. ****

Usage

Hello World:

	$ httpgo www.google.com

Synopsis:

	$ httpgo [flags] [METHOD] URL [ITEM [ITEM]]

See also "http --help".

Examples

[UNDER IMPLEMENTATION] Custom HTTP method, HTTP headers and JSON data:

    $ httpgo PUT example.org X-API-Token:123 name=John

[NOT IMPLEMENTED] Submitting forms:

    $ httpgo -f POST example.org hello=World

[NOT IMPLEMENTED] See the request that is being sent using one of the output options:

    $ httpgo -v example.org

[NOT IMPLEMENTED] Use Github API to post a comment on an issue with authentication:

    $ httpgo -a USERNAME POST https://api.github.com/repos/jkbr/httpie/issues/83/comments body='HTTPie is awesome!'

[NOT IMPLEMENTED] Upload a file using redirected input:

    $ httpgo example.org < file.json

[NOT IMPLEMENTED] Download a file and save it via redirected output:

    $ httpgo example.org/file > file

[NOT IMPLEMENTED] Download a file wget style:

    $ httpgo --download example.org/file

[NOT IMPLEMENTED] Use named sessions to make certain aspects or the communication persistent between requests to the same host:

    $ httpgo --session=logged-in -a username:password httpbin.org/get API-Key:123
    $ httpgo --session=logged-in httpbin.org/headers

[NOT IMPLEMENTED] Set a custom Host header to work around missing DNS records:

    $ httpgo localhost:8000 Host:example.com

HTTP Method

The name of the HTTP method comes right before the URL argument:

    $ httpgo DELETE example.org/todos/7

Which looks similar to the actual Request-Line that is sent:

    DELETE /todos/7 HTTP/1.1

When the METHOD argument is omitted from the command, HTTPie defaults to either GET (with no request data)
or POST (with request data).

Request URL

The only information HTTPie needs to perform a request is a URL.
The default scheme is, somewhat unsurprisingly, http://, and can be omitted from the argument
– http example.org works just fine.

If find yourself manually constructing URLs with querystring parameters on the terminal,
you may appreciate the param==value syntax for appending URL parameters so that you don't have to
worry about escaping the & separators.
To search for HTTPie on Google Images you could use this command:

    $ http GET www.google.com search==HTTPie tbm==isch

    GET /?search=HTTPie&tbm=isch HTTP/1.1

*/
package httpgo
