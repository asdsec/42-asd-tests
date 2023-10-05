package main

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// HiOp describes the operation to print a friendly message to the user.
type HiOp struct{}

func (h HiOp) Run(stdout, _ io.Writer) error {
	return printHi(stdout)
}

// Run prints a friendly message to the user.
func printHi(out io.Writer) error {
	hi := `
        :::       ::: ::::::::::            :::     :::::::::  ::::::::::           ::: :::::::::::           :::     ::::::::    
       :+:       :+: :+:                 :+: :+:   :+:    :+: :+:                :+: :+:   :+:              :+:     :+:    :+:    
      +:+       +:+ +:+                +:+   +:+  +:+    +:+ +:+               +:+   +:+  +:+             +:+ +:+        +:+      
     +#+  +:+  +#+ +#++:++#          +#++:++#++: +#++:++#:  +#++:++#         +#++:++#++: +#+            +#+  +:+      +#+         
    +#+ +#+#+ +#+ +#+               +#+     +#+ +#+    +#+ +#+              +#+     +#+ +#+           +#+#+#+#+#+  +#+            
    #+#+# #+#+#  #+#               #+#     #+# #+#    #+# #+#              #+#     #+# #+#                 #+#   #+#              
    ###   ###   ##########        ###     ### ###    ### ##########       ###     ### ###                 ###  ##########         
	`

	_, err := fmt.Fprintf(out, "%s\n", hi)
	if err == nil {
		err = printUsage(out)
	}
	return errors.Wrap(err, "write error")
}
