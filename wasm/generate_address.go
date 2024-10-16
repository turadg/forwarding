//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"

	"github.com/cosmos/cosmos-sdk/types/address"
)

const ModuleName = "forwarding"

func generateAddress(_ js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		err := "Error: Two arguments (channel and recipient) are required"
		js.Global().Get("console").Call("error", err)
		return err
	}

	channel := args[0].String()
	recipient := args[1].String()

	js.Global().Get("console").Call("log", "Channel:", channel)
	js.Global().Get("console").Call("log", "Recipient:", recipient)

	bz := []byte(channel + recipient)
	addrBytes := address.Derive([]byte(ModuleName), bz)[12:]
	
	// Convert bytes to hex string
	addrHex := make([]byte, len(addrBytes)*2)
	for i, b := range addrBytes {
		addrHex[i*2] = "0123456789abcdef"[b>>4]
		addrHex[i*2+1] = "0123456789abcdef"[b&15]
	}
	addrString := string(addrHex)

	js.Global().Get("console").Call("log", "Generated Address:", addrString)

	return addrString
}

func main() {
	c := make(chan struct{})
	js.Global().Set("generateAddress", js.FuncOf(generateAddress))
	<-c
}
