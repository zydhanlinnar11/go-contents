package embedfile

import _ "embed"

//go:embed file-to-import.txt
var File string
