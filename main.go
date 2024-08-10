package main

import (
	"fmt"
	"os"
	"rogchap.com/v8go"
)

var (
	isolate *v8go.Isolate
	context *v8go.Context
)

func init() {
	isolate = v8go.NewIsolate()
	context = v8go.NewContext(isolate)
	global := context.Global()
	console, _ := global.Get("console")
	consoleObj := console.Object()
	logFunc := v8go.NewFunctionTemplate(isolate, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		if len(info.Args()) > 0 {
			output := info.Args()[0].String()
			fmt.Println(output)
		}
		return nil
	})
	_ = consoleObj.Set("log", logFunc.GetFunction(context))
}

func executeJS(filePath string) {
	jsCode, _ := os.ReadFile(filePath)
	_, _ = context.RunScript(string(jsCode), "")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tjs <file>")
		return
	}
	filePath := os.Args[1]
	executeJS(filePath)
}
