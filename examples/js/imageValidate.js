var left = 0;
var id = "";
$(function () {
    // 初始化图片验证码
    initImageValidate();
    /* 初始化按钮拖动事件 */
    // 鼠标点击事件
    $("#sliderInner").mousedown(function () {
        // 鼠标移动事件
        document.onmousemove = function (ev) {
            left = ev.clientX;
            if (left >= 100 && left <= 563) {
                $("#sliderInner").css("left", (left - 100) + "px");
                $("#slideImage").css("left", (left - 100) + "px");
            }
        };
        // 鼠标松开事件
        document.onmouseup = function () {
            document.onmousemove = null;
            checkImageValidate();
        };
    });

});

function initImageValidate() {
    $.ajax({
        async: false,
        type: "GET",
        url: "http://localhost:8080/getImgTest",
        dataType: "json",
        data: {
            telephone: telephone
        },
        success: function (data) {
            // 设置图片的src属性
            $("#validateImage").attr("src", "data:image/png;base64," + data.im);
            $("#slideImage").attr("src", "data:image/png;base64," + data.imSlide);
            $("#slideImage").css("top", data.y);
            id = data.id;
        },
        error: function () {
        }
    });
}

function exchange() {
    initImageValidate();
}

// 校验
function checkImageValidate() {
    $.ajax({
        async: false,
        type: "Get",
        url: "http://localhost:8080/check?id=" + id + "&left=" + left,
        dataType: "json",
        success: function (data) {
            if (data.code < 400) {
                $("#operateResult").html(data.info).css("color", "#28a745");
                // 校验通过，调用发送短信的函数
                console.log("验证通过");
            } else {
                $("#operateResult").html(data.info).css("color", "#dc3545");
                // 验证未通过，将按钮和拼图恢复至原位置
                $("#sliderInner").animate({"left": "0px"}, 200);
                $("#slideImage").animate({"left": "0px"}, 200);
            }
        },
        error: function () {
        }
    });
}