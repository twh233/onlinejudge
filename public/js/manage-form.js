(function(){

    root = "/maintenance/"

    function get_time_interval(val) {
        if (val.length <= 0) {
            alert("亲，至少给个时间啊！")
            return null
        };

        var num = Number(val)
        
        if (isNaN(num)) {
            alert("亲，你确定你输入的是数字？")
            return null
        };


        if (num < 0 ) {
            alert("哥，为啥是负数？")
            return null
        }
    
        return num
    };

    function filter_record(data) {
        if (data == "true" ){
            $("#Record").val("成功！")
        } else {
            resp = JSON.parse(data)
            var str = ""
            for ( i in resp){
                str = str + "【" + i + "】:" + resp[i] + "\n"
            }

            $("#Record").val(str)
        }
    };

    function CloseServer() {
        var t = $("#close_time").val();
        var num = get_time_interval(t)
        if (num == null) {
            alert("缺少参数：时间")
            return
        }

        var servers = get_checked();
        if (servers.length <= 0) {
            alert("缺少参数：目标服务器")
            return 
        }

        var args = {
            servers : JSON.stringify(servers),
            num     : num,
            cmd     : "close_login"
        }

        $("#OP").val("关服倒计时（禁玩法）. 参数:" + num);
        $("#Record").val("wait  close_login!")
        $.get(root + "Notify",args ,function(data, b){
            filter_record(data)
        })
    };

    function NotifyServer() {
        var t = $("#notify_time").val();
        var num = get_time_interval(t)
        if (num == null) {
            alert("缺少参数：时间")
            return
        }

        var servers = get_checked();
        if (servers.length <= 0) {
            alert("缺少参数：目标服务器")
            return 
        }

        var args = {
            servers : JSON.stringify(servers),
            num     : num,
            cmd     : "notify_offline"
        }

        $("#OP").val("发送关服倒计时公告. 参数:" + num);
        $("#Record").val("wait  notify result!")
        $.get(root + "Notify",args ,function(data, b){
            filter_record(data)
        })
    };

    function get_checked(){
        var list = new Array()
        var muti_select = false
        $("table :checkbox").each(function(key, val){
                muti_select = true
                if( $(val).prop('checked')) {
                    list.push(parseInt($(val).prop("id")))
                }
            });

        if( !muti_select) {
            var val = $("#serverid").val();
            //alert("serverid :" + val)
            list.push(parseInt(val))
        }
        return list
    };

    function Shutdown() {
        var servers = get_checked();
        if (servers.length <= 0) {
            alert("缺少参数：目标服务器")
            return 
        }

        var args = {
            servers : JSON.stringify(servers),
            num     : 0,
            cmd     : "shutdown"
        }

        $("#OP").val("关服");
        $("#Record").val("wait  shit_down!")
        $.get(root + "Notify",args ,function(data, b){
            filter_record(data)
        })
    };

    function CoverSync() {
        var servers = get_checked();
        if (servers.length <= 0) {
            alert("缺少参数：目标服务器")
            return 
        }

        var args = {
            servers : JSON.stringify(servers),
            package : $("#package").val(),
        }

        $("#OP").val("同步代码");
        $("#Record").val("wait  sync code!")
        $.get(root + "CoverSync", args ,function(data, b){
            if (data != "running") {
                filter_record(data)
                return
            }
        });

        
        var interval = window.setInterval(function(){
            $.get(root + "Result", {}, function(data){
                if (data != "running") {
                    filter_record(data);
                    window.clearInterval(interval);
                    return
                }
            })
        }, 5000);
    }


    function AllSync() {
        var servers = get_checked();
        if (servers.length <= 0) {
            alert("缺少参数：目标服务器")
            return 
        }

        var args = {
            servers : JSON.stringify(servers),
            package : $("#package").val(),
        }

        $("#OP").val("同步代码");
        $("#Record").val("wait  sync code!")
        $.get(root + "AllSync", args ,function(data, b){
            if (data != "running") {
                filter_record(data)
                return
            }
        });

        
        var interval = window.setInterval(function(){
            $.get(root + "Result", {}, function(data){
                if (data != "running") {
                    filter_record(data);
                    window.clearInterval(interval);
                    return
                }
            })
        }, 5000);
    }

    function Start() {
        var servers = get_checked();
        if (servers.length <= 0) {
            alert("缺少参数：目标服务器")
            return 
        }

        var args = {
            servers : JSON.stringify(servers),
        }

        $("#OP").val("开服");
        $("#Record").val("wait  run!")
        $.get(root + "Start", args ,function(data, b){
            if (data != "running") {
                filter_record(data)
                return
            }
        })

        var interval = window.setInterval(function(){
            $.get(root + "Result", {}, function(data){
                if (data != "running") {
                    filter_record(data);
                    window.clearInterval(interval);
                    return
                }
            })
        }, 5000);
    };

    $("#btn-close").unbind('click').click(function(){
        CloseServer();
    });

    $("#btn-notify").unbind('click').click(function() {
        NotifyServer()
    });

    $("#close").unbind('click').click(function(){
        Shutdown()
    });

    $("#all-sync").unbind('click').click(function(){
        AllSync()
    });
    
    $("#cover-sync").unbind('click').click(function(){
        CoverSync()
    });

    $("#open").unbind('click').click(function() {
        Start()
    })

})(this);


