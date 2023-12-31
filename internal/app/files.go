package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type URLEvent struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type Producer struct {
	file   *os.File
	writer *bufio.Writer
}

type Consumer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewProducer(fileName string) (*Producer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &Producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (p *Producer) CloseFile() error {
	return p.file.Close()
}

func (p *Producer) WriteEvent(event *URLEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	if _, err := p.writer.Write(data); err != nil {
		return err
	}

	if err := p.writer.WriteByte('\n'); err != nil {
		return err
	}

	return p.writer.Flush()
}

func NewConsumer(filename string) (*Consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

func (c *Consumer) ReadEvent() (*URLEvent, error) {
	if !c.scanner.Scan() {
		return nil, c.scanner.Err()
	}
	data := c.scanner.Bytes()

	event := URLEvent{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (c *Consumer) Close() error {
	return c.file.Close()
}
