// Go Speak
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

func savePollyVoice(text string, folderName string, fileName string, voiceId string) {
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
		VoiceId:      aws.String(voiceId), // Replace with your desired voice
	}

	output, err := svc.SynthesizeSpeech(input)

	// Save as MP3
	mp3File := "tmp/" + folderName + "/" + fileName + ".mp3"

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
	fmt.Println("✅ --> temp/" + folderName + "/" + fileName + ".mp3")
}

var say string
var voiceId string
var audioFolderName string

func main() {

	flag.StringVar(&say, "say", "さよなら", "enter the sentence to speak")
	flag.StringVar(&voiceId, "voiceId", "Takumi", "enter the voice Id")
	flag.StringVar(&audioFolderName, "audioFolderName", "audio", "enter the name for the audio to be saved into")
	flag.Parse()
	fmt.Println("Converting text into audio for: [ " + say + "]  with voiceId: ["+ voiceId +"] to be saved into folder ["+ audioFolderName +"] inside tmp folder.")

	//create tmp folder
	os.Mkdir("tmp", os.ModePerm)
	//create voice folder
	os.Mkdir("tmp/"+audioFolderName, os.ModePerm)
	//create audio file
	savePollyVoice(say, audioFolderName, say, voiceId)

}
