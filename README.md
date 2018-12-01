# cid

[![Circle CI](https://circleci.com/gh/s32x/cid/tree/master.svg?style=svg)](https://circleci.com/gh/s32x/cid/tree/master)

cid (short for custom-import-domains) is a go package that can be used to host an import proxy service pointing to your host go repositories.

## Usage

```go
package main

import "s32x.com/cid"

func main() {
	cid.Start(
		// The page to redirect to when the index is requested
		"https://swolfe.me",
		// Where to retrieve the requested repository from
		"github.com/s32x",
		// The domain this service will be being hosted on
		"s32x.com",
		// The port this service will be hosted on
		"8080",
	)
}
```

The MIT License (MIT)
=====================

Copyright © 2018 Steven Wolfe

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation
files (the “Software”), to deal in the Software without
restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following
conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.