
# Speak

Go lang console text to amazon polly speech audio

## Install AWS CLI
You need AWS CLI on your computer and access to polly by default Im running polly in us-east-1

1)

[https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)

2)
[https://aws.amazon.com/polly/](https://aws.amazon.com/polly/)


## 🚀 Run
By default im using the voiceId Takumi to translate to japanese but it can be changed to other languages keep reading to learn how

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
      └── さよなら.mp3

cd audio/

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
      └── bonjour je m'appelle Celine.mp3

cd audio/

```
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://bmc.link/tquickbrownfox)
