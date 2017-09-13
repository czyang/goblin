{
    "title": "Strange Bug in iOS Webview",
    "author": "Chengzhi Yang",
    "createDate": "2017-04-21",
    "modifyDate": "2017-04-21",
    "permanent":"pixi-ios-webview-bug"
}

# Strange Bug in iOS Third Party Browser

I'm working on an HTML 5 web game, need to compatible with the iOS device. It's used the PIXI.js for render engine. Game detect the browser window size to suit different resolutions and different orientations for mobile. Everything works well on develop phase, but when I test it in real iOS device weird bug occurs...

This game is full page display, works well on Safari browser. But in another browser like Chrome, it will get the wrong size of the current window. This happened on resize event triggered, for the mobile browser, it's orientation change. For my situation I get the windows size [width1, height1] for the first time this is correct values, the game works fine on the point. But after an orientation change I get [width2, height2] it's wrong values!! it's far more small than the real window size, and if device rotates back [width1, height1] is expected but it will get [width3, height3] also smaller wrong values. It's wired and confusing!

With my friend's help who is a web dev guy, I solved the bug by a magic way. He gives me a thought wich use an invisible huge size container to inflate the page. Puff! The bug is gone!

Juse add this magic code to the page:
```html
<!--Deal with iOS WebView Bug.-->
<div id="Magic" style="width:2000px;height:2000px;overflow:hidden;"></div>
```

This is a full test page code, test_rotate.html
```html
<!DOCTYPE HTML>
<html>
<head>
  <title>test-rotate</title>
  <meta charset="utf-8">
  <meta name="apple-mobile-web-app-capable" content="yes" />
  <meta name="viewport" content="width=device-width,height=device-height,initial-scale=1.0"/>

  <script src="pixi.min.js"></script>

</head>
<body style='margin:0;'>
  <script>
    var size = [1280, 600];
    var rsize = [0, 0];

    size[0] = window.innerWidth;
    size[1] = window.innerHeight;
    var renderer = PIXI.autoDetectRenderer(size[0], size[1], {antialias: false, transparent: false, resolution: 1});
    document.body.appendChild(renderer.view);

    var stage = new PIXI.Stage(0x000000, true);

    function resize() {
      var w = window.innerWidth;
      var h = window.innerHeight;
      
      renderer.resize(w, h);
      //alert("resize1 " + w + " " + h +  ' ?' + screen.width + ' ' + screen.height);

      console.log("resize", w, h);
    }
    window.onresize = resize;

    requestAnimationFrame(gameLoop);
    resize();

    function gameLoop() {
      requestAnimationFrame(gameLoop);
      renderer.render(stage);
    }

  </script>
  <!--Deal with iOS WebView Bug.-->
  <div id="Magic" style="width:2000px;height:2000px;overflow:hidden;"></div>
</body>
</html>
```

