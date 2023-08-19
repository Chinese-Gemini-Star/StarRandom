function showFormRange() {
    let value = $("#count").val();
    $("#showRange").text(value);
}

function initFormRange() {
    $("#count").val('1');
    showFormRange();
}

$(function () {
    $("#result").hide();
    initFormRange();
});

function ajaxGet() {
    // 获取表单数据
    let data = {};
    $.each($("#randomInfo").serializeArray(), function (index, item) {
        let value = item.value;

        // 类型转换
        if (item.name === "left" || item.name === "right" || item.name === "number") {
            value = parseInt(value)
        } else if (item.name === "isRepeat") {
            value = value === "true"
        }

        data[item.name] = value;
    });

    // 表单验证
    if (data.left > data.right) {
        alert("最小值不应当大于最大值");
        return false;
    }

    // 转为json格式
    let json = JSON.stringify(data);
    // console.log(json);

    // ajax get请求服务器
    $.ajax({
        url: "/postRandom",
        type: "POST",
        async: true,
        data: json,
        success: function (result) {
            // console.log(result.randoms);
            $("#randomInfo").fadeOut("slow", function () {
                $("#result").fadeIn("slow");
            });

            // 显示随机数结果
            $("#randoms").text("");
            $.each(result.randoms, function (index, item) {
                $("#randoms").text($("#randoms").text() + " " + item);
            })
        },
        error: function (xhr) {
            // console.log(xhr);
            if (xhr.status === 404) {
                window.location.href = "/404.html";
            } else if (xhr.status === 502) {
                window.location.href = "/502.html";
            } else if (xhr.status === 500) {
                window.location.href = "/500.html";
            } else {
                alert("请求发送失败,请检查您的浏览器状态与网络状态.");
            }
        },
    })

    // 阻止表单提交
    return false;
}

function back() {
    $("#result").fadeOut("slow", function () {
        $("#randomInfo").fadeIn("slow");
    });
}