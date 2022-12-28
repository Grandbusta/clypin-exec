package main

import "clypin-exec/internal"

func init() {
	internal.InitializeClipboard()
}

func main() {
	internal.RunWatcher()
}
