// Go Speak japanese
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

func savePollyVoice(text string, fileName string) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		fmt.Println(err)
	}
	svc := polly.New(sess)

	txt := "<speak>"+text+".<break time=\"1s\"/></speak>"

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"), // Specify the desired output format
		Text:         aws.String(txt),  // Set the SSML text
		TextType:     aws.String("ssml"), // Set the text type as SSML
		VoiceId:      aws.String("Takumi"), // Replace with your desired voice
	}

	output, err := svc.SynthesizeSpeech(input)

	// Save as MP3
	names := strings.Split(fileName, ".")
	name := names[0]
	mp3File := name + ".mp3"

	outFile, err := os.Create(mp3File)
	if err != nil {
		fmt.Println("Got error creating " + mp3File + ":")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer outFile.Close()
	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		fmt.Println("Got error saving MP3:")
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println("✅ -->" + name + ".mp3")
}

var say string

func main() {
	flag.StringVar(&say, "say", say, "enter a sentence to speak")
	flag.Parse()
	fmt.Println("converting japanese text into audio for: [ " + say + "] ")

	//create voice folder
	os.Mkdir("audio", os.ModePerm)
	//create audio file
	savePollyVoice(say, "audio/"+say)

}
