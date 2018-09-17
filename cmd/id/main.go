// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file

package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
)

func main() {
	// read csv file
	file, err := os.Open("NodeIds.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create temporary buffer
	var b bytes.Buffer

	b.WriteString("// Code generated by cmd/id; DO NOT EDIT\n\n")
	b.WriteString("package id\n// NodeId definitions, generated automatically by cmd/id.\n const(")

	// loop over each row
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		b.WriteString(fmt.Sprintf("%s = %s\n", record[0], record[1]))
	}

	// close const(...) bracket
	b.Write([]byte(")"))

	// format file
	fmt, err := format.Source(b.Bytes())
	if err != nil {
		panic(err)
	}

	// write formatted code to file
	out, err := os.Create("../../id/id.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	out.Write(fmt)

	log.Println("done")
}