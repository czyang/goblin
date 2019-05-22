{
    "title": "How to Make a Voice Control Game on iPhone",
    "author": "Chengzhi Yang",
    "createDate": "2011-08-02",
    "modifyDate": "2011-08-02",
    "permanent":"how-to-make-a-voice-control-game-on-iphone"
}

# How to Make a Voice Control Game on iPhone

Some of you might play several game using voice to control the game, say “ahhh” to move the role and shout “pah” to shoot something out…

If you hadn’t played a game like that, you must try my game: [Voice Fighter](http://itunes.apple.com/us/app/voice-fighter/id449643608?ls=1&mt=8)

This article will teach you how to use the iPhone SDK to develop the voice control part of these kinds of game. You can use the code from this article to make your own game. But this article will not show you how to make the graphic part of the game. You can use [Cocos2d](http://www.cocos2d-iphone.org/) or other game engine to make the rest of the game.

“Voice Control” in this point is not about the “speech recognition” but like the “voice operated switch”. These games receive the information from microphone and do some simple analysis，We will use AVAudioRecorder to receive the information. You must be aware that doing this work has a lot of ways, you can find more information in iphone SDK documents, as I know use AVAudioRecorder is the simplest way, so let’s start!

Here is the class(Objective-C) from my game Voice Fighter:

AVMicInput.h
```objectivec
//
//  AVMicInput.h
//  voicefighter
//
//  Created by Chengzhi Yang on 11-6-25.
//  http://codeand.me.
//  Copyright 2011 All rights reserved.
//
#import <Foundation/Foundation.h>
#import <AVFoundation/AVFoundation.h>
#import <CoreAudio/CoreAudioTypes.h>

@interface AVMicInput : NSObject {
    AVAudioRecorder *recorder;
    NSTimer *levelTimer;
}

-(float) micAveragePower;

@end
```

AVMicInput.m

```objectivec
//
//  AVMicInput.m
//  voicefighter
//
//  Created by Chengzhi Yang on 11-6-25.
//  http://codeand.me.
//  Copyright 2011 All rights reserved.
#import "AVMicInput.h"
@implementation AVMicInput
-(id) init
{
    if( (self=[super init]) ) {
        NSURL *url = [NSURL fileURLWithPath:@"/dev/null"];

        NSDictionary *settings = [NSDictionary dictionaryWithObjectsAndKeys:
                                  [NSNumber numberWithFloat: 44100.0],                 AVSampleRateKey,
                                  [NSNumber numberWithInt: kAudioFormatAppleLossless], AVFormatIDKey,
                                  [NSNumber numberWithInt: 1],                         AVNumberOfChannelsKey,
                                  [NSNumber numberWithInt: AVAudioQualityMin],         AVEncoderAudioQualityKey,
                                  nil];
        NSError *error;

        recorder = [[AVAudioRecorder alloc] initWithURL:url settings:settings error:&error];

        if (recorder) {
            [recorder prepareToRecord];
            recorder.meteringEnabled = YES;
            [recorder record];
        } else
            NSLog(@"%@",[error description]);
    }
    return self;
}

- (void) dealloc {
    [recorder dealloc];
    [super dealloc];
}

-(float) micAveragePower
{
    [recorder updateMeters];
    float avgPower = [recorder averagePowerForChannel:0];
    //NSLog(@"%f",avgPower);
    return avgPower;
}

@end
```

You can add AVMicInput.h and AVMicInput.m to your project without any change. This class will setup the AVAudioRecorder for you. You may notice the strings “/dev/null”. In Unix-like operating systems, “/dev/null” is a special file that discards all data written to it (but reports that the write operation succeeded) and provides no data to any process that reads from it [1](http://en.wikipedia.org/wiki//dev/null), iPhone OS is one of Unix-like operating systems. Other codes will explain in rest of this article.

Put the following code to your initializer to initialize the AVMicInput class:

```objectivec
AVMicInput *avMicIn;
avMicIn = [[AVMicInput alloc] init];
```

In your main loop or where you want to get the information from microphone add the following code:
```objectivec
float micIn = [avMicIn micAveragePower];
```

If the AVMicInput was correctly initialized, micAveragePower will return a float that is the immediate dB(Decibel) information(you can also understand the information into digital signal) from microphone, the range of the return value micIn is [-120, 0]. -120 is the absolutely quiet, 0 is the maximum value(if the voice is very loud, such as a shout, you could get 0). I make a list to help you understand the micIn:

1. If your microphone was broken, iPhone or your Mac, you will always get -120 on your iphone or simulator.

2. In most cases, a silent environment. The range of micIn is about -30 to -50, you can test by yourself.

3. If you say “ahhh” or say something softly, the range of micIn is about -30 to -10.

4. If you speak loudly or shout, the range of micIn is -10 to 0.

In my game [Voice Fighter](http://itunes.apple.com/us/app/voice-fighter/id449643608?ls=1&mt=8), through a lot of test and tweak, I used the below code to determine which state of value I got, is “ahhh” or “pah”(It’s not necessary to use these two incantation, “HA”, “PEW”, “BANG”, “KAMEHAMEHA”… is available)

```objectivec
if ( micIn > -6 ){
    //We got a "Pah"!
} else if ( micIn >= -25 && micIn < -6  ) {
    //we got a "Ahhhhhh"
}
```

You can tweak the threshold by your self.

Others you need know:

1. For now time, most iOS device has a microphone. I’m write the Voice Fighter use a ipod touch 4, but the older generation ipod may need a plug-in microphone to let you play with the microphone.

2. If you use Cocos2d to make your game or in your own engine you used the cocosDension to play the sound like this: [[SimpleAudioEngine sharedEngine] playEffect:@”1.aifc”]; You will noticed if you play a sound the AVAudioRecorder will stop work. For this situation, You can solve the problem by this method:

(1). Add #import “CDAudioManager.h” to AppDelegate.m .

(2). Add [CDAudioManager initAsynchronously:kAMM_PlayAndRecord]; to – (void) applicationDidFinishLaunching:(UIApplication*)application .

And you can find more solution in here. In that forum czyang is me :)

1. If you want to make a voice control game on Android platform please think twice. Theoretically make a voice control game on Android could use the same way, but I was made the Voice Fighter for Android(here) in same way. The input of microphone was totally different! Thanks to separate hardware on android. Each type of android phone have different hardware, and information from microphone on these device may have a huge different. The Voice Fighter for Android just working normally on a bit type of android phone, so I got a lot of 1 star in market. For now I don’t have a solution about this problem.

Now you got all of the “voice control” things. You can use these code to make a voice control game by your self.

If you feel this article is good for you, don’t forget to buy a Voice Fighter or tell your friends about this game for support me :)

BTW, Now I have a few promo code of voice fighter left. I will giveaway 5 of them. If you want to get one, please fill your email correctly and leave a reply follow this article with “Show me your code”. I will send the promo code to your email in 24 hours.
