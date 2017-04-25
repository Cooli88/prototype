package tcpserver

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"telparse/dto"
	"telparse/parse"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "7610"
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
			deviceImei := handleRequestImei(conn)
			handleRequestDataPayload(conn, deviceImei)
			conn.Close()
		}(conn)
	}
}

// Handles incoming requests.
func handleRequestImei(conn net.Conn) string {
	var deviceImei string
	buf := make([]byte, 4096)
	countBytes, err := conn.Read(buf)

	if err == io.EOF {
		conn.Close()
	}

	message := buf[:countBytes]

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
	buf := make([]byte, 4096)
	countBytes, err := conn.Read(buf)

	if err == io.EOF {
		conn.Close()
	}

	message := buf[:countBytes]

	allRecords := parse.CreateData(message)
	fmt.Printf("deviceImei: %+v\n", deviceImei)
	answerBuf := make([]byte, 4)

	if allRecords.Error == nil {
		for _, record := range allRecords.Records {
			dataToRabbit := dto.CreateDataToRabbitFromRecord(deviceImei, record)
			fmt.Printf("dataToRabbit: %+v\n", dataToRabbit)
		}
		binary.BigEndian.PutUint32(answerBuf, uint32(allRecords.NumberOfData))
	}
	conn.Write(answerBuf)
}
