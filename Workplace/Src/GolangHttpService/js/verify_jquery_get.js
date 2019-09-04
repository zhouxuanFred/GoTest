function verifyJqueryGet() {
    $.get("http://localhost:9001/index?username=zhouxuan&password=zhouxuan",
        null,
        function (data) {
            $('#result').html("<p>"+data+"</p>");
        });
}