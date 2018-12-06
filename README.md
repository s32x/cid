# cid

[![Circle CI](https://circleci.com/gh/s32x/cid/tree/master.svg?style=svg)](https://circleci.com/gh/s32x/cid/tree/master)
[![GoDoc](https://godoc.org/github.com/s32x/cid/cid?status.svg)](https://godoc.org/github.com/s32x/cid/cid)

cid (short for custom-import-domains) is a go package that can be used to host an import proxy service pointing to your hosted go repositories.

## Usage

You have two options for using the API: 
* Import the child dependency using the below go example.
* Run the docker image `s32x/cid`.

### Go example

```go
package main

import "s32x.com/cid/cid"

func main() {
	cid.Start(
		// Where to retrieve the requested repository from
		"https://github.com/s32x",
		// The domain this service will be being hosted on
		"s32x.com",
		// The port this service will be hosted on
		"8080",
	)
}
```

### Running with Docker
```
docker run -p 8080:8080 -e USER_URL=https://github.com/s32x -e DOMAIN=s32x.com -e PORT=8080 s23x/cid
```

### Installing
```
go get s32x.com/cid
cid
```

The BSD 3-clause License
========================

Copyright (c) 2018, Steven Wolfe. All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

 - Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

 - Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

 - Neither the name of cid nor the names of its contributors may
   be used to endorse or promote products derived from this software without
   specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.