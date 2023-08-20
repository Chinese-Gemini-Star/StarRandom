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
        let isAllNumber = new RegExp(/^\d{1,}$/);
        // 类型转换
        if (item.name === "left" || item.name === "right" || item.name === "number") {
            if (isAllNumber.test(value)) {
                value = parseInt(value)
            }
            else {
                value = NaN;
            }
        } else if (item.name === "isRepeat") {
            value = value === "true"
        }

        data[item.name] = value;
    });

    // 表单验证
    if (isNaN(data.left) || isNaN(data.right) || data.left === null || data.right === null) {
        // 输入不合法
        alert("请输入整数");
        return false;
    }
    if (data.left > data.right) {
        // 范围不成立
        alert("最小值不应当大于最大值");
        return false;
    }
    if (data.left < -1000 || data.right > 1000) {
        // 范围过大
        alert("请求的随机数范围过大,因为某人太懒了,请输入-1000到1000以内的范围");
        return false;
    }
    if (!data.isRepeat && data.number > data.right - data.left + 1) {
        // 范围内不可能不重复
        alert("此范围不可能不重复");
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