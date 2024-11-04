document.addEventListener("DOMContentLoaded", function() {
    const counterDisplay = document.getElementById("counter-display");
    const counterInput = document.getElementById("counter-input");

    function setCookie(name, value) {
        document.cookie = name + "=" + (value || "");
    }

    function getCookie(name) {
        var matches = document.cookie.match(new RegExp("(?:^|; )" + name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1') + "=([^;]*)"));
        return matches ? decodeURIComponent(matches[1]) : undefined;
    }

    let counter = parseInt(getCookie("counter")) || 1;

    document.getElementById("increment").addEventListener("click", function() {
        if (counter < 2147483646){
            counter++;
        }
        updateCounterDisplay();
    });

    document.getElementById("decrement").addEventListener("click", function() {
        if (counter > 1) {
            counter--;
        }
        updateCounterDisplay();
    });

    function updateCounterDisplay() {
        counterDisplay.textContent = "количество воркеров: " + counter;
        counterInput.value = counter;
        setCookie("counter", counter);
    }
    updateCounterDisplay();
});
