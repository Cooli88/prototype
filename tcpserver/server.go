package tcpserver

import (
    "device-listener-go-new-device/dto"
    "device-listener-go-new-device/parse"
    "device-listener-go-new-device/rabbit"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
    "go.uber.org/zap"
)

const (
	CONN_HOST = "0.0.0.0"
    CONN_PORT = "7066"
	CONN_TYPE = "tcp"
)

func ServerStart() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	countConnection := 0
	for {
		countConnection++
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%d: %v <-> %v\n", countConnection, conn.LocalAddr(), conn.RemoteAddr())
		// Handle connections in a new goroutine.
		go func(conn net.Conn) {
            handleRequestImei(conn)
            //handleRequestDataPayload(conn, deviceImei)
			conn.Close()
		}(conn)
	}
}

// Handles incoming requests.
func handleRequestImei(conn net.Conn) string {
    logger, _ := zap.NewProduction()
    sugar := logger.Sugar()
    sugar.Infow("App start.", )

    var deviceImei string = "123"
	buf := make([]byte, 4096)
	countBytes, err := conn.Read(buf)

	if err == io.EOF {
		conn.Close()
	}

	message := buf[:countBytes]

    sugar.Infow("buf.", buf)
    sugar.Infow("message.", message)
    sugar.Infow("countBytes.", countBytes)
    fmt.Printf("err answer: %+v\n", buf)
    fmt.Printf("err answer: %+v\n", message)
    fmt.Printf("err answer: %+v\n", countBytes)

	if countBytes == 17 {
		deviceImei = string(message)
		count, err := conn.Write([]byte{01})
		fmt.Printf("err answer: %+v\n", err)
		fmt.Printf("err answer: %+v\n", count)
	}

	return deviceImei
}

// Handles incoming requests.
func handleRequestDataPayload(conn net.Conn, deviceImei string) {
    buf := make([]byte, 1800)
	countBytes, err := conn.Read(buf)

	if err == io.EOF {
		conn.Close()
	}

	message := buf[:countBytes]

	allRecords := parse.CreateData(message)
	fmt.Printf("deviceImei: %+v\n", deviceImei)
	answerBuf := make([]byte, 4)

	if allRecords.Error == nil {
		dataToRabbitArr := dto.CreateDataToRabbitFromRecords(allRecords, deviceImei)
		b, _ := json.Marshal(dataToRabbitArr)
		rabbit.SendMessage(b)
		binary.BigEndian.PutUint32(answerBuf, uint32(allRecords.NumberOfData))
	}
	conn.Write(answerBuf)
}
