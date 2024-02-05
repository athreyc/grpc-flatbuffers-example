package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/athreyc/grpc-flatbuffers-example/api/models"
	"google.golang.org/grpc"
)

var (
	addr = "3000"
	name = flag.String("name", "Flatbuffers", "name to be sent go server :D")
	name2 = flag.String("name2", "test", "name to be sent go server :D")
)

func printSayHello(client models.GreeterClient, name string, name2 string) {

    c := flatbuffers.NewBuilder(1024)
    ni := c.CreateString("nestedname2")
    nj := c.CreateString("nestedmisc2")
    nk := c.CreateString("nestedtemp2")

    models.HelloRequestStartInnerVector(c, 0)
    ninners := c.EndVector(0)

    models.HelloRequestStart(c)
    models.HelloRequestAddName(c, ni)
    models.HelloRequestAddMisc(c, nj)
    models.HelloRequestAddTemp(c, nk)
    models.HelloRequestAddNumbr(c, 4)
    models.HelloRequestAddInner(c, ninners)
    nestedBuf := models.HelloRequestEnd(c)
    fmt.Printf("offset %v\n",  nestedBuf)
    c.Finish(nestedBuf)

    nbuf := c.FinishedBytes()
    b := flatbuffers.NewBuilder(1024)
    
    models.HelloRequestStartInnerVector(b, len(nbuf))
    fmt.Printf("offset %v\n",  nbuf)
    for i := len(nbuf)-1; i>=0; i-- {
         b.PrependByte(nbuf[i])
    }
    inners := b.EndVector(len(nbuf))

    i := b.CreateString("name2")
    j := b.CreateString("misc2")
    k := b.CreateString("temp2")
    models.HelloRequestStart(b)
    models.HelloRequestAddName(b, i)
    models.HelloRequestAddMisc(b, j)
    models.HelloRequestAddTemp(b, k)
    models.HelloRequestAddNumbr(b, 2)

    models.HelloRequestAddInner(b, inners)

    b.Finish(models.HelloRequestEnd(b))

    buf := b.FinishedBytes()
    fmt.Printf("size: %v\n", len(buf))

    replay := models.GetRootAsHelloRequest(buf, 0)

    v := replay.Name()
    fmt.Printf("replay name: %s\n", string(v))

    innerreplay := models.GetRootAsHelloRequest(replay.InnerBytes(), 0)

    vn := innerreplay.Name()
    fmt.Printf("replay name: %s\n", string(vn))


    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    request, err := client.SayHello(ctx, b, grpc.CallContentSubtype("flatbuffers"))
    if err != nil {
        log.Fatalf("%v.SayHello(_) = _, %v: ", client, err)
    }
    log.Printf("server said %q", request.Message())

/*
	log.Printf("Name to be sent (%s)", name)
	b := flatbuffers.NewBuilder(0)
	i := b.CreateString(name)
	models.HelloRequestStart(b)
	models.HelloRequestAddName(b, i)
	b.Finish(models.HelloRequestEnd(b))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request, err := client.SayHello(ctx, b, grpc.CallContentSubtype("flatbuffers"))
	if err != nil {
		log.Fatalf("%v.SayHello(_) = _, %v: ", client, err)
	}
	log.Printf("server said %q", request.Message())
*/
}

func printSayManyHello(client models.GreeterClient, name string, num int32) {
	log.Printf("Name to be sent (%s), num to be sent (%d)", name, num)
	b := flatbuffers.NewBuilder(0)
	i := b.CreateString(name)
	models.ManyHellosRequestStart(b)
	models.ManyHellosRequestAddName(b, i)
	models.ManyHellosRequestAddNumGreetings(b, num)
	b.Finish(models.ManyHellosRequestEnd(b))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.SayManyHellos(ctx, b, grpc.CallContentSubtype("flatbuffers"))
	if err != nil {
		log.Fatalf("%v.SayManyHellos(_) = _, %v", client, err)
	}
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.SayManyHellos(_) = _, %v", client, err)
		}
		log.Printf("server said %q", request.Message())
	}
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", addr), grpc.WithInsecure(), grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	client := models.NewGreeterClient(conn)
	printSayHello(client, *name, *name2)

	num := rand.Int31()
	printSayManyHello(client, *name, num)
}
