/**
 * Created by goku on 24/01/2017.
 */

var didScroll;
var lastScrollTop = 0;
var delta = 5;
var navbarHeight = $('header').outerHeight();

$(window).scroll(function(event) {
    didScroll = true;
});

function windowResize(event) {
    var w = window.innerWidth;
    if (w < 512) {
        $("#nav-right-icons").hide();
    }
    else {
        $("#nav-right-icons").show();
    }
}
window.onresize = windowResize;

windowResize(null);

setInterval(function() {
    if (didScroll) {
        hasScrolled();
        didScroll = false;
    }
}, 250);

function hasScrolled() {
    var st = $(this).scrollTop();

    if (Math.abs(lastScrollTop - st) <= delta)
        return;

    if (st > lastScrollTop && st > navbarHeight) {
        $('header').removeClass('nav-down').addClass('nav-up');
    } else {
        if (st + $(window).height() < $(document).height()) {
            $('header').removeClass('nav-up').addClass('nav-down');
        }
    }

    lastScrollTop = st;
}