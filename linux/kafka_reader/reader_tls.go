package main

import (
    "fmt"
    "./helpers"
    "io/ioutil"
    "context"
    "crypto/tls"
    "crypto/x509"
    "flag"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"
    kafka "github.com/segmentio/kafka-go"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main () {
//  TLS Configuration
    var infile []byte
    var rootPEM string
    infile, _ = ioutil.ReadFile("certs/rootPEM")
    rootPEM = string(infile)
    fmt.Println(rootPEM)
    roots := x509.NewCertPool()

    ok := roots.AppendCertsFromPEM([]byte(rootPEM))
    if !ok {
        panic("failed to parse root certificate")
    }

//  Kafka Configuration
    var configFile string
    flag.StringVar(&configFile, "config", "config.toml", "configuration file")
    flag.Parse()
    conf := helpers.GetConfig(configFile)
    fmt.Printf("Topic:  %s\n", conf.KafkaConfig.TopicName)

    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

//  Setup the client
    bootstrapServers := strings.Split(conf.KafkaConfig.BootstrapServers, ",")
    topic := conf.KafkaConfig.TopicName

//  TLS configuration
    dialer := &kafka.Dialer{
        Timeout:   10 * time.Second,
        DualStack: true,
        TLS: &tls.Config{
                RootCAs: roots,
                InsecureSkipVerify : true},
        }

    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   bootstrapServers,
        Topic:     topic,
        Partition: 0,
        MinBytes:  10, // 10KB
        MaxBytes:  10, // 10MB
        Dialer: dialer,
    })
//     r.SetOffset(42)

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            break
        }
        fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    }

    r.Close()
}
