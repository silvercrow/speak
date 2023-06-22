
# Speak

Go lang console japanese text to amazon polly speech audio

## Install AWS CLI
You need AWS CLI on your computer and access to polly by default Im running polly in us-east-1

1)

[https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

2)
[https://aws.amazon.com/polly/](https://aws.amazon.com/polly/)


## 🚀 Run

```
go run speak.go -say こんにちわ

```
or

```
go run speak.go -say="こんにちわ"

```

## 🏁 check the result inside the audio folder

```
.
└── tmp
    └── audio
        └── さよなら.mp3

cd tmp/audio/

```

## 🚀 Run with other languages
check [https://docs.aws.amazon.com/polly/latest/dg/voicelist.html](https://docs.aws.amazon.com/polly/latest/dg/voicelist.html) for more voiceIds

```
go run speak.go -say="bonjour je m'appelle Celine" -voiceId Celine

```
## 🏁 check the result inside the audio folder

```
.
└── tmp
    └── audio
        └── bonjour je m'appelle Celine.mp3

cd tmp/audio/

```
