function showFormRange() {
    let value = $("#count").val();
    $("#showRange").text(value);
}

function initFormRange() {
    $("#count").val('1');
    showFormRange();
}

$(function () {
    initFormRange();
});

function ajaxGet() {
    // 获取表单数据
    let data = {};
    $.each($("#main").serializeArray(), function (index, item) {
        data[item.name] = item.value;
    });

    // 转为json格式
    let json = JSON.stringify(data);
    // console.log(json);

    // ajax get请求服务器
    $.ajax({
        // url: "/postRandom",
        type: "POST",
        async: true,
        data: json,
        success: function (result) {

        },
        error: function (xhr) {
            // console.log(xhr);
            if (xhr.status === 404) {
                window.location.href = "/404.html";
            } else if (xhr.status === 502) {
                window.location.href = "/502.html";
            } else {
                alert("请求发送失败,请检查您的浏览器状态与网络状态.");
            }
        },
    })

    // 阻止表单提交
    return false;
}